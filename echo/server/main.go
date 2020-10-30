package main

import (
	"flag"
	"log"
	"net"
)

var (
	local = flag.String("local", "0.0.0.0:12345", "Local listening address.")
)

func tcpEchoHandler(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Println("read from", conn.RemoteAddr().String(), "error", err)
			break
		}
		log.Println("tcp", conn.RemoteAddr().String())

		reply := buf[:n]
		reply = append(reply[:], []byte(conn.RemoteAddr().String())...)
		_, err = conn.Write(reply[:])
		if err != nil {
			log.Println("send echo message error", err)
			break
		}
	}
}

func udpEchoHandler(conn *net.UDPConn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("udp", addr.String())

		reply := "(" + addr.String() + ")"
		copy(buffer[n:], reply[:])
		_, err = conn.WriteToUDP(buffer[:], addr)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func main() {
	flag.Parse()

	tcpListener, err := net.Listen("tcp", *local)
	if err != nil {
		log.Fatalln(err)
	}
	lAddr, err := net.ResolveUDPAddr("udp", *local)
	if err != nil {
		log.Fatalln(err)
	}
	udpListener, err := net.ListenUDP("udp", lAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer tcpListener.Close()
	defer udpListener.Close()

	log.Println("echo server running on", *local)
	go func() {
		for {
			conn, err := tcpListener.Accept()
			if err != nil {
				log.Println(err)
				continue
			}

			go tcpEchoHandler(conn)
		}
	}()
	go func() {
		go udpEchoHandler(udpListener)
	}()

	select {}
}
