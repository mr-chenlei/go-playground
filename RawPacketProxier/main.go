package main

import (
	"encoding/binary"
	"flag"
	"log"
	"net"
	"time"

	"golang.org/x/net/ipv4"
)

var (
	network     = flag.String("network", "", "Network type TCP/UDP")
	source      = flag.String("source", "", "Source address.")
	destination = flag.String("dest", "", "Destination address.")
	payload     = flag.String("payload", "", "Payload")
)

func checksum(b []byte) uint16 {
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

func udpChecksum(b []byte) uint16 {
	b[18] = 0
	b[19] = 0
	csum := checksum(b[:])
	return 0
}

func ipChecksum(b []byte) uint16 {
	b[10] = 0
	b[11] = 0
	csum := checksum(b[:])
}

func newUDPPacket(src, dst string, payload []byte) ([]byte, error) {
	srcAddr, err := net.ResolveUDPAddr("udp", src)
	if err != nil {
		return nil, err
	}
	dstAddr, err := net.ResolveUDPAddr("udp", dst)
	if err != nil {
		return nil, err
	}
	b := make([]byte, 20)
	// source ip
	b = append(b, srcAddr.IP...)
	// destination ip
	b = append(b, dstAddr.IP...)
	// protocol
	binary.BigEndian.PutUint16(b[8:10], 17)
	// udp header length
	binary.BigEndian.PutUint16(b[10:12], uint16(len(payload)+20))
	// source port
	binary.BigEndian.PutUint16(b[12:14], uint16(srcAddr.Port))
	// destination port
	binary.BigEndian.PutUint16(b[14:16], uint16(dstAddr.Port))
	// udp header length
	binary.BigEndian.PutUint16(b[16:18], uint16(len(payload)+20))
	// checksum
	udpChecksum(b[:])
	return b, nil
}

func newIPPacket(network, src, dst string, payload []byte) (*ipv4.Header, error) {
	var protocol int
	switch network {
	case "client":
		protocol = 6
	case "udp":
		protocol = 17
	}
	srcAddr, err := net.ResolveUDPAddr("udp", src)
	if err != nil {
		return nil, err
	}
	dstAddr, err := net.ResolveUDPAddr("udp", dst)
	if err != nil {
		return nil, err
	}
	ip := &ipv4.Header{
		Version:  ipv4.Version,
		Len:      ipv4.HeaderLen,
		TOS:      0x00,
		TotalLen: ipv4.HeaderLen + len(payload),
		TTL:      64,
		Flags:    ipv4.DontFragment,
		FragOff:  0,
		Protocol: protocol,
		Checksum: 0,
		Src:      srcAddr.IP,
		Dst:      dstAddr.IP,
	}
	h, err := ip.Marshal()
	if err != nil {
		return nil, err
	}
	ip.Checksum = int(ipChecksum(h))
	return ip, nil
}

func main() {
	flag.Parse()

	ip, err := newIPPacket(*network, *source, *destination, []byte(*payload))
	if err != nil {
		log.Fatalln(err)
	}
	udp, err := newUDPPacket(*source, *destination, []byte(*payload))
	if err != nil {
		log.Fatalln(err)
	}

	listener, err := net.ListenPacket("ip4:udp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	r, err := ipv4.NewRawConn(listener)
	for {
		if err = r.WriteTo(ip, udp, nil); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}
