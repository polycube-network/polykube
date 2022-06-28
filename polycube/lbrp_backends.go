package polycube

import (
	"context"
	"errors"
	"fmt"
	lbrp "github.com/ekoops/polykube-operator/polycube/clients/lbrp"
	"github.com/ekoops/polykube-operator/types"
	"strings"
)

func backendToLbrpServiceBackend(b *types.Backend, epsId string) *lbrp.ServiceBackend {
	return &lbrp.ServiceBackend{
		Name:   epsId,
		Ip:     b.Ip,
		Port:   b.Port,
		Weight: b.Weight,
	}
}

func lbrpServiceBackendToBackend(sb *lbrp.ServiceBackend) *types.Backend {
	return &types.Backend{
		Ip:     sb.Ip,
		Port:   sb.Port,
		Weight: sb.Weight,
	}
}

func extractLbrpServiceBackendsToAddAndDel(
	epsId string, oldLbSbList []lbrp.ServiceBackend, newBackendsSet types.BackendsSet,
) ([]lbrp.ServiceBackend, []lbrp.ServiceBackend) {
	var lbBkdToAdd []lbrp.ServiceBackend
	var lbBkdToDel []lbrp.ServiceBackend
	bs := make(types.BackendsSet)
	for _, lbSb := range oldLbSbList {
		if lbSb.Name != epsId {
			continue
		}
		b := lbrpServiceBackendToBackend(&lbSb)
		if !newBackendsSet.Contains(*b) {
			lbBkdToDel = append(lbBkdToDel, lbSb)
		} else {
			// add to the set the services with name = epsId and that have not to be deleted
			bs.Add(*b)
		}
	}
	for b := range newBackendsSet {
		if !bs.Contains(b) {
			sb := backendToLbrpServiceBackend(&b, epsId)
			lbBkdToAdd = append(lbBkdToAdd, *sb)
		}
	}
	return lbBkdToAdd, lbBkdToDel
}
func createLbrpServiceBackends(lbName string, svc *lbrp.Service, sbs []lbrp.ServiceBackend) error {
	log := log.WithValues(
		"lbrp", lbName, "service", fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto),
	)

	resp, err := lbrpAPI.CreateLbrpServiceBackendListByID(context.TODO(), lbName, svc.Vip, svc.Vport, svc.Proto, sbs)
	if err != nil {
		log.Error(err, "failed to add lbrp service backends", "response", fmt.Sprintf("%+v", resp))
		return fmt.Errorf("failed to add lbrp service backends")
	}
	log.V(1).Info("added lbrp service backends")
	return nil
}

func deleteLbrpServiceBackend(lbName string, svc *lbrp.Service, sb *lbrp.ServiceBackend) error {
	log := log.WithValues(
		"lbrp", lbName,
		"service", fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto),
		"backend", fmt.Sprintf("%s:%d", sb.Ip, sb.Port),
	)

	resp, err := lbrpAPI.DeleteLbrpServiceBackendByID(context.TODO(), lbName, svc.Vip, svc.Vport, svc.Proto, sb.Ip)
	if err != nil {
		log.Error(err, "failed to delete lbrp service backend", "response", fmt.Sprintf("%+v", resp))
		return fmt.Errorf("failed to delete lbrp service backend")
	}
	log.V(1).Info("deleted lbrp service backend")
	return nil
}

func SyncLbrpsServicesBackends(epsDetail *types.EndpointsDetail) (bool, error) {
	epsId := epsDetail.EndpointsId
	log := log.WithValues("epsId", epsId)

	lbs, err := getLbrps()
	if err != nil {
		log.Error(err, "failed to retrieve the lbrp list")
		return false, errors.New("failed to retrieve the lbrp list")
	}

	ilbName := conf.intLbrpName
	actualClusterIPServiceNum := 0
	expectedClusterIPServiceNum := len(epsDetail.ClusterIPServiceToBackends)
	actualNodePortServiceNum := 0
	expectedNodePortServiceNum := len(epsDetail.NodePortServiceToBackends)

	for _, lb := range lbs {
		lbName := lb.Name
		log := log.WithValues("lbrp", lbName)
		for _, svc := range lb.Service {
			if svc.Name == epsId {
				var backendsSet types.BackendsSet
				svcId := fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, strings.ToUpper(svc.Proto))
				log := log.WithValues("svcId", svcId)
				if lbName == ilbName {
					backendsSet = epsDetail.ClusterIPServiceToBackends.GetBackendsSet(svcId)
					actualClusterIPServiceNum++
				} else {
					backendsSet = epsDetail.NodePortServiceToBackends.GetBackendsSet(svcId)
					actualNodePortServiceNum++
				}
				log.V(1).Info("retrieved proper backend set",
					"set", fmt.Sprintf("%+v", backendsSet),
				)
				log.V(1).Info("retrieved actual lbrp service backends",
					"backends", fmt.Sprintf("%+v", svc.Backend),
				)

				lbSbsToAdd, lbSbsToDel := extractLbrpServiceBackendsToAddAndDel(epsId, svc.Backend, backendsSet)
				log.V(1).Info("evaluated lbrp service backends to add and delete",
					"toAdd", fmt.Sprintf("%+v", lbSbsToAdd),
					"toDel", fmt.Sprintf("%+v", lbSbsToDel),
				)

				lenToAdd := len(lbSbsToAdd)
				lenToDel := len(lbSbsToDel)
				if lenToAdd == 0 && lenToDel == 0 {
					log.V(1).Info("lbrp service backends already synced")
					continue
				}

				if lenToAdd > 0 {
					if err := createLbrpServiceBackends(lbName, &svc, lbSbsToAdd); err != nil {
						log.Error(err, "error during lbrp service backends addition")
						return false, errors.New("error during lbrp service backends addition")
					}
				}

				for _, lbSb := range lbSbsToDel {
					if err := deleteLbrpServiceBackend(lbName, &svc, &lbSb); err != nil {
						log.Error(err, "error during lbrp service backend deletion")
						return false, errors.New("error during lbrp service backend deletion")
					}
				}
			}
		}
	}
	if actualClusterIPServiceNum < expectedClusterIPServiceNum || actualNodePortServiceNum < expectedNodePortServiceNum {
		log.V(1).Info("some services have not been registered yet")
		return true, nil
	}
	log.V(1).Info("synced lbrps services backends")
	return false, nil
}
