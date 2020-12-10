package main

import "strings"

//请你设计一个可以解释字符串 command 的 Goal 解析器 。command 由 "G"、"()" 和/或 "(al)" 按某种顺序组成。Goal 解析器会将 "G" 解释为字符串 "G"、"()" 解释为字符串 "o" ，"(al)" 解释为字符串 "al" 。然后，按原顺序将经解释得到的字符串连接成一个字符串。
//
//给你字符串 command ，返回 Goal 解析器 对 command 的解释结果。
//不要用replace方法
//
//
//示例 1：
//
//输入：command = "G()(al)"
//输出："Goal"
//解释：Goal 解析器解释命令的步骤如下所示：
//G -> G
//() -> o
//(al) -> al
//最后连接得到的结果是 "Goal"
//示例 2：
//
//输入：command = "G()()()()(al)"
//输出："Gooooal"
//示例 3：
//
//输入：command = "(al)G(al)()()G"
//输出："alGalooG"
//
//
//提示：
//
//1 <= command.length <= 100
//command 由 "G"、"()" 和/或 "(al)" 按某种顺序组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/goal-parser-interpretation
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func parseString(s string) string {
	var result string
	length := len(s)
	pos := 0
	for i := 0; i < length; i++ {
		pos2 := strings.Index(s[pos:], "G") + pos
		pos3 := strings.Index(s[pos:], "()") + pos
		pos4 := strings.Index(s[pos:], "(al)") + pos
		if (pos2 != -1) && (pos >= pos2) {
			pos = pos2
		} else if (pos3 != -1) && (pos >= pos3) {
			pos = pos3
		} else if (pos4 != -1) && (pos >= pos4) {
			pos = pos4
		}

		if pos == pos2 {
			result += "G"
			pos += 1
		} else if pos == pos3 {
			result += "o"
			pos += 2
		} else if pos == pos4 {
			result += "al"
			pos += 3
		}
	}
	return result
}

func main() {}
