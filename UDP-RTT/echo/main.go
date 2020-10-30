package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"time"
)

type Receiver struct {
	sync.RWMutex
	recvList map[uint64][]byte
}

func NewReceiver() *Receiver {
	return &Receiver{
		recvList: make(map[uint64][]byte, 0),
	}
}

func (r *Receiver) Start(addr string) {
	pc, err := net.ListenPacket("udp", addr)
	if err != nil {
		fmt.Println("ListenPacket error:", err)
		return
	}
	fmt.Println("ListenPacket on:", pc.LocalAddr())
	go func() {
		buf := make([]byte, 1024)
		for {
			n, raddr, err := pc.ReadFrom(buf)
			if err != nil {
				fmt.Println("ReadFrom error:", err)
				return
			}
			time.Sleep(time.Millisecond * 5)
			binary.BigEndian.PutUint64(buf[n:], uint64(time.Now().UnixNano()))

			r.decomposePackage(buf[:])

			_, err = pc.WriteTo(buf, raddr)
			if err != nil {
				fmt.Println("Write buffer back error:", err)
				return
			}
		}
	}()
}

// decomposePackage ...
// index: package index
// t1: sender -> echo
// t2: echo -> sender
// rtt:
func (r *Receiver) decomposePackage(buf []byte) (uint32, uint64, uint64, uint64) {
	index := binary.BigEndian.Uint32(buf[:4])
	// Timestamp of this side
	t1 := binary.BigEndian.Uint64(buf[4:12])
	size := binary.BigEndian.Uint32(buf[12:16])
	// Timestamp of other side
	t2 := binary.BigEndian.Uint64(buf[size:])

	t1 = t2 - t1
	t2 = uint64(time.Now().UnixNano() - int64(t2))
	rtt := uint64(time.Now().UnixNano() - int64(t1))
	fmt.Println("------", "t1:", t1/1e6, "t2:", t2/1e6, "rtt:", rtt/1e6)
	return index, t1, t2 / 1e6, rtt / 1e6
}

func main() {
	r := NewReceiver()
	r.Start(":9090")
	select {}
}
