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
	"github.com/ekoops/polykube-operator/pkg/node"
	"github.com/ekoops/polykube-operator/pkg/polycube"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
)

// ServiceReconciler reconciles a Service object
type ServiceReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// buildServiceDetail builds a struct containing the info about two frontends sets: one for the ClusterIP Service
// feature and the other for the NodePort one (if the service is of type NodePort; otherwise this will be empty)
func buildServiceDetail(s *corev1.Service) *ServiceDetail {
	nodeIP := node.Conf.ExtIface.IPNet.IP.String()
	st := s.Spec.Type

	sd := &ServiceDetail{
		ClusterIPFrontendsSet: make(FrontendsSet),
		NodePortFrontendsSet:  make(FrontendsSet),
	}

	for _, p := range s.Spec.Ports {
		for _, vip := range s.Spec.ClusterIPs {
			f := Frontend{
				Vip:   vip,
				Vport: p.Port,
				Proto: string(p.Protocol),
			}
			sd.ClusterIPFrontendsSet[f] = struct{}{}
		}

		if st == corev1.ServiceTypeNodePort {
			f := Frontend{
				Vip:   nodeIP,
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
	l := ctrllog.FromContext(ctx, "service", sId)

	s := &corev1.Service{}
	if err := r.Get(ctx, req.NamespacedName, s); err != nil {
		if errors.IsNotFound(err) {
			// the following work for both the internal and the external load balancers.
			// In case o services other than ClusterIP and NodePort ones, the following
			// will not delete anything
			if err := polycube.CleanupLbrpsServices(sId); err != nil {
				l.Error(err, "something went wrong during related lbrp services cleanup")
				return ctrl.Result{}, err
			}
			l.Info("cleanup performed on service object deletion event")
			return ctrl.Result{}, nil
		}
		l.Error(err, "failed to retrieve the service object")
		return ctrl.Result{}, err
	}

	l.Info("service object retrieved")

	st := s.Spec.Type
	// not able to handle services other than ClusterIP and NodePort ones
	if st != corev1.ServiceTypeClusterIP && st != corev1.ServiceTypeNodePort {
		l.Info("service type not supported. Skipped.", "type", st)
		return ctrl.Result{}, nil
	}

	sd := buildServiceDetail(s)
	if err := polycube.SyncLbrpsServices(sId, sd); err != nil {
		l.Error(err, "something went wrong during lbrps services sync")
		return ctrl.Result{}, err
	}

	l.Info("cluster status reconciled for the service object")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ServiceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Service{}).
		Complete(r)
}
