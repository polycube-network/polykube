package node

import (
	"errors"
	"github.com/polycube-network/polykube/types"
	"github.com/vishvananda/netlink"
	"net"
)

// createPolykubeVethPair creates the polykube veth pair on the node. This function returns a slice containing two
// links: the veth pair host end in the first position and the veth pair net end in the second position
func createPolykubeVethPair() ([]netlink.Link, error) {
	prefix := Env.PolykubeVethPairNamePrefix
	hostEndName := prefix + "_host"
	netEndName := prefix + "_net"
	log := log.WithValues("hostInterface", hostEndName, "netInterface", netEndName)

	// defining the polykube veth pair properties
	link := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{Name: hostEndName},
		PeerName:  netEndName,
	}

	// creating the polykube veth pair interface
	if err := netlink.LinkAdd(link); err != nil {
		log.Error(err, "failed to create the cluster node polykube veth pair")
		return nil, errors.New("failed to create the cluster node polykube veth pair")
	}

	// retrieving the polykube veth pair host end
	hostEnd, err := netlink.LinkByName(hostEndName)
	if err != nil {
		log.Error(err, "failed to retrieve the cluster node polykube veth pair host end after its creation")
		return nil, errors.New("failed to retrieve the cluster node polykube veth pair host end after its creation")
	}
	// retrieving the polykube veth pair net end
	netEnd, err := netlink.LinkByName(netEndName)
	if err != nil {
		log.Error(err, "failed to retrieve the cluster node polykube veth pair net end after its creation")
		return nil, errors.New("failed to retrieve the cluster node polykube veth pair net end after its creation")
	}

	log.Info("created cluster node polykube veth pair")
	return []netlink.Link{hostEnd, netEnd}, nil
}

func ensurePolykubeVethPairEnds() (netlink.Link, netlink.Link, error) {
	prefix := Env.PolykubeVethPairNamePrefix
	hostEndName := prefix + "_host"
	netEndName := prefix + "_net"

	var vethEnds []netlink.Link

	for _, endName := range []string{hostEndName, netEndName} {
		log := log.WithValues("interface", endName)
		// checking end existance
		end, err := netlink.LinkByName(endName)
		if err != nil {
			if _, notFound := err.(netlink.LinkNotFoundError); notFound {
				// for now, only logging
				log.V(1).Info("no polykube veth pair end")
			} else {
				log.Error(err, "failed polykube veth pair end lookup")
				return nil, nil, errors.New("failed polykube veth pair end lookup")
			}
		} else {
			// the polykube veth pair end is present: checking if its type is correct
			if endType := end.Type(); endType != "veth" {
				log.Error(
					err, "wrong polykube veth pair end type", "found", endType, "required", "veth",
				)
				return nil, nil, errors.New("wrong polykube veth pair end type")
			}
			// the polykube veth pair end is present and its type is correct: adding it to the ends slice
			log.V(1).Info("found polykube veth pair end")
			vethEnds = append(vethEnds, end)
		}
	}

	vethEndsLen := len(vethEnds)

	// if only one veth ends has been found, abort the operation
	if vethEndsLen == 1 {
		log.Error(
			errors.New("uncompleted veth pair configuration"),
			"found uncompleted polykube veth pair configuration",
		)
		return nil, nil, errors.New("found uncompleted polykube veth pair configuration")
	}

	// if no veth ends are present, creating it
	if vethEndsLen == 0 {
		log.V(1).Info("no cluster node polykube veth pair: creating it...", "prefix", prefix)
		var err error
		// notice that the return value is assigned to the already existing vethEnds slice
		vethEnds, err = createPolykubeVethPair()
		if err != nil {
			return nil, nil, err
		}
	}

	return vethEnds[0], vethEnds[1], nil
}

// ensurePolykubeVethPairIPv4Configuration ensures that the polykube veth pair ends have a proper IPv4 configuration and
// returns a types.PolykubeVeth struct containing the proper information associated to each end
func ensurePolykubeVethPairIPv4Configuration(hostEnd, netEnd netlink.Link) (*types.PolykubeVeth, error) {
	addrs, err := netlink.AddrList(hostEnd, netlink.FAMILY_V4)
	if err != nil {
		log.Error(err, "failed to retrieve the cluster node polykube veth pair host end IPv4 addresses list")
		return nil, errors.New("failed to retrieve the cluster node polykube veth pair host end IPv4 addresses list")
	}

	hostEndIPNet := Conf.polykubeVethHostIPNet
	netEndIPNet := Conf.polykubeVethNetIPNet
	hostIPNetStr := hostEndIPNet.String()
	netEndIPNetStr := netEndIPNet.String()
	addrFound := false
	for _, addr := range addrs {
		if addr.IPNet.String() == hostIPNetStr && addr.Peer.String() == netEndIPNetStr {
			addrFound = true
			break
		}
	}

	// assigning the IPv4 address if it has not been found
	if !addrFound {
		log.V(1).Info("IPv4 address not set on cluster node polykube veth pair host end: setting it...")
		addr := &netlink.Addr{
			IPNet: hostEndIPNet,
			Label: "",
			Peer:  netEndIPNet,
		}

		log := log.WithValues("address", hostIPNetStr, "peer", netEndIPNetStr)

		if err := netlink.AddrAdd(hostEnd, addr); err != nil {
			log.Error(err, "failed to set the IPv4 address on the cluster node polykube veth pair end")
			return nil, errors.New("failed to set the IPv4 address on the cluster node polykube veth pair end")
		}
		log.V(1).Info("set IPv4 address on the cluster node polykube veth pair end")
	}

	return &types.PolykubeVeth{
		Host: &types.Iface{
			IPNet: hostEndIPNet,
			Link:  hostEnd,
		},
		Net: &types.Iface{
			IPNet: netEndIPNet,
			Link:  netEnd,
		},
	}, nil
}

