package main

import (
	"strings"
)

//最大单词长度乘积
//
//
//给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。
//
//示例 1:
//
//输入: ["abcw","baz","foo","bar","xtfn","abcdef"]
//输出: 16
//解释: 这两个单词为 "abcw", "xtfn"。
//示例 2:
//
//输入: ["a","ab","abc","d","cd","bcd","abcd"]
//输出: 4
//解释: 这两个单词为 "ab", "cd"。
//示例 3:
//
//输入: ["a","aa","aaa","aaaa"]
//输出: 0
//解释: 不存在这样的两个单词。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-product-of-word-lengths
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func maxProduct(input []string) int {
	var result int
	length := len(input)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if !strings.ContainsAny(input[i], input[j]) {
				swap := len(input[i]) * len(input[j])
				if swap > result {
					result = swap
				}
			}
		}
	}
	return result
}

func maxProductV2(input []string) int {
	var result int
	checker := make(map[string]map[int32]struct{}, 0)
	for _, v := range input {
		for _, v2 := range v {
			if _, ok := checker[v][v2]; !ok {
				checker[v] = make(map[int32]struct{}, 0)
				checker[v][v2] = struct{}{}
			}
		}
	}
	for k, v := range checker {
		isExited := false
		for _, v2 := range input {
			if k == v2 {
				continue
			}
			for _, v3 := range v2 {
				if _, ok := v[v3]; ok {
					isExited = true
					break
				}
			}
			if isExited == false {
				swap := len(k) * len(v2)
				if swap > result {
					result = swap
				}
			}
		}
	}
	return result
}

func main() {

}
