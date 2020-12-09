package main

//旋转图像
//
//给定一个 n × n 的二维矩阵表示一个图像。
//
//将图像顺时针旋转 90 度。
//
//说明：
//
//你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
//
//示例 1:
//
//给定 matrix =
//[
//[1,2,3],
//[4,5,6],
//[7,8,9]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//[7,4,1],
//[8,5,2],
//[9,6,3]
//]
//示例 2:
//
//给定 matrix =
//[
//[ 5, 1, 9,11],
//[ 2, 4, 8,10],
//[13, 3, 6, 7],
//[15,14,12,16]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//[15,13, 2, 5],
//[14, 3, 4, 1],
//[12, 6, 8, 9],
//[16, 7,10,11]
//]
//
//https://leetcode-cn.com/classic/problems/rotate-image/description/

import (
	"log"
	"time"
)

func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	if n == 1 {
		return matrix
	}
	var i, j int
	// mirror by diagonal
	for i = 0; i < n; i++ {
		for j = i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// reverse the element order in each row
	for i = 0; i < n; i++ {
		for j = 0; j < n-j-1; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
	return matrix
}

func main() {
	matrix := [][]int{
		{1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3},
		{4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5},
	}

	t1 := time.Now()
	log.Println(rotate(matrix), time.Since(t1))
}
