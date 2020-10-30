package main

import "log"

func monthConvert(y, m int) (int, int) {
	y2 := y
	m2 := m
	if m < 3 {
		switch m {
		case 2:
			m2 = 14
		case 1:
			m2 = 13
		}
		y2 = y - 1
	}
	return y2, m2
}

func isSunday(y, m, d int) bool {
	y2, m2 := monthConvert(y, m)
	c := y2 / 100
	y2 %= 100
	w := (y2 + y2/4 + c/4 - 2*c + (26 * (m2 + 1) / 10) + d - 1) % 7
	if w == 0 {
		return true
	}
	return false
}

func main() {
	sum := 0
	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			if isSunday(y, m, 1) {
				sum++
			}
		}
	}
	log.Println(sum)
}
