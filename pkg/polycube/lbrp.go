package polycube

import (
	"context"
	"errors"
	"fmt"
	"github.com/ekoops/polykube-operator/controllers"
	lbrp "github.com/ekoops/polykube-operator/pkg/clients/lbrp"
	"sync"
)

func toFrontend(s lbrp.Service) controllers.Frontend {
	return controllers.Frontend{
		Vip:   s.Vip,
		Vport: s.Vport,
		Proto: s.Proto,
	}
}
func toLbrpService(f *controllers.Frontend, sn string) lbrp.Service {
	return lbrp.Service{
		Name:  sn,
		Vip:   f.Vip,
		Vport: f.Vport,
		Proto: f.Proto,
	}
}

func extractLbrpServicesToAddAndDel(
	svcId string, oldLbSvcList []lbrp.Service, newFrontendsSet controllers.FrontendsSet,
) ([]lbrp.Service, []lbrp.Service) {
	var lbSvcsToAdd []lbrp.Service
	var lbSvcsToDel []lbrp.Service
	fs := make(controllers.FrontendsSet)
	for _, lbSvc := range oldLbSvcList {
		if lbSvc.Name != svcId {
			continue
		}
		f := toFrontend(lbSvc)
		if !newFrontendsSet.Contains(f) {
			lbSvcsToDel = append(lbSvcsToDel, lbSvc)
		} else {
			// add to the set the services with name = svcId and that has not to be deleted
			fs.Add(f)
		}
		if _, ok := newFrontendsSet[f]; !ok {
			lbSvcsToDel = append(lbSvcsToDel, lbSvc)
		} else {

		}
	}
	for f := range newFrontendsSet {
		if !fs.Contains(f) {
			lbSvc := toLbrpService(&f, svcId)
			lbSvcsToAdd = append(lbSvcsToAdd, lbSvc)
		}
	}
	return lbSvcsToAdd, lbSvcsToDel
}

func createLbrpService(lbName string, svc *lbrp.Service) error {
	l := log.WithValues(
		"lbrp", lbName,
		"service", fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto),
	)
	resp, err := lbrpAPI.CreateLbrpServiceByID(context.TODO(), lbName, svc.Vip, svc.Vport, svc.Proto, *svc)
	if err != nil && resp.StatusCode != 409 {
		log.V(2).Error(
			err,
			"failed to create lbrp service",
			"response", fmt.Sprintf("%+v", resp),
		)
		return fmt.Errorf("failed to create lbrp service")
	}
	l.V(2).Info("created lbrp service")
	return nil
}

func deleteLbrpService(lbName string, svc *lbrp.Service) error {
	l := log.WithValues(
		"lbrp", lbName,
		"service", fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto),
	)
	resp, err := lbrpAPI.DeleteLbrpServiceByID(context.TODO(), lbName, svc.Vip, svc.Vport, svc.Proto)
	if err != nil && resp.StatusCode != 409 {
		l.V(2).Error(
			err,
			"failed to delete lbrp service",
			"response", fmt.Sprintf("%+v", resp),
		)
		return fmt.Errorf("failed to delete lbrp service")
	}
	l.V(2).Info("deleted lbrp service")
	return nil
}

