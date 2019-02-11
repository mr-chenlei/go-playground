package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Printf("Current time in second: %02d:%02d:%02d\n",
				time.Now().Hour(), time.Now().Minute(), time.Now().Second())
			time.Sleep(1 * time.Second)
		}
	}()

	select {}
}
