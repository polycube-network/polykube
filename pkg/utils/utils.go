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
