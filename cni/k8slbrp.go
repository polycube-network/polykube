package main

import (
	"context"
	"fmt"
)

func checkK8sLbrpPort(name string, portName string, peerName string, contIP string) error {
	port, resp, err := k8sLbrpAPI.ReadK8sLbrpPortsByID(context.TODO(), name, portName)
	// checking if status code != 200 because the api is broken
	if err != nil && resp.StatusCode != 200 {
		return fmt.Errorf("failed to retrieve k8slbrp port - error: %s, response: %+v", err, resp)
	}

	if port.Type_ != "frontend" {
		return fmt.Errorf("wrong k8slbrp port type - required: frontend, found: %q", port.Type_)
	}

	if port.Peer != peerName {
		return fmt.Errorf("wrong k8slbrp port peer - required: %q, found: %q", peerName, port.Peer)
	}
	if port.Status != "UP" {
		return fmt.Errorf("wrong k8slbrp port status - required: UP, found: DOWN")
	}

	if port.Ip_ != contIP {
		return fmt.Errorf("wrong k8slbrp port associated IP - required: %q, found: %q", contIP, port.Ip_)
	}

	return nil
}
