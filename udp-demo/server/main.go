package main

import (
	"flag"
	"log"
	"net"
)

var (
	local = flag.String("local", "0.0.0.0:12345", "Local listen address")
)

var (
	udpConnections = make(map[string]interface{}, 0)
)

func tcpWorker(local string) {
	lddr, err := net.ResolveTCPAddr("tcp", local)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", lddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Println("tcp worker:", lddr.String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}
		log.Println("accept new connection from", conn.RemoteAddr())

		go func() {
			conn.Write([]byte("hello world!"))
			conn.Close()
		}()
	}
}

func main() {
	flag.Parse()

	go tcpWorker(*local)

	select {}
}
