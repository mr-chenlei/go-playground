package main

import (
	"log"
	"time"
)

func isMultiple(n int, start, maxDivisor int) bool {
	for i := start; i <= maxDivisor; i++ {
		if (n % i) != 0 {
			return false
		}
	}
	return true
}

func findMultipleWithDivisor(d int) int {
	i := 1
	for {
		if isMultiple(i, 1, d) {
			return i
		}
		i++
	}
}

func findMultipleWithDivisorV2(d int) int {
	i := 1
	for {
		n := i * d
		if isMultiple(n, 1, d) {
			return n
		}
		i++
	}
}

func main() {
	t1 := time.Now()
	ret := findMultipleWithDivisorV2(20)
	log.Println(ret, time.Since(t1))
}
