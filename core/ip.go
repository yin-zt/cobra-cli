package core

import (
	"net"
	"regexp"
	"strings"
)

// GetNetworkIP 作用是获取与"外网"通信的网卡IP
func (this *Common) GetNetworkIP() string {
	var (
		err  error
		conn net.Conn
	)
	if conn, err = net.Dial("udp", "8.8.8.8:80"); err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]

}

// GetAllIps 获取主机的所有IP信息
func (this *Common) GetAllIps() []string {
	var response any
	ips := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		response = err
		panic(response)
	}
	for _, addr := range addrs {
		ip := addr.String()
		pos := strings.Index(ip, "/")
		if match, _ := regexp.MatchString("(\\d+\\.){3}\\d+", ip); match {
			if pos != -1 {
				ips = append(ips, ip[0:pos])
			}
		}
	}
	return ips
}
