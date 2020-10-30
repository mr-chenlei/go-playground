package main

import (
	"strconv"
	"sync"
	"time"
)

type data struct {
	key     string
	version int64
}

var (
	db    map[string]*data
	mutex sync.RWMutex
)

func Add2DB(d *data) {
	mutex.Lock()
	defer mutex.Unlock()
	db[d.key] = d
}

func GetFromDB(k string) *data {
	mutex.RLock()
	defer mutex.RUnlock()
	if _, ok := db[k]; !ok {
		return nil
	}
	return db[k]
}

func write() {
	var i int
	for {
		key := strconv.Itoa(i)
		i++

		d := &data{
			key:     key,
			version: time.Now().Unix(),
		}
		Add2DB(d)

		time.Sleep(1 * time.Millisecond)
	}
}

func read() {
	for {
		var i int
		for {
			key := strconv.Itoa(i)
			d := GetFromDB(key)
			if d == nil {
				break
			}
			d.version = time.Now().Unix()
			Add2DB(d)
			i++
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	db = make(map[string]*data, 0)
	go write()
	go write()
	go write()
	go read()
	go read()
	go read()

	select {}
}
