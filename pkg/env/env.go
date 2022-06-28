package env

import (
	"errors"
	"fmt"
	"github.com/containernetworking/plugins/pkg/ip"
	"github.com/ekoops/polykube-operator/pkg/node"
	"net"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"strconv"
)

const (
	CNIConfTemplate = `
{
	"cniVersion": "0.4.0",
	"name": "mynet",
	"type": "polykube-cni-plugin",
	"mtu": %d,
	"vclustercidr": "%s",
	"bridge": "%s",
	"gateway": {
		"ip": "%s",
		"mac": "%s"
	},
	"ipam": {
		"type": "host-local",
		"ranges": [
			[
				{
					"subnet": "%s",
					"rangeStart": "%s",
					"rangeEnd": "%s",
					"gateway": "%s"
				}
			]
		],
		"dataDir": "/var/lib/cni/networks/mynet",
		"resolvConf": "/etc/resolv.env"
	}
}
`
)

var (
	// logger used throughout the package
	log = ctrl.Log.WithName("env-pkg")
	Conf *Configuration
)

func getEnvVar(envVar string, defaultVal string) string {
	env := os.Getenv(envVar)
	if env == "" {
		env = defaultVal
		log.Info("env variable not found. Default value applied",
			"envVar", envVar, "default", defaultVal,
		)
	}
	return env
}

// LoadConfig create a env object taking values from environment variables and in some cases, if the environment
// variable is not defined, defaulting them
func LoadConfig() error {
	// NodeName
	conf := &Configuration{}
	conf.NodeName = os.Getenv("NODE_K8S_NAME")
	if conf.NodeName == "" {
		log.Error(
			errors.New("env variable not found"),
			"failed to parse env variable",
			"envVar", "K8S_NODE_NAME",
		)
		return errors.New("K8S_NODE_NAME env variable not found")
	}

	// VxlanIfName
	conf.VxlanIfName = getEnvVar("NODE_VXLAN_IFACE_NAME", "vxlan0")

	// VtepCIDR
	vtepCIDRStr := getEnvVar("NODE_VTEP_CIDR", "10.18.0.0/16")
	_, vtepCIDR, err := net.ParseCIDR(vtepCIDRStr)
	if err != nil {
		log.Error(
			err,
			"failed to parse env variable",
			"envVar", "NODE_VTEP_CIDR",
			"value", vtepCIDRStr,
		)
		return errors.New("failed to parse NODE_VTEP_CIDR")
	}
	conf.VtepCIDR = vtepCIDR

	// NodePortRange
	conf.NodePortRange = getEnvVar("NODE_K8S_NODEPORT_RANGE", "30000-32767")

	// CNIConfFilePath
	conf.CNIConfFilePath = getEnvVar("CNI_CONF_FILE_PATH", "/etc/cni/net.d/00-polykube.json")

	// VClusterCIDR
	vClusterCIDRStr := getEnvVar("POLYCUBE_VPODS_RANGE", "10.96.0.0/12")
	_, vClusterCIDR, err := net.ParseCIDR(vClusterCIDRStr)
	if err != nil {
		log.Error(
			err,
			"failed to parse env variable",
			"envVar", "POLYCUBE_VPODS_RANGE",
			"value", vClusterCIDRStr,
		)
		return errors.New("failed to parse POLYCUBE_VPODS_RANGE")
	}
	conf.VClusterCIDR = vClusterCIDR

	// MTU
	MTUStr := getEnvVar("POLYCUBE_MTU", "1450")
	MTU, err := strconv.Atoi(MTUStr)
	if err != nil {
		log.Error(
			err,
			"failed to parse env variable",
			"envVar", "POLYCUBE_MTU",
			"value", MTUStr,
		)
		return errors.New("failed to parse POLYCUBE_MTU")
	}
	conf.MTU = MTU

	// BridgeName
	conf.BridgeName = getEnvVar("POLYCUBE_BRIDGE_NAME", "br0")

	// RouterName
	conf.RouterName = getEnvVar("POLYCUBE_ROUTER_NAME", "r0")

	// LbrpName
	conf.LbrpName = getEnvVar("POLYCUBE_LBRP_NAME", "lbrp0")

	// K8sDispName
	conf.K8sDispName = getEnvVar("POLYCUBE_K8SDISP_NAME", "k0")

	Conf = conf
	log.V(1).Info("environment configuration loaded")
	return nil
}

// CreateCNIConfFile creates the configuration file for the CNI plugin
func CreateCNIConfFile() error {
	fPath := Conf.CNIConfFilePath
	f, err := os.Create(fPath)
	if err != nil {
		log.Error(
			err,
			"failed to create file",
			"path", fPath,
			)
		return errors.New("failed to create file")
	}
	defer f.Close()

	podCIDR := node.Conf.PodCIDR
	podGwIP := node.Conf.PodGwInfo.IPNet.IP
	podGwMAC := node.Conf.PodGwInfo.MAC

	if _, err := fmt.Fprintf(f,
		CNIConfTemplate,
		Conf.MTU,
		Conf.VClusterCIDR,
		Conf.BridgeName,
		podGwIP.String(),
		podGwMAC.String(),
		podCIDR.String(),
		ip.NextIP(podCIDR.IP).String(), // .1
		ip.PrevIP(podGwIP).String(),    // .253
		podGwIP.String(),
	); err != nil {
		log.Error(
			err,
			"failed to write to file",
			"path", fPath,
			)
		return errors.New("failed to write to file")
	}
	log.Info("CNI configuration file created")
	return nil
}
