package main

//给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。
//
//让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：
//
//A.length >= 3
//在 0 < i < A.length - 1 条件下，存在 i 使得：
//A[0] < A[1] < ... A[i-1] < A[i]
//A[i] > A[i+1] > ... > A[A.length - 1]
//
//<C0F55369-89BF-4F32-A351-80F3CA975974.png>
//
//示例 1：
//
//输入：[2,1]
//输出：false
//示例 2：
//
//输入：[3,5,5]
//输出：false
//示例 3：
//
//输入：[0,3,2,1]
//输出：true
//
//
//提示：
//
//0 <= A.length <= 10000
//0 <= A[i] <= 10000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-mountain-array

func isMountainArray(input []int) bool {
	length := len(input)
	max := 0
	pos := 0
	swap := -1
	// Find max number and it's pos
	for k, v := range input {
		if max < v {
			pos = k
			max = v
		}
	}
	if max == input[0] || max == input[length-1] {
		return false
	}
	// Compare from beginning to max number pos
	for _, v := range input[:pos] {
		if swap >= v {
			return false
		}
		swap = v
	}
	// Compare from max number pos to the end
	swap = input[pos]
	for _, v := range input[pos+1:] {
		if swap <= v {
			return false
		}
		swap = v
	}

	return true
}

func main() {}
