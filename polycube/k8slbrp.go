package polycube

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	k8slbrp "github.com/ekoops/polykube-operator/polycube/clients/k8slbrp"
	"github.com/ekoops/polykube-operator/types"
	"github.com/ekoops/polykube-operator/utils"
	"net"
	"strings"
)

func CreateIntK8sLbrpFrontendPort(portId string, IP net.IP) error {
	iklName := conf.intK8sLbrpName
	port := k8slbrp.Ports{
		Name:  portId,
		Peer:  utils.GetHostIfaceName("eth0", portId),
		Type_: "frontend",
		Ip_:   IP.String(),
	}
	l := log.WithValues("name", iklName, "port", fmt.Sprintf("%+v", port))
	resp, err := k8sLbrpAPI.CreateK8sLbrpPortsByID(context.TODO(), iklName, portId, port)
	if err != nil {
		l.Error(
			err, "failed to create internal k8s lbrp frontend port",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to create internal k8s lbrp frontend port")
	}
	l.V(1).Info("created internal k8s lbrp frontend port")
	return nil
}

// GetIntK8sLbrpFrontendPortsMap returns a map containing all the internal k8s lbrp frontend ports indexed
// by the hexadecimal representation of the port associated IP
func GetIntK8sLbrpFrontendPortsMap() (map[string]*k8slbrp.Ports, error) {
	iklName := conf.intK8sLbrpName
	l := log.WithValues("name", iklName)
	ports, resp, err := k8sLbrpAPI.ReadK8sLbrpPortsListByID(context.TODO(), iklName)
	if err != nil {
		l.Error(
			err, "failed to retrieve internal k8s lbrp ports list",
			"response", fmt.Sprintf("%+v", resp),
		)
		return nil, errors.New("failed to retrieve internal k8s lbrp ports list")
	}
	var m map[string]*k8slbrp.Ports
	for _, port := range ports {
		if port.Type_ == "backend" {
			continue
		}
		portIp := net.ParseIP(port.Ip_)
		if portIp == nil {
			l.Error(
				err, "failed to parse internal k8s lbrp frontend port IP",
				"response", fmt.Sprintf("%+v", resp),
			)
			return nil, errors.New("failed to parse internal k8s lbrp frontend port IP")
		}

		portIpHexStr := hex.EncodeToString(portIp)
		m[portIpHexStr] = &port
	}
	l.V(1).Info("retrieved k8s internal lbrp frontend ports IPs set")
	return m, nil
}

func k8sLbrpServiceToFrontend(s k8slbrp.Service) types.Frontend {
	return types.Frontend{
		Vip:   s.Vip,
		Vport: s.Vport,
		Proto: strings.ToUpper(s.Proto),
	}
}
func frontendToK8sLbrpService(f *types.Frontend, sn string) k8slbrp.Service {
	return k8slbrp.Service{
		Name:  sn,
		Vip:   f.Vip,
		Vport: f.Vport,
		Proto: strings.ToUpper(f.Proto),
	}
}

func extractK8sLbrpServicesToAddAndDel(
	svcId string, oldKlSvcList []k8slbrp.Service, newFrontendsSet types.FrontendsSet,
) ([]k8slbrp.Service, []k8slbrp.Service) {
	var klSvcsToAdd []k8slbrp.Service
	var klSvcsToDel []k8slbrp.Service
	fs := make(types.FrontendsSet)
	for _, klSvc := range oldKlSvcList {
		if klSvc.Name != svcId {
			continue
		}
		f := k8sLbrpServiceToFrontend(klSvc)
		if !newFrontendsSet.Contains(f) {
			klSvcsToDel = append(klSvcsToDel, klSvc)
		} else {
			// add to the set the services with name = svcId and that have not to be deleted
			fs.Add(f)
		}
	}
	for f := range newFrontendsSet {
		if !fs.Contains(f) {
			klSvc := frontendToK8sLbrpService(&f, svcId)
			klSvcsToAdd = append(klSvcsToAdd, klSvc)
		}
	}
	return klSvcsToAdd, klSvcsToDel
}

func createK8sLbrpService(klName string, svc *k8slbrp.Service) error {
	log := log.WithValues(
		"k8slbrp", klName,
		"service", fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto),
	)
	resp, err := k8sLbrpAPI.CreateK8sLbrpServiceByID(context.TODO(), klName, svc.Vip, svc.Vport, svc.Proto, *svc)
	if err != nil && resp.StatusCode != 409 {
		log.Error(err, "failed to create k8s lbrp service", "response", fmt.Sprintf("%+v", resp))
		return fmt.Errorf("failed to create k8s lbrp service")
	}
	log.V(1).Info("created k8s lbrp service")
	return nil
}

