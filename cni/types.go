package main

import (
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"net"
)

type NetConf struct {
	types.NetConf
	MTU          int    `json:"mtu"`
	VClusterCIDR string `json:"vclustercidr"`
	BridgeName   string `json:"bridge"`
	Gw           GwInfo `json:"gateway"`
}

type GwInfo struct {
	IP     net.IP           `json:"ip"`
	RawMAC string           `json:"mac"`
	MAC    net.HardwareAddr `json:"-"`
}

type IFaceConf struct {
	ResultIndex int
	Interface   *current.Interface
	IPConf      *current.IPConfig
}
