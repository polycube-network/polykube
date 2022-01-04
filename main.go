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
	"os"
	"os/signal"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"syscall"
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

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	//+kubebuilder:scaffold:scheme
}

func ensurePodsAttachment() error {
	pods, err := node.GetPods(node.Conf.Node.Name)
	if err != nil {
		return err
	}

	portsMap, err := polycube.GetIntK8sLbrpFrontendPortsMap()
	if err != nil {
		setupLog.Error(err, "failed to retrieve retrieve actual infrastructure ports info")
		return err
	}

	if err := polycube.CreateIntK8sLbrpMissingFrontendPorts(pods.Items, portsMap); err != nil {
		setupLog.Error(err, "failed to reattach cluster node pods to the infrastructure")
		return err
	}

	setupLog.Info("ensured cluster node pods attachment")
	return nil
}

func exit(code int) {
	node.CleanupFdb()
	node.DeleteVxlanIface()
	os.Exit(code)
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

	// ---
	// obtaining environment configuration by taking it from env variables
	if err := node.LoadEnvironment(); err != nil {
		setupLog.Error(err, "failed to load environment configuration from environment variables")
		os.Exit(1)
	}

	if err := node.LoadConfig(); err != nil {
		setupLog.Error(err, "failed to load node configuration")
		os.Exit(2)
	}

	if err := node.EnsureVxlanIface(); err != nil {
		setupLog.Error(err, "failed to ensure vxlan interface")
		os.Exit(3)
	}
	if err := node.CleanupFdb(); err != nil {
		setupLog.Error(err, "failed to cleanup fdb")
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		exit(16)
	}()

	polycube.InitConf()

	if err := polycube.EnsureConnection(); err != nil {
		setupLog.Error(err, "failed to ensure polycubed connection")
		exit(4)
	}

	if err := polycube.EnsureDeployment(); err != nil {
		setupLog.Error(err, "failed to ensure polycube infrastructure deployment")
		exit(5)
	}

	podGwMAC, err := polycube.GetRouterToIntK8sLbrpPortMAC()
	if err != nil {
		setupLog.Error(err, "failed to load cluster node pods default gateway mac")
		exit(6)
	}
	node.Conf.PodGwInfo.MAC = podGwMAC

	if err := node.EnsureCNIConf(); err != nil {
		setupLog.Error(err, "failed to ensure CNI configuration")
		exit(7)
	}

	if err := ensurePodsAttachment(); err != nil {
		exit(8)
	}

	//time.Sleep(24 * 60 * 60 * time.Second)

	// ---

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         false,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		exit(9)
	}

	if err = (&controllers.NodeReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Node")
		exit(10)
	}
	if err = (&controllers.ServiceReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Service")
		exit(11)
	}
	if err = (&controllers.EndpointsReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Endpoints")
		exit(12)
	}

	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		exit(13)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		exit(14)
	}

	polycube.PollPolycube()

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		exit(15)
	}
	// the following is used in order to stop the main goroutine waiting for the second main one to invoke os.Exit
	s := make(chan int)
	<-s
}
