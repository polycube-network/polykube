package types

import (
	"github.com/vishvananda/netlink"
	"net"
)

type GwInfo struct {
	IPNet *net.IPNet
	MAC   net.HardwareAddr
}

type Iface struct {
	IPNet *net.IPNet
	Link  netlink.Link
}

type PolykubeVeth struct {
	Host *Iface
	Net  *Iface
}
