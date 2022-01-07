package node

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/ekoops/polykube-operator/types"
	"github.com/vishvananda/netlink"
	"math"
	"net"
	"os/exec"
	"strings"
)

// CalcVtepIPNet calculates the vxlan tunnel endpoint IP and prefix length starting from the pod CIDR.
// Starting from the clusterCIDR, only the part that identify the pod CIDR is isolated and is placed at the end of
// the vtep CIDR to create the new address. Example:
// given clusterIP = 192.178.0.0/16, podCIDR = 192.178.1.0/24 and vtepCIDR = 10.18.0.0/16
// the .1. part of the pod CIDR is extracted and placed at the end of the vtepCIDR tp create the new address:
// 10.18.0.1/24
func CalcVtepIPNet(podCIDR *net.IPNet) (*net.IPNet, error) {
	log := log.WithValues("podCIDR", podCIDR)

	podCIDROnes, _ := podCIDR.Mask.Size()             // 24
	clusterCIDROnes, _ := Env.ClusterCIDR.Mask.Size() // 16
	vtepCIDROnes, _ := Env.VtepCIDR.Mask.Size()       // 16
	lenDiff := podCIDROnes - clusterCIDROnes          // 24 - 16 = 8
	if lenDiff > 32-vtepCIDROnes {                    // 8 > 32 - 16 ? if yes, not enough capacity in vtep CIDR
		log.Error(
			errors.New("the vtep CIDR capacity is not enough"),
			"failed to evaluate the vtep IP for the cluster node",
			"minCapacity", math.Pow(2, float64(lenDiff)),
		)
		return nil, errors.New("failed to evaluate the vtep IP for the cluster node")
	}
	n := binary.BigEndian.Uint32(podCIDR.IP)
	n >>= 32 - podCIDROnes
	n &= (1 << lenDiff) - 1
	addr := make(net.IP, 4)
	binary.BigEndian.PutUint32(addr, binary.BigEndian.Uint32(Env.VtepCIDR.IP)|n)

	vtepIPNet := &net.IPNet{
		IP:   addr,
		Mask: Env.VtepCIDR.Mask,
	}
	log.V(1).Info("calculated vtep IP for cluster node", "IP", vtepIPNet)
	return vtepIPNet, nil
}

// createVxlanIface creates a vxlan interface on the node associating it with the node external interface
func createVxlanIface() (netlink.Link, error) {
	name := Env.VxlanIfaceName
	log := log.WithValues("interface", name)

	// defining the vxlan interface properties
	link := &netlink.Vxlan{
		LinkAttrs:    netlink.LinkAttrs{Name: name},
		VxlanId:      42, // TODO mocked
		VtepDevIndex: Conf.ExtIface.Link.Attrs().Index,
		Port:         4789,
	}

	// creating the vxlan interface
	if err := netlink.LinkAdd(link); err != nil {
		log.Error(err, "failed to create the cluster node vxlan interface")
		return nil, errors.New("failed to create the cluster node vxlan interface")
	}

	// retrieving the vxlan interface
	// TODO is it really necessary?
	vxlan, err := netlink.LinkByName(name)
	if err != nil {
		log.Error(err, "failed to retrieve the cluster node vxlan interface after its creation")
		return nil, errors.New("failed to retrieve the cluster node vxlan interface after its creation")
	}

	log.Info("created cluster node vxlan interface")
	return vxlan, nil
}

// EnsureVxlanIface ensures that the vxlan interface is available and properly configured
func EnsureVxlanIface() error {
	name := Env.VxlanIfaceName
	log := log.WithValues("interface", name)

	// checking vxlan interface existence
	vxlan, err := netlink.LinkByName(name)
	if err != nil {
		if _, notFound := err.(netlink.LinkNotFoundError); notFound {
			// creating the vxlan interface
			log.V(1).Info("no cluster node vxlan interface: creating it...")
			vxlan, err = createVxlanIface()
			if err != nil {
				return err
			}
		} else {
			log.Error(err, "failed vxlan interface lookup")
			return errors.New("failed vxlan interface lookup")
		}
	} else {
		// cleaning up node bridge fdb if the vxlan interface is already present
		if err := cleanupFdb(); err != nil {
			return err
		}
	}

	// checking that vxlan interface is up
	if vxlan.Attrs().Flags&net.FlagUp == 0 {
		log.V(1).Info("cluster node vxlan interface is not up: setting it up...")
		// setting up the vxlan interface
		if err := netlink.LinkSetUp(vxlan); err != nil {
			log.Error(err, "failed to set the cluster node vxlan interface up")
			return errors.New("failed to set the cluster node vxlan interface up")
		}
		log.V(1).Info("set cluster node vxlan interface up")
	}

	// checking IPv4 address configuration on vxlan interface
	addrs, err := netlink.AddrList(vxlan, netlink.FAMILY_V4)
	if err != nil {
		log.Error(err, "failed to retrieve the cluster node vxlan interface IPv4 addresses list")
		return errors.New("failed to retrieve the cluster node vxlan interface IPv4 addresses list")
	}

	vtepIPNet := Conf.VtepIPNet
	vtepIPNetStr := vtepIPNet.String()
	addrFound := false

	for _, addr := range addrs {
		if addr.IPNet.String() == vtepIPNetStr {
			addrFound = true
			break
		}
	}
	if !addrFound {
		log.V(1).Info("IPv4 address not set on cluster node vxlan interface: setting it...")
		addr := &netlink.Addr{
			IPNet: vtepIPNet,
			Label: "",
		}
		if err := netlink.AddrAdd(vxlan, addr); err != nil {
			log.Error(err, "failed to set the IPv4 address on the cluster node vxlan interface")
			return errors.New("failed to set the IPv4 address on the cluster node vxlan interface")
		}
		log.V(1).Info("set IPv4 address on the cluster node vxlan interface")
	}

	Conf.VxlanIface = &types.Iface{
		IPNet: vtepIPNet,
		Link:  vxlan,
	}

	log.Info("ensured cluster node vxlan interface")
	return nil
}

