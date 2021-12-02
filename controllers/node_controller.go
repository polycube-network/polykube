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
	"github.com/ekoops/polykube-operator/pkg/node"
	"github.com/ekoops/polykube-operator/pkg/polycube"
	"github.com/ekoops/polykube-operator/pkg/utils"
	"github.com/go-logr/logr"
	"github.com/prometheus/common/log"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"net"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	"sync"
)

// NodeReconciler reconciles a Node object
type NodeReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

func buildNodeDetail(n *corev1.Node) (*NodeDetail, error) {
	IP := node.GetIP(n)
	if IP == nil {
		return nil, errors.New("failed to parse node IP")
	}

	_, podCIDR, err := net.ParseCIDR(n.Spec.PodCIDR)
	if err != nil {
		return nil, errors.New("failed to parse node Pod CIDR")
	}

	vtepIPNet, err := node.CalcVtepIPNet(n)
	if err != nil {
		return nil, errors.New("failed to calculate node vtep")
	}
	return &NodeDetail{
		IP:        IP,
		podCIDR:   podCIDR,
		vtepIPNet: vtepIPNet,
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
	l := ctrllog.FromContext(ctx, "node", nId)

	n := &corev1.Node{}
	if err := r.Get(ctx, req.NamespacedName, n); err != nil {
		l.Error(err, "failed to retrieve the service object")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	l.Info("node object retrieved")

	// name of the custom finalizer
	finalizer := "polykube.com/finalizer"

	// examine DeletionTimestamp to determine if object is under deletion
	if !n.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is being deleted
		if utils.ContainsString(n.GetFinalizers(), finalizer) {
			// the finalizer is present, so perform cleanup steps
			nodeInfo, err := buildNodeDetail(n)
			if err != nil {
				log.Error(err, "failed to get node info")
				return ctrl.Result{}, err
			}

			if err := polycube.DeleteRouteToNodePodCIDR(nodeInfo.podCIDR, nodeInfo.vtepIPNet.IP); err != nil {
				return ctrl.Result{}, err
			}

			if err := node.DeleteFdbEntry(nodeInfo.IP); err != nil {
				return ctrl.Result{}, err
			}

			// remove our finalizer from the list and update it.
			controllerutil.RemoveFinalizer(n, finalizer)
			if err := r.Update(ctx, n); err != nil {
				return ctrl.Result{}, err
			}
		}
		// Stop reconciliation as the item is being deleted
		return ctrl.Result{}, nil
	}

	// the object is not under deletion: if the finalizers is not present, Add it
	if !utils.ContainsString(n.GetFinalizers(), finalizer) {
		controllerutil.AddFinalizer(n, finalizer)
		if err := r.Update(ctx, n); err != nil {
			return ctrl.Result{}, err
		}
	}

	if !node.IsReady(n) {
		// stop reconciliation since a new event will be fired when the node will be ready
		log.Info("node not ready")
		return ctrl.Result{}, nil
	}

	nodeInfo, err := buildNodeDetail(n)
	if err != nil {
		log.Error(err, "failed to build node detail")
		return ctrl.Result{}, err
	}

	var lastErr error
	wg := sync.WaitGroup{}
	wg.Add(2)
	c := make(chan error)

	go func() {
		defer wg.Done()
		routeExist, err := polycube.CheckRouteExistence(nodeInfo.podCIDR, nodeInfo.vtepIPNet.IP)
		if err != nil {
			c <- err
		}
		if routeExist {
			c <- nil
		}
		c <- polycube.CreateRouteToNodePodCIDR(nodeInfo.podCIDR, nodeInfo.vtepIPNet.IP)
	}()

	go func() {
		defer wg.Done()
		entryExist, err := node.CheckFdbEntryExistence(nodeInfo.IP)
		if err != nil {
			c <- err
		}
		if entryExist {
			c <- nil
		}
		c <- node.CreateFdbEntry(nodeInfo.IP)
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for err := range c {
		if err != nil {
			log.Error(err, "error during cluster status reconciliation for the node object")
			lastErr = err
		}
	}

	if lastErr != nil {
		return ctrl.Result{}, lastErr
	}

	log.Info("cluster status reconciled for the node object")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Node{}).
		Complete(r)
}
