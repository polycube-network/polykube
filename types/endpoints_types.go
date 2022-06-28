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

package types

type Backend struct {
	Ip     string
	Port   int32
	Weight int32
}
type BackendsSet map[Backend]struct{}

func (bs BackendsSet) Contains(b Backend) bool {
	_, ok := bs[b]
	return ok
}
func (bs BackendsSet) Add(b Backend) {
	bs[b] = struct{}{}
}

type ServiceToBackends map[string]BackendsSet

func (stb ServiceToBackends) Add(s string, b Backend) {
	if stb[s] == nil {
		stb[s] = make(BackendsSet)
	}
	stb[s][b] = struct{}{}
}

func (stb ServiceToBackends) GetBackendsSet(s string) BackendsSet {
	return stb[s]
}

type EndpointsDetail struct {
	EndpointsId                string
	ClusterIPServiceToBackends ServiceToBackends
	NodePortServiceToBackends  ServiceToBackends
}
