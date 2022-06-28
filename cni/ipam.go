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
	"errors"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/plugins/pkg/ipam"
	"net"
)

func allocIP(ipamType string, stdin []byte) (*net.IPNet, error) {
	// running IPAM plugin and get back the config to apply
	r, err := ipam.ExecAdd(ipamType, stdin)
	if err != nil {
		return nil, err
	}

	// invoking ipam del if err to avoid ip leak
	defer func() {
		if err != nil {
			ipam.ExecDel(ipamType, stdin)
		}
	}()

	// converting the IPAM result into the current Result type
	result, err := current.NewResultFromResult(r)
	if err != nil {
		return nil, err
	}

	if len(result.IPs) == 0 {
		// this assignment is done in order to allow the deferred function to work
		err = errors.New("missing ip config")
		return nil, err
	}

	for _, ipI := range result.IPs {
		// checking if ipI is an IPv4 address
		if address := ipI.Address; address.IP.To4() != nil {
			address.IP = address.IP.To4()
			return &address, nil
		}
	}
	// this assignment is done in order to allow the deferred function to properly work
	err = errors.New("no IPv4 address found")
	return nil, err
}
