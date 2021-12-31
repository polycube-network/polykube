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
	"encoding/hex"
	"flag"
	"github.com/ekoops/polykube-operator/controllers"
	"github.com/ekoops/polykube-operator/node"
	"github.com/ekoops/polykube-operator/polycube"
	"github.com/ekoops/polykube-operator/utils"
	v1 "k8s.io/api/core/v1"
	"net"
	"os"
	"os/signal"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sync"
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

func attachPods() error {
	pods, err := node.GetPods(node.Conf.Node.Name)
	if err != nil {
		setupLog.Error(err, "failed to retrieve node pods")
		os.Exit(9)
	}
	portsMap, err := polycube.GetIntK8sLbrpFrontendPortsMap()
	if err != nil {
		setupLog.Error(err, "failed to retrieve k8s lbrp frontend ports IPs set")
		os.Exit(9)
	}

	c := make(chan error)
	wg := sync.WaitGroup{}
	wg.Add(len(pods.Items))
	for _, pod := range pods.Items {
		go func(pod *v1.Pod) {
			defer wg.Done()
			if pod.Name == "this_pod_name" { // TODO mocked
				c <- nil
			}
			podIP := net.ParseIP(pod.Status.PodIP)
			if podIP == nil {
				return
			}
			portId := hex.EncodeToString(podIP)

			port, ok := portsMap[portId]
			if !ok {
				if err := polycube.CreateIntK8sLbrpFrontendPort(portId, podIP); err != nil {
					c <- err
					return
				}
				c <- nil
				return
			}

			hostIfaceName := utils.GetHostIfaceName("eth0", portId)
			if port.Peer != hostIfaceName {
				// TODO cu ccu cazzu è ncucciatu?
			}

			if port.Status == "DOWN" {
				// TODO how is it possible?
			}
			c <- nil
		}(&pod)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for err := range c {
		if err != nil {
			return err
		}
	}

	return nil
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
		os.Exit(3)
	}

	if err := node.LoadConfig(); err != nil {
		setupLog.Error(err, "failed to load node configuration")
		os.Exit(4)
	}

	if err := node.EnsureVxlanIface(); err != nil {
		setupLog.Error(err, "failed to ensure vxlan interface")
		os.Exit(5)
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		node.CleanupFdb()
		node.DeleteVxlanIface()
	}()

	polycube.InitConf()

	if err := polycube.EnsureConnection(); err != nil {
		setupLog.Error(err, "failed to ensure polycubed connection")
		os.Exit(6)
	}

	if err := polycube.EnsureDeployment(); err != nil {
		setupLog.Error(err, "failed to ensure polycube infrastructure deployment")
		os.Exit(6)
	}

	podGwMAC, err := polycube.GetRouterToIntK8sLbrpPortMAC()
	if err != nil {
		setupLog.Error(err, "failed to load node pod default gateway mac")
		os.Exit(7)
	}
	node.Conf.PodGwInfo.MAC = podGwMAC

	if err := node.EnsureCNIConf(); err != nil {
		setupLog.Error(err, "failed to ensure CNI configuration")
		os.Exit(8)
	}

	//time.Sleep(24 * 60 * 60 * time.Second)

	//if err := attachPods(); err != nil {
	//	os.Exit(9)
	//}
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
		os.Exit(1)
	}

	//if err = (&controllers.NodeReconciler{
	//	Client: mgr.GetClient(),
	//	Scheme: mgr.GetScheme(),
	//}).SetupWithManager(mgr); err != nil {
	//	setupLog.Error(err, "unable to create controller", "controller", "Node")
	//	os.Exit(1)
	//}
	if err = (&controllers.ServiceReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Service")
		os.Exit(1)
	}
	if err = (&controllers.EndpointsReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Endpoints")
		os.Exit(1)
	}

	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	polycube.PollPolycube()

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
