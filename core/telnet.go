package core

import (
	"fmt"
	"net"
	"time"
)

func TelnetCheck(host string, tout int) error {
	timeout := time.Duration(tout) * time.Second
	fmt.Printf("Start port connectivity detection, destination address:%s,timeout time:%v\n", host, timeout.String())
	t1 := time.Now()
	_, err := net.DialTimeout("tcp", host, timeout)
	fmt.Println("Time consumption :", time.Now().Sub(t1))
	if err != nil {
		return err
	}
	fmt.Printf("%v Server connection successful\n", host)
	return nil
}
