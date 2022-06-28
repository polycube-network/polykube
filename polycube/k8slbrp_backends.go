package polycube

import (
	"context"
	"errors"
	"fmt"
	k8slbrp "github.com/ekoops/polykube-operator/polycube/clients/k8slbrp"
	"github.com/ekoops/polykube-operator/types"
	"strings"
)

func areBackendsSynced(actualBackendsList []k8slbrp.ServiceBackend, expectedBackendsSet types.BackendsSet) bool {
	if len(actualBackendsList) != len(expectedBackendsSet) {
		return false
	}
	for _, sb := range actualBackendsList {
		b := types.Backend{Ip: sb.Ip, Port: sb.Port, Weight: sb.Weight}
		if !expectedBackendsSet.Contains(b) {
			return false
		}
	}
	return true
}

func replaceK8sLbrpServiceBackends(klName string, svc *k8slbrp.Service, newBackendsSet types.BackendsSet, epsId string) error {
	log := log.WithValues(
		"k8slbrp", klName, "service", fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto),
	)

	var newSvcBackends []k8slbrp.ServiceBackend
	for b := range newBackendsSet {
		newSvcBackends = append(newSvcBackends, k8slbrp.ServiceBackend{
			Name:   epsId,
			Ip:     b.Ip,
			Port:   b.Port,
			Weight: b.Weight,
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
			err, "failed to delete all k8s lbrp service backends", "response", fmt.Sprintf("%+v", resp),
		)
		return fmt.Errorf("failed to all delete k8s lbrp service backends")
	}
	log.V(1).Info("deleted all k8s lbrp service backends")
	return nil
}

func backendToK8sLbrpServiceBackend(b *types.Backend, epsId string) *k8slbrp.ServiceBackend {
	return &k8slbrp.ServiceBackend{
		Name:   epsId,
		Ip:     b.Ip,
		Port:   b.Port,
		Weight: b.Weight,
	}
}

func k8sLbrpServiceBackendToBackend(sb *k8slbrp.ServiceBackend) *types.Backend {
	return &types.Backend{
		Ip:     sb.Ip,
		Port:   sb.Port,
		Weight: sb.Weight,
	}
}

func extractK8sLbrpServiceBackendsToAddAndDel(
	epsId string, oldKlSbList []k8slbrp.ServiceBackend, newBackendsSet types.BackendsSet,
) ([]k8slbrp.ServiceBackend, []k8slbrp.ServiceBackend) {
	var klBkdToAdd []k8slbrp.ServiceBackend
	var klBkdToDel []k8slbrp.ServiceBackend
	bs := make(types.BackendsSet)
	for _, klSb := range oldKlSbList {
		if klSb.Name != epsId {
			continue
		}
		b := k8sLbrpServiceBackendToBackend(&klSb)
		if !newBackendsSet.Contains(*b) {
			klBkdToDel = append(klBkdToDel, klSb)
		} else {
			// add to the set the services with name = epsId and that have not to be deleted
			bs.Add(*b)
		}
	}
	for b := range newBackendsSet {
		if !bs.Contains(b) {
			sb := backendToK8sLbrpServiceBackend(&b, epsId)
			klBkdToAdd = append(klBkdToAdd, *sb)
		}
	}
	return klBkdToAdd, klBkdToDel
}
func addK8sLbrpServiceBackends(klName string, svc *k8slbrp.Service, sbs []k8slbrp.ServiceBackend) error {
	log := log.WithValues(
		"k8slbrp", klName, "service", fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto),
	)

	resp, err := k8sLbrpAPI.CreateK8sLbrpServiceBackendListByID(context.TODO(), klName, svc.Vip, svc.Vport, svc.Proto, sbs)
	if err != nil {
		log.Error(err, "failed to add k8s lbrp service backends", "response", fmt.Sprintf("%+v", resp))
		return fmt.Errorf("failed to add k8s lbrp service backends")
	}
	log.V(1).Info("added k8s lbrp service backends")
	return nil
}

func deleteK8sLbrpServiceBackend(klName string, svc *k8slbrp.Service, sb *k8slbrp.ServiceBackend) error {
	log := log.WithValues(
		"k8slbrp", klName,
		"service", fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto),
		"backend", fmt.Sprintf("%s:%d", sb.Ip, sb.Port),
	)

	resp, err := k8sLbrpAPI.DeleteK8sLbrpServiceBackendByID(context.TODO(), klName, svc.Vip, svc.Vport, svc.Proto, sb.Ip)
	if err != nil {
		log.Error(err, "failed to delete k8s lbrp service backend", "response", fmt.Sprintf("%+v", resp))
		return fmt.Errorf("failed to delete k8s lbrp service backend")
	}
	log.V(1).Info("deleted k8s lbrp service backend")
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
		klName := kl.Name
		log := log.WithValues("k8slbrp", klName)
		for _, svc := range kl.Service {
			if svc.Name == epsId {
				var backendsSet types.BackendsSet
				svcId := fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, strings.ToUpper(svc.Proto))
				log := log.WithValues("svcId", svcId)
				if klName == iklName {
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

				klSbsToAdd, klSbsToDel := extractK8sLbrpServiceBackendsToAddAndDel(epsId, svc.Backend, backendsSet)
				log.V(1).Info("evaluated k8s lbrp service backends to add and delete",
					"toAdd", fmt.Sprintf("%+v", klSbsToAdd),
					"toDel", fmt.Sprintf("%+v", klSbsToDel),
				)

				lenToAdd := len(klSbsToAdd)
				lenToDel := len(klSbsToDel)
				if lenToAdd == 0 && lenToDel == 0 {
					log.V(1).Info("k8s lbrp service backends already synced")
					continue
				}

				if lenToAdd > 0 {
					if err := addK8sLbrpServiceBackends(klName, &svc, klSbsToAdd); err != nil {
						log.Error(err, "error during k8s lbrp service backends addition")
						return false, errors.New("error during k8s lbrp service backends addition")
					}
				}

				for _, klSb := range klSbsToDel {
					if err := deleteK8sLbrpServiceBackend(klName, &svc, &klSb); err != nil {
						log.Error(err, "error during k8s lbrp service backend deletion")
						return false, errors.New("error during k8s lbrp service backend deletion")
					}
				}

				//if areBackendsSynced(svc.Backend, backendsSet) {
				//	log.V(1).Info("k8s lbrp service backends already synced")
				//	continue
				//}
				//if len(backendsSet) == 0 {
				//	if err := deleteK8sLbrpServiceBackends(kl.Name, &svc); err != nil {
				//		log.Error(err, "error during k8s lbrp service backends deletion")
				//		return false, errors.New("error during k8s lbrp service backends deletion")
				//	}
				//} else {
				//	if err := replaceK8sLbrpServiceBackends(kl.Name, &svc, backendsSet, epsId); err != nil {
				//		log.Error(err, "error during k8s lbrp service backends replacement")
				//		return false, errors.New("error during k8s lbrp service backends replacement")
				//	}
				//}
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
