package utils

import (
	"net"
	"strings"
)


var (
	localIPPrefix = [...]string{"192.168", "10.0", "169.254", "172.16"}
)

func IsLocalIP(ip string) bool {
	for i := 0; i < len(localIPPrefix); i++ {
		if strings.HasPrefix(ip, localIPPrefix[i]) {
			return true
		}
	}
	return false
}

func GetInterfaceIP() string {
	addrs, _ := net.InterfaceAddrs()

	// get internet ip first
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			if !IsLocalIP(ipnet.IP.String()) {
				return ipnet.IP.String()
			}
		}
	}

	// get internet ip
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}

	return ""
}
