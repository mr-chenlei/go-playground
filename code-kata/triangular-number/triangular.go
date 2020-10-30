package main

import "log"

func Sum(n uint32) uint32 {
	sum := uint32(0)
	for i := uint32(0); i <= n; i++ {
		sum += i
	}
	return sum
}

func Triangular(t uint32) []uint32 {
	var list []uint32
	t1 := t / 2
	for i := uint32(1); i <= t1; i++ {
		if (t % i) == 0 {
			list = append(list, i)
		}
	}
	return list
}

func main() {
	list := Triangular(76576500)
	log.Println(list)
}
