package env

import "net"

type Configuration struct {
	NodeName        string
	VxlanIfName     string
	VtepCIDR        *net.IPNet
	CNIConfFilePath string
	VClusterCIDR    *net.IPNet
	NodePortRange   string
	MTU             int
	BridgeName      string
	RouterName      string
	LbrpName        string
	K8sDispName     string
}
