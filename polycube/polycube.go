package polycube

import (
	"context"
	"errors"
	"fmt"
	"github.com/ekoops/polykube-operator/node"
	k8sdispatcher "github.com/ekoops/polykube-operator/polycube/clients/k8sdispatcher"
	k8slbrp "github.com/ekoops/polykube-operator/polycube/clients/k8slbrp"
	"github.com/ekoops/polykube-operator/polycube/clients/router"
	"github.com/ekoops/polykube-operator/utils"
	"math"
	"net"
	"net/http"
	ctrl "sigs.k8s.io/controller-runtime"
	"time"
)

const (
	basePath = "http://127.0.0.1:9000/polycube/v1"
)

var (
	// logger used throughout the package
	log              = ctrl.Log.WithName("polycube-pkg")
	k8sLbrpAPI       *k8slbrp.K8sLbrpApiService
	routerAPI        *router.RouterApiService
	k8sDispatcherAPI *k8sdispatcher.K8sdispatcherApiService
	conf             *configuration
)

func init() {
	// init k8slbrp API
	cfgK8sLbrp := k8slbrp.Configuration{BasePath: basePath}
	srK8sLbrp := k8slbrp.NewAPIClient(&cfgK8sLbrp)
	k8sLbrpAPI = srK8sLbrp.K8sLbrpApi

	// init router API
	cfgRouter := router.Configuration{BasePath: basePath}
	srRouter := router.NewAPIClient(&cfgRouter)
	routerAPI = srRouter.RouterApi

	// init k8sdispatcher API
	cfgK8sdispatcher := k8sdispatcher.Configuration{BasePath: basePath}
	srK8sdispatcher := k8sdispatcher.NewAPIClient(&cfgK8sdispatcher)
	k8sDispatcherAPI = srK8sdispatcher.K8sdispatcherApi
}

