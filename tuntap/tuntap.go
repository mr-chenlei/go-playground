package main

import (
	"encoding/binary"
	"flag"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/songgao/water"
)

const (
	// I use TUN interface, so only plain IP packet, no ethernet header + mtu is set to 1300
	BufferSize = 1500
	MTU        = "1300"

	NetworkTypeTCP = 6
	NetworkTypeUDP = 17

	RoleTypeClient = "client"
	RoleTypeServer = "server"
)

var (
	remoteIP = flag.String("remote", "", "Remote server (external) IP like 8.8.8.8")
	localIP  = flag.String("local", "", "Source IP in IP datagram")
	replace  = flag.String("replace", "", "Replace destination address in IP")
	role     = flag.String("role", "", "Client or Server")
)

func runIP(args ...string) {
	cmd := exec.Command("ip", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if nil != err {
		log.Fatalln("Error running ip:", err)
	}
}

func parseAddr(b []byte) (uint8, string, string) {
	log.Println("------>", b[:28])
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

func verifyIPChecksum(b []byte) uint16 {
	return binary.BigEndian.Uint16(b[10:12])
}

func replaceIPDestAddr(b []byte, addr []byte) {
	copy(b[16:20], addr[:])
	ipChecksum(b[:])
}

func ipChecksum(b []byte) uint16 {
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

func badCheckSum(b []byte) uint16 {
	b[10] = 0
	b[11] = 0

	sum := 0
	for n := 1; n < len(b)-1; n += 2 {
		sum += int(b[n])*256 + int(b[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	return uint16(^sum)
}

func switchAddress(b []byte) {
	// replace address
	sAddr := net.IPv4(b[12], b[13], b[14], b[15])
	rAddr := net.IPv4(b[16], b[17], b[18], b[19])
	copy(b[12:16], rAddr.To4())
	copy(b[16:20], sAddr.To4())
	// replace ports
	sPort := binary.BigEndian.Uint16(b[20:22])
	rPort := binary.BigEndian.Uint16(b[22:24])
	binary.BigEndian.PutUint16(b[20:22], rPort)
	binary.BigEndian.PutUint16(b[22:24], sPort)
}

func replaceDest(b []byte, ip []byte, port int) {
	copy(b[16:20], ip[:])
	binary.BigEndian.PutUint16(b[22:24], uint16(port))
}

func newInterface(addr string) (*water.Interface, error) {
	// create TUN interface
	iface, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if nil != err {
		return nil, err
	}
	// set interface parameters
	runIP("link", "set", "dev", iface.Name(), "mtu", MTU)
	runIP("addr", "add", addr, "dev", iface.Name())
	runIP("link", "set", "dev", iface.Name(), "up")
	return iface, nil
}

func addIPRoute(from, via string) {
	runIP("rule", "add", "from", from, "to", "all", "table", "100", "prio", "100")
	runIP("route", "add", "default", "via", via, "table", "100")
}

func main() {
	flag.Parse()

	tun0 := "192.168.2.1/24"
	// create TUN interface
	tunDev, err := newInterface(tun0)
	if nil != err {
		log.Fatalln("allocate TUN0 interface error:", err)
	}
	// All traffic from 192.168.1.0/24 via tun0
	if strings.ToLower(*role) == RoleTypeClient {
		//addIPRoute("192.168.1.0/24", "192.168.2.1")
	}
	log.Println("Interface allocated:", tunDev.Name())
	log.Println("start reading from tun...")

	b := make([]byte, BufferSize)
	for {
		n, err := tunDev.Read(b)
		if err != nil {
			log.Fatalln("read from tun interface error:", err)
		}

		parseAddr(b[:n])
	}
}
