package main

import (
	"log"
	"strconv"
)

//二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。
//
//每个 LED 代表一个 0 或 1，最低位在右侧。
//Hour：            x x x x
//
//Minute:      x x x x x x
//
//例如，上面的二进制手表（红色代表置位）读取 “3:25”。（可参考原题图片）
//
//给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。
//
//示例：
//
//输入: n = 1
//返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
//
//
//提示：
//
//输出的顺序没有要求。
//小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
//分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。
//超过表示范围（小时 0-11，分钟 0-59）的数据将会被舍弃，也就是说不会出现 "13:00", "0:61" 等时间。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-watch
//
//
//
//题解提示：
//官方题解为回溯算法，需进行递归与剪枝。
//快速解法可暴力计算时间后对比确认。

func binaryWatch(n int) []string {
	var result []string
	for i := 0; i < n; i++ {
		minute := (59 << i) & 59
		result = append(result, strconv.Itoa(minute))
	}
	return result
}

func main() {
	log.Println(binaryWatch(1))
}
