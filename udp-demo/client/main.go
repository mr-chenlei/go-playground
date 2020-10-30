package main

import (
	"crypto/rand"
	"encoding/binary"
	"flag"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

var (
	local    = flag.String("local", "0.0.0.0:0", "Remove host address.")
	target   = flag.String("target", "150.109.41.205:12345", "Remote server IP.")
	runEvery = flag.String("run-every", "1m", "Echo message frequency(ms).")
	keepRun  = flag.String("keep-run", "30s", "Keep running specific time.")
	worker   = flag.Int("worker", 100, "TCP worker number.")
)

var (
	logger         *log.Logger
	udpConnections = make(map[string]interface{}, 0)
)

const (
	minPacketLength = 5
	maxPacketLength = 1400
)

func random(x int64) int64 {
	if n, err := rand.Int(rand.Reader, big.NewInt(x)); err != nil {
		return 0
	} else {
		return n.Int64()
	}
}

func randomBytes(size int64) (blk []byte, err error) {
	l := int64(0)
	for {
		l = random(size)
		if l > minPacketLength {
			break
		}
	}
	blk = make([]byte, l)
	_, err = rand.Read(blk)
	return
}

func newPacket(index uint64) []byte {
	buffer, err := randomBytes(maxPacketLength)
	if err != nil {
		return nil
	}
	header := make([]byte, 8+2+8) // index + length
	binary.BigEndian.PutUint64(header[:8], index)
	binary.BigEndian.PutUint16(header[8:10], uint16(len(buffer)))
	binary.BigEndian.PutUint64(header[10:18], uint64(time.Now().UnixNano()))
	header = append(header, buffer...)
	return header
}

func decodePacket(b []byte) (uint64, uint16, int64) {
	return binary.BigEndian.Uint64(b[:8]), binary.BigEndian.Uint16(b[8:10]), int64(binary.BigEndian.Uint64(b[10:18]))
}

func tcpWorker(local, target string) {
	laddr, err := net.ResolveTCPAddr("tcp", local)
	if err != nil {
		logger.Println(err)
		return
	}
	raddr, err := net.ResolveTCPAddr("tcp", target)
	if err != nil {
		logger.Println(err)
		return
	}

	t1 := time.Now()
	conn, err := net.DialTCP("tcp", laddr, raddr)
	if err != nil {
		logger.Println(err)
		return
	}
	b := make([]byte, 32)
	conn.Read(b)
	elapse := time.Since(t1)
	d, _ := time.ParseDuration("1s")
	if elapse > d {
		logger.Printf("%v dial timeout, cost: %v", conn.LocalAddr(), elapse)
	}

	conn.Close()
}

func main() {
	flag.Parse()

	_ = os.Mkdir("./", 0755)
	logFile, err := os.OpenFile("./tcp-client.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if nil != err {
		panic(err)
	}
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	logger.Println("start working....")
	re, _ := time.ParseDuration(*runEvery)
	for range time.Tick(re) {
		logger.Println("start running tcp dialer")
		kr, _ := time.ParseDuration(*keepRun)

		go func() {
			t := time.NewTimer(kr)
			for range time.Tick(time.Second) {
				for j := 0; j < *worker; j++ {
					go tcpWorker(*local, *target)
				}

				select {
				case <-t.C:
					t.Stop()
					return
				}
			}
		}()
	}

	select {}
}
