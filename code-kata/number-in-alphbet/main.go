package main

import (
	"log"
	"strings"
)

var numberToLetter = map[int]string{
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	100:  "hundred",
	1000: "thousand",
}

func alphabetLength(s string) int {
	if s == "" {
		return 0
	}
	s = strings.Replace(s, "-", " ", -1)
	sub := strings.Split(s, " ")
	sum := 0
	for _, v := range sub {
		sum += len(v)
	}
	return sum
}

func fillTheGap() {
	for i := 0; i <= 1000; i++ {
		if _, ok := numberToLetter[i]; !ok {
			if i < 100 {
				n := i / 10 * 10
				latter := numberToLetter[n] + "-" + numberToLetter[i-n]
				numberToLetter[i] = latter
			} else if i > 100 {
				hundred := i / 100 * 100
				ten := (i - hundred) / 10 * 10
				one := i - hundred - ten
				var latter string
				if ten == 0 {
					latter = numberToLetter[hundred/100] + "-" + numberToLetter[hundred] + "-and-" + numberToLetter[one]
				} else if (ten + one) < 20 {
					latter = numberToLetter[hundred/100] + "-" + numberToLetter[hundred] + "-and-" + numberToLetter[ten+one]
				} else if one == 0 {
					latter = numberToLetter[hundred/100] + "-" + numberToLetter[hundred] + "-and-" + numberToLetter[ten]
				} else {
					latter = numberToLetter[hundred/100] + "-" + numberToLetter[hundred] + "-and-" + numberToLetter[ten] + "-" + numberToLetter[one]
				}
				numberToLetter[i] = latter
			}
		}
	}
}

func main() {
	fillTheGap()

	sum := 0
	for _, v := range numberToLetter {
		sum += alphabetLength(v)
	}
	log.Println(sum)
}
