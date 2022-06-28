package polycube

import "net"

type Configuration struct {
	brName           string
	rName            string
	lbName           string
	k8sDispName      string
	brToRPortName    string
	rToBrPortName    string
	rToVxlanPortName string
	rToLbPortName    string
	lbToRPortName    string
	lbToKPortName    string
	kToLbPortName    string
	kToIntPortName   string
	vClusterCIDR     *net.IPNet
	nodePortRange    string
}