// DeleteVxlanIface deletes the node vxlan interface and cleans up all the node bridge fdb entry related to it.
func DeleteVxlanIface() error {
	name := Env.VxlanIfaceName
	log := log.WithValues("interface", name)

	link, err := netlink.LinkByName(name)
	if err != nil {
		log.Error(err, "failed to retrieve the cluster node vxlan interface")
		return errors.New("failed to retrieve the cluster node vxlan interface")
	}

	// no need to handle the possible returned error since this function is called during operator shutdown
	cleanupFdb()

	if err := netlink.LinkDel(link); err != nil {
		log.Error(err, "failed to delete the cluster node vxlan interface")
		return errors.New("failed to delete the cluster node vxlan interface")
	}
	log.V(1).Info("deleted cluster node vxlan interface")
	return nil
}

// CheckFdbEntryExistence allows to verify the presence of a node bridge fdb entry related to the vxlan interface
// for the provided cluster node IP.
func CheckFdbEntryExistence(nodeIP net.IP) (bool, error) {
	vxlanIfName := Env.VxlanIfaceName
	log := log.WithValues("interface", vxlanIfName, "IP", nodeIP)

	cmd := exec.Command("bash", "-c", fmt.Sprintf("bridge fdb show dev %s", vxlanIfName))
	stdout, err := cmd.Output()
	if err != nil {
		errMsg := "failed to retrieve bridge fdb entry related to the vxlan interface for the node IP"
		log.Error(err, errMsg)
		return false, errors.New(errMsg)
	}
	for _, e := range strings.Split(string(stdout), "\n") {
		if strings.Contains(e, "00:00:00:00:00:00") && strings.Contains(e, nodeIP.String()) {
			log.V(1).Info("found a node bridge fdb entry related to the vxlan interface for the node IP")
			return true, nil
		}
	}
	log.V(1).Info("bridge fdb entry related to the vxlan interface for the node IP not found")
	return false, nil
}

// CreateFdbEntry creates a node bridge fdb entry in order to allow the communication with the node provided IP
// through the vxlan interface.
func CreateFdbEntry(nodeIP net.IP) error {
	vxlanIfName := Env.VxlanIfaceName
	log := log.WithValues("interface", vxlanIfName, "IP", nodeIP)

	cmd := exec.Command(
		"bash", "-c",
		fmt.Sprintf("bridge fdb append to 00:00:00:00:00:00 dst %s dev %s", nodeIP, vxlanIfName),
	)
	if _, err := cmd.Output(); err != nil {
		errMsg := "failed to create a node bridge fdb entry related to the vxlan interface for provided cluster node IP"
		log.Error(err, errMsg)
		return errors.New(errMsg)
	}
	log.V(1).Info("created a node bridge fdb entry related to the vxlan interface for the provided cluster node IP")
	return nil
}

// DeleteFdbEntry deletes the node bridge fdb entry related to the vxlan interface for the provided cluster node IP.
func DeleteFdbEntry(nodeIP net.IP) error {
	vxlanIfName := Env.VxlanIfaceName
	log := log.WithValues("interface", vxlanIfName, "IP", nodeIP)

	cmd := exec.Command(
		"bash", "-c",
		fmt.Sprintf("bridge fdb delete to 00:00:00:00:00:00 dst %s dev %s", nodeIP, vxlanIfName),
	)

	_, err := cmd.Output()
	if err != nil {
		errMsg := "failed to delete the node bridge fdb entry related to the vxlan interface for the provided cluster node IP"
		log.Error(err, errMsg)
		return errors.New(errMsg)
	}
	log.V(1).Info("deleted the node bridge fdb entry related to the vxlan interface for the provided cluster node IP")
	return nil
}

// cleanupFdb deletes all the node bridge fdb entries related to the vxlan interface
func cleanupFdb() error {
	vxlanIfName := Env.VxlanIfaceName
	log := log.WithValues("interface", vxlanIfName)

	cmd := exec.Command(
		"bash", "-c", fmt.Sprintf("bridge fdb delete to 00:00:00:00:00:00 dev %s", vxlanIfName),
	)

	_, err := cmd.Output()
	if err != nil {
		errMsg := "failed to delete all the node bridge fdb entries related to the vxlan interface"
		log.Error(err, errMsg)
		return errors.New(errMsg)
	}
	log.V(1).Info("deleted all the node bridge fdb entries related to the vxlan interface")
	return nil
}
