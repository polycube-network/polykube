package node

import (
	"github.com/ekoops/polykube-operator/pkg/types"
	v1 "k8s.io/api/core/v1"
	"net"
)

type Configuration struct {
	Node          *v1.Node
	PodCIDR       *net.IPNet
	PodGwInfo     *types.GwInfo
	VxlanIface    *types.Iface
	ExtIface      *types.Iface
	NodeVtepIPNet *net.IPNet
	NodeGwInfo    *types.GwInfo
}
