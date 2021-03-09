package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"code.lstaas.com/lightspeed/atom"
	"code.lstaas.com/lightspeed/atom/app/gis/localgis"
	"code.lstaas.com/lightspeed/atom/features/gis"

	ping "github.com/caucy/batch_ping"
)

var (
	source = flag.String("source", "", "IP list source, split by space.")
)

var (
	ipList    = make(map[string]string, 0)
	afterPing = make(map[string]*ping.Statistics, 0)
)

func executePing(addr []string) {
	bp, err := ping.NewBatchPinger(addr, true) // true will need to be root

	if err != nil {
		panic("new batch ping err")
	}
	bp.SetDebug(false) // debug == true will fmt debug log

	bp.SetSource("") // if hava multi source ip, can use one isp
	bp.SetCount(5)
	bp.SetInterval(time.Second)
	bp.SetTimeout(time.Second * 15)

	bp.OnFinish = func(stMap map[string]*ping.Statistics) {
		for ip, st := range stMap {
			fmt.Printf("\n--- %s ping statistics ---\n", st.Addr)
			fmt.Printf("ip %s, %d packets transmitted, %d packets received, %v%% packet loss\n", ip,
				st.PacketsSent, st.PacketsRecv, st.PacketLoss)
			fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
				st.MinRtt, st.AvgRtt, st.MaxRtt, st.StdDevRtt)
			fmt.Printf("%v rtts is %v \n", st.Addr, st.Rtts)

			if st.PacketLoss < 100.0 {
				if _, ok := afterPing[st.Addr]; !ok {
					afterPing[st.Addr] = st
					fmt.Println("new reachable ip:", st.Addr)
				}
			}
		}
	}

	err = bp.Run()
	if err != nil {
		fmt.Printf("run err %v \n", err)
	}
	bp.OnFinish(bp.Statistics())
}

func readIPFromFile(filename string) ([]string, error) {
	var result []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func main() {
	flag.Parse()

	if result, err := readIPFromFile(*source); err != nil {
		panic(err)
	} else {
		for _, v := range result {
			if strings.Contains(v, "\t") {
				v = strings.Replace(v, "\t", " ", -1)
			}
			sub := strings.Split(v, " ")
			if len(sub) == 0 {
				continue
			}

			key := sub[len(sub)-1]
			if _, ok := ipList[key]; !ok {
				ipList[key] = v
			}
		}
		fmt.Printf("read %v(s) ip from file %v", len(ipList), *source)
	}

	fmt.Println("start ping...")

	steps := 10
	index := 0
	var list []string
	for k := range ipList {
		if len(list) < steps && index < len(ipList) {
			list = append(list, k)
			index++
			continue
		}
		t1 := time.Now()
		executePing(list)
		fmt.Printf("ping %v complete, cost: %v\n", list, time.Since(t1))
		list = []string{}
	}
	fmt.Println("ping task complete!")

	fmt.Println("start verify ip GEO")
	// Verify ip GEO
	config := &localgis.Config{}
	c, err := atom.CreateObject(context.Background(), config)
	if err != nil {
		panic(err)
	}
	g, ok := c.(gis.Client)
	if !ok {
		panic("not a feature")
	}
	_ = g.Start()

	for k := range afterPing {
		v, ok := ipList[k]
		if !ok {
			fmt.Println(k, "not existed in", *source)
		}
		gi, err := g.LookupGeoInfo(k)
		if err != nil {
			fmt.Printf("%v not found in IPDB, error: %v", k, err)
			continue
		}
		fmt.Printf("%v compare with IPDB, source [%v], in IPDB [%v %v]\n", k, v, gi.CountryCode, gi.ProvinceCode)
	}
}
