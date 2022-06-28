package types

import "net"

type NodeDetail struct {
	IP        net.IP
	PodCIDR   *net.IPNet
	VtepIPNet *net.IPNet
}
