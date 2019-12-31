package main

import (
	"encoding/binary"
	"flag"
	"log"
	"net"
	"time"
)

var (
	remote    = flag.String("remote", "", "Echo server address.")
	frequency = flag.Int("frequency", 1000, "Echo message frequency(ms).")
)

func echo(conn *net.TCPConn, elapse int) {
	defer conn.Close()

	t := time.Duration(elapse)
	for range time.Tick(t * time.Millisecond) {
		buf := make([]byte, 1024)
		binary.BigEndian.PutUint64(buf[:8], uint64(time.Now().Unix()))
		_, err := conn.Write(buf[:])
		if err != nil {
			log.Println(err)
			break
		}
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Println(err)
			break
		}
		msg := buf[8 : n-1]
		log.Println("from", conn.RemoteAddr().String(), "->", string(msg))
	}
}

func main() {
	flag.Parse()

	if *remote == "" {
		log.Fatalln("echo server address must not be nil.")
	}
	lAddr, err := net.ResolveTCPAddr("tcp", ":0")
	if err != nil {
		log.Fatalln(err)
	}
	rAddr, err := net.ResolveTCPAddr("tcp", *remote)
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := net.DialTCP("tcp", lAddr, rAddr)
	if err != nil {
		log.Fatalln(err)
	}

	echo(conn, *frequency)
}
