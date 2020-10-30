package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ip2location/ip2location-go"

	"github.com/MrVegeta/go-playground/log-extractor/geoquery"
	_ "github.com/ip2location/ip2location-go"
)

type SourceInfo struct {
	ip      []byte
	port    int
	network int
	time    []byte
	counter int
}

var (
	log        = flag.String("log", "", "Local log file path name.")
	query      = flag.String("query", "", "Query in ipip.ipdb.")
	showSource = flag.Bool("source", false, "Show source address and counts.")
	showTarget = flag.Bool("target", false, "Show target address and counts.")
	matchRule  = flag.Bool("match-rule", false, "Show match/no match rules.")
	// galaxy product server list
	exclude = map[string]bool{
		"103.101.204.155": true,
		"103.192.214.70":  true,
		"124.156.215.53":  true,
		"129.226.118.21":  true,
		"134.175.139.123": true,
		"139.224.26.46":   true,
		"150.109.238.19":  true,
		"152.136.175.63":  true,
		"212.64.95.156":   true,
		"39.106.32.116":   true,
		"39.97.253.178":   true,
		"39.97.254.158":   true,
		"47.101.180.117":  true,
		"47.244.246.88":   true,
		"47.245.56.93":    true,
		"47.75.215.147":   true,
		"47.88.29.108":    true,
		"49.51.154.43":    true,
		"49.51.244.94":    true,
		"8.209.73.162":    true,
		"129.211.105.170": true,
		"49.232.140.221":  true,
	}

	targetList = make(map[string]int, 0)
)

const (
	keywordSourceAddress = "Address"
	keywordTargetAddress = "target"
	keywordMatchRule     = "match rule"
)

func extractIP(line []byte) (*SourceInfo, error) {
	info := &SourceInfo{
		counter: 1,
	}
	if bytes.Contains(line, []byte("source")) {
		info.ip = line[bytes.LastIndex(line, []byte("["))+1 : bytes.LastIndex(line, []byte("]"))]
		info.ip = bytes.ReplaceAll(info.ip, []byte(","), []byte("."))
		return info, nil
	}
	return nil, errors.New("no source info contained")
}

func extractByKeyword(line, keyword []byte) ([]byte, error) {
	if bytes.Contains(line, keyword) == false {
		return nil, errors.New("keyword not found")
	}
	match := make([]byte, 32)
	switch string(keyword) {
	case keywordTargetAddress:
		pos1 := bytes.LastIndex(line, []byte(keywordTargetAddress)) + 12
		pos2 := bytes.LastIndex(line, []byte("."))
		match = line[pos1:pos2]
		bytes.Replace(match, []byte(","), []byte("."), -1)
	}
	return match, nil
}

func main() {
	flag.Parse()
	err := geoquery.Init("./ipip.ipdb")
	if err != nil {
		fmt.Println("geoquery init error:", err)
		return
	}
	ip2location.Open("./IP-COUNTRY-REGION-CITY-ISP-SAMPLE.BIN")

	if *query != "" {
		geo, err := geoquery.GetGEOInfo(*query)
		if err != nil {
			fmt.Println("query ipip.ipdb error:", err)
			return
		}
		fmt.Println("query", *query, "success:", geo)
		return
	}
	fmt.Println("Log file:", *log)
	if *log == "" {
		fmt.Println("Log must not empty.")
		return
	}
	fi, err := os.Open(*log)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		target, err := extractByKeyword(line, []byte(keywordTargetAddress))
		if err == nil {
			if _, ok := targetList[string(target)]; !ok {
				targetList[string(target)] = 1
			} else {
				targetList[string(target)]++
			}
		}
	}
	for k, v := range targetList {
		s := strings.Split(k, ":")
		geo1, _ := geoquery.GetGEOInfo(s[0])
		geo2 := ip2location.Get_all(s[0])
		fmt.Println("------>target:", k, "counts:", v, "geo:", geo1, ",", geo2.Country_long, geo2.City)
	}
}