// ensureHostToPodsRoute ensures the presence of the route towards the Cluster CIDR through the polykube veth pair
// in order to allow the communication between the host and pods on the cluster nodes
func ensureHostToPodsRoute() error {
	hostEndLink := Conf.PolykubeVeth.Host.Link
	log := log.WithValues("interface", hostEndLink.Attrs().Name)

	// retrieving the IPv4 routes associated with the polykube veth pair host end
	routes, err := netlink.RouteList(hostEndLink, netlink.FAMILY_V4)
	if err != nil {
		log.Error(err, "failed to set retrieve IPv4 host routes through the polykube veth pair host end")
		return errors.New("failed to set retrieve IPv4 host routes through the polykube veth pair host end")
	}

	// preparing some variables for a later reuse
	clusterCIDRStr := Env.ClusterCIDR.String()
	netEndIP := Conf.PolykubeVeth.Net.IPNet.IP
	hostEndLinkIndex := hostEndLink.Attrs().Index

	routeFound := false

	// checking the presence of a route towards the Cluster CIDR through the polykube host end
	for _, route := range routes {
		if route.Dst.String() == clusterCIDRStr && route.Gw.Equal(netEndIP) && route.LinkIndex == hostEndLinkIndex {
			routeFound = true
			break
		}
	}

	// creating the route if it has not been found
	if !routeFound {
		log := log.WithValues("network", clusterCIDRStr, "nexthop", netEndIP)
		route := &netlink.Route{
			LinkIndex: hostEndLinkIndex,
			Dst:       Env.ClusterCIDR,
			Gw:        netEndIP,
		}
		if err := netlink.RouteAdd(route); err != nil {
			log.Error(err, "failed to create the host route towards the Cluster CIDR through the polykube veth pair")
			return errors.New("failed to create the host route towards the Cluster CIDR through the polykube veth pair")
		}
		log.V(1).Info("created the host route towards the Cluster CIDR through the polykube veth pair")
	}
	return nil
}

// setPolykubeVethPairEndsUp sets up provided polykube veth pair ends
func setPolykubeVethPairEndsUp(hostEnd, netEnd netlink.Link) error {
	for _, end := range []netlink.Link{hostEnd, netEnd} {
		if end.Attrs().Flags&net.FlagUp == 0 {
			log := log.WithValues("interface", end.Attrs().Name)
			log.V(1).Info("cluster node polykube veth pair end is not up: setting it up...")
			// setting up the polykube veth pair end up
			if err := netlink.LinkSetUp(end); err != nil {
				log.Error(err, "failed to set the cluster node polykube veth pair end up")
				return errors.New("failed to set the cluster node polykube veth pair end up")
			}
			log.V(1).Info("set cluster node polykube veth pair end up")
		}
	}
	return nil
}

// EnsurePolykubeVethPair ensures that the polykube veth pair is available and properly configured for enabling
// host agents to pods communication
func EnsurePolykubeVethPair() error {
	// ensuring that the polykube veth pairs ends are present on the node
	hostEnd, netEnd, err := ensurePolykubeVethPairEnds()
	if err != nil {
		return err
	}

	// setting the polykube veth pair ends up
	if err := setPolykubeVethPairEndsUp(hostEnd, netEnd); err != nil {
		return err
	}

	// ensuring that the polykube veth pair ends have the proper IPv4 configuration
	polykubeVeth, err := ensurePolykubeVethPairIPv4Configuration(hostEnd, netEnd)
	if err != nil {
		return err
	}

	// setting the needed internal package information
	Conf.PolykubeVeth = polykubeVeth

	// checking the presence of the route towards the Cluster CIDR
	if err := ensureHostToPodsRoute(); err != nil {
		return err
	}

	log.Info("ensured cluster node polykube veth pair")
	return nil
}

// DeletePolykubeVethPair deletes the node polykube veth pair
func DeletePolykubeVethPair() error {
	hostEndName := Conf.PolykubeVeth.Host.Link.Attrs().Name
	log := log.WithValues("hostInterface", hostEndName)

	link, err := netlink.LinkByName(hostEndName)
	if err != nil {
		log.Error(err, "failed to retrieve the cluster node polykube veth pair")
		return errors.New("failed to retrieve the cluster node polykube veth pair")
	}

	// no need to delete the host route toward the Cluster CIDR as the veth pair deletion will cause also the deletion
	// of the related routes

	if err := netlink.LinkDel(link); err != nil {
		log.Error(err, "failed to delete the cluster node polykube veth pair")
		return errors.New("failed to delete the cluster node polykube veth pair")
	}

	log.V(1).Info("deleted cluster node polykube veth pair")
	return nil
}
