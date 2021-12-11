package node

import (
	"errors"
	"fmt"
	"github.com/containernetworking/plugins/pkg/ip"
	"net"
	"os"
	"strconv"
)

const (
	CNIConfTemplate = `
{
	"cniVersion": "0.4.0",
	"name": "mynet",
	"type": "polykube-cni-plugin",
	"mtu": %d,
	"intK8sLbrp": "%s",
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
		"resolvConf": "/etc/resolv.conf"
	}
}
`
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

// LoadEnvironment initializes an env object taking values from environment variables and in some cases, if the environment
// variable is not defined, defaulting them
func LoadEnvironment() error {
	// NodeName
	env := &Environment{}
	env.NodeName = os.Getenv("NODE_K8S_NAME")
	if env.NodeName == "" {
		log.Error(
			errors.New("env variable not found"),
			"failed to parse env variable",
			"envVar", "K8S_NODE_NAME",
		)
		return errors.New("K8S_NODE_NAME env variable not found")
	}

	// VxlanIfaceName
	env.VxlanIfaceName = getEnvVar("VXLAN_IFACE_NAME", "vxlan0")

	// VtepCIDR
	vtepCIDRStr := getEnvVar("VTEP_CIDR", "10.18.0.0/16")
	_, vtepCIDR, err := net.ParseCIDR(vtepCIDRStr)
	if err != nil {
		log.Error(
			err,
			"failed to parse env variable",
			"envVar", "VTEP_CIDR",
			"value", vtepCIDRStr,
		)
		return errors.New("failed to parse VTEP_CIDR")
	}
	env.VtepCIDR = vtepCIDR

	// VirtualPodCIDR
	virtualPodCIDRStr := getEnvVar("VIRTUAL_POD_CIDR", "3.3.0.0/16")
	_, virtualPodCIDR, err := net.ParseCIDR(virtualPodCIDRStr)
	if err != nil {
		log.Error(
			err,
			"failed to parse env variable",
			"envVar", "VIRTUAL_POD_CIDR",
			"value", virtualPodCIDRStr,
		)
		return errors.New("failed to parse VIRTUAL_POD_CIDR")
	}
	env.VirtualPodCIDR = virtualPodCIDR

	// ClusterIPCIDR
	clusterIPCIDRStr := getEnvVar("CLUSTER_IP_CIDR", "10.96.0.0/12")
	_, clusterIPCIDR, err := net.ParseCIDR(clusterIPCIDRStr)
	if err != nil {
		log.Error(
			err,
			"failed to parse env variable",
			"envVar", "CLUSTER_IP_CIDR",
			"value", clusterIPCIDRStr,
		)
		return errors.New("failed to parse CLUSTER_IP_CIDR")
	}
	env.ClusterIPCIDR = clusterIPCIDR

	// NodePortRange
	env.NodePortRange = getEnvVar("NODEPORT_RANGE", "30000-32767")

	// CNIConfFilePath
	env.CNIConfFilePath = getEnvVar("CNI_CONF_FILE_PATH", "/etc/cni/net.d/00-polykube.json")

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
	env.MTU = MTU

	// IntK8sLbrpName
	env.IntK8sLbrpName = getEnvVar("POLYCUBE_INT_K8S_LBRP_NAME", "ikl0")

	// RouterName
	env.RouterName = getEnvVar("POLYCUBE_ROUTER_NAME", "r0")

	// elbName
	env.ExtK8sLbrpName = getEnvVar("POLYCUBE_EXT_K8S_LBRP_NAME", "ekl0")

	// K8sDispName
	env.K8sDispName = getEnvVar("POLYCUBE_K8SDISP_NAME", "k0")

	Env = env
	log.V(1).Info("environment configuration loaded")
	return nil
}

// EnsureCNIConf creates the configuration file for the CNI plugin
func EnsureCNIConf() error {
	fPath := Env.CNIConfFilePath
	f, err := os.Create(fPath)
	if err != nil {
		log.Error(err, "failed to create file", "path", fPath)
		return errors.New("failed to create file")
	}
	defer f.Close()

	podCIDR := Conf.PodCIDR
	podGwIP := Conf.PodGwInfo.IPNet.IP
	podGwMAC := Conf.PodGwInfo.MAC

	if _, err := fmt.Fprintf(f,
		CNIConfTemplate,
		Env.MTU,
		Env.IntK8sLbrpName,
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
