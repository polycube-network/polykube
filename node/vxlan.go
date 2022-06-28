package node

import (
	"errors"
	"fmt"
	"github.com/containernetworking/plugins/pkg/ip"
	"github.com/ekoops/polykube-operator/types"
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
	nodeVtepIP := Env.VtepCIDR.IP
	for i := 0; i < k; i++ {
		nodeVtepIP = ip.NextIP(nodeVtepIP)
	}
	nodeVtepIPNet := &net.IPNet{
		IP:   nodeVtepIP,
		Mask: Env.VtepCIDR.Mask,
	}
	l.V(1).Info(
		"calculated Vtep for cluster node",
		"vtep", fmt.Sprintf("%+v", nodeVtepIPNet),
	)
	return nodeVtepIPNet, nil
}

// createVxlanIface creates a vxlan interface on the node associating it with the node external interface
func createVxlanIface() (netlink.Link, error) {
	name := Env.VxlanIfaceName

	l := log.WithValues("interface", name)
	// defining the vxlan interface properties
	link := &netlink.Vxlan{
		LinkAttrs:    netlink.LinkAttrs{Name: name},
		VxlanId:      42,                               // TODO mocked
		VtepDevIndex: Conf.ExtIface.Link.Attrs().Index, // TODO this is the index of the associated link?
		Port:         4789,
	}

	// creating the vxlan interface
	if err := netlink.LinkAdd(link); err != nil {
		l.Error(err, "failed to create the cluster node vxlan interface")
		return nil, errors.New("failed to create the cluster node vxlan interface")
	}

	// retrieving the vxlan interface
	// TODO is it really necessary?
	vxlan, err := netlink.LinkByName(name)
	if err != nil {
		l.Error(err, "failed to retrieve the cluster node vxlan interface after its creation")
		return nil, errors.New("failed to retrieve the cluster node vxlan interface after its creation")
	}

	l.Info("created cluster node vxlan interface")
	return vxlan, nil
}

// EnsureVxlanIface ensures that the vxlan interface is available and properly configured
func EnsureVxlanIface() error {
	name := Env.VxlanIfaceName

	l := log.WithValues("interface", name)

	// checking vxlan interface existence
	vxlan, err := netlink.LinkByName(name)
	if err != nil {
		if _, notFound := err.(netlink.LinkNotFoundError); notFound {
			// creating the vxlan interface
			l.Info("no cluster node vxlan interface: creating it...")
			vxlan, err = createVxlanIface()
			if err != nil {
				return err
			}
		} else {
			l.Error(err, "failed vxlan interface lookup")
			return errors.New("failed vxlan interface lookup")
		}
	}

	// checking that vxlan interface is up
	if vxlan.Attrs().Flags&net.FlagUp == 0 {
		l.Info("cluster node vxlan interface is not up: setting it up...")
		// setting up the vxlan interface
		if err := netlink.LinkSetUp(vxlan); err != nil {
			l.Error(err, "failed to set the cluster node vxlan interface up")
			return errors.New("failed to set the cluster node vxlan interface up")
		}
		l.Info("set cluster node vxlan interface up")
	}

	// checking IPv4 address configuration on vxlan interface
	addrs, err := netlink.AddrList(vxlan, netlink.FAMILY_V4)
	if err != nil {
		l.Error(err, "failed to retrieve the cluster node vxlan interface IPv4 address list")
		return errors.New("failed to retrieve the cluster node vxlan interface IPv4 address list")
	}

	vtepIPNet := Conf.NodeVtepIPNet
	vtepIPNetStr := vtepIPNet.String()
	addrFound := false

	for _, addr := range addrs {
		if addr.IPNet.String() == vtepIPNetStr {
			addrFound = true
			break
		}
	}
	if !addrFound {
		l.Info("IPv4 address not set on cluster node vxlan interface: setting it...")
		addr := &netlink.Addr{
			IPNet: vtepIPNet,
			Label: "",
		}
		if err := netlink.AddrAdd(vxlan, addr); err != nil {
			l.Error(err, "failed to set the IPv4 address on the cluster node vxlan interface")
			return errors.New("failed to set the IPv4 address on the cluster node vxlan interface")
		}
		l.Info("set IPv4 address on the cluster node vxlan interface")
	}

	Conf.VxlanIface = &types.Iface{
		IPNet: vtepIPNet,
		Link:  vxlan,
	}

	l.Info("ensured cluster node vxlan interface")
	return nil
}

func DeleteVxlanIface() error {
	name := Env.VxlanIfaceName
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
	vxlanIfName := Env.VxlanIfaceName

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
	l.V(1).Info("no fdb entry related to the vxlan interface for the node IP")
	return false, nil
}

// CreateFdbEntry creates a node fdb entry in order to allow the communication with the node provided IP
// through the vxlan interface
func CreateFdbEntry(nodeIP net.IP) error {
	vxlanIfName := Env.VxlanIfaceName

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
	vxlanIfName := Env.VxlanIfaceName

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
	vxlanIfName := Env.VxlanIfaceName

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
