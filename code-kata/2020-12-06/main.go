package main

import (
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/google/go-cmp/cmp"
)

//125874 是一个特殊的数，它的两倍数 251748 包含了相同的数字，但是顺序不同。
//
//请找出最小的正整数 x，符合 2x，3x，4x，5x 和 6x 都包含相同的数字。
//
//tips：
//1. 直接使用暴力计算即可
//2. 效率的瓶颈应该是排序算法

func isFullyContained(n1, n2 int) bool {
	result := false
	s1 := strconv.Itoa(n1)
	s2 := strconv.Itoa(n2)
	var sort1, sort2 []string
	for _, v := range s1 {
		sort1 = append(sort1, string(v))
	}
	for _, v := range s2 {
		sort2 = append(sort2, string(v))
	}
	sort.Strings(sort1)
	sort.Strings(sort2)
	if cmp.Equal(sort1, sort2) {
		result = true
	}
	return result
}

func find() int {
	result := 0
	for {
		result++
		if !isFullyContained(result, result*2) {
			continue
		}
		if !isFullyContained(result, result*3) {
			continue
		}
		if !isFullyContained(result, result*4) {
			continue
		}
		if !isFullyContained(result, result*5) {
			continue
		}
		if !isFullyContained(result, result*6) {
			continue
		}
		log.Println(result, result*2, result*3, result*4, result*5, result*6)
		break
	}
	return result
}

func main() {
	t1 := time.Now()
	log.Println(find(), time.Since(t1))

	// 142857,285714,428571,571428,714285,857142
}
