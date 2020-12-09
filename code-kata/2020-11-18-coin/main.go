package main

//你总共有 n 枚硬币，你需要将它们摆成一个阶梯形状，第 k 行就必须正好有 k 枚硬币。
//给定一个数字 n，找出可形成完整阶梯行的总行数。
//n 是一个非负整数，并且在32位有符号整型的范围内。
//请给出尽可能高效的解法。
//
//示例 1:
//n = 5
//硬币可排列成以下几行:
//¤
//¤ ¤
//¤ ¤
//因为第三行不完整，所以返回2.
//
//示例 2:
//n = 8
//硬币可排列成以下几行:
//¤
//¤ ¤
//¤ ¤ ¤
//¤ ¤
//因为第四行不完整，所以返回3.
//
//可能会用到的：
//1. 等差数列通项公式：an=a1+(n-1)*d。首项a1=1，公差d=2 ???
//2. 等差数列前n项和：Sn=[n*(a1+an)]/2

func doMathV1(n int) int {
	var result, sum int
	for i := 1; i <= n; i++ {
		sum += i
		next := n - sum
		if (n-sum) > 0 && (next <= i) {
			result = i
			break
		}
	}
	return result
}

func doMathV2(n int) int {
	var result, sum int
	for i := 1; i <= n; i++ {
		sum += i
		next := n - sum
		if next <= i {
			result = i
			break
		}
	}
	return result
}

func main() {
}
