package polycube

import (
	"context"
	"errors"
	"fmt"
	"github.com/ekoops/polykube-operator/node"
	k8sdispatcher "github.com/ekoops/polykube-operator/polycube/clients/k8sdispatcher"
	lbrp "github.com/ekoops/polykube-operator/polycube/clients/lbrp"
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
	lbrpAPI          *lbrp.LbrpApiService
	routerAPI        *router.RouterApiService
	k8sDispatcherAPI *k8sdispatcher.K8sdispatcherApiService
	conf             *configuration
)

func init() {
	// init lbrp API
	cfgLbrp := lbrp.Configuration{BasePath: basePath}
	srLbrp := lbrp.NewAPIClient(&cfgLbrp)
	lbrpAPI = srLbrp.LbrpApi

	// init router API
	cfgRouter := router.Configuration{BasePath: basePath}
	srRouter := router.NewAPIClient(&cfgRouter)
	routerAPI = srRouter.RouterApi

	// init k8sdispatcher API
	cfgK8sdispatcher := k8sdispatcher.Configuration{BasePath: basePath}
	srK8sdispatcher := k8sdispatcher.NewAPIClient(&cfgK8sdispatcher)
	k8sDispatcherAPI = srK8sdispatcher.K8sdispatcherApi
}

// createIntLbrp creates the polycube internal lbrp cube.
func createIntLbrp() error {
	ilbName := conf.intLbrpName

	// defining internal lbrp port that will be connected to the router
	ilbToRPort := lbrp.Ports{
		Name:  conf.ilbToRPortName,
		Type_: "backend",
	}
	ilbPorts := []lbrp.Ports{ilbToRPort}
	ilb := lbrp.Lbrp{
		Name:      ilbName,
		Loglevel:  conf.cubesLogLevel,
		Ports:     ilbPorts,
		PortMode_: "MULTI",
	}

	// creating internal lbrp
	if resp, err := lbrpAPI.CreateLbrpByID(context.TODO(), ilbName, ilb); err != nil {
		log.Error(
			err, "failed to create internal lbrp",
			"lbrp", fmt.Sprintf("%+v", ilb),
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to create internal lbrp")
	}
	log.V(1).Info("created internal lbrp", "name", ilbName)
	return nil
}

// createRouter creates a polycube router cube.
func createRouter() error {
	rName := conf.rName

	// defining the router port that will be connected to the internal lbrp
	podsGwIPNetStr := node.Conf.PodGwInfo.IPNet.String()
	podsGwMACStr := node.Conf.PodGwInfo.MAC.String()
	rToIlbPort := router.Ports{
		Name: conf.rToIlbPortName,
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
		Mask: net.IPMask{255, 255, 255, 252}, // /30 for allowing host veth pair end reachability
	}
	rToHostPort := router.Ports{
		Name: conf.rToHostPortName,
		Ip:   vethNetEndIfaceIPNet.String(),
		Mac:  vethNetEndIface.Link.Attrs().HardwareAddr.String(),
		// TODO: assigning the peer in port creation
		// Notice that it is not possible to set the peer here, since in this way polycube will set the IP on the net end
		//Peer: vethNetEndIface.Link.Attrs().Name,
	}

	// defining the router port that will be connected to the external lbrp
	extIface := node.Conf.ExtIface
	rToElbPort := router.Ports{
		Name: conf.rToElbPortName,
		Ip:   extIface.IPNet.String(),
		Mac:  extIface.Link.Attrs().HardwareAddr.String(),
	}
	rPorts := []router.Ports{rToIlbPort, rToVxlanPort, rToHostPort, rToElbPort}

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
			Interface_: conf.rToElbPortName,
		},
		{ // traffic for vpod address must go towards the k8sdispatcher
			Network:    node.Conf.VPodIPNet.String(),
			Nexthop:    nodeGwIPStr,
			Interface_: conf.rToElbPortName,
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
			Interface_: conf.rToElbPortName,
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

// createExtLbrp creates the polycube external lbrp cube.
func createExtLbrp() error {
	elbName := conf.extLbrpName

	// defining the external lbrp port that will be connected to the router interface
	elbToRPort := lbrp.Ports{
		Name:  conf.elbToRPortName,
		Type_: "backend",
	}
	// defining the external lbrp port that will be connected to the k8sdispatcher interface
	elbToKPort := lbrp.Ports{
		Name:  conf.elbToKPortName,
		Type_: "frontend",
	}

	elbPorts := []lbrp.Ports{elbToRPort, elbToKPort}
	elb := lbrp.Lbrp{
		Name:      elbName,
		Loglevel:  conf.cubesLogLevel,
		Ports:     elbPorts,
		PortMode_: "SINGLE",
	}

	// creating external lbrp
	if resp, err := lbrpAPI.CreateLbrpByID(context.TODO(), elbName, elb); err != nil {
		log.Error(
			err, "failed to create external lbrp",
			"lbrp", fmt.Sprintf("%+v", elb),
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to create external lbrp")
	}
	log.V(1).Info("created external lbrp", "name", elbName)
	return nil
}

// createK8sDispatcher creates a polycube k8s dispatcher cube for managing incoming connection.
func createK8sDispatcher() error {
	kdName := conf.k8sDispName

	// defining the k8sdispatcher port that will be connected to the external lbrp interface
	kToElbPort := k8sdispatcher.Ports{
		Name:  conf.kToElbPortName,
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
	kPorts := []k8sdispatcher.Ports{kToElbPort, kToIntPort}

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
	ilbName := conf.intLbrpName
	rName := conf.rName
	elbName := conf.extLbrpName
	kName := conf.k8sDispName

	// updating internal lbrp ports
	// updating internal lbrp "to_r0" port in order to set peer=r0:to_ilb0
	ilbToRPortName := conf.ilbToRPortName
	ilbToRPortPeer := utils.CreatePeer(rName, conf.rToIlbPortName)
	l := log.WithValues("name", ilbName, "port", ilbToRPortName, "peer", ilbToRPortPeer)
	ilbToRPort := lbrp.Ports{
		Name: ilbToRPortName,
		Peer: ilbToRPortPeer,
	}
	if resp, err := lbrpAPI.UpdateLbrpPortsByID(context.TODO(), ilbName, ilbToRPortName, ilbToRPort); err != nil {
		l.Error(
			err, "failed to set internal lbrp port peer",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to set internal lbrp port peer")
	}
	l.V(1).Info("set internal lbrp port peer")

	// updating router ports
	// updating router "to_ilb0" port in order to set peer=ilb0:to_r0
	rToIlbPortName := conf.rToIlbPortName
	rToIlbPortPeer := utils.CreatePeer(ilbName, conf.ilbToRPortName)
	podsGwInfo := node.Conf.PodGwInfo
	l = log.WithValues("name", rName, "port", rToIlbPortName, "peer", rToIlbPortPeer)
	rToIlbPort := router.Ports{
		Name: rToIlbPortName,
		Ip:   podsGwInfo.IPNet.String(),
		Mac:  podsGwInfo.MAC.String(),
		Peer: rToIlbPortPeer,
	}
	if resp, err := routerAPI.UpdateRouterPortsByID(context.TODO(), rName, rToIlbPortName, rToIlbPort); err != nil {
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
		Mask: net.IPMask{255, 255, 255, 252}, // /30 for allowing host veth pair end reachability
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

	// updating router "to_elb0" port in order to set peer=elb0:to_r0
	rToElbPortName := conf.rToElbPortName
	rToElbPortPeer := utils.CreatePeer(elbName, conf.elbToRPortName)
	extIface := node.Conf.ExtIface
	l = log.WithValues("name", rName, "port", rToElbPortName, "peer", rToElbPortPeer)
	rToElbPort := router.Ports{
		Name: conf.rToElbPortName,
		Ip:   extIface.IPNet.String(),
		Mac:  extIface.Link.Attrs().HardwareAddr.String(),
		Peer: rToElbPortPeer,
	}
	if resp, err := routerAPI.UpdateRouterPortsByID(context.TODO(), rName, rToElbPortName, rToElbPort); err != nil {
		l.Error(err, "failed to set router port peer", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to set router port peer")
	}
	l.V(1).Info("set router port peer")

	// updating external lbrp ports
	// updating external lbrp "to_r0" port in order to set peer=r0:to_elb0
	elbToRPortName := conf.elbToRPortName
	elbToRPortPeer := utils.CreatePeer(rName, conf.rToElbPortName)
	l = log.WithValues("name", elbName, "port", elbToRPortName, "peer", elbToRPortPeer)
	elbToRPort := lbrp.Ports{
		Name: elbToRPortName,
		Peer: elbToRPortPeer,
	}
	if resp, err := lbrpAPI.UpdateLbrpPortsByID(context.TODO(), elbName, elbToRPortName, elbToRPort); err != nil {
		l.Error(
			err, "failed to set external lbrp port peer",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to set external lbrp port peer")
	}
	l.V(1).Info("set external lbrp port peer")

	// updating external lbrp "to_k0" port in order to set peer=k0:to_elb0
	elbToKPortName := conf.elbToKPortName
	elbToKPortPeer := utils.CreatePeer(kName, conf.kToElbPortName)
	l = log.WithValues("name", elbName, "port", elbToKPortName, "peer", elbToKPortPeer)
	elbToKPort := lbrp.Ports{
		Name: elbToKPortName,
		Peer: elbToKPortPeer,
	}
	if resp, err := lbrpAPI.UpdateLbrpPortsByID(context.TODO(), elbName, elbToKPortName, elbToKPort); err != nil {
		l.Error(
			err, "failed to set external lbrp port peer",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to set external lbrp port peer")
	}
	l.V(1).Info("set external lbrp port peer")

	// updating k8sdispatcher ports
	// updating k8sdispatcher "to_elb0" port in order to set peer=elb0:to_k0
	kToElbPortName := conf.kToElbPortName
	kToElbPortPeer := utils.CreatePeer(elbName, conf.elbToKPortName)
	l = log.WithValues("name", kName, "port", kToElbPortName, "peer", kToElbPortPeer)
	kToElbPort := k8sdispatcher.Ports{
		Name: kToElbPortName,
		Peer: kToElbPortPeer,
	}
	if resp, err := k8sDispatcherAPI.UpdateK8sdispatcherPortsByID(
		context.TODO(), kName, kToElbPortName, kToElbPort,
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
	// retrieving the lbrp list
	lbs, resp, err := lbrpAPI.ReadLbrpListByID(context.TODO())
	if err != nil {
		log.Error(err, "failed to retrieve the lbrp list", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to retrieve the lbrp list")
	}

	// checking lbrps existence
	ilbName := conf.intLbrpName
	elbName := conf.extLbrpName
	ilbExist := false
	elbExist := false
	for _, lb := range lbs {
		if lb.Name == ilbName {
			ilbExist = true
		} else if lb.Name == elbName {
			elbExist = true
		}
	}
	// creating internal lbrp if it doesn't exist
	l := log.WithValues("name", ilbName)
	if !ilbExist {
		l.V(1).Info("failed to find internal lbrp: creating it...")
		if err := createIntLbrp(); err != nil {
			return err
		}
	} else {
		l.V(1).Info("cleaning up services of the already existing internal lbrp...")
		if err := CleanupLbrpServices(ilbName); err != nil {
			return err
		}
	}
	// creating external lbrp if it doesn't exist
	l = log.WithValues("name", elbName)
	if !elbExist {
		l.V(1).Info("failed to find external lbrp: creating it...")
		if err := createExtLbrp(); err != nil {
			return err
		}
	} else {
		l.V(1).Info("cleaning up services of the already existing external lbrp...")
		if err := CleanupLbrpServices(elbName); err != nil {
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
		intLbrpName:      ec.IntLbrpName,
		rName:            ec.RouterName,
		extLbrpName:      ec.ExtLbrpName,
		k8sDispName:      ec.K8sDispName,
		ilbToRPortName:   "to_" + ec.RouterName,
		rToIlbPortName:   "to_" + ec.IntLbrpName,
		rToVxlanPortName: "to_" + ec.VxlanIfaceName,
		rToHostPortName:  "to_host",
		rToElbPortName:   "to_" + ec.ExtLbrpName,
		elbToRPortName:   "to_" + ec.RouterName,
		elbToKPortName:   "to_" + ec.K8sDispName,
		kToElbPortName:   "to_" + ec.ExtLbrpName,
		kToIntPortName:   "to_int",
	}
}
