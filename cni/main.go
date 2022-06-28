package main

import (
	"context"
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
	lbrp "github.com/ekoops/polykube-operator/pkg/clients/lbrp"
	simplebridge "github.com/ekoops/polykube-operator/pkg/clients/simplebridge"
	"github.com/ekoops/polykube-operator/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
	"io/ioutil"
	"net"
	"os"
	"runtime"
)

const (
	basePath = "http://127.0.0.1:9000/polycube/v1"
)

var (
	simplebridgeAPI *simplebridge.SimplebridgeApiService
	lbrpAPI         *lbrp.LbrpApiService
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

	// init simplebridge API
	cfgSimplebridge := simplebridge.Configuration{BasePath: basePath}
	srSimplebridge := simplebridge.NewAPIClient(&cfgSimplebridge)
	simplebridgeAPI = srSimplebridge.SimplebridgeApi

	// init lbrp API
	cfgLbrp := lbrp.Configuration{BasePath: basePath}
	srLbrp := lbrp.NewAPIClient(&cfgLbrp)
	lbrpAPI = srLbrp.LbrpApi
}

func setupVeth(netns ns.NetNS, contIfName string, hostIfName string, mtu int) (*current.Interface, *current.Interface, error) {
	contIface := &current.Interface{}
	hostIface := &current.Interface{}

	err := netns.Do(func(hostNS ns.NetNS) error {

		// create the veth pair in the container and move host end into host netns
		hostVeth, contVeth, err := ip.SetupVethWithName(
			contIfName,
			hostIfName,
			mtu,
			"",
			hostNS,
		)
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
		return nil, nil, fmt.Errorf("failed to lookup %q: %v", hostIface.Name, err)
	}
	hostIface.Mac = hostVeth.Attrs().HardwareAddr.String()

	return hostIface, contIface, nil
}

// loadNetConf load the network configuration from stdin (including the prevResult)
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

	if conf.BridgeName == "" {
		return nil, errors.New("bridge name must be specified")
	}

	if conf.VClusterCIDR == "" {
		return nil, errors.New("VClusterCIDR must be specified")
	}

	if conf.Gw.IP == nil || conf.Gw.IP.To4() == nil {
		return nil, errors.New("the gateway IP must be an ipv4 address")
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

func configureNetns(netns ns.NetNS, ifName string, address *net.IPNet, gwInfo *GwInfo) error {
	if err := netns.Do(func(_ ns.NetNS) error {
		// setting up the veth interface
		link, err := netlink.LinkByName(ifName)
		if err != nil {
			return fmt.Errorf("failed to lookup iface: %v", err)
		}
		if err := netlink.LinkSetUp(link); err != nil {
			return fmt.Errorf("failed to set iface up: %v", err)
		}

		// adding ip4 address to the interface
		addr := &netlink.Addr{IPNet: address, Label: ""}
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

func cmdAdd(args *skel.CmdArgs) error {
	// defining the attachment identifier and the base logger
	att := utils.Truncate(fmt.Sprintf("%s_%s", args.IfName, args.ContainerID[0:10]), 15)
	l := log.WithField("id", fmt.Sprintf("ADD_%s", att))

	// parsing configuration
	conf, err := loadNetConf(args.StdinData)
	if err != nil {
		l.WithFields(log.Fields{
			"subject": "netconf",
			"detail":  err,
		}).Fatal("parsing failed")
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
		l.WithFields(log.Fields{
			"netns":  args.Netns,
			"detail": err,
		}).Fatal("failed to retrieve netns")
		return fmt.Errorf("failed to open netns %q: %v", args.Netns, err)
	}
	defer netns.Close() // TODO why?

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
		l.WithFields(log.Fields{
			"netns":  args.Netns,
			"iface":  args.IfName,
			"detail": err,
		}).Fatal("error during iface existence checking")
		return fmt.Errorf("error during checking iface %q existence into netns %q: %v", args.IfName, args.Netns, err)
	}

	// getting ip from ipam plugin
	addr, err := allocIP(conf.IPAM.Type, args.StdinData)
	if err != nil {
		l.WithFields(log.Fields{
			"scope":  "ipam",
			"detail": err,
		}).Fatal("failed to get ip")
		return fmt.Errorf("failed to get ip through ipam plugin: %v", err)
	}
	l.WithFields(log.Fields{
		"ip":      addr.IP,
		"netmask": addr.Mask,
	}).Info("ip allocated")

	// setting up the veth pair
	// using a truncation of ifName_containerId[0:10] up to 15 characters since it is the max possibile
	// link name dimensione
	hostIface, contIface, err := setupVeth(
		netns,
		args.IfName,
		att,
		conf.MTU,
	)
	if err != nil {
		l.WithFields(log.Fields{
			"netns":  args.Netns,
			"iface":  args.IfName,
			"mtu":    conf.MTU,
			"detail": err,
		}).Fatal("failed to setup veth pair")
		return fmt.Errorf("failed to setup veth pair: %v", err)
	}
	l.WithFields(log.Fields{
		"hostIface": fmt.Sprintf("%+v", hostIface),
		"contIface": fmt.Sprintf("%+v", contIface),
		"mtu":       conf.MTU,
	}).Info("veth pair created")

	// configuring netns
	netnsLgr := l.WithFields(log.Fields{
		"netns":   args.Netns,
		"iface":   args.IfName,
		"address": fmt.Sprintf("%+v", addr),
		"gateway": fmt.Sprintf("%+v", conf.Gw),
	})
	if err := configureNetns(netns, args.IfName, addr, &conf.Gw); err != nil {
		netnsLgr.WithField("detail", err).Fatal("failed to configure the netns")
		return fmt.Errorf("failed to configure the netns %q: %v", args.Netns, err)
	}
	netnsLgr.Info("netns configured")

	// creating lbrp (using pod ip as id, so it can be referenced by operator)
	// and connecting the frontend port to hostInterface
	//lbrpName := fmt.Sprintf("lbrp-%s", addr.IP.String())
	lbName := "lbrp_" + hostIface.Name
	llog := l.WithField("lbrp", lbName)
	if err := createLbrp(lbName, hostIface); err != nil {
		llog.WithField("detail", err).Fatal("failed to create lbrp")
		return fmt.Errorf("failed to create lbrp %q: %v", lbName, err)
	}
	llog.WithField(
		"connection", fmt.Sprintf("%s <-> %s", utils.CreatePeer(lbName, "to_pod"), hostIface.Name),
	).Info("lbrp created and connected to pod")

	// creating bridge port and connect it to the lbrp
	brName := conf.BridgeName
	conlog := l.WithFields(log.Fields{
		"lbrp":   lbName,
		"bridge": brName,
	})
	lbPort, brPort, err := connectLbrpToBridge(lbName, brName)
	if err != nil {
		conlog.WithField("detail", err).Fatal("failed to connect lbrp to bridge")
		return fmt.Errorf("failed to connect %q lbrp to %q bridge: %v", lbName, brName, err)
	}
	conlog.WithField(
		"connection", fmt.Sprintf("%s <-> %s", brPort.Peer, lbPort.Peer),
	).Info("lbrp connected to bridge")

	// setting up the plugin result
	result := &current.Result{}
	if prevResult != nil {
		result = prevResult
	}
	contIp := &current.IPConfig{
		Interface: current.Int(len(result.Interfaces)), // 0 if unchained
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
	result.Interfaces = append(result.Interfaces, contIface, hostIface) // the order is important!
	result.IPs = append(result.IPs, contIp)
	result.Routes = append(result.Routes, contRoute)

	return types.PrintResult(result, conf.CNIVersion)
}

func checkIface(l *log.Entry, netns string, iface *IFaceConf) error {
	name := iface.Interface.Name
	// obtaining interface corresponding link
	link, err := netlink.LinkByName(name)
	if err != nil {
		if _, notFound := err.(netlink.LinkNotFoundError); notFound {
			l.WithField("detail", err).Fatal("iface doesn't exist")
			return fmt.Errorf("%q iface doesn't exist into %q netns: %v", name, netns, err)
		}
		l.WithField("detail", err).Fatal("failed iface lookup")
		return fmt.Errorf("failed %q iface lookup into %q netns: %v", name, netns, err)
	}

	// if no ip configuration are expected to be configured on link, simply return
	if iface.IPConf == nil {
		return nil
	}

	// obtaining addresses configured on link
	addrs, err := netlink.AddrList(link, netlink.FAMILY_V4)
	if err != nil {
		l.WithField("detail", err).Fatal("failed iface addresses lookup")
		return fmt.Errorf("failed %q iface addresses lookup into %q netns: %v", name, netns, err)
	}

	// checking if the ip addresses are correctly configured on link
	for _, a := range addrs {
		if a.IPNet.String() == iface.IPConf.Address.String() {
			return nil
		}
	}
	l.WithField("detail", err).Fatal("iface ip misconfiguration")
	return fmt.Errorf("%q iface ip misconfiguration into %q netns: %v", name, netns, err)
}

// getIfaceConfs scans the prevResult.Interfaces in order to find the expected container and host interface created
// during the ADD operation. If the two interfaces are found, they are returned in association with their IPConf
func getIfaceConfs(contIfName, hostIfName, netns string, prevResult *current.Result) (*IFaceConf, *IFaceConf, error) {
	var contIfaceConf, hostIfaceConf *IFaceConf
	// scanning all prevResult interfaces
	for i, iface := range prevResult.Interfaces {
		// is this the container interface?
		if iface.Name == contIfName && iface.Sandbox == netns {
			contIfaceConf = &IFaceConf{
				ResultIndex: i,
				Interface:   iface,
			}
		}
		// is this the host interface?
		if iface.Name == hostIfName && iface.Sandbox == "" {
			hostIfaceConf = &IFaceConf{
				ResultIndex: i,
				Interface:   iface,
			}
		}
	}
	// if at least one of the two interfaces is not found return an error
	if contIfaceConf == nil || hostIfaceConf == nil {
		return nil, nil, errors.New("unexpected interfaces: wrong or missing") // TODO
	}
	// scanning prevResult.IPs in order to associate the container interface to its ip configuration
	for _, ipConf := range prevResult.IPs {
		if *ipConf.Interface == hostIfaceConf.ResultIndex {
			hostIfaceConf.IPConf = ipConf
		}
		if *ipConf.Interface == contIfaceConf.ResultIndex {
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

// cmdCheck is called for CHECK requests
func cmdCheck(args *skel.CmdArgs) error {
	// defining the attachment identifier and the base logger
	att := utils.Truncate(fmt.Sprintf("%s_%s", args.IfName, args.ContainerID[0:10]), 15)
	l := log.WithField("id", fmt.Sprintf("CHK_%s", att))

	// parsing configuration
	conf, err := loadNetConf(args.StdinData)
	if err != nil {
		l.WithFields(log.Fields{
			"subject": "netconf",
			"detail":  err,
		}).Fatal("parsing failed")
		return err
	}

	// checking the presence of prevResult (its presence is made mandatory by the CNI specification
	// in order to check the container networking)
	if conf.PrevResult == nil {
		l.WithField(
			"detail", "prevResult must be specified",
		).Fatal("missing configuration")
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
		l.WithFields(log.Fields{
			"scope":  "ipam",
			"detail": err,
		}).Fatal("CHECK operation failed")
		return fmt.Errorf("CHECK operation failed on ipam plugin: %v", err)
	}
	l.Info("ip checked")

	// getting netns handle
	netns, err := ns.GetNS(args.Netns)
	if err != nil {
		l.WithFields(log.Fields{
			"netns":  args.Netns,
			"detail": err,
		}).Fatal("failed to retrieve netns")
		return fmt.Errorf("failed to open netns %q: %v", args.Netns, err)
	}
	defer netns.Close()

	// extracting the container interface and the host interface with their own ip configurations
	contIfaceConf, hostIfaceConf, err := getIfaceConfs(args.IfName, att, args.Netns, prevResult)

	if err != nil {
		l.WithFields(log.Fields{
			"prevResult": fmt.Sprintf("%+v", *prevResult),
			"detail":     err,
		}).Fatal("unexpected prevResult")
		return fmt.Errorf("unexpected prevResult: %v", err)
	}

	// checking args.Netns netns interface and routes
	nlog := l.WithField("netns", args.Netns)
	if err := netns.Do(func(_ ns.NetNS) error {
		ilog := nlog.WithField("iface", args.IfName)
		if err := checkIface(ilog, args.Netns, contIfaceConf); err != nil {
			return err
		}
		ilog.Info("netns iface checked")

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
	ilog := nlog.WithField("iface", hostIfaceConf.Interface.Name)
	if err := checkIface(ilog, "root", hostIfaceConf); err != nil {
		return err
	}
	ilog.Info("netns iface checked")
	nlog.Info("netns checked")

	// checking lbrp
	lbName := "lbrp_" + att
	lbFPeer := att                                             // lbrp frontend port peer
	lbBPeer := utils.CreatePeer(conf.BridgeName, "to_"+lbName) // lbrp backend port peer
	llog := l.WithField("lbrp", lbName)                        // load balancer logger
	if err := checkLbrp(
		lbName,
		lbFPeer,
		lbBPeer,
	); err != nil {
		llog.WithField("detail", err).Fatal("failed lbrp checking")
		return fmt.Errorf("failed %q lbrp checking: %v", lbName, err)
	}
	llog.Info("lbrp checked")

	// checking bridge port
	brName := conf.BridgeName
	brPortName := "to_" + lbName
	brPeer := utils.CreatePeer(lbName, "to_bridge")
	brlog := l.WithFields(log.Fields{
		"bridge": brName,
		"port":   brPortName,
	})
	port, resp, err := simplebridgeAPI.ReadSimplebridgePortsByID(context.TODO(), brName, brPortName)
	if err != nil {
		brlog.WithField("detail", fmt.Sprintf(
			"failed to retrieve %q bridge - error: %s, response: %+v", brName, err, resp,
		)).Fatal("failed bridge port checking")
		return fmt.Errorf("failed to retrieve %q bridge %q port - error: %s, response: %+v", brName, brPortName, err, resp)
	}
	if port.Peer != brPeer {
		brlog.WithField("detail", fmt.Sprintf(
			"wrong %q bridge %q port peer - required: %q, found: %q", brName, brPortName, brPeer, port.Peer,
		)).Fatal("failed bridge port checking")
		return fmt.Errorf("wrong %q bridge %q port peer - required: %q, found: %q", brName, brPortName, brPeer, port.Peer)
	}
	if port.Status != "UP" {
		brlog.WithField("detail", fmt.Sprintf(
			"wrong %q bridge %q port status - required: UP, found: DOWN", brName, brPortName,
		)).Fatal("failed bridge port checking")
		return fmt.Errorf("wrong %q bridge %q port status - required: UP, found: DOWN", brName, brPortName)
	}
	brlog.Info("bridge port checked")

	return nil
}

// cmdDel is called for DELETE requests
func cmdDel(args *skel.CmdArgs) error {
	// defining the attachment identifier and the base logger
	att := utils.Truncate(fmt.Sprintf("%s_%s", args.IfName, args.ContainerID[0:10]), 15)
	l := log.WithField("id", fmt.Sprintf("DEL_%s", att))

	// parsing configuration
	conf, err := loadNetConf(args.StdinData)
	if err != nil {
		l.WithFields(log.Fields{
			"subject": "netconf",
			"detail":  err,
		}).Fatal("parsing failed")
		return err
	}

	// actually, this implementation of the DELETE operation doesn't need to access the
	// information about prevResult, so it will not be checked

	// releasing IP address
	if err := ipam.ExecDel(conf.IPAM.Type, args.StdinData); err != nil {
		l.WithFields(log.Fields{
			"scope":  "ipam",
			"detail": err,
		}).Fatal("DEL operation failed")
		return fmt.Errorf("DEL operation failed on ipam plugin: %v", err)
	}
	l.Info("ip released")

	// deleting netns iface and related stuff (routes, arpentry, etc...)
	if args.Netns != "" {
		nlog := l.WithFields(log.Fields{
			"netns": args.Netns,
			"iface": args.IfName,
		})
		// There is a netns so try to clean up. Delete can be called multiple times
		// so don't return an error if the device is already removed.
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

	// deleting load balancer
	lbName := "lbrp_" + att
	llog := l.WithField("lbrp", lbName)
	if resp, err := lbrpAPI.DeleteLbrpByID(context.TODO(), lbName); err != nil && resp.StatusCode != 409 {
		llog.WithField("detail", err).Fatal("failed to delete lbrp")
		return fmt.Errorf("failed to delete lbrp %q - error: %s, response: %+v", lbName, err, resp)
	}
	llog.Info("lbrp deleted")

	// deleting bridge port
	brName := conf.BridgeName
	brPortName := "to_" + lbName
	brlog := l.WithFields(log.Fields{
		"bridge": brName,
		"port":   brPortName,
	})
	if resp, err := simplebridgeAPI.DeleteSimplebridgePortsByID(context.TODO(), brName, brPortName); err != nil && resp.StatusCode != 409 {
		brlog.WithField("detail", err).Fatal("failed to delete bridge port")
		return fmt.Errorf(
			"failed to delete port %q on bridge %q: - error: %s, response: %+v", brPortName, brName, err, resp,
		)
	}
	brlog.Info("bridge port deleted")

	return nil
}

func main() {
	skel.PluginMain(cmdAdd, cmdCheck, cmdDel, version.All, "polykube-cni-plugin")
}
