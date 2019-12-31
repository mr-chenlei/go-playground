package main

import (
	"C"
	"fmt"
)
import "time"

//export DisplayTime
func DisplayTime(msg string) {
	fmt.Println(msg, time.Now())
}

func main() {
	DisplayTime("DisplayTime: ")
}
