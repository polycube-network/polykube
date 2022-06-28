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

// createIntK8sLbrp creates the polycube internal k8s lbrp cube
func createIntK8sLbrp() error {
	iklName := conf.intK8sLbrpName
	l := log.WithValues("name", iklName)

	// defining internal k8s lbrp port that will be connected to the router
	iklToRPort := k8slbrp.Ports{
		Name:  conf.iklToRPortName,
		Type_: "backend",
	}
	iklPorts := []k8slbrp.Ports{iklToRPort}
	ikl := k8slbrp.K8sLbrp{
		Name:      iklName,
		Loglevel:  "TRACE",
		Ports:     iklPorts,
		PortMode_: "MULTI",
	}

	l = l.WithValues("k8slbrp", fmt.Sprintf("%+v", ikl))
	// creating internal k8s lbrp
	if resp, err := k8sLbrpAPI.CreateK8sLbrpByID(context.TODO(), iklName, ikl); err != nil {
		l.Error(err, "failed to create internal k8s lbrp", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to create internal k8s lbrp")
	}
	l.Info("internal k8s lbrp created")
	return nil
}

// createRouter creates a polycube router cube
func createRouter() error {
	extIface := node.Conf.ExtIface
	vxlanIface := node.Conf.VxlanIface
	podsGwInfo := node.Conf.PodGwInfo
	nodeGwInfo := node.Conf.NodeGwInfo
	rName := conf.rName
	l := log.WithValues("name", rName)

	// defining the router port that will be connected to the internal k8s lbrp
	rToIklPort := router.Ports{
		Name: conf.rToIklPortName,
		Ip:   podsGwInfo.IPNet.String(),
		Mac:  podsGwInfo.MAC.String(),
	}
	// defining the router port that will be connected to the vxlan interface
	rToVxlanPort := router.Ports{
		Name: conf.rToVxlanPortName,
		Ip:   vxlanIface.IPNet.String(),
		Mac:  vxlanIface.Link.Attrs().HardwareAddr.String(),
		Peer: vxlanIface.Link.Attrs().Name,
	}
	// defining the router port that will be connected to the external k8s lbrp
	rToEklPort := router.Ports{
		Name: conf.rToEklPortName,
		Ip:   extIface.IPNet.String(),
		Mac:  extIface.Link.Attrs().HardwareAddr.String(),
	}
	rPorts := []router.Ports{rToIklPort, rToVxlanPort, rToEklPort}

	// defining router default route and setting static arp table entry for the default gateway
	routes := []router.Route{
		{
			Network:    "0.0.0.0/0",
			Nexthop:    nodeGwInfo.IPNet.IP.String(),
			Interface_: conf.kToEklPortName,
		},
	}
	arptable := []router.ArpTable{
		{
			Address:    nodeGwInfo.IPNet.IP.String(),
			Mac:        nodeGwInfo.MAC.String(),
			Interface_: conf.kToEklPortName,
		},
	}
	r := router.Router{
		Name:  rName,
		Ports: rPorts,
		//Loglevel: "TRACE",
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

// createExtK8sLbrp creates the polycube external k8s lbrp cube
func createExtK8sLbrp() error {
	eklName := conf.extK8sLbrpName
	l := log.WithValues("name", eklName)

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
		Loglevel:  "TRACE",
		Ports:     eklPorts,
		PortMode_: "SINGLE",
	}

	l = l.WithValues("k8slbrp", fmt.Sprintf("%+v", ekl))
	// creating external k8s lbrp
	if resp, err := k8sLbrpAPI.CreateK8sLbrpByID(context.TODO(), eklName, ekl); err != nil {
		l.Error(err, "failed to create external k8s lbrp", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to create external k8s lbrp")
	}
	l.Info("external k8s lbrp created")
	return nil
}

// createK8sDispatcher creates a polycube k8sdispatcher cube for managing incoming connection
func createK8sDispatcher() error {
	kName := conf.k8sDispName
	l := log.WithValues("name", kName)

	// defining the k8sdispatcher port that will be connected to the external k8s lbrp interface
	kToEklPort := k8sdispatcher.Ports{
		Name:  conf.kToEklPortName,
		Type_: "BACKEND",
	}
	// defining the k8sdispatcher port that will be connected to the node external interface
	kToIntPort := k8sdispatcher.Ports{
		Name:  conf.kToIntPortName,
		Type_: "FRONTEND",
		Ip_:   node.Conf.ExtIface.IPNet.IP.String(),
		Peer:  node.Conf.ExtIface.Link.Attrs().Name,
	}
	kPorts := []k8sdispatcher.Ports{kToEklPort, kToIntPort}

	k := k8sdispatcher.K8sdispatcher{
		Name: kName,
		//Loglevel:      "TRACE",
		Ports:         kPorts,
		InternalSrcIp: "3.3.1.3", // TODO mocked
		NodeportRange: conf.nodePortRange,
	}

	l = l.WithValues("k8sdispatcher", fmt.Sprintf("%+v", k))
	// creating k8sdispatcher
	if resp, err := k8sDispatcherAPI.CreateK8sdispatcherByID(context.TODO(), kName, k); err != nil {
		l.Error(err, "failed to create k8sdispatcher", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to create k8sdispatcher")
	}
	l.Info("k8sdispatcher created")
	return nil
}

// ensureCubesConnections connect each port of the already deployed polycube infrastructure with the right peer
func ensureCubesConnections() error {
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
			err,
			"failed to set internal k8s lbrp port peer",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to set internal k8s lbrp port peer")
	}
	l.Info("internal k8s lbrp port peer set")

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
	l.Info("router port peer set")

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
	l.Info("router port peer set")

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
			err,
			"failed to set external k8s lbrp port peer",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to set external k8s lbrp port peer")
	}
	l.Info("external k8s lbrp port peer set")

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
			err,
			"failed to set external k8s lbrp port peer",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to set external k8s lbrp port peer")
	}
	l.Info("external k8s lbrp port peer set")

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
		l.Error(err, "failed to set k8sdispatcher port peer", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to set k8sdispatcher port peer")
	}
	l.Info("k8sdispatcher port peer set")

	return nil
}

func ensureCubes() error {
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
		l.Info("failed to find internal k8s lbrp: creating it...")
		if err := createIntK8sLbrp(); err != nil {
			return err
		}
	}
	// creating external k8s lbrp if it doesn't exist
	l = log.WithValues("name", eklName)
	if !eklExist {
		l.Info("failed to find external k8s lbrp: creating it...")
		if err := createExtK8sLbrp(); err != nil {
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
		l.Info("failed to find router: creating it...")
		if err := createRouter(); err != nil {
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
		l.Info("failed to find k8s dispatcher: creating it...")
		if err := createK8sDispatcher(); err != nil {
			return err
		}
	}
	return nil
}

func EnsureDeployment() error {
	if err := ensureCubes(); err != nil {
		return err
	}
	if err := ensureCubesConnections(); err != nil {
		return err
	}
	return nil
}

func InitConf() {
	ec := node.Env
	conf = &configuration{
		intK8sLbrpName:   ec.IntK8sLbrpName,
		rName:            ec.RouterName,
		extK8sLbrpName:   ec.ExtK8sLbrpName,
		k8sDispName:      ec.K8sDispName,
		iklToRPortName:   "to_" + ec.RouterName,
		rToIklPortName:   "to_" + ec.IntK8sLbrpName,
		rToVxlanPortName: "to_" + ec.VxlanIfaceName,
		rToEklPortName:   "to_" + ec.ExtK8sLbrpName,
		eklToRPortName:   "to_" + ec.RouterName,
		eklToKPortName:   "to_" + ec.K8sDispName,
		kToEklPortName:   "to_" + ec.ExtK8sLbrpName,
		kToIntPortName:   "to_int",
		vClusterCIDR:     ec.ClusterIPCIDR,
		nodePortRange:    ec.NodePortRange,
	}
}

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
	return errors.New("polycubed is unreachable")
}

func PollPolycube() {
	go func() {
		for {
			time.Sleep(10 * time.Second)
			if _, err := http.Get(basePath); err != nil {
				log.Info("lost polycubed connection, waiting for it...")
				panic("polycubed is not ready anymore")
				// TODO stop manager
				//if err := stopControllers(); err != nil {
				//	// TODO panic?
				//}
				//if err := EnsureConnection(); err != nil {
				//	// TODO panic?
				//}
			}
		}
	}()
}
