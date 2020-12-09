package main

//验证回文字符串 Ⅱ
//
//
//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
//
//示例 1:
//输入: “aba"
//输出: True
//
//示例 2:
//输入: "abca"
//输出: True
//解释: 你可以删除c字符。
//
//注意:
//字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-palindrome-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isPalindrome(s string) bool {
	for i, l := 0, len(s); i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

func tryDeleteOne(s string) bool {
	for i := 0; i < len(s); i++ {
		tmp := ""
		if i == 0 {
			tmp = s[1:]
		} else {
			tmp = s[0:i]
			tmp += s[i+1:]
		}
		if tmp != "" && isPalindrome(tmp) {
			return true
		}
	}
	return false
}

func main() {

}
