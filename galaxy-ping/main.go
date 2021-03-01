package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"runtime"
	"time"

	pinger "github.com/go-ping/ping"
)

const (
	pingEchoLength      = 28 // socks5(10 bytes) + sequence(2 bytes) + struct timeval(16 bytes)
	pingEchoReplyLength = pingEchoLength + 11
)

var (
	local = flag.Int("local", 10070, "Local listen port.")
)

type pingReply struct {
	Addr  string
	Reply []byte
}

func makeReply(b []byte, pkt *pinger.Packet) *pingReply {
	// Ping echo reply
	// +--------+----------+------+-------+----+------+
	// | SOCKS5 | Sequence | time | Bytes | RTT | TTL |
	// +--------+----------+------+-------+----+------+
	// |   10   |    2     |  16  |   2   |  8  |  1  |
	// +--------+----------+------+-------+----+------+
	//
	// Bytes: data size
	// RTT:
	// TTL:
	// Note: Both RTT & TTL nil means 100% packet loss
	reply := make([]byte, pingEchoReplyLength)
	copy(reply[:pingEchoLength], b[:pingEchoLength]) // Original data
	// Put Bytes
	binary.BigEndian.PutUint16(reply[pingEchoLength:pingEchoLength+2], uint16(pkt.Nbytes))
	// Put RTT
	binary.BigEndian.PutUint64(reply[pingEchoLength+2:pingEchoLength+2+8], uint64(pkt.Rtt))
	// Put TTL
	reply[pingEchoReplyLength-1] = byte(pkt.Ttl)
	return &pingReply{
		Addr:  pkt.Addr,
		Reply: reply,
	}
}

func reply(reply *pingReply, conn *net.UDPConn, src *net.UDPAddr) {
	p := make([]byte, pingEchoReplyLength)
	copy(p[:], reply.Reply[:])
	seq := binary.BigEndian.Uint16(p[10:12])
	t1 := binary.BigEndian.Uint64(p[12:20])
	t2 := binary.BigEndian.Uint64(p[20:28])
	log.Printf("reply ping result icmp_seq=%v send-time1: [%v] send-time2: [%v]\n",
		seq, t1, t2)
	if l, err := conn.WriteTo(p[:], src); err != nil {
		log.Printf("send ping reply error: %v\n", err)
	} else {
		log.Printf("send ping reply success, send data size: %d\n", l)
	}
}

func executePing(b []byte, conn *net.UDPConn, src *net.UDPAddr) error {
	if len(b) < pingEchoLength {
		return fmt.Errorf("in complete ping request")
	}
	addr := net.IP{b[4], b[5], b[6], b[7]}.String()
	seq := binary.BigEndian.Uint16(b[10:12])
	t1 := binary.BigEndian.Uint64(b[12:20])
	t2 := binary.BigEndian.Uint64(b[20:28])
	log.Printf("receive ping request to target %s, icmp_seq=%v send-time1: [%v] send-time2: [%v]\n", addr, seq, t1, t2)

	p, err := pinger.NewPinger(addr)
	if err != nil {
		return err
	}
	p.Count = 1
	p.Interval = time.Second
	p.Timeout = time.Second * 5
	switch runtime.GOOS {
	case "linux":
		p.SetPrivileged(true)
	default:
		p.SetPrivileged(false)
	}

	p.OnRecv = func(pkt *pinger.Packet) {
		log.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v\n", pkt.Nbytes, pkt.IPAddr, seq, pkt.Rtt, pkt.Ttl)

		pkt.Seq = int(seq)

		reply(makeReply(b, pkt), conn, src)
	}
	p.OnDuplicateRecv = func(pkt *pinger.Packet) {
		log.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n", pkt.Nbytes, pkt.IPAddr, seq, pkt.Rtt, pkt.Ttl)

		pkt.Seq = int(seq)
		reply(makeReply(b, pkt), conn, src)
	}
	p.OnFinish = func(stats *pinger.Statistics) {
		log.Printf("--- %s icmp_seq=%v ping statistics ---\n", stats.Addr, seq)
		log.Printf("icmp_seq=%v %d packets transmitted, %d packets received, %d duplicates, %f%% packet loss\n",
			seq, stats.PacketsSent, stats.PacketsRecv, stats.PacketsRecvDuplicates, stats.PacketLoss)
		log.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
		if stats.PacketLoss >= 100 {
			reply(makeReply(b, &pinger.Packet{
				Addr: stats.Addr,
			}), conn, src)
		}
	}

	return p.Run()
}

func main() {
	flag.Parse()
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: *local,
		IP:   net.ParseIP("0.0.0.0"),
	})
	if err != nil {
		log.Panic(err)
	}

	log.Println("start working...")
	for {
		buf := make([]byte, 128)
		l, remote, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v receive message from %s\n", time.Now(), remote.String())
		go executePing(buf[:l], conn, remote)
	}
}