// createIntK8sLbrp creates the polycube internal k8s lbrp cube.
func createIntK8sLbrp() error {
	iklName := conf.intK8sLbrpName

	// defining internal k8s lbrp port that will be connected to the router
	iklToRPort := k8slbrp.Ports{
		Name:  conf.iklToRPortName,
		Type_: "backend",
	}
	iklPorts := []k8slbrp.Ports{iklToRPort}
	ikl := k8slbrp.K8sLbrp{
		Name:      iklName,
		Loglevel:  conf.cubesLogLevel,
		Ports:     iklPorts,
		PortMode_: "MULTI",
	}

	// creating internal k8s lbrp
	if resp, err := k8sLbrpAPI.CreateK8sLbrpByID(context.TODO(), iklName, ikl); err != nil {
		log.Error(
			err, "failed to create internal k8s lbrp",
			"k8slbrp", fmt.Sprintf("%+v", ikl),
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to create internal k8s lbrp")
	}
	log.V(1).Info("created internal k8s lbrp", "name", iklName)
	return nil
}

// createRouter creates a polycube router cube.
func createRouter() error {
	rName := conf.rName

	// defining the router port that will be connected to the internal k8s lbrp
	podsGwIPNetStr := node.Conf.PodGwInfo.IPNet.String()
	podsGwMACStr := node.Conf.PodGwInfo.MAC.String()
	rToIklPort := router.Ports{
		Name: conf.rToIklPortName,
		Ip:   podsGwIPNetStr,
		Mac:  podsGwMACStr,
	}
	// defining the router port that will be connected to the vxlan interface
	vxlanIface := node.Conf.VxlanIface
	rToVxlanPort := router.Ports{
		Name: conf.rToVxlanPortName,
		Ip:   vxlanIface.IPNet.String(),
		Mac:  vxlanIface.Link.Attrs().HardwareAddr.String(),
		Peer: vxlanIface.Link.Attrs().Name,
	}

	// defining the router port that will be connected to the polykube veth pair net end interface
	vethNetEndIface := node.Conf.PolykubeVeth.Net
	vethNetEndIfaceIPNet := &net.IPNet{
		IP:   vethNetEndIface.IPNet.IP,
		Mask: node.Env.PolykubeVethPairCIDR.Mask, // this CIDR is used in order to allow to reach the host veth pair end
	}
	rToHostPort := router.Ports{
		Name: conf.rToHostPortName,
		Ip:   vethNetEndIfaceIPNet.String(),
		Mac:  vethNetEndIface.Link.Attrs().HardwareAddr.String(),
		// TODO: it is not possibile to set the peer here since in this way polycube will set the IP on the net end
		//Peer: vethNetEndIface.Link.Attrs().Name,
	}

	// defining the router port that will be connected to the external k8s lbrp
	extIface := node.Conf.ExtIface
	rToEklPort := router.Ports{
		Name: conf.rToEklPortName,
		Ip:   extIface.IPNet.String(),
		Mac:  extIface.Link.Attrs().HardwareAddr.String(),
	}
	rPorts := []router.Ports{rToIklPort, rToVxlanPort, rToHostPort, rToEklPort}

	// defining router default route and setting static arp table entry for the default gateway
	nodeGwIPStr := node.Conf.NodeGwInfo.IPNet.IP.String()
	nodeGwMACStr := node.Conf.NodeGwInfo.MAC.String()

	nodeIPNet := &net.IPNet{
		IP:   node.Conf.ExtIface.IPNet.IP,
		Mask: net.IPMask{255, 255, 255, 255}, //  /32 for allowing communication only with processes listening on the node physical interface
	}

	routes := []router.Route{
		{ // default route through node gateway
			Network:    "0.0.0.0/0",
			Nexthop:    nodeGwIPStr,
			Interface_: conf.rToEklPortName,
		},
		{ // traffic for vpod address must go towards the k8sdispatcher
			Network:    node.Conf.VPodIPNet.String(),
			Nexthop:    nodeGwIPStr,
			Interface_: conf.rToEklPortName,
		},
		{ // traffic for node physical interface address must go towards the veth host end
			Network:    nodeIPNet.String(),
			Nexthop:    node.Conf.PolykubeVeth.Host.IPNet.IP.String(),
			Interface_: conf.rToHostPortName,
		},
	}
	arptable := []router.ArpTable{
		{
			Address:    nodeGwIPStr,
			Mac:        nodeGwMACStr,
			Interface_: conf.rToEklPortName,
		},
	}
	r := router.Router{
		Name:     rName,
		Ports:    rPorts,
		Loglevel: conf.cubesLogLevel,
		Route:    routes,
		ArpTable: arptable,
	}

	// creating router
	if resp, err := routerAPI.CreateRouterByID(context.TODO(), rName, r); err != nil {
		log.Error(err, "failed to create router",
			"router", fmt.Sprintf("%+v", r),
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to create router")
	}
	log.V(1).Info("created router", "name", rName)
	return nil
}

// createExtK8sLbrp creates the polycube external k8s lbrp cube.
func createExtK8sLbrp() error {
	eklName := conf.extK8sLbrpName

	// defining the external k8s lbrp port that will be connected to the router interface
	eklToRPort := k8slbrp.Ports{
		Name:  conf.eklToRPortName,
		Type_: "backend",
	}
	// defining the external k8s lbrp port that will be connected to the k8sdispatcher interface
	eklToKPort := k8slbrp.Ports{
		Name:  conf.eklToKPortName,
		Type_: "frontend",
	}

	eklPorts := []k8slbrp.Ports{eklToRPort, eklToKPort}
	ekl := k8slbrp.K8sLbrp{
		Name:      eklName,
		Loglevel:  conf.cubesLogLevel,
		Ports:     eklPorts,
		PortMode_: "SINGLE",
	}

	// creating external k8s lbrp
	if resp, err := k8sLbrpAPI.CreateK8sLbrpByID(context.TODO(), eklName, ekl); err != nil {
		log.Error(
			err, "failed to create external k8s lbrp",
			"k8slbrp", fmt.Sprintf("%+v", ekl),
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to create external k8s lbrp")
	}
	log.V(1).Info("created external k8s lbrp", "name", eklName)
	return nil
}

// createK8sDispatcher creates a polycube k8s dispatcher cube for managing incoming connection.
func createK8sDispatcher() error {
	kdName := conf.k8sDispName

	// defining the k8sdispatcher port that will be connected to the external k8s lbrp interface
	kToEklPort := k8sdispatcher.Ports{
		Name:  conf.kToEklPortName,
		Type_: "BACKEND",
	}
	// defining the k8sdispatcher port that will be connected to the node external interface
	extIface := node.Conf.ExtIface
	kToIntPort := k8sdispatcher.Ports{
		Name:  conf.kToIntPortName,
		Type_: "FRONTEND",
		Ip_:   extIface.IPNet.IP.String(),
		Peer:  extIface.Link.Attrs().Name,
	}
	kPorts := []k8sdispatcher.Ports{kToEklPort, kToIntPort}

	kd := k8sdispatcher.K8sdispatcher{
		Name:          kdName,
		Loglevel:      conf.cubesLogLevel,
		Ports:         kPorts,
		InternalSrcIp: node.Conf.VPodIPNet.IP.String(),
		NodeportRange: node.Env.NodePortRange,
	}

	// creating k8s dispatcher
	if resp, err := k8sDispatcherAPI.CreateK8sdispatcherByID(context.TODO(), kdName, kd); err != nil {
		log.Error(
			err, "failed to create k8s dispatcher",
			"k8sdisp", fmt.Sprintf("%+v", kd),
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to create k8s dispatcher")
	}
	log.V(1).Info("created k8s dispatcher", "name", kdName)
	return nil
}

// EnsureCubesConnections ensures that each port of the already deployed polycube infrastructure is connected with
// the right peer.
func EnsureCubesConnections() error {
	iklName := conf.intK8sLbrpName
	rName := conf.rName
	eklName := conf.extK8sLbrpName
	kName := conf.k8sDispName

	// updating internal k8s lbrp ports
	// updating internal k8s lbrp "to_r0" port in order to set peer=r0:to_ikl0
	iklToRPortName := conf.iklToRPortName
	iklToRPortPeer := utils.CreatePeer(rName, conf.rToIklPortName)
	l := log.WithValues("name", iklName, "port", iklToRPortName, "peer", iklToRPortPeer)
	iklToRPort := k8slbrp.Ports{
		Name: iklToRPortName,
		Peer: iklToRPortPeer,
	}
	if resp, err := k8sLbrpAPI.UpdateK8sLbrpPortsByID(context.TODO(), iklName, iklToRPortName, iklToRPort); err != nil {
		l.Error(
			err, "failed to set internal k8s lbrp port peer",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to set internal k8s lbrp port peer")
	}
	l.V(1).Info("set internal k8s lbrp port peer")

	// updating router ports
	// updating router "to_ikl0" port in order to set peer=ikl0:to_r0
	rToIklPortName := conf.rToIklPortName
	rToIklPortPeer := utils.CreatePeer(iklName, conf.iklToRPortName)
	podsGwInfo := node.Conf.PodGwInfo
	l = log.WithValues("name", rName, "port", rToIklPortName, "peer", rToIklPortPeer)
	rToIklPort := router.Ports{
		Name: rToIklPortName,
		Ip:   podsGwInfo.IPNet.String(),
		Mac:  podsGwInfo.MAC.String(),
		Peer: rToIklPortPeer,
	}
	if resp, err := routerAPI.UpdateRouterPortsByID(context.TODO(), rName, rToIklPortName, rToIklPort); err != nil {
		l.Error(err, "failed to set router port peer", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to set router port peer")
	}
	l.V(1).Info("set router port peer")

	// updating router "to_host" port in order to set peer=polykube_net
	// Notice that the peer cannot be assigned in port creation since polycube try to enforce the IP address on the
	// veth end: this is an unwanted behavior
	// TODO: assigning the peer in port creation
	rToHostPortName := conf.rToHostPortName
	vethNetEndIface := node.Conf.PolykubeVeth.Net
	rToHostPortPeer := vethNetEndIface.Link.Attrs().Name
	vethNetEndIfaceIPNet := &net.IPNet{
		IP:   vethNetEndIface.IPNet.IP,
		Mask: node.Env.PolykubeVethPairCIDR.Mask, // this CIDR is used in order to allow to reach the host veth pair end
	}
	l = log.WithValues("name", rName, "port", rToHostPortName, "peer", rToHostPortPeer)
	rToHostPort := router.Ports{
		Name: rToHostPortName,
		Ip:   vethNetEndIfaceIPNet.String(),
		Mac:  vethNetEndIface.Link.Attrs().HardwareAddr.String(),
		Peer: rToHostPortPeer,
	}
	if resp, err := routerAPI.UpdateRouterPortsByID(context.TODO(), rName, rToHostPortName, rToHostPort); err != nil {
		l.Error(err, "failed to set router port peer", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to set router port peer")
	}
	l.V(1).Info("set router port peer")

	// updating router "to_ekl0" port in order to set peer=ekl0:to_r0
	rToEklPortName := conf.rToEklPortName
	rToEklPortPeer := utils.CreatePeer(eklName, conf.eklToRPortName)
	extIface := node.Conf.ExtIface
	l = log.WithValues("name", rName, "port", rToEklPortName, "peer", rToEklPortPeer)
	rToEklPort := router.Ports{
		Name: conf.rToEklPortName,
		Ip:   extIface.IPNet.String(),
		Mac:  extIface.Link.Attrs().HardwareAddr.String(),
		Peer: rToEklPortPeer,
	}
	if resp, err := routerAPI.UpdateRouterPortsByID(context.TODO(), rName, rToEklPortName, rToEklPort); err != nil {
		l.Error(err, "failed to set router port peer", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to set router port peer")
	}
	l.V(1).Info("set router port peer")

	// updating external k8s lbrp ports
	// updating external k8s lbrp "to_r0" port in order to set peer=r0:to_ekl0
	eklToRPortName := conf.eklToRPortName
	eklToRPortPeer := utils.CreatePeer(rName, conf.rToEklPortName)
	l = log.WithValues("name", eklName, "port", eklToRPortName, "peer", eklToRPortPeer)
	eklToRPort := k8slbrp.Ports{
		Name: eklToRPortName,
		Peer: eklToRPortPeer,
	}
	if resp, err := k8sLbrpAPI.UpdateK8sLbrpPortsByID(context.TODO(), eklName, eklToRPortName, eklToRPort); err != nil {
		l.Error(
			err, "failed to set external k8s lbrp port peer",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to set external k8s lbrp port peer")
	}
	l.V(1).Info("set external k8s lbrp port peer")

	// updating external k8s lbrp "to_k0" port in order to set peer=k0:to_ekl0
	eklToKPortName := conf.eklToKPortName
	eklToKPortPeer := utils.CreatePeer(kName, conf.kToEklPortName)
	l = log.WithValues("name", eklName, "port", eklToKPortName, "peer", eklToKPortPeer)
	eklToKPort := k8slbrp.Ports{
		Name: eklToKPortName,
		Peer: eklToKPortPeer,
	}
	if resp, err := k8sLbrpAPI.UpdateK8sLbrpPortsByID(context.TODO(), eklName, eklToKPortName, eklToKPort); err != nil {
		l.Error(
			err, "failed to set external k8s lbrp port peer",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to set external k8s lbrp port peer")
	}
	l.V(1).Info("set external k8s lbrp port peer")

	// updating k8sdispatcher ports
	// updating k8sdispatcher "to_ekl0" port in order to set peer=ekl0:to_k0
	kToEklPortName := conf.kToEklPortName
	kToEklPortPeer := utils.CreatePeer(eklName, conf.eklToKPortName)
	l = log.WithValues("name", kName, "port", kToEklPortName, "peer", kToEklPortPeer)
	kToEklPort := k8sdispatcher.Ports{
		Name: kToEklPortName,
		Peer: kToEklPortPeer,
	}
	if resp, err := k8sDispatcherAPI.UpdateK8sdispatcherPortsByID(
		context.TODO(), kName, kToEklPortName, kToEklPort,
	); err != nil {
		l.Error(
			err, "failed to set k8s dispatcher port peer",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to set k8s dispatcher port peer")
	}
	l.V(1).Info("set k8s dispatcher port peer")
	return nil
}

// EnsureCubes ensures that the polycube infrastructure is correctly deployed: in order to do that, it creates
// the cubes that don't exist yet and cleans up the configuration of the already existing ones.
func EnsureCubes() error {
	// retrieving the k8s lbrp list
	kls, resp, err := k8sLbrpAPI.ReadK8sLbrpListByID(context.TODO())
	if err != nil {
		log.Error(err, "failed to retrieve the k8s lbrp list", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to retrieve the k8s lbrp list")
	}

	// checking k8s lbrps existence
	iklName := conf.intK8sLbrpName
	eklName := conf.extK8sLbrpName
	iklExist := false
	eklExist := false
	for _, kl := range kls {
		if kl.Name == iklName {
			iklExist = true
		} else if kl.Name == eklName {
			eklExist = true
		}
	}
	// creating internal k8s lbrp if it doesn't exist
	l := log.WithValues("name", iklName)
	if !iklExist {
		l.V(1).Info("failed to find internal k8s lbrp: creating it...")
		if err := createIntK8sLbrp(); err != nil {
			return err
		}
	} else {
		l.V(1).Info("cleaning up services of the already existing internal k8s lbrp...")
		if err := CleanupK8sLbrpServices(iklName); err != nil {
			return err
		}
	}
	// creating external k8s lbrp if it doesn't exist
	l = log.WithValues("name", eklName)
	if !eklExist {
		l.V(1).Info("failed to find external k8s lbrp: creating it...")
		if err := createExtK8sLbrp(); err != nil {
			return err
		}
	} else {
		l.V(1).Info("cleaning up services of the already existing external k8s lbrp...")
		if err := CleanupK8sLbrpServices(eklName); err != nil {
			return err
		}
	}

	// retrieving the router list
	rName := conf.rName
	l = log.WithValues("name", rName)
	rs, resp, err := routerAPI.ReadRouterListByID(context.TODO())
	if err != nil {
		log.Error(err, "failed to retrieve the router list", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to retrieve the router list")
	}

	// checking router existence
	rExist := false
	for _, r := range rs {
		if r.Name == rName {
			rExist = true
			break
		}
	}
	if !rExist {
		l.V(1).Info("failed to find router: creating it...")
		if err := createRouter(); err != nil {
			return err
		}
	} else {
		l.V(1).Info("cleaning up routes of the already existing router...")
		if err := CleanupRouterRoutes(); err != nil {
			return err
		}
	}

	// retrieving the k8s dispatcher list
	kdName := conf.k8sDispName
	l = log.WithValues("name", kdName)
	kds, resp, err := k8sDispatcherAPI.ReadK8sdispatcherListByID(context.TODO())
	if err != nil {
		log.Error(err, "failed to retrieve the k8s dispatcher list", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to retrieve the k8s dispatcher list")
	}

	// checking k8s dispatcher existence
	kdExist := false
	for _, kd := range kds {
		if kd.Name == kdName {
			kdExist = true
			break
		}
	}
	if !kdExist {
		l.V(1).Info("failed to find k8s dispatcher: creating it...")
		if err := createK8sDispatcher(); err != nil {
			return err
		}
	} else {
		l.V(1).Info("cleaning up NodePort rules of the already existing k8s dispatcher...")
		if err := CleanupK8sDispatcherNodePortRules(); err != nil {
			return err
		}
	}

	log.Info("ensured polycube cubes")
	return nil
}

// EnsureConnection probes polycubed in order to be sure that it is up and running. If it is not immediately reachable,
// up to 5 further attempts with an exponential backoff delay are made. If all the attempts fail, an error is returned.
func EnsureConnection() error {
	for i := 0; i < 6; i++ {
		log.Info("waiting for polycubed...")
		if i != 0 {
			backoff := time.Duration(math.Pow(2, float64(i-1))) * time.Second
			time.Sleep(backoff)
		}
		if _, err := http.Get(basePath); err == nil {
			log.Info("polycubed is ready")
			return nil
		}
	}
	log.Error(errors.New("polycubed is unreachable"), "failed to connect to polycubed")
	return errors.New("failed to connect to polycubed")
}

// InitConf initializes the internal polycube configuration in order to allow easier polycube components name resolution.
func InitConf() {
	ec := node.Env
	conf = &configuration{
		cubesLogLevel:    ec.CubesLogLevel,
		intK8sLbrpName:   ec.IntK8sLbrpName,
		rName:            ec.RouterName,
		extK8sLbrpName:   ec.ExtK8sLbrpName,
		k8sDispName:      ec.K8sDispName,
		iklToRPortName:   "to_" + ec.RouterName,
		rToIklPortName:   "to_" + ec.IntK8sLbrpName,
		rToVxlanPortName: "to_" + ec.VxlanIfaceName,
		rToHostPortName:  "to_host",
		rToEklPortName:   "to_" + ec.ExtK8sLbrpName,
		eklToRPortName:   "to_" + ec.RouterName,
		eklToKPortName:   "to_" + ec.K8sDispName,
		kToEklPortName:   "to_" + ec.ExtK8sLbrpName,
		kToIntPortName:   "to_int",
	}
}
