package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

var slice1024 = []byte("90123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234")
var redundancy = make(map[uint64]uint32)

func reader(conn *net.UDPConn) {
	buffer := make([]byte, 4096)
	_, raddr, err := conn.ReadFrom(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Receive message from: %s\n", raddr.String())
	go writer(conn, raddr)

	var readCounter = uint64(1)
	for {
		buf := make([]byte, 4096)
		l, raddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			break
		}

		fmt.Println("reader:", raddr.String(), "len:", l)

		readCounter++
		index := binary.BigEndian.Uint64(buf[0:8])
		if _, ok := redundancy[index]; !ok {
			redundancy[index] = 1
			fmt.Println(time.Now(), raddr, "Total received package:", readCounter, "package index:", index)
		} else {
			redundancy[index]++
			fmt.Println(time.Now(), raddr, "Total received package:", readCounter, "redundancy package index:", index, "repeat times:", redundancy[index])
		}
	}
}

func writer(conn *net.UDPConn, raddr net.Addr) {
	var counter = uint64(30000)
	for i := counter; i < 60000; i++ {
		buffer := make([]byte, 0)
		// Package index
		counter++
		wc := make([]byte, 8)
		binary.BigEndian.PutUint64(wc, counter)
		buffer = append(buffer, wc...)
		// Data
		buffer = append(buffer, slice1024...)

		conn.WriteTo(buffer[:], raddr)
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println(time.Now(), "--------->Writer DONE!")
}

func main() {
	addr := &net.UDPAddr{
		IP:   net.IP{0, 0, 0, 0},
		Port: 9090,
	}
	fmt.Println(addr)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	go reader(conn)

	select {}
}
