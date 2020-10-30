package main

import (
	"log"
	"strconv"
)

func isPalindrome(n int) bool {
	if n <= 0 {
		return false
	}
	str := strconv.Itoa(n)
	for i, l := 0, len(str); i < l/2; i++ {
		if str[i] != str[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	largest := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			n := i * j
			if isPalindrome(n) {
				if largest < n {
					largest = n
				}
			}
		}
	}
	log.Println(largest)
}
