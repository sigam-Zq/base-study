package main

import "log"

/*
3258. 统计满足 K 约束的子字符串数量 I
给你一个 二进制 字符串 s 和一个整数 k。

如果一个 二进制字符串 满足以下任一条件，则认为该字符串满足 k 约束：

字符串中 0 的数量最多为 k。
字符串中 1 的数量最多为 k。
返回一个整数，表示 s 的所有满足 k 约束 的
子字符串
的数量。



示例 1：

输入：s = "10101", k = 1

输出：12

解释：

s 的所有子字符串中，除了 "1010"、"10101" 和 "0101" 外，其余子字符串都满足 k 约束。

示例 2：

输入：s = "1010101", k = 2

输出：25

解释：

s 的所有子字符串中，除了长度大于 5 的子字符串外，其余子字符串都满足 k 约束。

示例 3：

输入：s = "11111", k = 1

输出：15

解释：

s 的所有子字符串都满足 k 约束。



提示：

1 <= s.length <= 50
1 <= k <= s.length
s[i] 是 '0' 或 '1'。
*/

func main() {
	log.Println(countKConstraintSubstrings("10101", 1))
}
func countKConstraintSubstrings(s string, k int) int {
	if len(s) == 1 {
		if k == 1 {
			return 1
		}
		return 0
	}

	leftIdx, rightIdx := 0, 1

	res := 0

	for rightIdx < len(s) || leftIdx < len(s) {

		log.Printf(" sub Str idx %d  %d \n", leftIdx, rightIdx)
		if k >= countFactor(s[leftIdx:rightIdx]) {
			res++
		} else {
			leftIdx++
			rightIdx = leftIdx + 1
			continue
		}

		if rightIdx < len(s) {
			rightIdx++
		} else if leftIdx < len(s) {
			leftIdx++
			rightIdx = leftIdx + 1
		}

		log.Printf("after sub Str idx %d  %d  len(s) %d\n", leftIdx, rightIdx, len(s))
	}

	return res
}

// 检查 字符串中 1 和0 的最小
func countFactor(s string) int {
	log.Printf(" sub Str %s \n", s)
	oneCount := 0
	zeroCount := 0
	for _, v := range s {
		if v == '1' {
			oneCount++
		} else if v == '0' {
			zeroCount++
		}
	}

	if oneCount < zeroCount {
		return oneCount
	} else {
		return zeroCount
	}
}
