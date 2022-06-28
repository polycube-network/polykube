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
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"net"
)

type NetConf struct {
	types.NetConf
	MTU      int    `json:"mtu"`
	LbrpName string `json:"intLbrp"`
	LogLevel string `json:"logLevel"`
	Gw       GwInfo `json:"gateway"`
}

type GwInfo struct {
	IP     net.IP           `json:"ip"`
	RawMAC string           `json:"mac"`
	MAC    net.HardwareAddr `json:"-"`
}

type IFaceConf struct {
	ResultIndex int
	Interface   *current.Interface
	IPConf      *current.IPConfig
}
