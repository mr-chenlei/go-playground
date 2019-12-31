package main

import (
	"net"
	"time"
)

func main() {
	laddr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:8192")
	raddr1, _ := net.ResolveUDPAddr("udp", "192.168.101.76:9001")
	raddr2, _ := net.ResolveUDPAddr("udp", "192.168.101.76:9002")
	conn, _ := net.ListenUDP("udp", laddr)

	for {
		conn.WriteTo([]byte("Hello from tun-client"), raddr1)
		conn.WriteTo([]byte("Hello from tun-client"), raddr2)
		time.Sleep(1 * time.Second)
	}
}
