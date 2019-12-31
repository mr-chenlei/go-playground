package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Sender ...
type Sender struct {
	sync.RWMutex
	sendList map[int][]byte
	recvList map[int][]byte
}

func NewSender() *Sender {
	return &Sender{
		sendList: make(map[int][]byte, 0),
		recvList: make(map[int][]byte, 0),
	}
}

// Start ...
// addr: 目标地址
// number：数据报数量
// elapse: 发送间隔
// min, max：随机大小数据报范围
func (s *Sender) Start(addr string, number, min, max int) {
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return
	}

	go func() {
		for i := 0; i < number; i++ {
			buf := s.composePackage(i, min, max)
			conn.Write(buf)
			time.Sleep(time.Millisecond * 5)
		}
	}()
	go func() {
		var everageRTT uint64
		var counter uint64
		for {
			buf := make([]byte, 1024)
			l, _, err := conn.ReadFromUDP(buf)
			if err != nil {
				fmt.Println("ReadFromUDP error:", err)
				return
			}
			index, t1, t2, rtt := s.decomposePackage(buf[:l])
			if _, ok := s.recvList[int(index)]; !ok {
				s.recvList[int(index)] = buf[:l]
			} else {
			}

			everageRTT += rtt
			everageRTT = everageRTT / uint64(len(s.recvList))
			fmt.Println("Received package", counter, "t1:", t1, "t2:", t2, "RTT:", rtt, "Everage RTT:", everageRTT)
			counter++
		}
	}()
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (s *Sender) composePackage(index, min, max int) []byte {
	size := rand.Intn(max-min) + min
	buf := make([]byte, size+16)
	binary.BigEndian.PutUint32(buf[:4], uint32(index))
	binary.BigEndian.PutUint64(buf[4:], uint64(time.Now().UnixNano()))
	binary.BigEndian.PutUint32(buf[12:], uint32(size))
	tmp := []byte(RandStringRunes(size))
	copy(buf[16:], tmp[:])
	return buf
}

// decomposePackage ...
// index: package index
// t1: sender -> echo
// t2: echo -> sender
// rtt:
func (s *Sender) decomposePackage(buf []byte) (uint32, uint64, uint64, uint64) {
	index := binary.BigEndian.Uint32(buf[:4])
	// Timestamp of this side
	t1 := binary.BigEndian.Uint64(buf[4:12])
	size := binary.BigEndian.Uint32(buf[12:16])
	// Timestamp of other side
	t2 := binary.BigEndian.Uint64(buf[size:])

	t1 = t2 - t1
	t2 = uint64(time.Now().UnixNano() - int64(t2))
	rtt := uint64(time.Now().UnixNano() - int64(t1))
	return index, t1 / 1e6, t2 / 1e6, rtt / 1e6
}

func main() {
	s := NewSender()
	s.Start("47.101.180.117:9090", 10000, 32, 256)
	select {}
}
