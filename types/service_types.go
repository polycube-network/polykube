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

type Frontend struct {
	Vip   string
	Vport int32
	Proto string
}

type FrontendsSet map[Frontend]struct{}

func (fs FrontendsSet) Contains(f Frontend) bool {
	_, ok := fs[f]
	return ok
}
func (fs FrontendsSet) Add(f Frontend) {
	fs[f] = struct{}{}
}

type ServiceDetail struct {
	ServiceId             string
	NodePortFrontendsSet  FrontendsSet
	ClusterIPFrontendsSet FrontendsSet
	ExternalTrafficPolicy string
	InternalTrafficPolicy string
}
