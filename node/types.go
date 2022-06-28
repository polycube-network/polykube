package node

import (
	"fmt"
	"github.com/ekoops/polykube-operator/types"
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
	IntK8sLbrpName             string
	RouterName                 string
	ExtK8sLbrpName             string
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
