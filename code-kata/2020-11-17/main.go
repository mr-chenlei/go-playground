package main

//56. 合并区间
//
//给出一个区间的集合，请合并所有重叠的区间。
//
//
//示例 1:
//
//输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2:
//
//输入: intervals = [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
//注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。
//
//提示：
//intervals[i][0] <= intervals[i][1]

func sortArray(input [][]int) [][]int {
	length := len(input)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if input[i][0] > input[j][0] {
				tmp := input[j]
				input[j] = input[i]
				input[i] = tmp
			}
		}
	}
	return input
}

func combine(input [][]int) [][]int {
	var result [][]int
	after := sortArray(input)
	length := len(after)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			a2 := after[j][0]
			b1 := after[i][1]
			b2 := after[j][1]
			if (b1 > a2) && (b1 < b2) {
				result = append(result, []int{after[i][0], after[j][1]})
				break
			} else {
				result = append(result, []int{after[j][0], after[j][1]})
			}
		}
	}
	return result
}

func main() {}
