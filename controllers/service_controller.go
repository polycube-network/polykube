/*
Copyright 2021 ekoops.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"github.com/ekoops/polykube-operator/node"
	"github.com/ekoops/polykube-operator/polycube"
	"github.com/ekoops/polykube-operator/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"net"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
)

// ServiceReconciler reconciles a Service object
type ServiceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// buildServiceDetail builds a struct containing the info about two frontends sets: one for the ClusterIP Service
// feature and the other for the NodePort one (if the service is of type NodePort; otherwise this second one will
// be empty)
func buildServiceDetail(s *corev1.Service, sId string, nodeIP net.IP) *types.ServiceDetail {
	nip := nodeIP.String()
	st := s.Spec.Type
	itp := "CLUSTER"
	if s.Spec.InternalTrafficPolicy != nil && *s.Spec.InternalTrafficPolicy == corev1.ServiceInternalTrafficPolicyLocal {
		itp = "LOCAL"
	}
	etp := "CLUSTER"
	if s.Spec.ExternalTrafficPolicy == corev1.ServiceExternalTrafficPolicyTypeLocal {
		etp = "LOCAL"
	}

	sd := &types.ServiceDetail{
		ServiceId:             sId,
		ClusterIPFrontendsSet: make(types.FrontendsSet),
		NodePortFrontendsSet:  make(types.FrontendsSet),
		InternalTrafficPolicy: itp,
		ExternalTrafficPolicy: etp,
	}

	for _, p := range s.Spec.Ports {
		for _, vip := range s.Spec.ClusterIPs {
			f := types.Frontend{
				Vip:   vip,
				Vport: p.Port,
				Proto: string(p.Protocol),
			}
			sd.ClusterIPFrontendsSet[f] = struct{}{}
		}

		if st == corev1.ServiceTypeNodePort {
			f := types.Frontend{
				Vip:   nip,
				Vport: p.NodePort,
				Proto: string(p.Protocol),
			}
			sd.NodePortFrontendsSet[f] = struct{}{}
		}
	}
	return sd
}

//+kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=services/status,verbs=get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Polykube object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *ServiceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	sId := req.NamespacedName.String()
	log := ctrllog.FromContext(ctx)

	s := &corev1.Service{}
	if err := r.Get(ctx, req.NamespacedName, s); err != nil {
		// if the error is different from not found returns error...
		if !errors.IsNotFound(err) {
			log.Error(err, "failed to retrieve the service object")
			return ctrl.Result{}, err
		}
		// the following work for both the internal and the external load balancers.
		// In case o services other than ClusterIP and NodePort ones, the following
		// will not delete anything
		if err := polycube.CleanupK8sLbrpsServicesById(sId); err != nil {
			log.Error(err, "something went wrong during related k8s lbrp services cleanup")
			return ctrl.Result{}, err
		}
		if err := polycube.CleanupK8sDispatcherNodePortRulesById(sId); err != nil {
			log.Error(err, "something went wrong during related k8s dispatcher NodePort rules cleanup")
			return ctrl.Result{}, err
		}
		log.Info("cluster status reconciled after service object deletion event")
		return ctrl.Result{}, nil
	}

	log.Info("service object retrieved")

	st := s.Spec.Type
	// not able to handle services other than ClusterIP and NodePort ones
	if st != corev1.ServiceTypeClusterIP && st != corev1.ServiceTypeNodePort {
		log.Info("service type not supported. Skipped.", "type", st)
		return ctrl.Result{}, nil
	}

	nodeIP := node.Conf.ExtIface.IPNet.IP
	sd := buildServiceDetail(s, sId, nodeIP)

	log.V(1).WithValues("detail", fmt.Sprintf("%+v", sd)).Info("built service detail")

	if err := polycube.SyncK8sLbrpServices(sd); err != nil {
		log.Error(err, "something went wrong during k8s lbrp services sync")
		return ctrl.Result{}, err
	}

	if err := polycube.SyncK8sDispatcherNodePortRules(sd, nodeIP); err != nil {
		log.Error(err, "something went wrong during k8s dispatcher NodePort rules sync")
		return ctrl.Result{}, err
	}

	log.Info("cluster status reconciled for the service object")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ServiceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Service{}).
		Complete(r)
}
