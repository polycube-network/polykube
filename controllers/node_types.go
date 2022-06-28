package controllers

import "net"

type NodeDetail struct {
	IP        net.IP
	podCIDR   *net.IPNet
	vtepIPNet *net.IPNet
}