func SyncLbrpsServices(svcId string, svcDetail *controllers.ServiceDetail) error {
	type ExitStatus struct {
		subject string
		err     error
	}

	// retrieving the list of all the node load balancers
	lbs, resp, err := lbrpAPI.ReadLbrpListByID(context.TODO())
	if err != nil {
		log.Error(err, "failed to retrieve the lbrps list", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to retrieve the lbrps list")
	}

	var lastErr error

	c := make(chan ExitStatus)
	wg := &sync.WaitGroup{}
	wg.Add(len(lbs))

	for _, lb := range lbs {
		go func(lb *lbrp.Lbrp) {
			defer wg.Done()
			var frontendsSet controllers.FrontendsSet
			if lb.Name == conf.lbName {
				frontendsSet = svcDetail.NodePortFrontendsSet
			} else {
				frontendsSet = svcDetail.ClusterIPFrontendsSet
			}
			// extractLbrpServiceToAddAndDel
			lbSvcsToAdd, lbSvcsToDel := extractLbrpServicesToAddAndDel(svcId, lb.Service, frontendsSet)

			for _, lbSvc := range lbSvcsToAdd {
				if err := createLbrpService(lb.Name, &lbSvc); err != nil {
					c <- ExitStatus{subject: lb.Name, err: err}
					return
				}
			}
			for _, lbSvc := range lbSvcsToDel {
				if err := deleteLbrpService(lb.Name, &lbSvc); err != nil {
					c <- ExitStatus{subject: lb.Name, err: err}
					return
				}
			}
			c <- ExitStatus{subject: lb.Name, err: nil}
		}(&lb)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for exitStatus := range c {
		if exitStatus.err != nil {
			lastErr = err
			log.V(1).Error(exitStatus.err, "operation failed on lbrp", "lbrp", exitStatus.subject)
		}
	}

	return lastErr
}

// CleanupLbrpsServices deletes all the lbrp services associated with the provided service ID
func CleanupLbrpsServices(svcId string) error {
	type ExitStatus struct {
		subject string
		err     error
	}

	// retrieving the list of all the node load balancers
	lbs, resp, err := lbrpAPI.ReadLbrpListByID(context.TODO())
	if err != nil {
		log.Error(err, "failed to retrieve the lbrps list", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to retrieve the lbrps list")
	}

	var lastErr error

	c := make(chan ExitStatus)
	wg := &sync.WaitGroup{}
	wg.Add(len(lbs))

	for _, lb := range lbs {
		go func(lb *lbrp.Lbrp) {
			defer wg.Done()
			for _, svc := range lb.Service {
				if svc.Name == svcId {
					if err := deleteLbrpService(lb.Name, &svc); err != nil {
						c <- ExitStatus{subject: lb.Name, err: err}
						return
					}
				}
			}
			c <- ExitStatus{subject: lb.Name, err: nil}
		}(&lb)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for exitStatus := range c {
		if exitStatus.err != nil {
			lastErr = err
			log.V(1).Error(exitStatus.err, "operation failed on lbrp", "lbrp", exitStatus.subject)
		}
	}
	return lastErr
}

func areBackendsEqual(actualBackendsList []lbrp.ServiceBackend, expectedBackendsSet controllers.BackendsSet) bool {
	if len(actualBackendsList) != len(expectedBackendsSet) {
		return false
	}
	for _, b := range actualBackendsList {
		key := controllers.Backend{Ip: b.Ip, Port: b.Port}
		if _, ok := expectedBackendsSet[key]; !ok {
			return false
		}
	}
	return true
}

func replaceBackends(epName string, lbName string, svc *lbrp.Service, backendsSet controllers.BackendsSet) error {
	l := log.WithValues(
		"lbrp", lbName,
		"service", fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto),
	)

	var newBackends []lbrp.ServiceBackend
	for b := range backendsSet {
		newBackends = append(newBackends, lbrp.ServiceBackend{
			Name:   epName,
			Ip:     b.Ip,
			Port:   b.Port,
			Weight: 1, // TODO mocked
		})
	}
	resp, err := lbrpAPI.ReplaceLbrpServiceBackendListByID(
		context.TODO(),
		lbName,
		svc.Vip,
		svc.Vport,
		svc.Proto,
		newBackends,
	)
	if err != nil {
		l.V(2).Error(
			err,
			"failed to replace lbrp service backends",
			"response", fmt.Sprintf("%+v", resp),
		)
		return fmt.Errorf("failed to replace lbrp service backends")
	}
	l.V(2).Info("deleted lbrp service")
	return nil
}

func SyncLbrpsServicesBackends(epsName string, epsDetail *controllers.EndpointsDetail) error {
	type ExitStatus struct {
		subject string
		err     error
	}

	// retrieving the list of all the node load balancers
	lbs, resp, err := lbrpAPI.ReadLbrpListByID(context.TODO())
	if err != nil {
		log.Error(err, "failed to retrieve the lbrps list", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to retrieve the lbrps list")
	}

	var lastErr error

	c := make(chan ExitStatus)
	wg := &sync.WaitGroup{}
	wg.Add(len(lbs))

	for _, lb := range lbs {
		go func(lb *lbrp.Lbrp) {
			for _, svc := range lb.Service {
				if svc.Name == epsName {
					var backendsSet controllers.BackendsSet
					svcId := fmt.Sprintf("%s:%d:%s", svc.Vip, svc.Vport, svc.Proto)
					if lb.Name != conf.lbName {
						backendsSet = epsDetail.ClusterIPServiceToBackends.GetBackendsSet(svcId)
					} else {
						backendsSet = epsDetail.NodePortServiceToBackends.GetBackendsSet(svcId)
					}
					if areBackendsEqual(svc.Backend, backendsSet) {
						return
					}
					if err := replaceBackends(epsName, lb.Name, &svc, backendsSet); err != nil {
						c <- ExitStatus{subject: lb.Name, err: err}
						return
					}
				}
			}
			c <- ExitStatus{subject: lb.Name, err: nil}
		}(&lb)
	}

	go func() {
		wg.Wait()
		close(c)
	}()
	for exitStatus := range c {
		if exitStatus.err != nil {
			lastErr = err
			log.V(1).Error(exitStatus.err, "operation failed on lbrp", "lbrp", exitStatus.subject)
		}
	}
	return lastErr
}
