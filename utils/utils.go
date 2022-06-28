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

package utils

var cubeLogLevels = [...]string{"TRACE", "DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL", "OFF"}
var CNILogLevels = [...]string{"trace", "debug", "info", "warn", "warning", "fatal", "panic", "off"}

func CreatePeer(serviceName, servicePort string) string {
	return serviceName + ":" + servicePort
}

func Truncate(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

func CreateAttachment(ifaceName, containerID string) string {
	containerID = Truncate(containerID, 10)
	return Truncate(ifaceName+"_"+containerID, 15)
}

func GetHostIfaceName(contIface string, ipHexStr string) string {
	hostIface := contIface
	hostIfaceLen := len(hostIface)
	actualFinalLen := hostIfaceLen + len(ipHexStr) + 1
	if actualFinalLen > 15 {
		toTruncate := actualFinalLen - 15
		hostIface = hostIface[:hostIfaceLen-toTruncate]
	}
	return hostIface + "_" + ipHexStr
}

func IsValidCubeLogLevel(logLevel string) bool {
	for _, l := range cubeLogLevels {
		if logLevel == l {
			return true
		}
	}
	return false
}

func IsValidCNILogLevel(logLevel string) bool {
	for _, l := range CNILogLevels {
		if logLevel == l {
			return true
		}
	}
	return false
}
