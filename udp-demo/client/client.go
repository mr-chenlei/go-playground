package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

var (
	host      = flag.String("host", "", "Remove host address.")
	remote    = flag.String("remote", "", "Remote server IP.")
	frequency = flag.Int("frequency", 1000, "Echo message frequency(ms).")

	globalIPList = []string{
		"1.10.140.43",     // TG#TH-1_10_140_43
		"103.10.124.45",   // TG#SG-103_10_124_45
		"103.10.125.146",  // TG#AU-103_10_125_146
		"103.28.86.241",   // TG#NP-103_28_86_241
		"103.44.18.248",   // TG#IN-103_44_18_248
		"105.247.244.235", // TG#ZA-105_247_244_235
		"110.37.226.83",   // TG#PK-110_37_226_83
		"111.125.88.188",  // TG#PH-111_125_88_188
		"14.161.44.120",   // TG#VN-14_161_44_120
		"146.66.154.35",   // TG#LU-146_66_154_35
		"149.56.1.48",     // TG#CA-149_56_1_48
		"155.133.238.162", // TG#ZA-155_133_238_162
		"155.133.249.194", // TG#CL-155_133_249_194
		"155.133.252.34",  // TG#SE-155_133_252_34
		"162.254.196.66",  // TG#GB-162_254_196_66
		"162.254.197.70",  // TG#DE-162_254_197_70
		"175.144.198.226", // TG#MY-175_144_198_226
		"176.124.96.9",    // TG#AM-176_124_96_9
		"178.62.246.248",  // TG#NL-178_62_246_248
		"181.10.129.85",   // TG#AR-181_10_129_85
		"202.181.202.140", // TG#HK-202_181_202_140
		"202.57.47.122",   // TG#PH-202_57_47_122
		"203.194.21.241",  // TG#AU-203_194_21_241
		"203.202.243.198", // TG#BD-203_202_243_198
		"85.26.146.169",   // TG#RU_85.26_146_169
		"91.102.231.54",   // TG#RS-91_102_231_54
		"93.113.49.1",     // TG#ES-93_113_49_1
		"94.74.184.76",    // TG#IR-94_74_184_76
	}
)

func send(addr string, elapse int) {
	fmt.Println(time.Now(), "sending to:", addr)
	RemoteAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Fatalln(err)
		return
	}
	conn, err := net.DialUDP("udp4", nil, RemoteAddr)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()

	go func() {
		message := []byte{0, 0, 0, 1, 47, 102, 45, 160, 39, 96}
		message = append(message, "hello LightSpeed!"...)
		t := time.Duration(elapse)
		for range time.Tick(t * time.Millisecond) {
			_, err = conn.Write(message)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	reply := make([]byte, 1024)
	for {
		n, raddr, err := conn.ReadFromUDP(reply[:])
		if err != nil {
			break
		}
		log.Println(raddr.String(), "Read from udp server:", string(reply[:n]))
	}
}

func sendSingleUDPPacket(addr string) {
	log.Println("send single udp packet to:", addr)
	RemoteAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Fatalln(err)
		return
	}
	conn, err := net.DialUDP("udp4", nil, RemoteAddr)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("hello world!"))
}

func main() {
	flag.Parse()
	go func() {
		for {
			addrMap := make(map[string]interface{}, 0)
			for _, v := range globalIPList {
				addrMap[v] = nil
			}
			rand.Seed(time.Now().UnixNano())
			port := 12345 //rand.Intn(0xFFFF-0x2000) + 0x2000
			for k, _ := range addrMap {
				sendSingleUDPPacket(k + ":" + strconv.Itoa(port))
				time.Sleep(500 * time.Microsecond)
			}
			time.Sleep(15 * time.Second)
		}
	}()

	//go func() {
	//	send(*remote, *frequency)
	//}()

	select {}
}
