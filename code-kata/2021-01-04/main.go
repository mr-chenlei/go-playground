package main

//给你一个字符串 s ，请你根据下面的算法重新构造字符串：
//
//从 s 中选出 最小 的字符，将它 接在 结果字符串的后面。
//从 s 剩余字符中选出 最小 的字符，且该字符比上一个添加的字符大，将它 接在 结果字符串后面。
//重复步骤 2 ，直到你没法从 s 中选择字符。
//从 s 中选出 最大 的字符，将它 接在 结果字符串的后面。
//从 s 剩余字符中选出 最大 的字符，且该字符比上一个添加的字符小，将它 接在 结果字符串后面。
//重复步骤 5 ，直到你没法从 s 中选择字符。
//重复步骤 1 到 6 ，直到 s 中所有字符都已经被选过。
//在任何一步中，如果最小或者最大字符不止一个 ，你可以选择其中任意一个，并将其添加到结果字符串。
//
//请你返回将 s 中字符重新排序后的 结果字符串 。
//
//示例 1：
//输入：s = "aaaabbbbcccc"
//输出："abccbaabccba"
//解释：第一轮的步骤 1，2，3 后，结果字符串为 result = "abc"
//第一轮的步骤 4，5，6 后，结果字符串为 result = "abccba"
//第一轮结束，现在 s = "aabbcc" ，我们再次回到步骤 1
//第二轮的步骤 1，2，3 后，结果字符串为 result = "abccbaabc"
//第二轮的步骤 4，5，6 后，结果字符串为 result = "abccbaabccba"
//
//示例 2：
//输入：s = "rat"
//输出："art"
//解释：单词 "rat" 在上述算法重排序以后变成 "art"
//
//示例 3：
//输入：s = "leetcode"
//输出："cdelotee"
//
//示例 4：
//输入：s = "ggggggg"
//输出："ggggggg"
//
//示例 5：
//输入：s = "spo"
//输出："ops"
//
//
//提示：
//
//1 <= s.length <= 500
//s 只包含小写英文字母。
//
//
//
//==================================================
//解体思路：
//这道题是让从字符串s中先选出升序的字符，然后再选出降序字符……一直这样循环，直到选完为止。因为题中的提示中说了s只包含小写英文字符，我们可以申请一个大小为26的数组，相当于26个桶。
//
//把s中的每个字符分别放到对应的桶里，比如a放到第一个桶里，b放到第2个桶里……。
//第1次从左往右遍历26个桶，从每个桶里拿出一个字符(如果没有就不用拿)
//第2次从右往左遍历26个桶，从每个桶里拿出一个字符(如果没有就不用拿)
//……
//一直这样循环下去，直到所有的桶里的元素都拿完为止。
//这里以示例为例，来画个图看下
//<1606272762-LnxjnQ-image.png>
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/increasing-decreasing-string
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func pop(s string, pos int) string {
	var result string
	if pos == len(s)-1 {
		result = s[:pos]
	} else {
		result = s[:pos] + s[pos+1:]
	}
	return result
}

func increasingDecreasingString(s string) string {
	var result string
	for len(s) > 0 {
		pos := 0
		min := s[0]
		if len(result) == 0 {
			min = 0
		} else {
			min = result[len(result)-1]
		}
		tmp := uint8(0)
		for i := 1; i < len(s); i++ {
			if min == 0 && tmp > s[i] {
				tmp = s[i]
				pos = i
			} else if s[i] > min && s[i] < tmp {
				tmp = s[i]
				pos = i
			}
		}
		s = pop(s, pos)
	}
	return result
}

func main() {}
