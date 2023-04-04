package core

import (
	"net"
	"regexp"
	"strings"
)

// GetNetworkIP 作用是获取与"外网"通信的网卡IP
func (this *Cli) GetNetworkIP() string {
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
func (this *Cli) GetAllIps() []string {
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

// GetLocalIP 获取本地IP，先获取本地所有ip信息，并遍历每一个ip，如果ip中有10|172开头的ip直接返回；
// 否则，就返回127.0.0.1
func (this *Cli) GetLocalIP() string {

	ips := this.GetAllIps()
	for _, v := range ips {
		if strings.HasPrefix(v, "10.") || strings.HasPrefix(v, "172.") || strings.HasPrefix(v, "172.") {
			return v
		}
	}
	return "127.0.0.1"

}
