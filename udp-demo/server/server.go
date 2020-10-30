package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var (
	localAddr   = flag.String("local", "", "Local listen address")
	forwardAddr = flag.String("forward", "", "Remote server (external) IP like 8.8.8.8")
)

func handleUDPConnection(conn *net.UDPConn) {

	// here is where you want to do stuff like read or write to tun-client

	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)
	fmt.Println(time.Now(), "Received from UDP tun-client:", addr, "Message:", string(buffer[:n]))

	if err != nil {
		log.Fatal(err)
	}

	// NOTE : Need to specify tun-client address in WriteToUDP() function
	// otherwise, you will get this error message
	// write udp : write: destination address required if you use Write() function instead of WriteToUDP()

	// write message back to tun-client
	reply := "(" + addr.String() + ")"
	copy(buffer[n:], reply[:])
	_, err = conn.WriteToUDP(buffer[:], addr)

	if err != nil {
		log.Println(err)
	}
}

func main() {
	flag.Parse()

	udpAddr, err := net.ResolveUDPAddr("udp4", *localAddr)
	if err != nil {
		log.Fatal(err)
	}
	// setup listener for incoming UDP connection
	ln, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Println("UDP server up and listening on", *localAddr)
	for {
		// wait for UDP tun-client to connect
		handleUDPConnection(ln)
	}
}
