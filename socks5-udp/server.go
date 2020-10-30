package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func doReply(conn *net.UDPConn, saddr *net.UDPAddr) {
	for {
		reply := "This is server " + conn.LocalAddr().String()
		conn.WriteTo([]byte(reply), saddr)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	lport := flag.Int("port", 9090, "local port")
	flag.Parse()

	fmt.Println("Listening on local port", *lport)
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: *lport,
		IP:   net.ParseIP("0.0.0.0"),
	})
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	message := make([]byte, 4096)
	for {
		_, remote, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v Receive message %s, from %s\n", time.Now(), message, remote.String())
		go doReply(conn, remote)
	}
}
