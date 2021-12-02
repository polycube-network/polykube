package node

import (
	"errors"
	"fmt"
	"github.com/containernetworking/plugins/pkg/ip"
	"github.com/ekoops/polykube-operator/pkg/env"
	"github.com/ekoops/polykube-operator/pkg/types"
	"github.com/vishvananda/netlink"
	v1 "k8s.io/api/core/v1"
	"net"
	"os/exec"
	"strconv"
	"strings"
)

// CalcVtepIPNet calculates the ip and the prefix length of the Vxlan Tunnel Endpoint for the provided node.
// The address is extracted from the vtepCIDR range. It is calculated using a convention based on the node name (this
// is a temporary solution)
func CalcVtepIPNet(n *v1.Node) (*net.IPNet, error) {
	l := log.WithValues("node", n.Name)
	// extracting the worker number (this is possible since its worker node is called worker${k})
	// TODO this is a temporary solution
	k, err := strconv.Atoi(strings.TrimPrefix(n.Name, "worker"))
	if err != nil {
		l.Error(err, "failed to extract cluster node number for Vtep evaluation")
		return nil, errors.New("failed to extract cluster node number for Vtep evaluation")
	}
	nodeVtepIP := env.Conf.VtepCIDR.IP
	for i := 0; i < k; i++ {
		nodeVtepIP = ip.NextIP(nodeVtepIP)
	}
	nodeVtepIPNet := &net.IPNet{
		IP:   nodeVtepIP,
		Mask: env.Conf.VtepCIDR.Mask,
	}
	l.V(1).Info(
		"calculated Vtep for cluster node",
		"vtep", fmt.Sprintf("%+v", nodeVtepIPNet),
	)
	return nodeVtepIPNet, nil
}

// CreateVxlanIface creates a vxlan interface on the node associating it with the node external interface
func CreateVxlanIface() error {
	name := env.Conf.VxlanIfName
	vtepIPNet := Conf.NodeVtepIPNet

	l := log.WithValues("interface", name)
	// defining the vxlan interface properties
	vxlan := &netlink.Vxlan{
		LinkAttrs:    netlink.LinkAttrs{Name: name},
		VxlanId:      42,                               // TODO mocked
		VtepDevIndex: Conf.ExtIface.Link.Attrs().Index, // TODO this is the index of the associated link?
		Port:         4789,
	}

	// creating the vxlan interface
	if err := netlink.LinkAdd(vxlan); err != nil {
		l.Error(err, "failed to create the cluster node vxlan interface")
		return errors.New("failed to create the cluster node vxlan interface")
	}

	// retrieving the vxlan interface
	// TODO is it really necessary?
	link, err := netlink.LinkByName(name)
	if err != nil {
		l.Error(err, "failed to retrieve the cluster node vxlan interface after its creation")
		return errors.New("failed to retrieve the cluster node vxlan interface after its creation")
	}

	// setting up the vxlan interface
	if err := netlink.LinkSetUp(link); err != nil {
		l.Error(err, "failed to set the cluster node vxlan interface up")
		return errors.New("failed to set the cluster node vxlan interface up")
	}

	// adding IPv4 address to the interface
	addr := &netlink.Addr{
		IPNet: vtepIPNet,
		Label: "",
	}
	l = l.WithValues("address", fmt.Sprintf("%+v", vtepIPNet))
	if err = netlink.AddrAdd(link, addr); err != nil {
		l.Error(err, "failed to setup the IPv4 address to the cluster node vxlan interface")
		return errors.New("failed to setup the IPv4 address to the cluster node vxlan interface")
	}
	vxlanIface := &types.Iface{
		IPNet: vtepIPNet,
		Link:  link,
	}
	l.Info("cluster node vxlan interface created")

	Conf.VxlanIface = vxlanIface
	return nil
}

func DeleteVxlanIface() error {
	name := env.Conf.VxlanIfName
	l := log.WithValues("interface", name)

	link, err := netlink.LinkByName(name)
	if err != nil {
		l.Error(err, "failed to retrieve the cluster node vxlan interface")
		return errors.New("failed to retrieve the cluster node vxlan interface")
	}

	if err := netlink.LinkDel(link); err != nil {
		l.Error(err, "failed to delete the cluster node vxlan interface")
		return errors.New("failed to delete the cluster node vxlan interface")
	}
	l.V(1).Info("cluster node vxlan interface deleted")
	return nil
}

