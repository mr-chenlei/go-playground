package main

import "log"

//仅仅反转字母
//
//给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。
//
//示例 1：
//输入："ab-cd"
//输出："dc-ba”
//
//示例 2：
//输入："a-bC-dEf-ghIj"
//输出："j-Ih-gfE-dCba”
//
//示例 3：
//输入："Test1ng-Leet=code-Q!"
//输出："Qedo1ct-eeLg=ntse-T!"
//
//
//提示：
//
//• S.length <= 100
//• 33 <= S[i].ASCIIcode <= 122
//• S 中不包含 \ or "
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-only-letters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func reverse(input string) string {
	var result, tmp string
	for k := range input {
		v := input[k]
		if (v >= 64 && v <= 90) || (v >= 97 && v <= 122) {
			tmp += string(v)
		}
	}

	runes := []rune(tmp)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	tmp = string(runes)
	index := 0
	for k, v := range input {
		if (v >= 64 && v <= 90) || (v >= 97 && v <= 122) {
			result += string(tmp[index])
			index++
		} else {
			result += string(input[k])
		}
	}
	return result
}

func main() {
	log.Println("Test1ng-Leet=code-Q!")
}
