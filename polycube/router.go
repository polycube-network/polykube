package polycube

import (
	"context"
	"errors"
	"fmt"
	"github.com/ekoops/polykube-operator/polycube/clients/router"
	"net"
	"net/url"
)

// GetRouterToIntK8sLbrpPortMAC returns the MAC of the polycube router interface representing the default gateway for
// the pods subnet
func GetRouterToIntK8sLbrpPortMAC() (net.HardwareAddr, error) {
	rName := conf.rName
	rToIklPortName := conf.rToIklPortName
	l := log.WithValues("name", rName, "port", rToIklPortName)
	MACStr, resp, err := routerAPI.ReadRouterPortsMacByID(context.TODO(), rName, rToIklPortName)
	if err != nil {
		l.Error(err, "failed to retrieve router port MAC", "response", fmt.Sprintf("%+v", resp))
		return nil, errors.New("failed to retrieve router port MAC")
	}
	MAC, err := net.ParseMAC(MACStr)
	if err != nil {
		l.Error(err, "failed to parse router port MAC")
		return nil, errors.New("failed to parse router port MAC")
	}
	l.V(1).Info("retrieved router to internal k8s lbrp port MAC", "MAC", MAC.String())
	return MAC, nil
}

// CheckRouteExistence verify the existence of the route towards nodePodCIDR through nodeVtepIP on the polycube router
func CheckRouteExistence(nodePodCIDR *net.IPNet, nodeVtepIP net.IP) (bool, error) {
	rName := conf.rName
	network := nodePodCIDR.String()
	nexthop := nodeVtepIP.String()
	l := log.WithValues("router", rName, "network", network, "nexthop", nexthop)

	rrs, resp, err := routerAPI.ReadRouterRouteListByID(context.TODO(), rName)
	if err != nil {
		l.Error(err, "failed to retrieve router route list", "response", fmt.Sprintf("%+v", resp))
		return false, errors.New("failed to retrieve router route list")
	}
	for _, rr := range rrs {
		if rr.Network == network && rr.Nexthop == nexthop {
			l.V(1).Info("verified router route existence")
			return true, nil
		}
	}
	l.V(1).Info("no router route")
	return false, nil

	//l.V(1).Info("verified router route existence")
	//return true, nil
}

// CreateRouteToNodePodCIDR creates a new route on the polycube router in order to enable connectivity toward
// the node Pod CIDR
func CreateRouteToNodePodCIDR(nodePodCIDR *net.IPNet, nodeVtepIP net.IP) error {
	rName := conf.rName
	network := nodePodCIDR.String()
	nexthop := nodeVtepIP.String()
	l := log.WithValues("router", rName, "network", network, "nexthop", nexthop)

	route := router.Route{
		Network:    network,
		Nexthop:    nexthop,
		Interface_: conf.rToVxlanPortName,
	}

	// adding route to router in order to make node pod CIDR reachable throw vxlan interface
	if resp, err := routerAPI.CreateRouterRouteByID(
		context.TODO(), rName, url.QueryEscape(route.Network), route.Nexthop, route,
	); err != nil && resp.StatusCode != 409 {
		l.Error(
			err,
			"failed to create router route for enabling communication with the node Pod CIDR through the vxlan interface",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New(
			"failed to create router route for enabling communication with the node Pod CIDR through the vxlan interface",
		)
	}
	l.V(1).Info("created router route")
	return nil
}

// DeleteRouteToNodePodCIDR deletes the route on the polycube router in order to disable connectivity toward
// the node Pod CIDR
func DeleteRouteToNodePodCIDR(nodePodCIDR *net.IPNet, nodeVtepIP net.IP) error {
	rName := conf.rName
	network := nodePodCIDR.String()
	nexthop := nodeVtepIP.String()
	l := log.WithValues("router", rName, "network", network, "nexthop", nexthop)

	// delete route to router in order to make node pod CIDR unreachable throw vxlan interface
	if resp, err := routerAPI.DeleteRouterRouteByID(
		context.TODO(), rName, url.QueryEscape(network), nexthop,
	); err != nil && resp.StatusCode != 409 {
		l.Error(
			err,
			"failed to delete router route for disabling communication with the node Pod CIDR through the vxlan interface",
			"response", fmt.Sprintf("%+v", resp),
		)
		return errors.New(
			"failed to delete router route for disabling communication with the node Pod CIDR through the vxlan interface",
		)
	}
	l.V(1).Info("deleted router route")
	return nil
}

// DeleteRouterRoute deletes the route on the polycube router
func DeleteRouterRoute(network, nexthop string) error {
	rName := conf.rName
	log := log.WithValues("router", rName, "network", network, "nexthop", nexthop)

	// delete route to router in order to make node pod CIDR unreachable throw vxlan interface
	if resp, err := routerAPI.DeleteRouterRouteByID(
		context.TODO(), rName, url.QueryEscape(network), nexthop,
	); err != nil && resp.StatusCode != 409 {
		log.Error(err, "failed to delete router route", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to delete router route")
	}
	log.V(1).Info("deleted router route")
	return nil
}

func CleanupRouterRoutes() error {
	rName := conf.rName
	log := log.WithValues("router", rName)
	routes, resp, err := routerAPI.ReadRouterRouteListByID(context.TODO(), rName)
	if err != nil {
		log.Error(err, "failed to retrieve router routes", "response", fmt.Sprintf("%+v", resp))
		return errors.New("failed to retrieve router routes")
	}
	for _, route := range routes {
		if route.Nexthop != "local" {
			if err := DeleteRouterRoute(route.Network, route.Nexthop); err != nil {
				return err
			}
		}
	}
	log.V(1).Info("deleted all non-local router route")
	return nil
}
