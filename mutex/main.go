package main

import (
	"fmt"
	"sync"
)

var lock sync.RWMutex

func main() {
	for i := 0; i < 10; i++ {
		lock.Lock()
		fmt.Println("Locked", i)
		if i == 7 {
			continue
		}
		lock.Unlock()
		fmt.Println("Unlocked:", i)
	}
}
