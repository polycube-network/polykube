package polycube

import (
	"context"
	"errors"
	"fmt"
	k8sdispatcher "github.com/ekoops/polykube-operator/polycube/clients/k8sdispatcher"
	"github.com/ekoops/polykube-operator/types"
	"net"
	"strings"
)

func getK8sDispatcherNodePortRules() ([]k8sdispatcher.NodeportRule, error) {
	kdName := conf.k8sDispName
	log := log.WithValues("k8sdisp", kdName)
	nprs, resp, err := k8sDispatcherAPI.ReadK8sdispatcherNodeportRuleListByID(context.TODO(), kdName)
	if err != nil {
		log.Error(
			err, "failed to retrieve k8s dispatcher NodePort rules",
			"response", fmt.Sprintf("%+v", resp),
		)
		return nil, errors.New("failed to retrieve k8s dispatcher NodePort rules")
	}
	log.V(1).Info(
		"retrieved k8s dispatcher NodePort rules", "rules", fmt.Sprintf("%+v\n", nprs),
	)
	return nprs, nil

}

func nodePortRuleToFrontend(npr k8sdispatcher.NodeportRule, nodeIP net.IP) types.Frontend {
	return types.Frontend{
		Vip:   nodeIP.String(),
		Vport: npr.Port,
		Proto: strings.ToUpper(npr.Proto),
	}
}

func frontendToK8sDispatcherNodePortRule(f *types.Frontend, sn string, etp string) k8sdispatcher.NodeportRule {
	return k8sdispatcher.NodeportRule{
		Name:        sn,
		Port:        f.Vport,
		Proto:       strings.ToUpper(f.Proto),
		ServiceType: etp,
	}
}

func extractK8sDispatcherNodePortRulesToAddUpdAndDel(
	oldNprList []k8sdispatcher.NodeportRule, svcDetail *types.ServiceDetail, nodeIP net.IP,
) ([]k8sdispatcher.NodeportRule, []k8sdispatcher.NodeportRule, []k8sdispatcher.NodeportRule) {
	svcId := svcDetail.ServiceId
	newFrontendsSet := svcDetail.NodePortFrontendsSet
	etp := svcDetail.ExternalTrafficPolicy

	var nprToAdd []k8sdispatcher.NodeportRule
	var nprToUpd []k8sdispatcher.NodeportRule
	var nprToDel []k8sdispatcher.NodeportRule
	fs := make(types.FrontendsSet)

	for _, npr := range oldNprList {
		if npr.Name != svcId {
			continue
		}
		f := nodePortRuleToFrontend(npr, nodeIP)
		if !newFrontendsSet.Contains(f) {
			// the NodePort rule has to be deleted if is not present in the newFrontendsSet
			nprToDel = append(nprToDel, npr)
		} else {
			if strings.ToUpper(npr.ServiceType) != etp {
				// the NodePort rule has to be updated if it is present in the newFrontendsSet
				npr.ServiceType = etp
				nprToUpd = append(nprToUpd, npr)
			}
			// add to the set the NodePort rule with name = svcId and that have not to be deleted
			fs.Add(f)
		}
	}
	for f := range newFrontendsSet {
		if !fs.Contains(f) {
			npr := frontendToK8sDispatcherNodePortRule(&f, svcId, etp)
			nprToAdd = append(nprToAdd, npr)
		}
	}
	return nprToAdd, nprToUpd, nprToDel
}