// CheckFdbEntryExistence allows to verify the presence of a node fdb entry related to the vxlan interface
// for the provided cluster node IP
func CheckFdbEntryExistence(nodeIP net.IP) (bool, error) {
	vxlanIfName := env.Conf.VxlanIfName

	l := log.WithValues("interface", vxlanIfName, "IP", nodeIP)

	cmd := exec.Command(
		"bash", "-c",
		fmt.Sprintf("bridge fdb show dev %s", vxlanIfName),
	)
	stdout, err := cmd.Output()
	if err != nil {
		l.Error(err, "failed to retrieve fdb entry related to the vxlan interface for the node IP")
		return false, errors.New(
			"failed to retrieve fdb entry related to the vxlan interface for the node IP",
		)
	}
	for _, e := range strings.Split(string(stdout), "\n") {
		if strings.Contains(e, "00:00:00:00:00:00") && strings.Contains(e, nodeIP.String()) {
			l.V(1).Info("found a node fdb entry related to the vxlan interface for the node IP")
			return true, nil
		}
	}
	l.Error(
		errors.New("no entry found"),
		"failed to find a node fdb entry related to the vxlan interface for the node IP",
	)
	return false, errors.New(
		"failed to find a node fdb entry related to the vxlan interface for the node IP",
	)
}

// CreateFdbEntry creates a node fdb entry in order to allow the communication with the node provided IP
// through the vxlan interface
func CreateFdbEntry(nodeIP net.IP) error {
	vxlanIfName := env.Conf.VxlanIfName

	l := log.WithValues("interface", vxlanIfName, "IP", nodeIP)

	cmd := exec.Command(
		"bash", "-c",
		fmt.Sprintf("bridge fdb append to 00:00:00:00:00:00 dst %s dev %s", nodeIP, vxlanIfName),
	)
	if _, err := cmd.Output(); err != nil {
		l.Error(
			err,
			"failed to configure the node fdb entry for allowing communication with"+
				"the provided cluster node IP through the vxlan interface",
		)
		return errors.New("failed to configure the node fdb entry for allowing communication with" +
			"the provided cluster node IP through the vxlan interface",
		)
	}
	l.V(1).Info(
		"configured node fdb entry in order to allow communication with the cluster node IP through vxlan interface",
	)
	return nil
}

// DeleteFdbEntry deletes the node fdb entry related to the vxlan interface for the provided cluster node IP
func DeleteFdbEntry(nodeIP net.IP) error {
	vxlanIfName := env.Conf.VxlanIfName

	l := log.WithValues("interface", vxlanIfName, "IP", nodeIP)

	cmd := exec.Command(
		"bash", "-c",
		fmt.Sprintf("bridge fdb delete to 00:00:00:00:00:00 dst %s dev %s", nodeIP, vxlanIfName),
	)

	_, err := cmd.Output()
	if err != nil {
		l.Error(
			err,
			"failed to delete the node fdb entry related to the vxlan interface for the provided cluster node IP",
		)
		return errors.New(
			"failed to delete the node fdb entry related to the vxlan interface for the provided cluster node IP",
		)
	}
	l.V(1).Info("deleted the node fdb entry related to the vxlan interface for the provided cluster node IP")
	return nil
}

// CleanupFdb deletes all the node fdb entries related to the vxlan interface
func CleanupFdb() error {
	vxlanIfName := env.Conf.VxlanIfName

	l := log.WithValues("interface", vxlanIfName)

	cmd := exec.Command(
		"bash", "-c",
		fmt.Sprintf("bridge fdb delete to 00:00:00:00:00:00 dev %s", vxlanIfName),
	)

	_, err := cmd.Output()
	if err != nil {
		l.Error(err, "failed to delete all the node fdb entries related to the vxlan interface")
		return errors.New("failed to delete all the node fdb entries related to the vxlan interface")
	}
	l.V(1).Info("deleted all the node fdb entries related to the vxlan interface")
	return nil
}
