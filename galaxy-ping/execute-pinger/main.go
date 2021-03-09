package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var (
	period = flag.String("period", "5m", "Ping period.")
	target = flag.String("target", "43.231.186.20,185.78.104.15,180.178.75.144,69.28.52.109", "Target list, split by comma.")
	times  = flag.String("times", "30", "Ping times.")
)

func executeGalaxyPing(addr, times string) ([]byte, error) {
	var err error
	var result []byte
	cmd := exec.Command("./pinger", addr, times)
	if result, err = cmd.CombinedOutput(); err != nil && len(result) == 0 {
		return nil, err
	}
	return result, nil
}

func executePing(addr, times string) ([]byte, error) {
	var err error
	var result []byte
	var count string
	switch runtime.GOOS {
	case "darwin":
		count = "-t"
	case "linux":
		count = "-c"
	case "windows":
		count = "-n"
	default:
		count = "-c"
	}
	cmd := exec.Command("ping", addr, count, times)
	if result, err = cmd.CombinedOutput(); err != nil && len(result) == 0 {
		return nil, err
	}
	return result, nil
}

func write2File(prefix, suffix string, data []byte) error {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
	timestamp = strings.Replace(timestamp, "-", "", -1)
	timestamp = strings.Replace(timestamp, ":", "", -1)
	timestamp = strings.Replace(timestamp, ".", "_", -1)
	hostname := ""
	if name, err := os.Hostname(); err != nil {
		hostname = ""
	} else {
		hostname = name + "_"
	}
	filename := "./log/" + hostname + prefix + "_" + timestamp + "_" + suffix + ".log"
	return ioutil.WriteFile(filename, data, 0644)
}

func main() {
	flag.Parse()

	targetList := strings.Split(*target, ",")

	p, err := time.ParseDuration(*period)
	if err != nil {
		log.Panic(err)
	}
	log.Println("start ping with time period:", p)
	for range time.Tick(p) {
		for _, v := range targetList {
			log.Println("ping", v, *times)
			suffix := strings.Replace(v, ".", "_", -1)
			if result, err := executePing(v, *times); err != nil {
				log.Printf("execute ping on target %v failed, error: %v\n", v, err)
			} else if err := write2File("ping", suffix, result); err != nil {
				log.Println("write ping result to file failed, error:", err)
			}
			log.Println("ping", v, "complete")

			log.Println("galaxy-ping", v, *times)
			if result, err := executeGalaxyPing(v, *times); err != nil {
				log.Printf("execute ping on target %v failed, error: %v\n", v, err)
			} else if err := write2File("galaxy-ping", suffix, result); err != nil {
				log.Println("write ping result to file failed, error:", err)
			}
			log.Println("galaxy-ping", v, "compelte")
		}
	}
}
