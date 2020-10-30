package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

var (
	alphbet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

func sortInAlphabetical(intput []string) []string {
	sort.Strings(intput)
	return intput
}

func charToNumber(s string) int {
	sum := 0
	for k1 := range s {
		str := s[k1 : k1+1]
		for k2, v2 := range alphbet {
			if str == v2 {
				sum += (k2 + 1)
			}
		}
	}
	return sum
}

func main() {
	b, err := ioutil.ReadFile("/Users/c/Codes/go-playground/code-kata/alphabetical/names.txt")
	if err != nil {
		log.Panic(err)
	}
	list := strings.Split(string(b), ",")

	newList := sortInAlphabetical(list)
	for k, v := range newList {
		sum := charToNumber(v) * (k + 1)
		log.Println(v, k, sum)
	}
}
