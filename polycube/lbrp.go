package polycube

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/polycube-network/polykube/node"
	lbrp "github.com/polycube-network/polykube/polycube/clients/lbrp"
	"github.com/polycube-network/polykube/types"
	"github.com/polycube-network/polykube/utils"
	corev1 "k8s.io/api/core/v1"
	"net"
	"strings"
)

func createLbrpPorts(lbName string, ports []lbrp.Ports) error {
	log := log.WithValues("lbrp", lbName, "ports", fmt.Sprintf("%+v", ports))

	resp, err := lbrpAPI.CreateLbrpPortsListByID(context.TODO(), lbName, ports)
	if err != nil {
		log.Error(err, "failed to create lbrp ports", "response", fmt.Sprintf("%+v", resp))
		return fmt.Errorf("failed to create lbrp ports")
	}
	log.V(1).Info("created lbrp ports")
	return nil
}

// getIntLbrpFrontendPortsMap returns a map containing all the internal lbrp frontend ports indexed
// by the hexadecimal representation of the port associated IP
func getIntLbrpFrontendPortsMap() (map[string]*lbrp.Ports, error) {
	ilbName := conf.intLbrpName
	l := log.WithValues("name", ilbName)
	ports, resp, err := lbrpAPI.ReadLbrpPortsListByID(context.TODO(), ilbName)
	if err != nil {
		l.Error(
			err, "failed to retrieve internal lbrp ports list",
			"response", fmt.Sprintf("%+v", resp),
		)
		return nil, errors.New("failed to retrieve internal lbrp ports list")
	}
	m := make(map[string]*lbrp.Ports, len(ports)-1)
	for _, port := range ports {
		if port.Type_ == "backend" {
			continue
		}
		portIp := net.ParseIP(port.Ip_)
		if portIp == nil {
			l.Error(
				err, "failed to parse internal lbrp frontend port IP",
				"response", fmt.Sprintf("%+v", resp),
			)
			return nil, errors.New("failed to parse internal lbrp frontend port IP")
		}
		portIp = portIp.To4()
		if portIp == nil {
			l.Error(
				err, "failed to parse internal lbrp frontend port IP",
				"response", fmt.Sprintf("%+v", resp),
			)
			return nil, errors.New("failed to parse internal lbrp frontend port IP")
		}

		portIpHexStr := hex.EncodeToString(portIp)
		m[portIpHexStr] = &port
	}
	l.V(1).Info("retrieved internal lbrp frontend ports info")
	return m, nil
}

// EnsureIntLbrpMissingFrontendPorts ensures that all the pods deployed on the node are connected to the internal
// lbrp.
func EnsureIntLbrpMissingFrontendPorts(pods []corev1.Pod) error {
	portsMap, err := getIntLbrpFrontendPortsMap()
	if err != nil {
		return nil
	}

	thisPodName := node.Env.PodName

	var portsToAdd []lbrp.Ports
	for _, pod := range pods {
		if pod.Name == thisPodName ||
			pod.Spec.HostNetwork == true ||
			pod.ObjectMeta.DeletionTimestamp != nil { // TODO adjust check
			continue
		}
		podIP := net.ParseIP(pod.Status.PodIP)
		if podIP == nil {
			continue
		}
		podIP = podIP.To4()
		if podIP == nil {
			continue
		}
		portId := hex.EncodeToString(podIP)

		if _, ok := portsMap[portId]; !ok {
			portsToAdd = append(portsToAdd, lbrp.Ports{
				Name:  portId,
				Type_: "frontend",
				Ip_:   podIP.String(),
				Peer:  utils.GetHostIfaceName("eth0", portId), // TODO how to return eth0 from api-server?
			})
		}
	}

	if len(portsToAdd) == 0 {
		log.V(1).Info("internal lbrp frontend ports already synced")
		return nil
	}
	if err := createLbrpPorts(conf.intLbrpName, portsToAdd); err != nil {
		log.Error(err, "something went wrong during internal lbrp frontend ports creation")
		return errors.New("something went wrong during internal lbrp frontend ports creation")
	}
	log.V(1).Info("synced internal lbrp frontend ports")
	return nil
}

func getLbrps() ([]*lbrp.Lbrp, error) {
	ilb, resp, err := lbrpAPI.ReadLbrpByID(context.TODO(), conf.intLbrpName)
	if err != nil {
		log.Error(err, "failed to retrieve the internal lbrp", "response", fmt.Sprintf("%+v", resp))
		return nil, errors.New("failed to retrieve the internal lbrp")
	}

	elb, resp, err := lbrpAPI.ReadLbrpByID(context.TODO(), conf.extLbrpName)
	if err != nil {
		log.Error(err, "failed to retrieve the external lbrp", "response", fmt.Sprintf("%+v", resp))
		return nil, errors.New("failed to retrieve the external lbrp")
	}
	return []*lbrp.Lbrp{&ilb, &elb}, nil
}

func lbrpServiceToFrontend(s *lbrp.Service) types.Frontend {
	return types.Frontend{
		Vip:   s.Vip,
		Vport: s.Vport,
		Proto: strings.ToUpper(s.Proto),
	}
}
func frontendToLbrpService(f *types.Frontend, sn string) lbrp.Service {
	return lbrp.Service{
		Name:  sn,
		Vip:   f.Vip,
		Vport: f.Vport,
		Proto: strings.ToUpper(f.Proto),
	}
}

