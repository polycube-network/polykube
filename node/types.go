/*
Copyright 2022 The Polykube Authors.

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

package node

import (
	"fmt"
	"github.com/polycube-network/polykube/types"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"net"
)

type Environment struct {
	PodName                    string
	NodeName                   string
	APIServerIP                net.IP
	APIServerPort              int
	VxlanIfaceName             string
	PolykubeVethPairNamePrefix string
	VtepCIDR                   *net.IPNet
	ClusterCIDR                *net.IPNet
	NodePortRange              string
	CNIConfFilePath            string
	MTU                        int
	IntLbrpName                string
	RouterName                 string
	ExtLbrpName                string
	K8sDispName                string
	CubesLogLevel              string
	CNILogLevel                string
	IsCPNodesDeployAllowed     bool
}

func (e *Environment) APIServerEndpoint() string {
	return fmt.Sprintf("%s:%d", e.APIServerIP, e.APIServerPort)
}

type Configuration struct {
	clientset             *kubernetes.Clientset
	Node                  *v1.Node
	PodCIDR               *net.IPNet
	PodGwInfo             *types.GwInfo
	VPodIPNet             *net.IPNet
	vtepIPNet             *net.IPNet
	VxlanIface            *types.Iface
	ExtIface              *types.Iface
	polykubeVethHostIPNet *net.IPNet
	polykubeVethNetIPNet  *net.IPNet
	PolykubeVeth          *types.PolykubeVeth
	NodeGwInfo            *types.GwInfo
}
