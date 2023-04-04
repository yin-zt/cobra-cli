package core

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"log"
	"net"
	"os"
	"time"
)

func (this *Cli) Trace(host string) {
	// Tracing an IP packet route to www.baidu.com.
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	var dst net.IPAddr
	for _, ip := range ips {
		if ip.To4() != nil {
			dst.IP = ip
			fmt.Printf("using %v for tracing an IP packet route to %s\n", dst.IP, host)
			break
		}
	}
	if dst.IP == nil {
		log.Fatal("no A record found")
	}

	//  Create a raw socket with IPPROTO_ICMP for sending ICMP echo requests.
	c, err := net.ListenPacket("ip4:1", "0.0.0.0") // ICMP for IPv4
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	// Create an ICMP message buffer.
	p := ipv4.NewPacketConn(c)
	if err := p.SetControlMessage(ipv4.FlagTTL|ipv4.FlagSrc|ipv4.FlagDst|ipv4.FlagInterface, true); err != nil {
		log.Fatal(err)
	}

	// 构建ICMP报文，code为0，type为8，代表echo请求
	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}

	rb := make([]byte, 1500)
	for i := 1; i <= 64; i++ { // up to 64 hops
		wm.Body.(*icmp.Echo).Seq = i
		wb, err := wm.Marshal(nil)
		if err != nil {
			log.Fatal(err)
		}
		if err := p.SetTTL(i); err != nil {
			log.Fatal(err)
		}

		// In the real world usually there are several
		// multiple traffic-engineered paths for each hop.
		// You may need to probe a few times to each hop.
		begin := time.Now()
		if _, err := p.WriteTo(wb, nil, &dst); err != nil {
			log.Fatal(err)
		}
		if err := p.SetReadDeadline(time.Now().Add(3 * time.Second)); err != nil {
			log.Fatal(err)
		}
		n, cm, peer, err := p.ReadFrom(rb)
		if err != nil {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				fmt.Printf("%v\t*\n", i)
				continue
			}
			log.Fatal(err)
		}
		rm, err := icmp.ParseMessage(1, rb[:n])
		if err != nil {
			log.Fatal(err)
		}
		rtt := time.Since(begin)

		// In the real world you need to determine whether the
		// received message is yours using ControlMessage.Src,
		// ControlMessage.Dst, icmp.Echo.ID and icmp.Echo.Seq.
		switch rm.Type {
		case ipv4.ICMPTypeTimeExceeded:
			fmt.Printf("%d\t%v  %v\t%+v\n", i, peer, rtt, cm)
		case ipv4.ICMPTypeEchoReply:
			fmt.Printf("%d\t%v  %v\t%+v\n", i, peer, rtt, cm)
			return
		default:
			// fmt.Printf("%d\t%v  %v\t%+v\n", i, peer, rtt, cm)
		}
	}
}
