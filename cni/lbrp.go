package main

import (
	"context"
	"fmt"
	current "github.com/containernetworking/cni/pkg/types/100"
	lbrp "github.com/ekoops/polykube-operator/pkg/clients/lbrp"
	simplebridge "github.com/ekoops/polykube-operator/pkg/clients/simplebridge"
	"github.com/ekoops/polykube-operator/pkg/utils"
)

func createLbrp(name string, hostIface *current.Interface) error {
	lbFPort := lbrp.Ports{
		Name:  "to_pod",
		Type_: "frontend",
		Peer:  hostIface.Name,
	}
	lbBPort := lbrp.Ports{
		Name:  "to_bridge",
		Type_: "backend",
	}
	lbrpPorts := []lbrp.Ports{lbFPort, lbBPort}
	lb := lbrp.Lbrp{
		Name:     name,
		Ports:    lbrpPorts,
		Loglevel: "TRACE",
	}
	if resp, err := lbrpAPI.CreateLbrpByID(context.TODO(), name, lb); err != nil {
		return fmt.Errorf("error: %s, response: %+v", err, resp)
	}
	return nil
}

func connectLbrpToBridge(lb string, br string) (*lbrp.Ports, *simplebridge.Ports, error) {
	// Creating port on bridge
	brPortName := "to_" + lb
	brPort := simplebridge.Ports{
		Name: brPortName,
		Peer: utils.CreatePeer(lb, "to_bridge"),
	}
	if resp, err := simplebridgeAPI.CreateSimplebridgePortsByID(context.TODO(), br, brPortName, brPort); err != nil {
		return nil, nil, fmt.Errorf(
			"failed to create %q port on bridge - error: %s, response: %+v", brPortName, err, resp,
		)
	}

	// updating lbrp backend port "to_bridge" in order to set peer=br-name:to_lb-name
	lbPortName := "to_bridge"
	lbPort := lbrp.Ports{
		Peer: utils.CreatePeer(br, brPortName),
	}
	if resp, err := lbrpAPI.UpdateLbrpPortsByID(context.TODO(), lb, lbPortName, lbPort); err != nil {
		return nil, nil, fmt.Errorf("failed to update %q port on lbrp - error: %s, response: %+v", lbPortName, err, resp)
	}
	return &lbPort, &brPort, nil
}

func checkLbrp(name, fpeer, bpeer string) error {
	lb, resp, err := lbrpAPI.ReadLbrpByID(context.TODO(), name)
	// checking if status code != 200 because the api are broken
	if err != nil && resp.StatusCode != 200 {
		return fmt.Errorf("failed to retrieve lbrp - error: %s, response: %+v", err, resp)
	}

	if len(lb.Ports) != 2 {
		return fmt.Errorf("wrong port number - required: 2, found: %d", len(lb.Ports))
	}
	for _, port := range lb.Ports {
		if port.Type_ == "frontend" {
			if port.Name != "to_pod" {
				return fmt.Errorf("wrong FRONTEND port name - required: to_pod, found: %q", port.Name)
			}
			if port.Peer != fpeer {
				return fmt.Errorf("wrong FRONTEND port peer - required: %q, found: %q", fpeer, port.Peer)
			}
		} else { // BACKEND port
			if port.Name != "to_bridge" {
				return fmt.Errorf("wrong BACKEND port name - required: to_bridge, found: %q", port.Name)
			}
			if port.Peer != bpeer {
				return fmt.Errorf("wrong BACKEND port peer - required: %q, found: %q", bpeer, port.Peer)
			}
		}
		if port.Status != "UP" {
			return fmt.Errorf("wrong %q port status - required: UP, found: DOWN", port.Name)
		}
	}
	return nil
}
