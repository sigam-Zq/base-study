package main

import "log"

/*
3211. 生成不含相邻零的二进制字符串
给你一个正整数 n。

如果一个二进制字符串 x 的所有长度为 2 的
子字符串
中包含 至少 一个 "1"，则称 x 是一个 有效 字符串。

返回所有长度为 n 的 有效 字符串，可以以任意顺序排列。



示例 1：

输入： n = 3

输出： ["010","011","101","110","111"]

解释：

长度为 3 的有效字符串有："010"、"011"、"101"、"110" 和 "111"。

示例 2：

输入： n = 1

输出： ["0","1"]

解释：

长度为 1 的有效字符串有："0" 和 "1"。



提示：

1 <= n <= 18
*/

func main() {
	log.Println(validStrings(3))
}

const (
	Zero byte = '0'
	One  byte = '1'
)

func validStrings(n int) []string {

	res := []string{"0", "1"}

	for i := len(res[0]); i < n; i++ {
		for ii, v := range res {
			if v[len(v)-1] == Zero {
				res[ii] += "1"
			} else if v[len(v)-1] == One {
				res[ii] += "1"
				res = append(res, v+"0")
			}
		}
	}

	return res
}
