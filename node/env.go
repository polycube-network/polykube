package node

import (
	"errors"
	"fmt"
	"github.com/containernetworking/plugins/pkg/ip"
	"github.com/polycube-network/polykube/utils"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	CNIConfTemplate = `
{
	"cniVersion": "0.4.0",
	"name": "mynet",
	"type": "polykube-cni-plugin",
	"mtu": %d,
	"intLbrp": "%s",
	"logLevel": "%s",
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
	env := &Environment{}

	// PodName
	env.PodName = os.Getenv("POD_NAME")
	if env.PodName == "" {
		log.Error(
			errors.New("env variable not found"),
			"failed to parse env variable",
			"envVar", "POD_NAME",
		)
		return errors.New("POD_NAME env variable not found")
	}

	// NodeName
	env.NodeName = os.Getenv("NODE_K8S_NAME")
	if env.NodeName == "" {
		log.Error(
			errors.New("env variable not found"),
			"failed to parse env variable",
			"envVar", "K8S_NODE_NAME",
		)
		return errors.New("K8S_NODE_NAME env variable not found")
	}

	// APIServerIP
	APIServerIPStr := os.Getenv("API_SERVER_IP")
	APIServerIP := net.ParseIP(APIServerIPStr)
	if APIServerIP == nil {
		log.Error(
			errors.New("invalid API Server IP address"),
			"failed to parse env variable",
			"envVar", "API_SERVER_IP",
			"value", APIServerIPStr,
		)
		return errors.New("failed to parse API_SERVER_IP")
	}
	env.APIServerIP = APIServerIP

	// APIServerPort
	APIServerPortStr := os.Getenv("API_SERVER_PORT")
	APIServerPort, err := strconv.Atoi(APIServerPortStr)
	if err != nil || APIServerPort < 1 || APIServerPort > 65535 {
		log.Error(
			errors.New("invalid API Server port number"),
			"failed to parse env variable",
			"envVar", "API_SERVER_PORT",
			"value", APIServerPortStr,
		)
		return errors.New("failed to parse API_SERVER_PORT")
	}
	env.APIServerPort = APIServerPort

	// VxlanIfaceName
	env.VxlanIfaceName = getEnvVar("VXLAN_IFACE_NAME", "vxlan0")

	// PolykubeVethPairNamePrefix
	env.PolykubeVethPairNamePrefix = getEnvVar("POLYKUBE_VETH_PAIR_NAME_PREFIX", "polykube")

	// VtepCIDR
	vtepCIDRStr := getEnvVar("VTEP_CIDR", "10.18.0.0/16")
	_, vtepCIDR, err := net.ParseCIDR(vtepCIDRStr)
	if err != nil {
		log.Error(
			err, "failed to parse env variable",
			"envVar", "VTEP_CIDR",
			"value", vtepCIDRStr,
		)
		return errors.New("failed to parse VTEP_CIDR")
	}
	env.VtepCIDR = vtepCIDR

	// ClusterCIDR
	clusterCIDRStr := os.Getenv("CLUSTER_CIDR")
	_, clusterCIDR, err := net.ParseCIDR(clusterCIDRStr)
	if err != nil {
		log.Error(
			err, "failed to parse env variable",
			"envVar", "CLUSTER_CIDR",
			"value", clusterCIDRStr,
		)
		return errors.New("failed to parse CLUSTER_CIDR")
	}
	env.ClusterCIDR = clusterCIDR

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

	// IntLbrpName
	env.IntLbrpName = getEnvVar("POLYCUBE_INT_LBRP_NAME", "ilb0")

	// RouterName
	env.RouterName = getEnvVar("POLYCUBE_ROUTER_NAME", "r0")

	// elbName
	env.ExtLbrpName = getEnvVar("POLYCUBE_EXT_LBRP_NAME", "elb0")

	// K8sDispName
	env.K8sDispName = getEnvVar("POLYCUBE_K8SDISP_NAME", "k0")

	// CubesLogLevel
	defaultCubesLogLevel := "INFO"
	cubesLogLevel := strings.ToUpper(getEnvVar("POLYCUBE_CUBES_LOG_LEVEL", defaultCubesLogLevel))
	if !utils.IsValidCubeLogLevel(cubesLogLevel) {
		log.Info("invalid cubes log level. Default value applied",
			"POLYCUBE_CUBES_LOG_LEVEL", cubesLogLevel, "default", defaultCubesLogLevel,
		)
		cubesLogLevel = defaultCubesLogLevel
	}
	env.CubesLogLevel = cubesLogLevel

	// CNILogLevel
	defaultCNILogLevel := "info"
	CNILogLevel := strings.ToLower(getEnvVar("CNI_LOG_LEVEL", defaultCNILogLevel))
	if !utils.IsValidCNILogLevel(CNILogLevel) {
		log.Info("invalid CNI log level. Default value applied",
			"CNI_LOG_LEVEL", CNILogLevel, "default", defaultCNILogLevel,
		)
		CNILogLevel = defaultCNILogLevel
	}
	env.CNILogLevel = CNILogLevel

	// IsCPNodesDeployAllowed
	isCPNodesDeployAllowedStr := getEnvVar("IS_CP_NODES_DEPLOY_ALLOWED", "false")
	isCPNodesDeployAllowed, err := strconv.ParseBool(isCPNodesDeployAllowedStr)
	if err != nil {
		log.Error(
			err,
			"failed to parse env variable",
			"envVar", "IS_CP_NODES_DEPLOY_ALLOWED",
			"value", isCPNodesDeployAllowedStr,
		)
		return errors.New("failed to parse IS_CP_NODES_DEPLOY_ALLOWED")
	}
	env.IsCPNodesDeployAllowed = isCPNodesDeployAllowed

	Env = env
	log.Info("loaded environment configuration")
	return nil
}

// EnsureCNIConf creates the configuration file for the CNI plugin
func EnsureCNIConf() error {
	fPath := Env.CNIConfFilePath
	f, err := os.Create(fPath)
	if err != nil {
		log.Error(err, "failed to create CNI configuration file", "path", fPath)
		return errors.New("failed to create CNI configuration file")
	}
	defer f.Close()

	podGwIP := Conf.PodGwInfo.IPNet.IP
	podGwMAC := Conf.PodGwInfo.MAC

	if _, err := fmt.Fprintf(f,
		CNIConfTemplate,
		Env.MTU,
		Env.IntLbrpName,
		Env.CNILogLevel,
		podGwIP.String(),
		podGwMAC.String(),
		Conf.PodCIDR.String(),
		ip.NextIP(Conf.PolykubeVeth.Net.IPNet.IP).String(), // .4
		ip.PrevIP(podGwIP).String(),                        // .253
		podGwIP.String(),
	); err != nil {
		log.Error(err, "failed to write to CNI configuration file", "path", fPath)
		return errors.New("failed to write to CNI configuration file")
	}
	log.Info("ensured CNI configuration file")
	return nil
}
