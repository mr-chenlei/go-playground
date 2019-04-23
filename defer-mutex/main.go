package main

import (
	"fmt"
	"sync"
	"time"
)

var l sync.RWMutex

func doSomething() {
	for {
		l.RLock()
		defer l.RUnlock()

		fmt.Printf("%04d-%02d-%02d, %02d:%02d:%02d\n",
			time.Now().Year(), time.Now().Month(), time.Now().Day(),
			time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	}
}

func main() {
	go doSomething()
	go doSomething()

	select {}
}