func createK8sDispatcherNodePortRule(kdName string, npr *k8sdispatcher.NodeportRule) error {
	log := log.WithValues("k8sdisp", kdName, "nodePortRule", fmt.Sprintf("%+v", *npr))
	resp, err := k8sDispatcherAPI.CreateK8sdispatcherNodeportRuleByID(context.TODO(), kdName, npr.Port, npr.Proto, *npr)
	if err != nil && resp.StatusCode != 409 {
		log.Error(err, "failed to create k8s dispatcher NodePort rule", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to create k8s dispatcher NodePort rule")
	}
	log.V(1).Info("created k8s dispatcher NodePort rule")
	return nil
}

func updateK8sDispatcherNodePortRule(kdName string, npr *k8sdispatcher.NodeportRule) error {
	log := log.WithValues("k8sdisp", kdName, "nodePortRule", fmt.Sprintf("%+v", *npr))
	resp, err := k8sDispatcherAPI.UpdateK8sdispatcherNodeportRuleByID(context.TODO(), kdName, npr.Port, npr.Proto, *npr)
	if err != nil && resp.StatusCode != 409 {
		log.Error(err, "failed to update k8s dispatcher NodePort rule", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to update k8s dispatcher NodePort rule")
	}
	log.V(1).Info("updated k8s dispatcher NodePort rule")
	return nil
}

func addK8sDispatcherNodePortRules(kdName string, nprs []k8sdispatcher.NodeportRule) error {
	log := log.WithValues("k8sdisp", kdName, "rules", fmt.Sprintf("%+v", nprs))

	resp, err := k8sDispatcherAPI.CreateK8sdispatcherNodeportRuleListByID(context.TODO(), kdName, nprs)
	if err != nil {
		log.Error(err, "failed to add k8s dispatcher NodePort rules", "response", fmt.Sprintf("%+v", resp))
		return fmt.Errorf("failed to add k8s dispatcher NodePort rules")
	}
	log.V(1).Info("added k8s dispatcher NodePort rules")
	return nil
}
func updateK8sDispatcherNodePortRules(kdName string, nprs []k8sdispatcher.NodeportRule) error {
	log := log.WithValues("k8sdisp", kdName, "rules", fmt.Sprintf("%+v", nprs))

	resp, err := k8sDispatcherAPI.UpdateK8sdispatcherNodeportRuleListByID(context.TODO(), kdName, nprs)
	if err != nil {
		log.Error(err, "failed to update k8s dispatcher NodePort rules", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to update k8s dispatcher NodePort rules")
	}
	log.V(1).Info("updated k8s dispatcher NodePort rules")
	return nil
}

func deleteK8sDispatcherNodePortRule(kdName string, npr *k8sdispatcher.NodeportRule) error {
	log := log.WithValues("k8sdisp", kdName, "rule", fmt.Sprintf("%+v", *npr))
	resp, err := k8sDispatcherAPI.DeleteK8sdispatcherNodeportRuleByID(context.TODO(), kdName, npr.Port, npr.Proto)
	if err != nil && resp.StatusCode != 409 {
		log.Error(err, "failed to delete k8s dispatcher NodePort rule", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to delete k8s dispatcher NodePort rule")
	}
	log.V(1).Info("deleted k8s dispatcher NodePort rule")
	return nil
}

func SyncK8sDispatcherNodePortRules(svcDetail *types.ServiceDetail, nodeIP net.IP) error {
	kdName := conf.k8sDispName
	kdNprs, err := getK8sDispatcherNodePortRules()
	if err != nil {
		return err
	}
	log := log.WithValues("svcId", svcDetail.ServiceId)

	kdNprsToAdd, kdNprsToUpd, kdNprsToDel := extractK8sDispatcherNodePortRulesToAddUpdAndDel(kdNprs, svcDetail, nodeIP)
	log.V(1).Info("evaluated k8s dispatcher NodePort rules to add, update and delete",
		"toAdd", fmt.Sprintf("%+v", kdNprsToAdd),
		"toUpd", fmt.Sprintf("%+v", kdNprsToUpd),
		"toDel", fmt.Sprintf("%+v", kdNprsToDel),
	)

	lenToAdd := len(kdNprsToAdd)
	lenToUpd := len(kdNprsToUpd)
	lenToDel := len(kdNprsToDel)
	if lenToAdd == 0 && lenToUpd == 0 && lenToDel == 0 {
		log.V(1).Info("k8s dispatcher NodePort rule already synced")
		return nil
	}

	if lenToAdd > 0 {
		if err := addK8sDispatcherNodePortRules(kdName, kdNprsToAdd); err != nil {
			log.Error(err, "error during k8s dispatcher NodePort rules addition")
			return errors.New("error during k8s dispatcher NodePort rules addition")
		}
	}
	if lenToUpd > 0 {
		if err := updateK8sDispatcherNodePortRules(kdName, kdNprsToUpd); err != nil {
			log.Error(err, "error during k8s dispatcher NodePort rules update")
			return errors.New("error during k8s dispatcher NodePort rules update")
		}
	}

	//for _, npr := range kdNprsToAdd {
	//	if err := createK8sDispatcherNodePortRule(kdName, &npr); err != nil {
	//		log.Error(err, "error during k8s dispatcher NodePort rule creation")
	//		return errors.New("error during k8s dispatcher NodePort rule creation")
	//	}
	//}
	//for _, npr := range kdNprsToUpd {
	//	if err := updateK8sDispatcherNodePortRule(kdName, &npr); err != nil {
	//		log.Error(err, "error during k8s dispatcher NodePort rule update")
	//		return errors.New("error during k8s dispatcher NodePort rule update")
	//	}
	//}
	for _, npr := range kdNprsToDel {
		if err := deleteK8sDispatcherNodePortRule(kdName, &npr); err != nil {
			log.Error(err, "error during k8s dispatcher NodePort rule deletion")
			return errors.New("error during k8s dispatcher NodePort rule deletion")
		}
	}
	log.V(1).Info("synced k8s dispatcher NodePort rules")
	return nil
}

func CleanupK8sDispatcherNodePortRulesById(svcId string) error {
	kdName := conf.k8sDispName
	kdNprs, err := getK8sDispatcherNodePortRules()
	if err != nil {
		return err
	}
	log := log.WithValues("k8sdisp", kdName, "svcId", svcId)

	for _, npr := range kdNprs {
		if npr.Name == svcId {
			if err := deleteK8sDispatcherNodePortRule(kdName, &npr); err != nil {
				log.Error(err, "error during k8s dispatcher NodePort rule deletion")
				return errors.New("error during k8s dispatcher NodePort rule deletion")
			}
		}
	}
	log.V(1).Info("cleaned up k8s dispatcher NodePort rules")
	return nil
}

func CleanupK8sDispatcherNodePortRules() error {
	kdName := conf.k8sDispName
	log := log.WithValues("k8sdisp", kdName)

	resp, err := k8sDispatcherAPI.DeleteK8sdispatcherNodeportRuleListByID(context.TODO(), kdName)
	if err != nil {
		log.Error(
			err, "failed to delete all k8s dispatcher NodePort rules",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New("failed to delete all k8s dispatcher NodePort rules")
	}
	log.V(1).Info("deleted all k8s dispatcher NodePort rules")
	return nil
}
