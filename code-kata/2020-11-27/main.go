package main

import (
	"log"
	"strconv"
)

//给你一个字符串 path，其中 path[i] 的值可以是 'N'、'S'、'E' 或者 'W'，分别表示向北、向南、向东、向西移动一个单位。
//机器人从二维平面上的原点 (0, 0) 处开始出发，按 path 所指示的路径行走。
//如果路径在任何位置上出现相交的情况，也就是走到之前已经走过的位置，请返回 True ；否则，返回 False 。
//
//
//示例 1：
//
//输入：path = "NES"
//输出：false
//解释：该路径没有在任何位置相交。
//
//示例 2：
//
//输入：path = "NESWW"
//输出：true
//解释：该路径经过原点两次。
//
//
//提示：
//
//1 <= path.length <= 10^4
//path 仅由 {'N', 'S', 'E', 'W} 中的字符组成
//
//解题思路：
//记下每次走过的坐标，如果出现重复，则路径相交。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/path-crossing
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func direction2Coord(d string) (x, y int) {
	switch d {
	case "E":
		x = -1
		y = 0
	case "S":
		x = 0
		y = 1
	case "W":
		x = 1
		y = 0
	case "N":
		x = 0
		y = -1
	}
	return x, y
}

func isPathCrossed(path []string) bool {
	x := 0
	y := 0
	pathLog := make(map[string]struct{}, 0)
	pos := strconv.Itoa(x) + strconv.Itoa(y)
	pathLog[pos] = struct{}{}
	for _, v := range path {
		x2, y2 := direction2Coord(v)
		x += x2
		y += y2
		pos := strconv.Itoa(x) + strconv.Itoa(y)
		if _, ok := pathLog[pos]; !ok {
			pathLog[pos] = struct{}{}
		} else {
			return ok
		}
	}
	return false
}

func main() {
	log.Println([]string{"N", "E", "S", "W", "W"})
}