func deleteK8sLbrpService(klName string, svc *k8slbrp.Service) error {
	log := log.WithValues(
		"k8slbrp", klName,
		"service", fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto),
	)
	resp, err := k8sLbrpAPI.DeleteK8sLbrpServiceByID(context.TODO(), klName, svc.Vip, svc.Vport, svc.Proto)
	if err != nil && resp.StatusCode != 409 {
		log.Error(err, "failed to delete k8s lbrp service", "response", fmt.Sprintf("%+v", resp))
		return fmt.Errorf("failed to delete k8s lbrp service")
	}
	log.V(1).Info("deleted k8s lbrp service")
	return nil
}

func getK8sLbrps() ([]*k8slbrp.K8sLbrp, error) {
	ikl, resp, err := k8sLbrpAPI.ReadK8sLbrpByID(context.TODO(), conf.intK8sLbrpName)
	if err != nil {
		log.Error(err, "failed to retrieve the internal k8s lbrp", "response", fmt.Sprintf("%+v", resp))
		return nil, errors.New("failed to retrieve the internal k8s lbrp")
	}

	ekl, resp, err := k8sLbrpAPI.ReadK8sLbrpByID(context.TODO(), conf.extK8sLbrpName)
	if err != nil {
		log.Error(err, "failed to retrieve the external k8s lbrp", "response", fmt.Sprintf("%+v", resp))
		return nil, errors.New("failed to retrieve the external k8s lbrp")
	}
	return []*k8slbrp.K8sLbrp{&ikl, &ekl}, nil
}

func SyncK8sLbrpServices(svcDetail *types.ServiceDetail) error {
	svcId := svcDetail.ServiceId
	log := log.WithValues("svcId", svcId)

	kls, err := getK8sLbrps()
	if err != nil {
		log.Error(err, "failed to retrieve the k8s lbrp list")
		return errors.New("failed to retrieve the k8s lbrp list")
	}

	for _, kl := range kls {
		log := log.WithValues("k8slbrp", kl.Name)
		var frontendsSet types.FrontendsSet
		if kl.Name == conf.intK8sLbrpName {
			frontendsSet = svcDetail.ClusterIPFrontendsSet
		} else {
			frontendsSet = svcDetail.NodePortFrontendsSet
		}
		log.V(1).Info("retrieved proper frontend set",
			"set", fmt.Sprintf("%+v", frontendsSet),
		)
		log.V(1).Info("retrieved actual k8s lbrp services",
			"services", fmt.Sprintf("%+v", kl.Service),
		)
		klSvcsToAdd, klSvcsToDel := extractK8sLbrpServicesToAddAndDel(svcId, kl.Service, frontendsSet)
		log.V(1).Info("evaluated k8s lbrp services to add and delete",
			"toAdd", fmt.Sprintf("%+v", klSvcsToAdd),
			"toDel", fmt.Sprintf("%+v", klSvcsToDel),
		)

		for _, klSvc := range klSvcsToAdd {
			if err := createK8sLbrpService(kl.Name, &klSvc); err != nil {
				log.Error(err, "error during k8s lbrp service creation")
				return errors.New("error during k8s lbrp service creation")
			}
		}
		for _, klSvc := range klSvcsToDel {
			if err := deleteK8sLbrpService(kl.Name, &klSvc); err != nil {
				log.Error(err, "error during k8s lbrp service deletion")
				return errors.New("error during k8s lbrp service deletion")
			}
		}
	}
	log.V(1).Info("synced k8s lbrps services")
	return nil
}

// CleanupK8sLbrpsServices deletes all the lbrp services associated with the provided service ID
func CleanupK8sLbrpsServices(svcId string) error {
	log := log.WithValues("svcId", svcId)

	kls, err := getK8sLbrps()
	if err != nil {
		log.Error(err, "failed to retrieve the k8s lbrp list")
		return errors.New("failed to retrieve the k8s lbrp list")
	}

	for _, kl := range kls {
		log := log.WithValues("k8slbrp", kl.Name)
		log.V(1).Info("retrieved actual k8s lbrp services",
			"services", fmt.Sprintf("%+v", kl.Service),
		)
		for _, svc := range kl.Service {
			if svc.Name == svcId {
				if err := deleteK8sLbrpService(kl.Name, &svc); err != nil {
					log.Error(err, "error during k8s lbrp service deletion")
					return errors.New("error during k8s lbrp service deletion")
				}
			}
		}
	}
	log.V(1).Info("cleaned up k8s lbrps services")
	return nil
}

