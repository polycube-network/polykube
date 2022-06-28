/*
Copyright 2022 The Polykube Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package node

import (
	"context"
	"errors"
	"fmt"
	"github.com/containernetworking/plugins/pkg/ip"
	"github.com/polycube-network/polykube/types"
	"github.com/vishvananda/netlink"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"net"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	// logger used throughout the package
	log  = ctrl.Log.WithName("node-pkg")
	Conf *Configuration
	Env  *Environment
)

// getNode returns a node object describing the cluster node corresponding to the provided name.
// The request is performed using directly the provided cset (without using caching mechanism from
// the controller-runtime library
func getNode(cset *kubernetes.Clientset, name string) (*v1.Node, error) {
	log := log.WithValues("node", name)
	n, err := cset.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Error(err, "failed to retrieve cluster node info")
		return nil, fmt.Errorf("failed to retrieve cluster node info")
	}
	log.V(1).Info("cluster node info retrieved")
	return n, nil
}

// GetExtIface returns info about the external interface of the provided node. The external interface is
// found using the first parsable IPv4 address of type NodeInternalIP inside the provided node object.
func GetExtIface(n *v1.Node) (*types.Iface, error) {
	log := log.WithValues("node", n.Name)
	// extracting ip of the node external interface
	extIfaceIP := GetIP(n)
	if extIfaceIP == nil {
		log.Error(
			errors.New("no NodeInternalIP found inside node object"),
			"failed to extract the IP of the cluster node external interface",
		)
		return nil, errors.New("failed to extract the IP of the cluster node external interface")
	}

	// retrieving the list of all the interfaces in the node
	links, err := netlink.LinkList()
	if err != nil {
		log.Error(err, "failed to retrieve the list of all the cluster node interfaces")
		return nil, errors.New("failed to retrieve the list of all the cluster node interfaces")
	}

	// searching for the interface whose ip list contains the external interface ip
	for _, link := range links {
		linkName := link.Attrs().Name
		log := log.WithValues("interface", linkName)
		addrs, err := netlink.AddrList(link, netlink.FAMILY_V4)
		if err != nil {
			log.Error(err, "failed to retrieve the list of the node interface addresses")
			return nil, errors.New("failed to retrieve the list of the node interface addresses")
		}
		// scanning the list of addresses of the current interface in order to determine if the list contains
		// the external interface one
		for _, addr := range addrs {
			if addr.IP.Equal(extIfaceIP) {
				// interface found
				extIface := &types.Iface{
					IPNet: addr.IPNet,
					Link:  link,
				}
				log.V(1).Info("obtained cluster node external interface info", "IP", extIface.IPNet)
				return extIface, nil
			}
		}
	}
	log.Error(
		errors.New("no interface for the retrieved interface external IP"),
		"failed to retrieve cluster node external interface info", "IP", extIfaceIP)
	return nil, errors.New("failed to retrieve cluster node external interface info")
}

// GetDefaultGatewayIPNet returns the IP address and prefix length of the default gateway for the cluster node
// external interface.
func GetDefaultGatewayIPNet(extIface *types.Iface) (*net.IPNet, error) {
	extIfaceName := extIface.Link.Attrs().Name
	log := log.WithValues("interface", extIfaceName)

	// retrieving the default route by performing a query to an (hopefully) external IP address
	// TODO temporary solution
	routes, err := netlink.RouteGet(net.IPv4(1, 0, 0, 0))
	if err != nil {
		log.Error(err, "failed to retrieve the default route through the cluster node external interface")
		return nil, errors.New("failed to retrieve the default route through the cluster node external interface")
	}
	if len(routes) != 1 {
		log.Error(
			errors.New("multiple default route"),
			"failed to determine a single default route for the cluster node",
		)
		return nil, errors.New("failed to determine a single default route for the cluster node")
	}
	route := routes[0]

	// checking that the route link index is equal to the external interface index
	routeLI := route.LinkIndex
	extIfaceLI := extIface.Link.Attrs().Index
	if routeLI != extIfaceLI {
		log.Error(
			errors.New("the route link index doesn't match the external interface link index"),
			"link index mismatch", "routeLinkIndex", routeLI, "extIfaceLinkIndex", extIfaceLI,
		)
		return nil, errors.New("the route link index doesn't match the external interface link index")
	}

	gwIPNet := &net.IPNet{
		IP:   route.Gw,
		Mask: extIface.IPNet.Mask, // using the same prefix length of the external interface IP address
	}

	log.V(1).Info(
		"retrieved the IP address and prefix length of the default gateway for the cluster node external interface",
		"IP", fmt.Sprintf("%+v", gwIPNet),
	)
	return gwIPNet, nil
}

// GetDefaultGatewayMAC returns the MAC of the default gateway for the cluster node external interface.
func GetDefaultGatewayMAC(extIface *types.Iface, gwIP net.IP) (net.HardwareAddr, error) {
	extIfaceName := extIface.Link.Attrs().Name
	extIfaceLI := extIface.Link.Attrs().Index
	log := log.WithValues("interface", extIfaceName)

	// retrieving the neighbor list of the external interface
	neighs, err := netlink.NeighList(extIfaceLI, netlink.FAMILY_V4)
	if err != nil {
		log.Error(err, "failed to retrieve the external interface neighbors list")
		return nil, errors.New("failed to retrieve the external interface neighbors list")
	}

	// searching for a neighbor whose IP address is the default gateway for the external interface
	for _, neigh := range neighs {
		if neigh.IP.Equal(gwIP) {
			gwMAC := neigh.HardwareAddr
			log.V(1).Info(
				"retrieved the MAC of the default gateway for the cluster node external interface",
				"MAC", gwMAC,
			)
			return gwMAC, nil
		}
	}
	log.Error(
		errors.New("no ARP entry for default gateway"),
		"failed to retrieve the MAC of the default gateway for the cluster node external interface",
	)
	return nil, errors.New("failed to retrieve the MAC of the default gateway for the cluster node external interface")
}

// GetIP extracts the first parsable IPv4 address of type NodeInternalIP inside the provided node object.
func GetIP(n *v1.Node) net.IP {
	log := log.WithValues("node", n.Name)
	for _, addr := range n.Status.Addresses {
		if addr.Type == v1.NodeInternalIP {
			nodeIP := net.ParseIP(addr.Address)
			if nodeIP == nil {
				continue
			}
			nodeIP = nodeIP.To4()
			if nodeIP == nil {
				continue
			}
			log.V(1).Info("obtained node IPv4 address", "IP", nodeIP)
			return nodeIP
		}
	}
	log.Error(errors.New("not found"), "failed to obtain node IPv4 address")
	return nil
}

// IsReady returns true if the provided node is ready; false otherwise.
func IsReady(n *v1.Node) bool {
	log := log.WithValues("node", n.Name)
	for _, c := range n.Status.Conditions {
		if c.Type == v1.NodeReady && c.Status == v1.ConditionTrue {
			log.V(1).Info("the node is ready")
			return true
		}
	}
	log.V(1).Info("the node is not ready")
	return false
}

// LoadConfig load the node configuration in order to initialize the internal package configuration. Information like
// the cluster node name, the PodCIDR assigned to the node, the already configured addresses and so on are inferred.
func LoadConfig() error {
	// creating the in-cluster config
	clusterConfig, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	// changing the config host in such a way that the clientset (the one created starting from this config) can talk
	// directly with the API server without using the "kubernetes" ClusterIP Service
	clusterConfig.Host = Env.APIServerEndpoint()

	// creating the clientset
	cset, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		return err
	}

	name := Env.NodeName

	node, err := getNode(cset, name)
	if err != nil {
		return err
	}

	podCIDR, err := ParsePodCIDR(node)
	if err != nil {
		return err
	}

	podGwIPNet, err := CalcPodsDefaultGatewayIPNet(podCIDR)
	if err != nil {
		return err
	}
	// the following GwInfo struct has to be completed by adding the Mac address:
	// once the polycube infrastructure is ready, the polycube router port MAC has to be assigned to the Mac field
	podGwInfo := &types.GwInfo{
		IPNet: podGwIPNet,
	}

	// The first /30 subnet of the Pod CIDR is used in order to extract 3 IPv4 address: the vPod IP, the polykube
	// veth pair host end IP and the polykube veth pair net end IP.

	allOnesMask := net.IPv4Mask(255, 255, 255, 255)

	// calculating the IP address and prefix length that will be used for NATting external traffic directed to
	// NodePort services with externalTrafficPolicy=Cluster
	vPodIPNet := &net.IPNet{
		IP:   ip.NextIP(podCIDR.IP), // .1
		Mask: allOnesMask,           // /32
	}
	// calculating the IPv4 addresses and prefix length that will be setup on the node polykube veth pair.
	// The chosen convention is to use the first available IPv4 address after the vPod address for the host interface
	// and the second available one for the network interface (the one that will be connected to the polycube router).
	// A /32 IPv4 address will be configured on the host interface.
	// The evaluated net interface address will be used, with a /32 prefix, in order to define the host interface peer.
	// The net interface address will be also configured on the polycube router port to which the net interface
	// will be attached: in this context however, the address will have a /30 prefix length
	polykubeVethHostIPNet := &net.IPNet{
		IP:   ip.NextIP(vPodIPNet.IP), // .2
		Mask: allOnesMask,             // /32
	}
	polykubeVethNetIPNet := &net.IPNet{
		IP:   ip.NextIP(polykubeVethHostIPNet.IP), // .3
		Mask: allOnesMask,                         // /32
	}

	// calculating the IP address and prefix length that will be setup on the node vxlan interface
	vtepIPNet, err := CalcVtepIPNet(podCIDR)
	if err != nil {
		return err
	}

	// getting info about the cluster node external interface
	extIface, err := GetExtIface(node)
	if err != nil {
		return err
	}

	// retrieving info about the default gateway for the node external interface
	nodeGwIPNet, err := GetDefaultGatewayIPNet(extIface)
	if err != nil {
		return err
	}
	nodeGwMAC, err := GetDefaultGatewayMAC(extIface, nodeGwIPNet.IP)
	if err != nil {
		return err
	}
	nodeGwInfo := &types.GwInfo{
		IPNet: nodeGwIPNet,
		MAC:   nodeGwMAC,
	}

	Conf = &Configuration{
		clientset:             cset,
		Node:                  node,
		PodCIDR:               podCIDR,
		PodGwInfo:             podGwInfo,
		VPodIPNet:             vPodIPNet,
		vtepIPNet:             vtepIPNet,
		ExtIface:              extIface,
		polykubeVethHostIPNet: polykubeVethHostIPNet,
		polykubeVethNetIPNet:  polykubeVethNetIPNet,
		NodeGwInfo:            nodeGwInfo,
	}

	log.Info("loaded node configuration", "node", name)
	return nil
}
