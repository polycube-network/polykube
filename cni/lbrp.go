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