func areBackendsEqual(actualBackendsList []k8slbrp.ServiceBackend, expectedBackendsSet types.BackendsSet) bool {
	if len(actualBackendsList) != len(expectedBackendsSet) {
		return false
	}
	for _, sb := range actualBackendsList {
		b := types.Backend{Ip: sb.Ip, Port: sb.Port}
		if !expectedBackendsSet.Contains(b) {
			return false
		}
	}
	return true
}

func replaceK8sLbrpServiceBackends(klName string, svc *k8slbrp.Service, newBackendsSet types.BackendsSet, epName string) error {
	log := log.WithValues(
		"k8slbrp", klName, "service", fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto),
	)

	var newSvcBackends []k8slbrp.ServiceBackend
	for b := range newBackendsSet {
		newSvcBackends = append(newSvcBackends, k8slbrp.ServiceBackend{
			Name:   epName,
			Ip:     b.Ip,
			Port:   b.Port,
			Weight: 1, // TODO mocked
		})
	}

	if resp, err := k8sLbrpAPI.ReplaceK8sLbrpServiceBackendListByID(
		context.TODO(), klName, svc.Vip, svc.Vport, svc.Proto, newSvcBackends,
	); err != nil {
		log.V(1).Error(
			err, "failed to replace k8s lbrp service backends", "response", fmt.Sprintf("%+v", resp),
		)
		return fmt.Errorf("failed to replace k8s lbrp service backends")
	}
	log.V(1).Info("replaced k8s lbrp service backends")
	return nil
}

func deleteK8sLbrpServiceBackends(klName string, svc *k8slbrp.Service) error {
	log := log.WithValues(
		"k8slbrp", klName, "service", fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto),
	)

	if resp, err := k8sLbrpAPI.DeleteK8sLbrpServiceBackendListByID(
		context.TODO(), klName, svc.Vip, svc.Vport, svc.Proto,
	); err != nil {
		log.V(1).Error(
			err, "failed to delete k8s lbrp service backends", "response", fmt.Sprintf("%+v", resp),
		)
		return fmt.Errorf("failed to delete k8s lbrp service backends")
	}
	log.V(1).Info("deleted k8s lbrp service backends")
	return nil
}

func SyncK8sLbrpsServicesBackends(epsDetail *types.EndpointsDetail) (bool, error) {
	epsId := epsDetail.EndpointsId
	log := log.WithValues("epsId", epsId)

	kls, err := getK8sLbrps()
	if err != nil {
		log.Error(err, "failed to retrieve the k8s lbrp list")
		return false, errors.New("failed to retrieve the k8s lbrp list")
	}

	iklName := conf.intK8sLbrpName
	actualClusterIPServiceNum := 0
	expectedClusterIPServiceNum := len(epsDetail.ClusterIPServiceToBackends)
	actualNodePortServiceNum := 0
	expectedNodePortServiceNum := len(epsDetail.NodePortServiceToBackends)

	for _, kl := range kls {
		log := log.WithValues("k8slbrp", kl.Name)
		for _, svc := range kl.Service {
			if svc.Name == epsId {
				var backendsSet types.BackendsSet
				svcId := fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, strings.ToUpper(svc.Proto))
				log := log.WithValues("svcId", svcId)
				if kl.Name == iklName {
					backendsSet = epsDetail.ClusterIPServiceToBackends.GetBackendsSet(svcId)
					actualClusterIPServiceNum++
				} else {
					backendsSet = epsDetail.NodePortServiceToBackends.GetBackendsSet(svcId)
					actualNodePortServiceNum++
				}
				log.V(1).Info("retrieved proper backend set",
					"set", fmt.Sprintf("%+v", backendsSet),
				)
				log.V(1).Info("retrieved actual k8s lbrp service backends",
					"backends", fmt.Sprintf("%+v", svc.Backend),
				)
				if areBackendsEqual(svc.Backend, backendsSet) {
					log.V(1).Info("k8s lbrp service backends already synced")
					continue
				}
				if len(backendsSet) == 0 {
					if err := deleteK8sLbrpServiceBackends(kl.Name, &svc); err != nil {
						log.Error(err, "error during k8s lbrp service backends deletion")
						return false, errors.New("error during k8s lbrp service backends deletion")
					}
				} else {
					if err := replaceK8sLbrpServiceBackends(kl.Name, &svc, backendsSet, epsId); err != nil {
						log.Error(err, "error during k8s lbrp service backends replacement")
						return false, errors.New("error during k8s lbrp service backends replacement")
					}
				}
			}
		}
	}
	if actualClusterIPServiceNum < expectedClusterIPServiceNum || actualNodePortServiceNum < expectedNodePortServiceNum {
		log.V(1).Info("some services have not been registered yet")
		return true, nil
	}
	log.V(1).Info("synced k8s lbrps services backends")
	return false, nil
}
