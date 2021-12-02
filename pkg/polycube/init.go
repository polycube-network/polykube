package polycube

import (
	"context"
	"errors"
	"fmt"
	k8sdispatcher "github.com/ekoops/polykube-operator/pkg/clients/k8sdispatcher"
	lbrp "github.com/ekoops/polykube-operator/pkg/clients/lbrp"
	router "github.com/ekoops/polykube-operator/pkg/clients/router"
	simplebridge "github.com/ekoops/polykube-operator/pkg/clients/simplebridge"
	"github.com/ekoops/polykube-operator/pkg/env"
	"github.com/ekoops/polykube-operator/pkg/node"
	"github.com/ekoops/polykube-operator/pkg/utils"
	ctrl "sigs.k8s.io/controller-runtime"
)

const (
	basePath = "http://127.0.0.1:9000/polycube/v1"
)

var (
	// logger used throughout the package
	log              = ctrl.Log.WithName("polycube-pkg")
	simplebridgeAPI  *simplebridge.SimplebridgeApiService
	lbrpAPI          *lbrp.LbrpApiService
	routerAPI        *router.RouterApiService
	k8sdispatcherAPI *k8sdispatcher.K8sdispatcherApiService
	conf             *Configuration
)

func init() {
	// init simplebridge API
	cfgSimplebridge := simplebridge.Configuration{BasePath: basePath}
	srSimplebridge := simplebridge.NewAPIClient(&cfgSimplebridge)
	simplebridgeAPI = srSimplebridge.SimplebridgeApi

	// init router API
	cfgRouter := router.Configuration{BasePath: basePath}
	srRouter := router.NewAPIClient(&cfgRouter)
	routerAPI = srRouter.RouterApi

	// init lbrp API
	cfgLbrp := lbrp.Configuration{BasePath: basePath}
	srLbrp := lbrp.NewAPIClient(&cfgLbrp)
	lbrpAPI = srLbrp.LbrpApi

	// init k8sdispatcher API
	cfgK8sdispatcher := k8sdispatcher.Configuration{BasePath: basePath}
	srK8sdispatcher := k8sdispatcher.NewAPIClient(&cfgK8sdispatcher)
	k8sdispatcherAPI = srK8sdispatcher.K8sdispatcherApi
}

