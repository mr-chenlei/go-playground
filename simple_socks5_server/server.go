package main

import (
	"fmt"
	"io"
	"net"

	socks5 "github.com/armon/go-socks5"
)

func s5Remote(conn net.Conn, srcAddr, dstAddr string) {
	fmt.Printf("socks5 remote %v -> %v\n", srcAddr, dstAddr)
}

func s5Local(conn net.Conn) {
	// socks5 version
	s5Ver := make([]byte, 4)
	_, err := conn.Read(s5Ver)
	if err != nil {
		fmt.Println("error on tcpRelay read", err)
		return
	}
	fmt.Printf("Receive socks5 version from %v\n", conn.RemoteAddr().String())
	fmt.Println(s5Ver[0], s5Ver[1], s5Ver[2], s5Ver[3])
	// socks5 version replay
	s5VerReply := []byte{0x00, 0x00}
	_, err = conn.Write(s5VerReply)
	if err != nil {
		fmt.Println("error on socks5 version reply:", err)
		return
	}
	fmt.Printf("Writes socks5 version reply to %v\n", conn.RemoteAddr().String())
	fmt.Println(s5VerReply)
	// socks5 request
	s5Request := make([]byte, 10)
	_, err = io.ReadFull(conn, s5Request)
	if err != nil {
		fmt.Println("error on read socks5 request:", err)
		return
	}
	fmt.Printf("Receive socsk5 request from %v\n", conn.RemoteAddr().String())
	fmt.Println(s5Request)
	// analysis socks5 request type
	switch s5Request[2] {
	case 0x01: // Connect
		tcpAddr := net.TCPAddr{
			IP:   net.IP{s5Request[4], s5Request[5], s5Request[6], s5Request[7]},
			Port: (int(s5Request[8]) << 8) | int(s5Request[9]),
		}
		fmt.Println("Extract destination address from socks5 request:", tcpAddr.String())
		go s5Remote(conn, conn.RemoteAddr().String(), tcpAddr.String())
	case 0x02: // Bind
	case 0x03: // Associate
	}
}

func doTCPRelay(tcp *net.TCPListener) {
	for {
		conn, err := tcp.Accept()
		if err != nil {
			continue
		}
		go s5Local(conn)
	}
}

func main() {
	tcp, _ := net.ListenTCP(
		"client",
		&net.TCPAddr{
			IP:   net.IP{0, 0, 0, 0},
			Port: 10080,
		})
	s5Config := &socks5.Config{
		BindIP: net.IP{0, 0, 0, 0},
	}
	s5, err := socks5.New(s5Config)
	if err != nil {
		fmt.Println("", err)
	}
	err = s5.Serve(tcp)
	if err != nil {
		fmt.Println("Error on socks5 serve:", err)
	}

	select {}
}
