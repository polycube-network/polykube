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
	"errors"
	"fmt"
	"github.com/ekoops/polykube-operator/node"
	"github.com/ekoops/polykube-operator/polycube"
	"github.com/ekoops/polykube-operator/types"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"net"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// NodeReconciler reconciles a Node object
type NodeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

var NodeDetailMap map[string]*types.NodeDetail

const controlPlaneLabel = "node-role.kubernetes.io/control-plane"

func buildNodeDetail(n *corev1.Node) (*types.NodeDetail, error) {
	IP := node.GetIP(n)
	if IP == nil {
		return nil, errors.New("failed to parse node IP")
	}

	_, podCIDR, err := net.ParseCIDR(n.Spec.PodCIDR)
	if err != nil {
		return nil, errors.New("failed to parse node Pod CIDR")
	}

	vtepIPNet, err := node.CalcVtepIPNet(podCIDR)
	if err != nil {
		return nil, errors.New("failed to calculate node vtep")
	}

	return &types.NodeDetail{
		IP:        IP,
		PodCIDR:   podCIDR,
		VtepIPNet: vtepIPNet,
	}, nil
}

//+kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=nodes/status,verbs=get;list
//+kubebuilder:rbac:groups=core,resources=nodes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Polykube object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *NodeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	nId := req.NamespacedName.String()
	log := ctrllog.FromContext(ctx)

	n := &corev1.Node{}
	if err := r.Get(ctx, req.NamespacedName, n); err != nil {
		// if the error is different from not found returns error...
		if !apierrors.IsNotFound(err) {
			log.Error(err, "failed to retrieve the node object")
			return ctrl.Result{}, err
		}
		nodeDetail, found := NodeDetailMap[nId]
		if !found {
			return ctrl.Result{}, nil
		}
		// the node object was deleted, so it has to be removed from the infrastructure
		if err := polycube.DeleteRouterRoute(nodeDetail.PodCIDR, nodeDetail.VtepIPNet.IP); err != nil {
			log.Error(err, "something wrong during route deletion")
			return ctrl.Result{}, err
		}
		if err := node.DeleteFdbEntry(nodeDetail.IP); err != nil {
			log.Error(err, "something wrong during node bridge fdb entry deletion")
			return ctrl.Result{}, err
		}
		delete(NodeDetailMap, nId)
		log.Info("reconciled cluster status after node object deletion event")
		return ctrl.Result{}, nil
	}

	log.Info("retrieved node object")

	nodeDetail, err := buildNodeDetail(n)
	if err != nil {
		log.Error(err, "failed to build node detail")
		return ctrl.Result{}, err
	}

	if _, isControlPlaneNode := n.Labels[controlPlaneLabel]; isControlPlaneNode || !node.IsReady(n) {
		if err := polycube.DeleteRouterRoute(nodeDetail.PodCIDR, nodeDetail.VtepIPNet.IP); err != nil {
			log.Error(err, "something wrong during route deletion")
			return ctrl.Result{}, err
		}
		if err := node.DeleteFdbEntry(nodeDetail.IP); err != nil {
			log.Error(err, "something wrong during node bridge fdb entry deletion")
			return ctrl.Result{}, err
		}
		log.Info("reconciled cluster status for the node object")
		return ctrl.Result{}, nil
	}

	// the following will create or update the entry related to the node inside the map: this
	// is needed in order to handle a possible future node deletion
	NodeDetailMap[nId] = nodeDetail

	routeExist, err := polycube.CheckRouterRouteExistence(nodeDetail.PodCIDR, nodeDetail.VtepIPNet.IP)
	if err != nil {
		log.Error(err, "something wrong during route existence check")
		return ctrl.Result{}, err
	}
	if !routeExist {
		log.V(1).Info("route doesn't exist: creating it...")
		if err := polycube.CreateRouterRoute(nodeDetail.PodCIDR, nodeDetail.VtepIPNet.IP); err != nil {
			log.Error(err, "something wrong during route creation")
			return ctrl.Result{}, err
		}
	}

	entryExist, err := node.CheckFdbEntryExistence(nodeDetail.IP)
	if err != nil {
		log.Error(err, "something wrong during node bridge fdb entry existence check")
		return ctrl.Result{}, err
	}
	if !entryExist {
		log.V(1).Info("bridge fdb entry doesn't exist: creating it...")
		if err := node.CreateFdbEntry(nodeDetail.IP); err != nil {
			log.Error(err, "something wrong during node bridge fdb entry creation")
			return ctrl.Result{}, err
		}
	}

	log.Info("reconciled cluster status for the node object")

	return ctrl.Result{}, nil
}

func nodeControllerPredicate() predicate.Predicate {
	return predicate.Funcs{
		CreateFunc: func(e event.CreateEvent) bool {
			if e.Object.GetName() == node.Conf.Node.Name {
				return false
			}
			n, ok := e.Object.(*corev1.Node)
			if !ok {
				return false
			}
			_, isControlPlaneNode := n.Labels[controlPlaneLabel]
			return !isControlPlaneNode
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			if e.ObjectOld.GetName() == node.Conf.Node.Name {
				return false
			}
			oldNode, isOldNode := e.ObjectOld.(*corev1.Node)
			newNode, isNewNode := e.ObjectNew.(*corev1.Node)
			if !isOldNode || !isNewNode {
				return false
			}

			_, isControlPlaneOldNode := oldNode.Labels[controlPlaneLabel]
			_, isControlPlaneNewNode := newNode.Labels[controlPlaneLabel]
			if isControlPlaneOldNode != isControlPlaneNewNode {
				return true
			}

			if oldNode.Spec.PodCIDR != newNode.Spec.PodCIDR {
				return true
			}

			if fmt.Sprintf("%+v", oldNode.Status.Addresses) != fmt.Sprintf("%+v", newNode.Status.Addresses) {
				return true
			}

			isOldNodeReady := false
			isNewNodeReady := false
			for _, c := range oldNode.Status.Conditions {
				if c.Type == corev1.NodeReady && c.Status == corev1.ConditionTrue {
					isOldNodeReady = true
					break
				}
			}
			for _, c := range newNode.Status.Conditions {
				if c.Type == corev1.NodeReady && c.Status == corev1.ConditionTrue {
					isNewNodeReady = true
					break
				}
			}
			if isOldNodeReady != isNewNodeReady {
				return true
			}
			return false
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			if e.Object.GetName() == node.Conf.Node.Name {
				return false
			}
			n, ok := e.Object.(*corev1.Node)
			if !ok {
				return false
			}
			_, isControlPlaneNode := n.Labels[controlPlaneLabel]
			return !isControlPlaneNode
		},
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *NodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	NodeDetailMap = make(map[string]*types.NodeDetail)
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Node{}).
		WithEventFilter(nodeControllerPredicate()).
		Complete(r)
}
