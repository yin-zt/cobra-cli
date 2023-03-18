package utils

import "net"

func IsIp(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		return false
	}
	return true
}