func extractLbrpServicesToAddAndDel(
	svcId string, oldLbSvcList []lbrp.Service, newFrontendsSet types.FrontendsSet,
) ([]lbrp.Service, []lbrp.Service) {
	var lbSvcsToAdd []lbrp.Service
	var lbSvcsToDel []lbrp.Service
	fs := make(types.FrontendsSet)
	for _, lbSvc := range oldLbSvcList {
		if lbSvc.Name != svcId {
			continue
		}
		f := lbrpServiceToFrontend(&lbSvc)
		if !newFrontendsSet.Contains(f) {
			lbSvcsToDel = append(lbSvcsToDel, lbSvc)
		} else {
			// add to the set the services with name = svcId and that have not to be deleted
			fs.Add(f)
		}
	}
	for f := range newFrontendsSet {
		if !fs.Contains(f) {
			lbSvc := frontendToLbrpService(&f, svcId)
			lbSvcsToAdd = append(lbSvcsToAdd, lbSvc)
		}
	}
	return lbSvcsToAdd, lbSvcsToDel
}

func createLbrpServices(lbName string, svcs []lbrp.Service) error {
	log := log.WithValues("lbrp", lbName, "services", fmt.Sprintf("%+v", svcs))

	resp, err := lbrpAPI.CreateLbrpServiceListByID(context.TODO(), lbName, svcs)
	if err != nil {
		log.Error(err, "failed to add lbrp services", "response", fmt.Sprintf("%+v", resp))
		return fmt.Errorf("failed to add lbrp services")
	}
	log.V(1).Info("added lbrp services")
	return nil
}

func deleteLbrpService(lbName string, svc *lbrp.Service) error {
	log := log.WithValues(
		"lbrp", lbName,
		"service", fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto),
	)
	resp, err := lbrpAPI.DeleteLbrpServiceByID(context.TODO(), lbName, svc.Vip, svc.Vport, svc.Proto)
	if err != nil && resp.StatusCode != 409 {
		log.Error(err, "failed to delete lbrp service", "response", fmt.Sprintf("%+v", resp))
		return fmt.Errorf("failed to delete lbrp service")
	}
	log.V(1).Info("deleted lbrp service")
	return nil
}

func SyncLbrpServices(svcDetail *types.ServiceDetail) error {
	svcId := svcDetail.ServiceId
	log := log.WithValues("svcId", svcId)

	lbs, err := getLbrps()
	if err != nil {
		log.Error(err, "failed to retrieve the lbrp list")
		return errors.New("failed to retrieve the lbrp list")
	}

	for _, lb := range lbs {
		lbName := lb.Name
		lbSvcs := lb.Service
		log := log.WithValues("lbrp", lbName)
		var frontendsSet types.FrontendsSet
		if lbName == conf.intLbrpName {
			frontendsSet = svcDetail.ClusterIPFrontendsSet
		} else {
			frontendsSet = svcDetail.NodePortFrontendsSet
		}
		log.V(1).Info("retrieved proper frontend set",
			"set", fmt.Sprintf("%+v", frontendsSet),
		)
		log.V(1).Info("retrieved actual lbrp services",
			"services", fmt.Sprintf("%+v", lb.Service),
		)
		lbSvcsToAdd, lbSvcsToDel := extractLbrpServicesToAddAndDel(svcId, lbSvcs, frontendsSet)
		log.V(1).Info("evaluated lbrp services to add and delete",
			"toAdd", fmt.Sprintf("%+v", lbSvcsToAdd),
			"toDel", fmt.Sprintf("%+v", lbSvcsToDel),
		)

		lenToAdd := len(lbSvcsToAdd)
		lenToDel := len(lbSvcsToDel)
		if lenToAdd == 0 && lenToDel == 0 {
			log.V(1).Info("lbrp services already synced")
			continue
		}

		if lenToAdd > 0 {
			if err := createLbrpServices(lbName, lbSvcsToAdd); err != nil {
				log.Error(err, "error during lbrp services addition")
				return errors.New("error during lbrp services addition")
			}
		}

		for _, lbSvc := range lbSvcsToDel {
			if err := deleteLbrpService(lbName, &lbSvc); err != nil {
				log.Error(err, "error during lbrp service deletion")
				return errors.New("error during lbrp service deletion")
			}
		}
	}
	log.V(1).Info("synced lbrps services")
	return nil
}

// CleanupLbrpsServicesById deletes all the lbrp services associated with the provided service ID
func CleanupLbrpsServicesById(svcId string) error {
	log := log.WithValues("svcId", svcId)

	lbs, err := getLbrps()
	if err != nil {
		log.Error(err, "failed to retrieve the lbrp list")
		return errors.New("failed to retrieve the lbrp list")
	}

	for _, lb := range lbs {
		log := log.WithValues("lbrp", lb.Name)
		log.V(1).Info("retrieved actual lbrp services",
			"services", fmt.Sprintf("%+v", lb.Service),
		)
		for _, svc := range lb.Service {
			if svc.Name == svcId {
				if err := deleteLbrpService(lb.Name, &svc); err != nil {
					log.Error(err, "error during lbrp service deletion")
					return errors.New("error during lbrp service deletion")
				}
			}
		}
	}
	log.V(1).Info("cleaned up lbrps services")
	return nil
}

// CleanupLbrpServices deletes all the services of the specified lbrp
func CleanupLbrpServices(lbName string) error {
	log := log.WithValues("lbrp", lbName)
	resp, err := lbrpAPI.DeleteLbrpServiceListByID(context.TODO(), lbName)
	if err != nil {
		log.Error(err, "failed to delete all lbrp services", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to delete all lbrp services")
	}
	log.V(1).Info("deleted all lbrp services")
	return nil
}
