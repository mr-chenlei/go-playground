package main

import (
	"log"
	"math"
)

func LargestPrime(n uint64) uint64 {
	factor := uint64(3)
	maxFactor := float64(0)
	for {
		if n%factor == 0 {
			n /= factor
			maxFactor = math.Sqrt(float64(n))
			for {
				if n%factor == 0 {
					n /= factor
				} else {
					break
				}
			}
		}
		factor += 2
		if factor >= uint64(maxFactor) {
			break
		}
	}
	return uint64(maxFactor)
}

func main() {
	log.Println(LargestPrime(15))
}
