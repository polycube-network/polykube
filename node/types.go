package node

import (
	"github.com/ekoops/polykube-operator/types"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"net"
)

type Environment struct {
	NodeName        string
	VxlanIfaceName  string
	VtepCIDR        *net.IPNet
	VirtualPodCIDR  *net.IPNet
	ClusterIPCIDR   *net.IPNet
	NodePortRange   string
	CNIConfFilePath string
	MTU             int
	IntK8sLbrpName  string
	RouterName      string
	ExtK8sLbrpName  string
	K8sDispName     string
}

type Configuration struct {
	Node          *v1.Node
	PodCIDR       *net.IPNet
	PodGwInfo     *types.GwInfo
	VxlanIface    *types.Iface
	ExtIface      *types.Iface
	NodeVtepIPNet *net.IPNet
	NodeGwInfo    *types.GwInfo
	clientset     *kubernetes.Clientset
}
