package main

import (
	"fmt"
	"log"
	"math"
)

/*
633. 平方数之和
中等
相关标签
相关企业
给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。



示例 1：

输入：c = 5
输出：true
解释：1 * 1 + 2 * 2 = 5
示例 2：

输入：c = 3
输出：false


提示：

0 <= c <= 231 - 1
*/

func main() {
	// 2  = 1^2 +1 ^2  也属于 true
	// log.Println(judgeSquareSum(2))
	// log.Println(judgeSquareSum(4))
	log.Println(judgeSquareSum2(3))
	log.Println(judgeSquareSum(5))
	log.Println(judgeSquareSum2(5))
}

func judgeSquareSum2(c int) bool {

	// log.Println("math.Sqrt(float64(c)", math.Sqrt(float64(c)))
	limit := int(math.Sqrt(float64(c))) + 1

	start := 0
	log.Println("judgeSquareSum2 --", start, limit)
	for limit >= start {
		if diff := c - (limit * limit); diff == start*start {
			return true
		} else if diff-start*start > 0 {
			start++
		} else {
			limit--
		}

	}

	return false
}

var count int

// 当前 解法 在 这里有效 但是在leet code 针对 3 时返回 true
func judgeSquareSum(c int) bool {

	limit := math.Sqrt(float64(c))

	for i := int(limit); i > 0; i-- {
		// log.Printf("i %d  , c %d\n", i, c)
		if c >= (i * i) {
			c -= (i * i)
			fmt.Println("-- i", i)
			count++
			if c == 0 {
				return true
			} else if count == 2 {
				return false
			}
			return judgeSquareSum(c)
		}
	}
	return false
}
