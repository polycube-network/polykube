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
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"net"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
)

// EndpointsReconciler reconciles an Endpoints object
type EndpointsReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// buildServiceToBackends builds a map containing for each port (identified by the name
// or, in case of a single service port, by the placeholder "-") a set containing all the
// associated backends.
func buildServiceToBackends(ep *corev1.Endpoints) types.ServiceToBackends {
	stb := make(types.ServiceToBackends)
	for _, sub := range ep.Subsets {
		// each EndpointPort object is associated one-to-one with a virtual service
		for _, p := range sub.Ports {
			var svcId string
			if p.Name == "" { // this is the case of a single service port
				svcId = "-"
			} else {
				svcId = p.Name
			}
			// each address along with the port of an EndpointPort identifies a backend that has to be associated to the
			// virtual service
			for _, addr := range sub.Addresses {
				stb.Add(svcId, types.Backend{Ip: addr.IP, Port: p.Port})
			}
		}
	}
	return stb
}

// isExternalTrafficPolicyCompliant verifies that the provided backend is compliant with the provided external traffic
// policy. A backend is always compliant if the external traffic policy is set to "Cluster". In case it was
// set to "Local", the backend must have an IP contained in the local node Pod CIDR.
func isExternalTrafficPolicyCompliant(b *types.Backend, etp corev1.ServiceExternalTrafficPolicyType) bool {
	if etp == corev1.ServiceExternalTrafficPolicyTypeCluster {
		return true
	}
	ipa := net.ParseIP(b.Ip)
	if ipa == nil {
		return false
	}
	return node.Conf.PodCIDR.Contains(ipa)
}

// isInternalTrafficPolicyCompliant verifies that the provided backend is compliant with the provided internal traffic
// policy. A backend is always compliant if the external traffic policy is set to "Cluster" or is not defined.
// In case it was set to "Local", the backend must have an IP contained in the local node Pod CIDR.
func isInternalTrafficPolicyCompliant(b *types.Backend, itp *corev1.ServiceInternalTrafficPolicyType) bool {
	if itp == nil || *itp == corev1.ServiceInternalTrafficPolicyCluster {
		return true
	}
	ipa := net.ParseIP(b.Ip)
	if ipa == nil {
		return false
	}
	return node.Conf.PodCIDR.Contains(ipa)
}

// buildEndpointsDetail builds a struct containing the info about the backends sets organized per service.
func buildEndpointsDetail(s *corev1.Service, stb types.ServiceToBackends) *types.EndpointsDetail {
	itp := s.Spec.InternalTrafficPolicy
	etp := s.Spec.ExternalTrafficPolicy
	sd := &types.EndpointsDetail{
		ClusterIPServiceToBackends: make(types.ServiceToBackends),
		NodePortServiceToBackends:  make(types.ServiceToBackends),
	}

	// each Frontend represents a virtual service. It has to be associated
	// with the corresponding BackendsSet. The BackendsSet is extracted through
	// the ServiceToBackends struct using the port name or the "-" placeholder in
	// case of a single port.
	for _, p := range s.Spec.Ports {
		var backends map[types.Backend]struct{}
		if p.Name == "" {
			backends = stb.GetBackendsSet("-")
		} else {
			backends = stb.GetBackendsSet(p.Name)
		}
		for _, vip := range s.Spec.ClusterIPs {
			// the virtual service will be identified by the triple vip:vport:proto
			id := fmt.Sprintf("%s:%d:%s", vip, p.Port, p.Protocol)
			for b := range backends {
				if !isInternalTrafficPolicyCompliant(&b, itp) {
					continue
				}
				sd.ClusterIPServiceToBackends.Add(id, b)
			}
		}
		if s.Spec.Type == corev1.ServiceTypeNodePort {
			id := fmt.Sprintf("%s:%d:%s", node.Conf.ExtIface.IPNet.IP, p.NodePort, p.Protocol)
			for b := range backends {
				if !isExternalTrafficPolicyCompliant(&b, etp) {
					continue
				}
				sd.NodePortServiceToBackends.Add(id, b)
			}
		}
	}
	return sd
}

//+kubebuilder:rbac:groups=core,resources=endpoints,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=endpoints/status,verbs=get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Polykube object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *EndpointsReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	epsId := req.NamespacedName.String()
	l := ctrllog.FromContext(ctx, "endpoints", epsId)

	eps := &corev1.Endpoints{}
	if err := r.Get(ctx, req.NamespacedName, eps); err != nil {
		l.Error(err, "failed to retrieve the endpoints object")
		// if the endpoints resource is not found, it means that the associated
		// service was deleted: in this case, let the service controller handle
		// the cleanup
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	l.Info("endpoints object retrieved")

	s := &corev1.Service{}
	if err := r.Get(ctx, req.NamespacedName, s); err != nil {
		l.Error(err, "failed to retrieve the associated service object")
		// if the service resource is not found, the service controller will handle
		// the cleanup
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	l.Info("associated service object retrieved")

	st := s.Spec.Type
	// not able to handle services other than ClusterIP and NodePort ones
	if st != corev1.ServiceTypeClusterIP && st != corev1.ServiceTypeNodePort {
		l.Info("service type of the associated service not supported. Skipped.", "type", st)
		return ctrl.Result{}, nil
	}

	serviceToBackends := buildServiceToBackends(eps)
	endpointsDetail := buildEndpointsDetail(s, serviceToBackends)
	if err := polycube.SyncK8sLbrpsServicesBackends(eps.Name, endpointsDetail); err != nil {
		l.Error(err, "something went wrong during k8s lbrps services backends sync")
		return ctrl.Result{}, err
	}

	l.Info("cluster status reconciled for the endpoints object")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *EndpointsReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Service{}).
		Complete(r)
}
