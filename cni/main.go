package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/ip"
	"github.com/containernetworking/plugins/pkg/ipam"
	"github.com/containernetworking/plugins/pkg/ns"
	lbrp "github.com/ekoops/polykube-operator/polycube/clients/lbrp"
	"github.com/ekoops/polykube-operator/utils"
	log "github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"strings"
)

const (
	basePath = "http://127.0.0.1:9000/polycube/v1"
)

var (
	lbrpAPI *lbrp.LbrpApiService
)

func init() {
	log.SetLevel(log.DebugLevel)
	file, err := os.OpenFile("polykube-cni-plugin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		// I have not to log anything since stderr and stdout are reserved for CNI specification purposes.
		log.SetOutput(ioutil.Discard)
	}

	// this ensures that main runs only on main thread (thread group leader).
	// since namespace ops (unshare, setns) are done for a single thread, we
	// must ensure that the goroutine does not jump from OS thread to thread
	runtime.LockOSThread()

	// init lbrp API
	cfgLbrp := lbrp.Configuration{BasePath: basePath}
	srLbrp := lbrp.NewAPIClient(&cfgLbrp)
	lbrpAPI = srLbrp.LbrpApi
}

// setupVeth creates a veth pair using the provided ends names and puts the container end inside the provided netns. It
// returns a couple containing the container interface and the host interface info.
func setupVeth(netns ns.NetNS, contIfName, hostIfName string, mtu int) (*current.Interface, *current.Interface, error) {
	contIface := &current.Interface{}
	hostIface := &current.Interface{}

	err := netns.Do(func(hostNS ns.NetNS) error {
		// creating the veth pair in the container and moving host end into host netns
		hostVeth, contVeth, err := ip.SetupVethWithName(contIfName, hostIfName, mtu, "", hostNS)
		if err != nil {
			return err
		}
		contIface.Name = contVeth.Name
		contIface.Mac = contVeth.HardwareAddr.String()
		contIface.Sandbox = netns.Path()
		hostIface.Name = hostVeth.Name
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	// need to lookup hostVeth again as its index has changed during ns move
	hostVeth, err := netlink.LinkByName(hostIface.Name)
	if err != nil {
		return nil, nil, fmt.Errorf("failed %q lookup: %v", hostIface.Name, err)
	}
	hostIface.Mac = hostVeth.Attrs().HardwareAddr.String()

	return hostIface, contIface, nil
}

// loadNetConf loads the network configuration from stdin (including the prevResult)
func loadNetConf(stdin []byte) (*NetConf, error) {
	conf := &NetConf{}

	if err := json.Unmarshal(stdin, &conf); err != nil {
		return nil, fmt.Errorf("failed to parse network configuration: %v", err)
	}

	// parsing previous result
	if err := version.ParsePrevResult(&conf.NetConf); err != nil {
		return nil, fmt.Errorf("failed to parse prevResult: %v", err)
	}

	// validating netconf field
	if conf.MTU == 0 {
		return nil, errors.New("MTU must be specified")
	}

	if conf.LbrpName == "" {
		return nil, errors.New("internal lbrp name must be specified")
	}

	logLevelStr := strings.ToLower(conf.LogLevel)
	if logLevelStr == "off" {
		log.SetOutput(ioutil.Discard)
	} else {
		logLevel, err := log.ParseLevel(logLevelStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse log level: %v", err)
		}
		log.SetLevel(logLevel)
	}

	if conf.Gw.IP == nil || conf.Gw.IP.To4() == nil {
		return nil, errors.New("the gateway IP must be an IPv4 address")
	}

	// the mac address is put inside the field RawMAC as a string: it has to be parsed
	// and the result assigned to the field MAC
	hwAddr, err := net.ParseMAC(conf.Gw.RawMAC)
	if err != nil {
		return nil, fmt.Errorf("failed to parse gateway MAC address: %v", err)
	}
	conf.Gw.MAC = hwAddr

	return conf, nil
}

// configureNetns configures the provided address on the provided interface inside the provided netns. In addition,
// also the netns routing table and ARP table is properly configured.
func configureNetns(netns ns.NetNS, ifName string, address *net.IPNet, gwInfo *GwInfo) error {
	if err := netns.Do(func(_ ns.NetNS) error {
		// retrieving the veth interface
		link, err := netlink.LinkByName(ifName)
		if err != nil {
			return fmt.Errorf("failed to lookup iface: %v", err)
		}

		// adding a /32 IPv4 address to the interface and setting the peer to the default gateway address
		allOnesMask := net.IPv4Mask(255, 255, 255, 255)
		addr := &netlink.Addr{
			IPNet: &net.IPNet{
				IP:   address.IP,
				Mask: allOnesMask,
			},
			Label: "",
			Peer: &net.IPNet{
				IP:   gwInfo.IP,
				Mask: allOnesMask,
			},
		}
		if err = netlink.AddrAdd(link, addr); err != nil {
			return fmt.Errorf("failed to set IPv4 address on iface: %v", err)
		}

		// adding default route
		route := &netlink.Route{
			Dst: nil,
			Gw:  gwInfo.IP,
		}
		if err := netlink.RouteAdd(route); err != nil {
			return fmt.Errorf("failed to add default route: %v", err)
		}

		// adding arp entry for default gateway
		arpentry := &netlink.Neigh{
			LinkIndex:    link.Attrs().Index,
			State:        netlink.NUD_PERMANENT,
			IP:           gwInfo.IP,
			HardwareAddr: gwInfo.MAC,
		}
		if err := netlink.NeighAdd(arpentry); err != nil {
			return fmt.Errorf("failed to add static arp entry for default gateway: %v", err)
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// checkIface verifies that the provided interface is correctly configured.
func checkIface(lg *log.Entry, netnsName string, iface *IFaceConf) error {
	ifaceName := iface.Interface.Name
	l := lg.WithField("iface", ifaceName)

	// obtaining interface corresponding link
	link, err := netlink.LinkByName(ifaceName)
	if err != nil {
		if _, notFound := err.(netlink.LinkNotFoundError); notFound {
			l.WithField("detail", err).Fatal("iface doesn't exist")
			return fmt.Errorf("%q iface doesn't exist into %q netns: %v", ifaceName, netnsName, err)
		}
		l.WithField("detail", err).Fatal("failed iface lookup")
		return fmt.Errorf("failed %q iface lookup into %q netns: %v", ifaceName, netnsName, err)
	}

	// if no ip configuration are expected to be configured on link, simply return
	if iface.IPConf == nil {
		l.Info("netns iface checked")
		return nil
	}

	// obtaining addresses configured on link
	addrs, err := netlink.AddrList(link, netlink.FAMILY_V4)
	if err != nil {
		l.WithField("detail", err).Fatal("failed iface addresses lookup")
		return fmt.Errorf("failed %q iface addresses lookup into %q netns: %v", ifaceName, netnsName, err)
	}

	// checking if the ip addresses are correctly configured on link
	for _, addr := range addrs {
		if addr.IPNet.String() == iface.IPConf.Address.String() {
			l.Info("netns iface checked")
			return nil
		}
	}
	l.WithField("detail", err).Fatal("iface ip misconfiguration")
	return fmt.Errorf("%q iface ip misconfiguration into %q netns: %v", ifaceName, netnsName, err)
}

// getIfaceConfs scans the prevResult.Interfaces in order to find the expected container and host interface created
// during the ADD operation. If the two interfaces are found, they are returned in association with their IPConf
func getIfaceConfs(prevResult *current.Result, contIfName, netnsName string) (*IFaceConf, *IFaceConf, error) {
	var contIfaceConf, hostIfaceConf *IFaceConf
	// scanning all prevResult interfaces
	for i, iface := range prevResult.Interfaces {
		// is this the container interface?
		if iface.Name == contIfName && (netnsName == "" || netnsName == iface.Sandbox) {
			contIfaceConf = &IFaceConf{
				ResultIndex: i,
				Interface:   iface,
			}
			// thc host interface is placed immediately after the container interface
			hostIfaceConf = &IFaceConf{
				ResultIndex: i + 1,
				Interface:   prevResult.Interfaces[i+1],
			}
			break
		}
	}
	// if the two interfaces were not found return an error
	if contIfaceConf == nil {
		return nil, nil, errors.New("unexpected interfaces: wrong or missing")
	}
	// scanning prevResult.IPs in order to associate the container interface to its ip configuration
	for _, ipConf := range prevResult.IPs {
		if *ipConf.Interface == hostIfaceConf.ResultIndex {
			ipConf.Address.IP = ipConf.Address.IP.To4()
			if ipConf.Address.IP == nil {
				return nil, nil, errors.New("unexpected non IPv4 address configured on interface")
			}
			hostIfaceConf.IPConf = ipConf
		}
		if *ipConf.Interface == contIfaceConf.ResultIndex {
			ipConf.Address.IP = ipConf.Address.IP.To4()
			if ipConf.Address.IP == nil {
				return nil, nil, errors.New("unexpected non IPv4 address configured on interface")
			}
			contIfaceConf.IPConf = ipConf
		}
	}
	// checking that an ip configuration for the container interface and no configuration
	// for the host interface were found
	if contIfaceConf.IPConf == nil || hostIfaceConf.IPConf != nil {
		return nil, nil, errors.New("unexpected ip configurations: wrong or missing")
	}
	return contIfaceConf, hostIfaceConf, nil
}

func cmdAdd(args *skel.CmdArgs) error {
	// Defining the attachment identifier and the base logger. For the attachment, a truncation
	// of ifName_containerId[0:10] up to 15 characters is used (since it is the max possible)
	att := utils.CreateAttachment(args.IfName, args.ContainerID)
	l := log.WithField("id", "ADD_"+att)

	// parsing configuration
	conf, err := loadNetConf(args.StdinData)
	if err != nil {
		l.WithFields(log.Fields{"subject": "netconf", "detail": err}).Fatal("parsing failed")
		return fmt.Errorf("failed to parse netconf: %v", err)
	}

	// parsing prevResult, if present
	var prevResult *current.Result
	if conf.PrevResult != nil {
		if prevResult, err = current.NewResultFromResult(conf.PrevResult); err != nil {
			l.WithField("detail", err).Fatal("failed to convert prevResult to current version")
			return fmt.Errorf("failed to convert prevResult into current version: %v", err)
		}
	}

	// getting netns handle
	netns, err := ns.GetNS(args.Netns)
	if err != nil {
		l.WithFields(log.Fields{"netns": args.Netns, "detail": err}).Fatal("failed to open netns")
		return fmt.Errorf("failed to open %q netns: %v", args.Netns, err)
	}
	defer netns.Close()

	// checking if the specified iface already exists in the specified netns
	if err := netns.Do(func(_ ns.NetNS) error {
		_, err := netlink.LinkByName(args.IfName)
		if err == nil {
			return errors.New("iface already exists")
		}
		if _, notFound := err.(netlink.LinkNotFoundError); !notFound {
			return fmt.Errorf("failed iface lookup: %v", err)
		}
		return nil
	}); err != nil {
		l.WithFields(
			log.Fields{"netns": args.Netns, "iface": args.IfName, "detail": err},
		).Fatal("error during iface existence checking")
		return fmt.Errorf("error during checking %q iface existence into %q netns: %v", args.IfName, args.Netns, err)
	}

	// getting IP from ipam plugin. Even if the returned IPv4 has a /24 prefix length, eventually a /32 will
	// be configured on the container interface
	addr, err := allocIP(conf.IPAM.Type, args.StdinData)
	if err != nil {
		l.WithFields(log.Fields{"scope": "ipam", "detail": err}).Fatal("failed to get ip")
		return fmt.Errorf("failed to get ip through ipam plugin: %v", err)
	}
	l.WithFields(log.Fields{"ip": addr.IP, "netmask": addr.Mask}).Info("ip allocated")

	// setting up the veth pair
	contIfaceName := args.IfName
	ipHexStr := hex.EncodeToString(addr.IP)
	hostIfaceName := utils.GetHostIfaceName(contIfaceName, ipHexStr)

	hostIface, contIface, err := setupVeth(netns, contIfaceName, hostIfaceName, conf.MTU)
	if err != nil {
		l.WithFields(
			log.Fields{"netns": args.Netns, "iface": args.IfName, "mtu": conf.MTU, "detail": err},
		).Fatal("failed to setup veth pair")
		return fmt.Errorf("failed to setup veth pair: %v", err)
	}
	l.WithFields(log.Fields{
		"hostIface": fmt.Sprintf("%+v", hostIface),
		"contIface": fmt.Sprintf("%+v", contIface),
		"mtu":       conf.MTU,
	}).Info("veth pair created")

	// configuring netns
	nlog := l.WithFields(log.Fields{
		"netns":   args.Netns,
		"iface":   contIfaceName,
		"address": fmt.Sprintf("%+v", addr),
		"gateway": fmt.Sprintf("%+v", conf.Gw),
	})
	if err := configureNetns(netns, contIfaceName, addr, &conf.Gw); err != nil {
		nlog.WithField("detail", err).Fatal("failed to configure the netns")
		return fmt.Errorf("failed to configure the %q netns: %v", args.Netns, err)
	}
	nlog.Info("netns configured")

	// creating lbrp port and connecting it to hostInterface
	lbrpName := conf.LbrpName
	lbrpPortName := ipHexStr
	port := lbrp.Ports{
		Name:  lbrpPortName,
		Type_: "frontend",
		Ip_:   addr.IP.String(),
		Peer:  hostIfaceName,
	}
	nlog = l.WithFields(log.Fields{
		"lbrp": lbrpName,
		"port": fmt.Sprintf("%+v", port),
	})

	if resp, err := lbrpAPI.CreateLbrpPortsByID(context.TODO(), lbrpName, lbrpPortName, port); err != nil {
		nlog.WithFields(log.Fields{
			"response": fmt.Sprintf("%+v", resp), "detail": err,
		}).Fatal("failed to create and connect lbrp port")
		return fmt.Errorf(
			"failed to create and connect %q internal lbrp %q port: %v", lbrpName, lbrpPortName, err,
		)
	}
	nlog.WithField(
		"connection", fmt.Sprintf("%s:%s <-> %s", lbrpName, lbrpPortName, hostIfaceName),
	).Info("internal lbrp connected to the host interface")

	// setting up the plugin result
	result := &current.Result{}
	if prevResult != nil {
		result = prevResult
	}
	// The contIface will be placed at index len(result.Interfaces),
	// so the hostIface will be placed at index len(result.Interfaces) + 1
	contIp := &current.IPConfig{
		Interface: current.Int(len(result.Interfaces)), // 0 if the plugins is unchained
		Address:   *addr,
		Gateway:   conf.Gw.IP,
	}
	contRoute := &types.Route{
		Dst: net.IPNet{
			IP:   net.IPv4zero,
			Mask: net.IPv4Mask(0, 0, 0, 0),
		},
		GW: conf.Gw.IP,
	}
	// The order is important! The contIface must precede the hostIface,
	// since this will be exploited in the DEL and CHECK operations
	result.Interfaces = append(result.Interfaces, contIface, hostIface)
	result.IPs = append(result.IPs, contIp)
	result.Routes = append(result.Routes, contRoute)

	return types.PrintResult(result, conf.CNIVersion)
}

// cmdCheck is called for CHECK requests
func cmdCheck(args *skel.CmdArgs) error {
	// Defining the attachment identifier and the base logger. For the attachment, a truncation
	// of ifName_containerId[0:10] up to 15 characters is used (since it is the max possible)
	att := utils.CreateAttachment(args.IfName, args.ContainerID)
	l := log.WithField("id", "CHK_"+att)

	// parsing configuration
	conf, err := loadNetConf(args.StdinData)
	if err != nil {
		l.WithFields(log.Fields{"subject": "netconf", "detail": err}).Fatal("parsing failed")
		return err
	}

	// checking the presence of prevResult (its presence is made mandatory by the CNI specification
	// in order to check the container networking)
	if conf.PrevResult == nil {
		l.WithField("detail", "prevResult must be specified").Fatal("missing configuration")
		return errors.New("missing configuration: prevResult must be specified")
	}
	prevResult, err := current.NewResultFromResult(conf.PrevResult)
	if err != nil {
		l.WithField("detail", err).Fatal("failed to convert prevResult into current version")
		return fmt.Errorf("failed to convert prevResult into current version: %v", err)
	}

	// CHECK on ipam plugin
	err = ipam.ExecCheck(conf.IPAM.Type, args.StdinData)
	if err != nil {
		l.WithFields(log.Fields{"scope": "ipam", "detail": err}).Fatal("CHECK operation failed")
		return fmt.Errorf("CHECK operation failed on ipam plugin: %v", err)
	}
	l.Info("ip checked")

	// getting netns handle
	netns, err := ns.GetNS(args.Netns)
	if err != nil {
		l.WithFields(log.Fields{"netns": args.Netns, "detail": err}).Fatal("failed to open netns")
		return fmt.Errorf("failed to open %q netns: %v", args.Netns, err)
	}
	defer netns.Close()

	// extracting the container interface and the host interface with their own IP configurations
	contIfaceConf, hostIfaceConf, err := getIfaceConfs(prevResult, args.IfName, args.Netns)
	if err != nil {
		l.WithFields(log.Fields{
			"prevResult": fmt.Sprintf("%+v", *prevResult), "detail": err,
		}).Fatal("unexpected prevResult")
		return fmt.Errorf("unexpected prevResult: %v", err)
	}

	// checking args.Netns netns interface and routes
	nlog := l.WithField("netns", args.Netns)
	if err := netns.Do(func(_ ns.NetNS) error {
		if err := checkIface(nlog, args.Netns, contIfaceConf); err != nil {
			return err
		}

		// checking that routes are correctly configured
		if err := ip.ValidateExpectedRoute(prevResult.Routes); err != nil {
			nlog.WithField("detail", err).Fatal("failed netns routes checking")
			return fmt.Errorf("failed %q netns routes checking: %v", args.Netns, err)
		}
		nlog.Info("netns routes checked")
		return nil
	}); err != nil {
		return err
	}
	nlog.Info("netns checked")

	// checking root netns interface
	nlog = l.WithField("netns", "root")
	if err := checkIface(nlog, "root", hostIfaceConf); err != nil {
		return err
	}
	nlog.Info("netns checked")

	// checking lbrp port connection
	lbrpName := conf.LbrpName
	contIP := contIfaceConf.IPConf.Address.IP
	lbrpPortName := hex.EncodeToString(contIP)
	lbrpPortPeer := hostIfaceConf.Interface.Name
	lblog := l.WithField("lbrp", lbrpName) // lbrp logger
	if err := checkLbrpPort(lbrpName, lbrpPortName, lbrpPortPeer, contIP.String()); err != nil {
		lblog.WithField("detail", err).Fatal("failed lbrp checking")
		return fmt.Errorf("failed %q lbrp port checking: %v", lbrpName, err)
	}
	lblog.Info("lbrp checked")
	return nil
}

// cmdDel is called for DELETE requests
func cmdDel(args *skel.CmdArgs) error {
	// Defining the attachment identifier and the base logger. For the attachment, a truncation
	// of ifName_containerId[0:10] up to 15 characters is used (since it is the max possible)
	att := utils.CreateAttachment(args.IfName, args.ContainerID)
	l := log.WithField("id", "DEL_"+att)

	// parsing configuration
	conf, err := loadNetConf(args.StdinData)
	if err != nil {
		l.WithFields(log.Fields{"subject": "netconf", "detail": err}).Fatal("parsing failed")
		return err
	}

	// parsing prevResult
	if conf.PrevResult == nil {
		l.WithField("detail", "prevResult must be specified").Fatal("missing configuration")
		return errors.New("missing configuration: prevResult must be specified")
	}
	prevResult, err := current.NewResultFromResult(conf.PrevResult)
	if err != nil {
		l.WithField("detail", err).Fatal("failed to convert prevResult into current version")
		return fmt.Errorf("failed to convert prevResult into current version: %v", err)
	}

	// extracting the container interface with its own ip configurations
	contIfaceConf, _, err := getIfaceConfs(prevResult, args.IfName, args.Netns)
	if err != nil {
		l.WithFields(log.Fields{
			"prevResult": fmt.Sprintf("%+v", *prevResult), "detail": err,
		}).Fatal("unexpected prevResult")
		return fmt.Errorf("unexpected prevResult: %v", err)
	}
	contIP := contIfaceConf.IPConf.Address.IP

	// releasing IP address
	if err := ipam.ExecDel(conf.IPAM.Type, args.StdinData); err != nil {
		l.WithFields(log.Fields{"scope": "ipam", "detail": err}).Fatal("DEL operation failed")
		return fmt.Errorf("DEL operation failed on ipam plugin: %v", err)
	}
	l.Info("ip released")

	// deleting lbrp port
	lbrpName := conf.LbrpName
	lbrpPortName := hex.EncodeToString(contIP)
	lblog := l.WithFields(log.Fields{"lbrp": lbrpName, "port": lbrpPortName})
	if resp, err := lbrpAPI.DeleteLbrpPortsByID(
		context.TODO(), lbrpName, lbrpPortName,
	); err != nil && resp.StatusCode != 409 {
		lblog.WithFields(log.Fields{
			"response": fmt.Sprintf("%+v", resp), "detail": err,
		}).Fatal("failed to delete internal lbrp port")
		return fmt.Errorf(
			"failed to delete %q port on %q internal lbrp: - error: %s, response: %+v",
			lbrpPortName, lbrpName, err, resp,
		)
	}
	lblog.Info("internal lbrp port deleted")

	// deleting netns iface and related stuff (routes, arpentry, etc...)
	if args.Netns != "" {
		// There is a netns so try to clean up. Delete can be called multiple times
		// so don't return an error if the device is already removed.
		nlog := l.WithFields(log.Fields{"netns": args.Netns, "iface": args.IfName})
		if err := ns.WithNetNSPath(args.Netns, func(_ ns.NetNS) error {
			if err = ip.DelLinkByName(args.IfName); err != nil && err != ip.ErrLinkNotFound {
				// if there is an error different from ip.ErrLinkNotFound, returns error
				return err
			}
			return nil
		}); err != nil {
			// if netns is not found, continue anyway.
			if _, notFound := err.(ns.NSPathNotExistErr); !notFound {
				nlog.Fatal("failed to delete iface")
				return fmt.Errorf("failed to delete iface %q into netns %q: %v", args.IfName, args.Netns, err)
			}
		}
		nlog.Info("netns iface and related stuff (routes, arpentry, etc...) deleted")
	}

	return nil
}

func main() {
	skel.PluginMain(cmdAdd, cmdCheck, cmdDel, version.All, "polykube-cni-plugin")
}
