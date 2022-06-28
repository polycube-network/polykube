package main

import (
	"context"
	"fmt"
)

func checkLbrpPort(name string, portName string, peerName string, contIP string) error {
	port, resp, err := lbrpAPI.ReadLbrpPortsByID(context.TODO(), name, portName)
	// checking if status code != 200 because the api is broken
	if err != nil && resp.StatusCode != 200 {
		return fmt.Errorf("failed to retrieve lbrp port - error: %s, response: %+v", err, resp)
	}

	if port.Type_ != "frontend" {
		return fmt.Errorf("wrong lbrp port type - required: frontend, found: %q", port.Type_)
	}

	if port.Peer != peerName {
		return fmt.Errorf("wrong lbrp port peer - required: %q, found: %q", peerName, port.Peer)
	}
	if port.Status != "UP" {
		return fmt.Errorf("wrong lbrp port status - required: UP, found: DOWN")
	}

	if port.Ip_ != contIP {
		return fmt.Errorf("wrong lbrp port associated IP - required: %q, found: %q", contIP, port.Ip_)
	}

	return nil
}
