package main

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
	"time"

	"golang.org/x/net/ipv4"
)

const (
	udpHeaderLen = 8
)

func ipCheckSum(b []byte) uint16 {
	b[10] = 0
	b[11] = 0

	// to handle odd lengths, we loop to length - 1, incrementing by 2, then
	// handle the last byte specifically by checking against the original
	// length.
	var csum uint32
	length := len(b) - 1
	for i := 0; i < length; i += 2 {
		// For our test packet, doing this manually is about 25% faster
		// (740 ns vs. 1000ns) than doing it by calling binary.BigEndian.Uint16.
		csum += uint32(b[i]) << 8
		csum += uint32(b[i+1])
	}
	if len(b)%2 == 1 {
		csum += uint32(b[length]) << 8
	}
	for csum > 0xffff {
		csum = (csum >> 16) + (csum & 0xffff)
	}

	return ^uint16(csum)
}

func udpCheckSum(b []byte) uint16 {
	// to handle odd lengths, we loop to length - 1, incrementing by 2, then
	// handle the last byte specifically by checking against the original
	// length.
	var csum uint32
	length := len(b) - 1
	for i := 0; i < length; i += 2 {
		// For our test packet, doing this manually is about 25% faster
		// (740 ns vs. 1000ns) than doing it by calling binary.BigEndian.Uint16.
		csum += uint32(b[i]) << 8
		csum += uint32(b[i+1])
	}
	if len(b)%2 == 1 {
		csum += uint32(b[length]) << 8
	}
	for csum > 0xffff {
		csum = (csum >> 16) + (csum & 0xffff)
	}

	return ^uint16(csum)
}

func parseAddr(b []byte) (uint8, string, string) {
	network := uint8(b[9])

	sAddr := net.IPv4(b[12], b[13], b[14], b[15]).String()
	rAddr := net.IPv4(b[16], b[17], b[18], b[19]).String()
	sPort := binary.BigEndian.Uint16(b[20:22])
	rPort := binary.BigEndian.Uint16(b[22:24])

	sAddr = sAddr + ":" + strconv.Itoa(int(sPort))
	rAddr = rAddr + ":" + strconv.Itoa(int(rPort))
	log.Println("network:", network, "source:", sAddr, "remote:", rAddr)
	return network, sAddr, rAddr
}

func main() {
	buff := []byte("Hello LightSpeed!!")
	dst := net.IPv4(192, 168, 1, 1)
	src := net.IPv4(192, 168, 1, 100)
	iph := &ipv4.Header{
		Version:  ipv4.Version,
		Len:      ipv4.HeaderLen,
		TOS:      0x00,
		TotalLen: ipv4.HeaderLen + len(buff),
		TTL:      64,
		Flags:    ipv4.DontFragment,
		FragOff:  0,
		Protocol: 17,
		Checksum: 0,
		Src:      src,
		Dst:      dst,
	}
	h, err := iph.Marshal()
	if err != nil {
		log.Fatalln(err)
	}
	// 计算IP头部校验值
	iph.Checksum = int(ipCheckSum(h))

	// 填充udp首部
	// udp伪首部
	udph := make([]byte, 20)
	// 源ip地址
	copy(udph[0:4], src)
	// 目的ip地址
	copy(udph[4:8], dst)
	// 协议类型
	udph[8], udph[9] = 0x00, 0x11
	// udp头长度
	udph[10], udph[11] = 0x00, byte(len(buff)+udpHeaderLen)
	// 下面开始就真正的udp头部
	// 源端口号
	udph[12], udph[13] = 0xf9, 0x1e
	// 目的端口号
	udph[14], udph[15] = 0x30, 0x39
	// udp头长度
	udph[16], udph[17] = 0x00, byte(len(buff)+udpHeaderLen)
	// 校验和
	udph[18], udph[19] = 0x00, 0x00
	// 计算校验值
	check := udpCheckSum(udph)
	udph[18], udph[19] = byte(check>>8&255), byte(check&255)
	log.Println("checksum:", check, udph[18], udph[19])

	listener, err := net.ListenPacket("ip4:udp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	r, err := ipv4.NewRawConn(listener)
	for {
		if err = r.WriteTo(iph, append(udph, buff...), nil); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}
