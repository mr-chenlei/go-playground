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
	var readCounter uint64
	for {
		buffer := make([]byte, 4096)
		l, raddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			break
		}

		fmt.Println("reader:", raddr.String(), "len:", l)

		index := binary.BigEndian.Uint64(buffer[10:18])
		if _, ok := redundancy[index]; !ok {
			readCounter++
			fmt.Println(time.Now(), raddr, "Total received package:", readCounter, "package index:", index)
			redundancy[index] = 1
		} else {
			redundancy[index]++
			fmt.Println(time.Now(), raddr, "Total received package:", readCounter, "redundancy package index:", index, "repeat times:", redundancy[index])
		}
	}
}

func writer(conn *net.UDPConn) {
	//raddr := net.IP{47, 101, 180, 117} // sh01-dev
	raddr := net.IP{47, 102, 45, 160} // sh02-dev
	rport := []byte{0x23, 0x82}
	ss5 := make([]byte, 0)
	ss5 = append(ss5, 0x0, 0x0, 0x0)
	ss5 = append(ss5, 0x01)
	ss5 = append(ss5, raddr.To4()...)
	ss5 = append(ss5, rport...)

	var writeCounter uint64
	for i := 0; i < 30000; i++ {
		buffer := make([]byte, 0)
		// Socks5 header
		buffer = append(buffer, ss5...)
		// Package index
		writeCounter++
		wc := make([]byte, 8)
		binary.BigEndian.PutUint64(wc, writeCounter)
		buffer = append(buffer, wc...)
		// Data
		buffer = append(buffer, slice1024...)
		time.Now().Unix()

		conn.Write(buffer[:])
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println(time.Now(), "--------->Writer DONE!")
}

func main() {
	laddr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:9091")
	raddr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:10080")
	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	go writer(conn)
	go reader(conn)

	select {}
}
