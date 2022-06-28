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

package polycube

type configuration struct {
	cubesLogLevel    string
	intLbrpName      string
	rName            string
	extLbrpName      string
	k8sDispName      string
	ilbToRPortName   string
	rToIlbPortName   string
	rToVxlanPortName string
	rToHostPortName  string
	rToElbPortName   string
	elbToRPortName   string
	elbToKPortName   string
	kToElbPortName   string
	kToIntPortName   string
}
