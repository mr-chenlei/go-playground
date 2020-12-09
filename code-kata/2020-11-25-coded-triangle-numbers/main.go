package main

//Coded triangle numbers
//
//第 n 个三角数的计算方式是  tn = ½n(n+1) ；前十个三角数分别是：
//1, 3, 6, 10, 15, 21, 28, 36, 45, 55, …
//
//将一个单词中的每个字母转换为对应位置的数字，然后将他们相加，可以得到一个单词的值。比如 SKY 是 19+11+25 = 55 = t10.
//如果一个单词的值是一个三角数，我们称他为 三角词
//
//使用 words.txt，一个 16K 的文本文件，包含近2000个常用英语词汇，里面有多少个三角词？

import (
	"io/ioutil"
	"log"
	"strings"
)

func isTriangleNumber(n int) bool {
	index := 1
	for {
		t := 0.5 * float64(index*(index+1))
		if t == float64(n) {
			return true
		} else if t > float64(n) {
			break
		}
		index++
	}
	return false
}

func isTriangleString(input string) bool {
	sum := int32(0)
	for _, v := range input {
		if v >= 'a' && v <= 'z' {
			sum += v - 'a' + 1
		} else if v >= 'A' && v <= 'Z' {
			sum += v - 'A' + 1
		}
	}
	return isTriangleNumber(int(sum))
}

func main() {
	var err error
	var data []byte
	if data, err = ioutil.ReadFile("./words.txt"); err != nil {
		log.Panic(err)
	}

	after := strings.Split(strings.ReplaceAll(string(data), "\"", ""), ",")
	num := 0
	for _, v := range after {
		if isTriangleString(v) {
			num++
		}
	}
	log.Println("triangle number:", num)
}
