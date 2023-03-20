package core

import (
	"net"
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
