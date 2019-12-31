package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"v2ray.com/core/common/net"
)

type Log2 struct {
	sync.RWMutex
	addressWithTimeFileName    string
	addressWithTime            []string
	addressWithCounterFileName string
	addressWithCounter         map[net.Destination]int
}

// NewLog2 ...
func NewLog2(prefix string) *Log2 {
	l := &Log2{
		addressWithTimeFileName:    FileNameByHour("./", prefix+"_addr_with_time_"),
		addressWithTime:            make([]string, 0),
		addressWithCounterFileName: FileNameByHour("./", prefix+"_addr_with_counter_"),
		addressWithCounter:         make(map[net.Destination]int, 0),
	}
	go l.logTimer()
	return l
}

// FileNameByHour 将按当前UTC时间小时返回一个与时间有关的文件名
func FileNameByHour(folder string, prefix string) string {
	tm := time.Now().UTC()
	layout := "2006010215150405"
	filename := fmt.Sprintf("%s%s%s", prefix, tm.Format(layout), ".log")
	return path.Join(folder, filename)
}

func (l *Log2) LogDest(dst net.Destination) {
	l.Lock()
	defer l.Unlock()

	l.addressWithTime = append(l.addressWithTime, time.Now().String()+": "+dst.String())

	if _, ok := l.addressWithCounter[dst]; !ok {
		l.addressWithCounter[dst] = 1
	} else {
		l.addressWithCounter[dst]++
	}
}

func (l *Log2) logTimer() {
	for range time.Tick(time.Second * 5) {
		l.Lock()
		if len(l.addressWithTime) > 0 {
			f2, _ := os.OpenFile(l.addressWithTimeFileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
			for _, v := range l.addressWithTime {
				io.WriteString(f2, v+"\n")
				f2.Sync()
			}
			f2.Close()
			f1, _ := os.OpenFile(l.addressWithCounterFileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
			for k, v := range l.addressWithCounter {
				data := k.String() + ": " + strconv.Itoa(v)
				io.WriteString(f1, data+"\n")
				f1.Sync()
			}
			f1.Close()
		}
		l.Unlock()
	}
}

func (l *Log2) exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func main() {
	l := NewLog2("test")

	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.1")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.2")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.3")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.4")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.5")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.1")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.2")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.3")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.4")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.5")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.1")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.2")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.3")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.4")})
	l.LogDest(net.Destination{Address: net.ParseAddress("127.0.0.5")})
	select {}
}
