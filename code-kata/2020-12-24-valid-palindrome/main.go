package main

import "strings"

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
//
//示例 1:
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//示例 2:
//
//输入: "race a car"
//输出: false
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-palindrome
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isSymbolOrNumber(s int32) bool {
	if s >= 'a' && s <= 'z' {
		return true
	} else if s >= 'A' && s <= 'Z' {
		return true
	} else if s >= '0' && s <= '9' {
		return true
	}
	return false
}

func isValidPalindrome(s string) bool {
	s2 := ""
	for _, v := range s {
		if !isSymbolOrNumber(v) {
			continue
		}
		s2 += string(v)
	}
	s3 := strings.ToLower(s2)
	length := len(s3)
	for i := 0; i < length/2; i++ {
		if s3[i] != s3[length-i-1] {
			return false
		}
	}
	return true
}

func main() {}
