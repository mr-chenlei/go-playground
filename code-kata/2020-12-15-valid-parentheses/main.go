package main

//给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串，判断字符串是否有效。
//有效字符串需满足：
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//
//示例 1:
//输入: "()"
//输出: true

//示例2:
//输入: "()[]{}"
//输出: true

//示例3:
//输入: "(]"
//输出: false

//示例4:
//输入: "([)]"
//输出: false

//示例5:
//输入: "{[]}"
//输出: true
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-parentheses

func isValidParentheses(s string, pos int, next byte) bool {
	length := len(s)
	if next != 0 {
		for i := pos; i < length; i++ {
			switch s[i] {
			case '{':
				if !isValidParentheses(s, i+1, '}') {
					return false
				}
				i = i + 2
			case '[':
				if !isValidParentheses(s, i+1, ']') {
					return false
				}
				i = i + 2
			case '(':
				if !isValidParentheses(s, i+1, ')') {
					return false
				}
				i = i + 2
			}

			if s[i] == next {
				return true
			} else if s[i] == '}' || s[i] == ']' || s[i] == ')' {
				return false
			}
		}
	} else {
		for i := pos; i < length; i++ {
			switch s[i] {
			case '{':
				if !isValidParentheses(s, i+1, '}') {
					return false
				}
			case '[':
				if !isValidParentheses(s, i+1, ']') {
					return false
				}
			case '(':
				if !isValidParentheses(s, i+1, ')') {
					return false
				}
			}
		}
	}
	return true
}

func main() {

}
