package main

import (
	"log"
	"math"
	"time"
)

//找到唯一的正整数，其平方的形式为1_2_3_4_5_6_7_8_9_0， 其中每个“ _”是一个数字。
//
//Tips:
//1. 平方以0结尾的数字，其本身以0结尾。即，需要找的正整数最后一位为0，可以找到平方为1_2_3_4_5_6_7_8_9的数字后再乘以10
//2. 平方以9结尾的数字，其本身只能以3或7结尾。

func find(start, end uint64) uint64 {
	var result uint64
	for i := end; i >= start; i-- {
		integer, frac := math.Modf(math.Sqrt(float64(i)))
		if frac == 0 {
			result = uint64(integer)
			break
		}
	}
	return result
}
func main() {
	t1 := time.Now()
	log.Println(find(102030405060708090, 192939495969798999), time.Since(t1))
}
