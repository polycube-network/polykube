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
	// this assignment is done in order to allow the deferred function to work
	err = errors.New("no IPv4 address found")
	return nil, err
}
