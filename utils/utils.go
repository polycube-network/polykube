package utils

func CreatePeer(serviceName, servicePort string) string {
	return serviceName + ":" + servicePort
}

// ContainsString checks the presence of a string in a slice of strings
func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
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