// createBridge creates a polycube simplebridge cube
func createBridge() error {
	brName := conf.brName
	l := log.WithValues("name", brName)

	// defining bridge port that will be connected to the router
	brToRPort := simplebridge.Ports{
		Name: conf.brToRPortName,
	}
	brPorts := []simplebridge.Ports{brToRPort}
	br := simplebridge.Simplebridge{
		Name:     brName,
		Loglevel: "TRACE",
		Ports:    brPorts,
	}

	l = l.WithValues("bridge", fmt.Sprintf("%+v", br))
	// creating bridge
	if resp, err := simplebridgeAPI.CreateSimplebridgeByID(context.TODO(), brName, br); err != nil {
		l.Error(err, "failed to create bridge", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to create bridge")
	}
	l.Info("bridge created")
	return nil
}

// createRouter creates a polycube router cube
func createRouter() error {
	extIface := node.Conf.ExtIface
	podsGwInfo := node.Conf.PodGwInfo
	nodeGwInfo := node.Conf.NodeGwInfo
	rName := conf.rName
	l := log.WithValues("name", rName)

	// defining the router port that will be connected to the bridge
	rToBrPort := router.Ports{
		Name: conf.rToBrPortName,
		Ip:   podsGwInfo.IPNet.String(),
		Mac:  podsGwInfo.MAC.String(),
	}
	// defining the router port that will be connected to the vxlan interface
	rToVxlanPort := router.Ports{
		Name: conf.rToVxlanPortName,
	}
	// defining the router port that will be connected to the lbrp
	rToLbPort := router.Ports{
		Name: conf.rToLbPortName,
		Ip:   extIface.IPNet.String(),
		Mac:  extIface.Link.Attrs().HardwareAddr.String(),
	}
	rPorts := []router.Ports{rToBrPort, rToVxlanPort, rToLbPort}

	// defining router default route and setting static arp table entry for the default gateway
	routes := []router.Route{
		{
			Network:    "0.0.0.0/0",
			Nexthop:    nodeGwInfo.IPNet.IP.String(),
			Interface_: conf.kToLbPortName,
		},
	}
	arptable := []router.ArpTable{
		{
			Address:    nodeGwInfo.IPNet.IP.String(),
			Mac:        nodeGwInfo.MAC.String(),
			Interface_: conf.kToLbPortName,
		},
	}
	r := router.Router{
		Name:     rName,
		Ports:    rPorts,
		Loglevel: "TRACE",
		Route:    routes,
		ArpTable: arptable,
	}

	l = l.WithValues("router", fmt.Sprintf("%+v", r))
	// creating router
	if resp, err := routerAPI.CreateRouterByID(context.TODO(), rName, r); err != nil {
		l.Error(err, "failed to create router", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to create router")
	}
	l.Info("router created")
	return nil
}

// createLbrp creates a polycube lbrp cube for managing incoming connection
func createLbrp() error {
	lbName := conf.lbName
	l := log.WithValues("name", lbName)

	// defining the lbrp port that will be connected to the router interface
	lbToRPort := lbrp.Ports{
		Name:  conf.lbToRPortName,
		Type_: "backend",
	}
	// defining the lbrp port that will be connected to the k8sdispatcher interface
	lbToKPort := lbrp.Ports{
		Name:  conf.lbToKPortName,
		Type_: "frontend",
	}

	lbPorts := []lbrp.Ports{lbToRPort, lbToKPort}
	lb := lbrp.Lbrp{
		Name:     lbName,
		Loglevel: "TRACE",
		Ports:    lbPorts,
	}

	l = l.WithValues("lbrp", fmt.Sprintf("%+v", lb))
	// creating lbrp
	if resp, err := lbrpAPI.CreateLbrpByID(context.TODO(), lbName, lb); err != nil {
		l.Error(err, "failed to create lbrp", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to create lbrp")
	}
	l.Info("lbrp created")
	return nil
}

// createK8sDispatcher creates a polycube k8sdispatcher cube for managing incoming connection
func createK8sDispatcher() error {
	kName := conf.k8sDispName
	l := log.WithValues("name", kName)

	// defining the k8sdispatcher port that will be connected to the lbrp interface
	kToLbPort := k8sdispatcher.Ports{
		Name:  conf.kToLbPortName,
		Type_: "BACKEND",
	}
	// defining the k8sdispatcher port that will be connected to the node external interface
	kToIntPort := k8sdispatcher.Ports{
		Name:  conf.kToIntPortName,
		Type_: "FRONTEND",
	}
	kPorts := []k8sdispatcher.Ports{
		kToLbPort,
		//kToIntPort,
	}
	k := k8sdispatcher.K8sdispatcher{
		Name:            kName,
		Loglevel:        "TRACE",
		Ports:           kPorts,
		ClusterIpSubnet: conf.vClusterCIDR.String(), // TODO testing
		ClientSubnet:    node.Conf.PodCIDR.String(),
		InternalSrcIp:   "3.3.1.3",          // TODO mocked
		NodeportRange:   conf.nodePortRange, // TODO testing
	}

	l = l.WithValues("k8sdispatcher", fmt.Sprintf("%+v", k))
	// creating k8sdispatcher
	if resp, err := k8sdispatcherAPI.CreateK8sdispatcherByID(context.TODO(), kName, k); err != nil {
		l.Error(err, "failed to create k8sdispatcher", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to create k8sdispatcher")
	}
	// TODO trying to create in a single shot also the following port
	if resp, err := k8sdispatcherAPI.CreateK8sdispatcherPortsByID(
		context.TODO(), kName, conf.kToIntPortName, kToIntPort,
	); err != nil {
		l.Error(err,
			"failed to create k8sdispatcher port",
			"response", fmt.Sprintf("%+v", resp),
			"port", conf.kToIntPortName,
		)
		return errors.New("failed to create k8sdispatcher port")
	}
	l.Info("k8sdispatcher created")
	return nil
}

// connectCubes connect each port of the already deployed polycube infrastructure with the right peer
func connectCubes() error {
	brName := conf.brName
	rName := conf.rName
	lbName := conf.lbName
	kName := conf.k8sDispName

	// updating bridge ports
	// updating bridge "to_r0" port in order to set peer=r0:to_br0
	brToRPortName := conf.brToRPortName
	brToRPortPeer := utils.CreatePeer(rName, conf.rToBrPortName)
	l := log.WithValues("name", brName, "port", brToRPortName, "peer", brToRPortPeer)
	brToRPort := simplebridge.Ports{
		Peer: brToRPortPeer,
	}
	if resp, err := simplebridgeAPI.UpdateSimplebridgePortsByID(
		context.TODO(), brName, brToRPortName, brToRPort,
	); err != nil {
		l.Error(err, "failed to set bridge port peer", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to set bridge port peer")
	}
	l.Info("bridge port peer set")

	// updating router ports
	// updating router "to_br0" port in order to set peer=br0:to_r0
	rToBrPortName := conf.rToBrPortName
	rToBrPortPeer := utils.CreatePeer(brName, conf.brToRPortName)
	l = log.WithValues("name", rName, "port", rToBrPortName, "peer", rToBrPortPeer)
	rToBrPort := router.Ports{
		Peer: rToBrPortPeer,
	}
	if resp, err := routerAPI.UpdateRouterPortsByID(context.TODO(), rName, rToBrPortName, rToBrPort); err != nil {
		l.Error(err, "failed to set router port peer", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to set router port peer")
	}
	l.Info("router port peer set")

	// updating router "to_vxlan0" port in order to set peer=vxlan0
	rToVxlanPortName := conf.rToVxlanPortName
	rToVxlanPortPeer := node.Conf.VxlanIface.Link.Attrs().Name
	l = log.WithValues("name", rName, "port", rToVxlanPortName, "peer", rToVxlanPortPeer)
	rToVxlanPort := router.Ports{
		Peer: rToVxlanPortPeer,
	}
	if resp, err := routerAPI.UpdateRouterPortsByID(context.TODO(), rName, rToVxlanPortName, rToVxlanPort); err != nil {
		l.Error(err, "failed to set router port peer", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to set router port peer")
	}
	l.Info("router port peer set")

	// updating router "to_lbrp0" port in order to set peer=lbrp0:to_r0
	rToLbPortName := conf.rToLbPortName
	rToLbPortPeer := utils.CreatePeer(lbName, conf.lbToRPortName)
	l = log.WithValues("name", rName, "port", rToLbPortName, "peer", rToLbPortPeer)
	rToLbPort := router.Ports{
		Peer: rToLbPortPeer,
	}
	if resp, err := routerAPI.UpdateRouterPortsByID(context.TODO(), rName, rToLbPortName, rToLbPort); err != nil {
		l.Error(err, "failed to set router port peer", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to set router port peer")
	}
	l.Info("router port peer set")

	// updating lbrp ports
	// updating lbrp "to_r0" port in order to set peer=r0:to_lbrp0
	lbToRPortName := conf.lbToRPortName
	lbToRPortPeer := utils.CreatePeer(rName, conf.rToLbPortName)
	l = log.WithValues("name", lbName, "port", lbToRPortName, "peer", lbToRPortPeer)
	lbToRPort := lbrp.Ports{
		Peer: lbToRPortPeer,
	}
	if resp, err := lbrpAPI.UpdateLbrpPortsByID(context.TODO(), lbName, lbToRPortName, lbToRPort); err != nil {
		l.Error(err, "failed to set lbrp port peer", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to set lbrp port peer")
	}
	l.Info("lbrp port peer set")

	// updating lbrp "to_k0" port in order to set peer=k0:to_lbrp0
	lbToKPortName := conf.lbToKPortName
	lbToKPortPeer := utils.CreatePeer(kName, conf.kToLbPortName)
	l = log.WithValues("name", lbName, "port", lbToKPortName, "peer", lbToKPortPeer)
	lbToKPort := lbrp.Ports{
		Peer: lbToKPortPeer,
	}
	if resp, err := lbrpAPI.UpdateLbrpPortsByID(context.TODO(), lbName, lbToKPortName, lbToKPort); err != nil {
		l.Error(err, "failed to set lbrp port peer", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to set lbrp port peer")
	}
	l.Info("lbrp port peer set")

	// updating k8sdispatcher ports
	// updating k8sdispatcher "to_lbrp0" port in order to set peer=lbrp0:to_k0
	kToLbPortName := conf.kToLbPortName
	kToLbPortPeer := utils.CreatePeer(lbName, conf.lbToKPortName)
	l = log.WithValues("name", kName, "port", kToLbPortName, "peer", kToLbPortPeer)
	kToLbPort := k8sdispatcher.Ports{
		Peer: kToLbPortPeer,
	}
	if resp, err := k8sdispatcherAPI.UpdateK8sdispatcherPortsByID(
		context.TODO(), kName, kToLbPortName, kToLbPort,
	); err != nil {
		l.Error(err, "failed to set k8sdispatcher port peer", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to set k8sdispatcher port peer")
	}
	l.Info("k8sdispatcher port peer set")

	// updating k8sdispatcher "to_int" port in order to set peer=external_interface_name
	kToIntPortName := conf.kToIntPortName
	kToIntPortPeer := node.Conf.ExtIface.Link.Attrs().Name
	l = log.WithValues("name", kName, "port", kToIntPortName, "peer", kToIntPortPeer)
	kToIntPort := k8sdispatcher.Ports{
		Peer: kToIntPortPeer,
	}
	if resp, err := k8sdispatcherAPI.UpdateK8sdispatcherPortsByID(
		context.TODO(), kName, kToIntPortName, kToIntPort,
	); err != nil {
		l.Error(err, "failed to set k8sdispatcher port peer", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to set k8sdispatcher port peer")
	}
	l.Info("k8sdispatcher port peer set")

	return nil
}

func Deploy() error {
	if err := createBridge(); err != nil {
		return err
	}
	if err := createRouter(); err != nil {
		return err
	}
	if err := createLbrp(); err != nil {
		return err
	}
	if err := createK8sDispatcher(); err != nil {
		return err
	}
	if err := connectCubes(); err != nil {
		return err
	}
	return nil
}

func Init() {
	ec := env.Conf
	conf = &Configuration{
		brName:           ec.BridgeName,
		rName:            ec.RouterName,
		lbName:           ec.LbrpName,
		k8sDispName:      ec.K8sDispName,
		brToRPortName:    "to_" + ec.RouterName,
		rToBrPortName:    "to_" + ec.BridgeName,
		rToVxlanPortName: "to_" + ec.VxlanIfName,
		rToLbPortName:    "to_" + ec.LbrpName,
		lbToRPortName:    "to_" + ec.RouterName,
		lbToKPortName:    "to_" + ec.K8sDispName,
		kToLbPortName:    "to_" + ec.LbrpName,
		kToIntPortName:   "to_int",
		vClusterCIDR:     ec.VClusterCIDR,
		nodePortRange:    ec.NodePortRange,
	}
}
