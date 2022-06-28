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

package main

import (
	"flag"
	"github.com/ekoops/polykube-operator/controllers"
	"github.com/ekoops/polykube-operator/node"
	"github.com/ekoops/polykube-operator/polycube"
	"net/http"
	"os"
	"os/signal"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"syscall"
	"time"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	//+kubebuilder:scaffold:imports
)

const (
	polycubeBasePath = "http://127.0.0.1:9000/polycube/v1"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	//+kubebuilder:scaffold:scheme
}

func exit(code int) {
	node.DeleteVxlanIface()
	os.Exit(code)
}

// setupSignalHandler sets up a handler for handling cleanup if SIGTERM or SIGINT interrupts occur
func setupSignalHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		exit(4)
	}()
}

// ensureDeployment ensures the connection with polycubed, creates all the necessary polycube cubes and connects them
// together and with the existing pods in order to allow network communications
func ensureDeployment() error {
	if err := polycube.EnsureConnection(); err != nil {
		return err
	}
	if err := polycube.EnsureCubes(); err != nil {
		return err
	}
	if err := polycube.EnsureCubesConnections(); err != nil {
		return err
	}

	// the first time ensureDeployment is called, the Pod default gateway MAC must be retrieved from the infrastructure
	// and its value complete the information required for writing the CNI configuration file. If ensureDeployment
	// is called after a polycubed disconnection, simply the old MAC is set to the new polycube router port.
	podGwMAC := node.Conf.PodGwInfo.MAC
	if podGwMAC == nil {
		podGwMAC, err := polycube.GetRouterToIntK8sLbrpPortMAC()
		if err != nil {
			return err
		}
		node.Conf.PodGwInfo.MAC = podGwMAC

		if err := node.EnsureCNIConf(); err != nil {
			return err
		}
	} else {
		if err := polycube.SetRouterToIntK8sLbrpPortMAC(podGwMAC); err != nil {
			return err
		}
	}
	pods, err := node.GetPods(node.Conf.Node.Name)
	if err != nil {
		return err
	}
	if err := polycube.EnsureIntK8sLbrpMissingFrontendPorts(pods.Items); err != nil {
		return err
	}
	setupLog.Info("ensured network infrastructure deployment")
	return nil
}

// startPolycubedPolling starts an asynchronous task in order to periodically queries polycubed. If for some reason
// polycubed is not reachable, an attempts to restore the connection and the polycube network infrastructure is made
// by calling internally ensureDeployment
func startPolycubedPolling() {
	go func() {
		for {
			time.Sleep(10 * time.Second)
			if _, err := http.Get(polycubeBasePath); err != nil {
				setupLog.Info("lost polycubed connection, waiting for it...")
				if err := ensureDeployment(); err != nil {
					setupLog.Error(err, "failed to network infrastructure deployment")
					exit(12)
				}
				setupLog.Info("re-established network infrastructure deployment after connection loss")
			}
		}
	}()
}

func main() {
	var metricsAddr string
	var probeAddr string
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	// obtaining environment configuration by taking it from env variables
	if err := node.LoadEnvironment(); err != nil {
		setupLog.Error(err, "failed to load environment configuration from environment variables")
		os.Exit(1)
	}

	// loading node configuration details
	if err := node.LoadConfig(); err != nil {
		setupLog.Error(err, "failed to load cluster node configuration")
		os.Exit(2)
	}

	// ensuring vxlan interface on the node
	if err := node.EnsureVxlanIface(); err != nil {
		setupLog.Error(err, "failed to ensure vxlan interface")
		os.Exit(3)
	}

	setupSignalHandler()

	// initialize polycube package internal state
	polycube.InitConf()

	// ensure that the network infrastructure is correctly deployed and connected to the already existing entities
	if err := ensureDeployment(); err != nil {
		setupLog.Error(err, "failed to ensure network infrastructure deployment")
		exit(5)
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         false,
	})
	if err != nil {
		setupLog.Error(err, "failed to create manager")
		exit(6)
	}

	if err = (&controllers.NodeReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "failed to create controller", "controller", "node-controller")
		exit(7)
	}
	if err = (&controllers.ServiceReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "failed to create controller", "controller", "service-controller")
		exit(8)
	}
	if err = (&controllers.EndpointsReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "failed to create controller", "controller", "endpoints-controller")
		exit(9)
	}

	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "failed to set up health check")
		exit(10)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "failed to set up ready check")
		exit(11)
	}

	// the following is used in order to react to a possible polycubed crash
	startPolycubedPolling()

	setupLog.Info("starting manager...")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "failed to start manager")
		exit(13)
	}
	// the following is used in order to stop the main goroutine waiting for the second main one to invoke os.Exit
	<-make(chan struct{})
}
