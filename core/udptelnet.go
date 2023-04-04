package core

import (
	"net"
	"time"
)

func (this *Cli) UdpTelnetChekck(addr string, intime int) error {
	timeout := time.Duration(intime) * time.Second
	_, err := net.DialTimeout("udp", addr, timeout)
	if err != nil {
		return err
	}
	return nil
}
