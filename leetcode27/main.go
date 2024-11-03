package main

import (
	"fmt"
	"log"
)

/*
3226. 使两个整数相等的位更改次数
给你两个正整数 n 和 k。

你可以选择 n 的 二进制表示 中任意一个值为 1 的位，并将其改为 0。

返回使得 n 等于 k 所需要的更改次数。如果无法实现，返回 -1。



示例 1：

输入： n = 13, k = 4

输出： 2

解释：
最初，n 和 k 的二进制表示分别为 n = (1101)2 和 k = (0100)2，

我们可以改变 n 的第一位和第四位。结果整数为 n = (0100)2 = k。

示例 2：

输入： n = 21, k = 21

输出： 0

解释：
n 和 k 已经相等，因此不需要更改。

示例 3：

输入： n = 14, k = 13

输出： -1

解释：
无法使 n 等于 k。



提示：

1 <= n, k <= 106

*/

func main() {

	log.Println(minChanges(14, 13))
}

func minChanges(n int, k int) int {
	diff := n - k
	if diff == 0 {
		return 0
	} else if diff < 0 {
		return -1
	}

	// 获取大于等于n的开始节点
	start := 0
	i := 0
	for n >= start {
		start = 2 << i
		i++

		fmt.Println("in for i", i)
	}
	// 回撤一步
	fmt.Println("i", i)
	i--
	start = 2 << i

	// 获取n可以变化的位数
	var factorList []int
	for n > 0 {
		if n >= start {
			n -= start
			factorList = append(factorList, start)
		}
		i--
		fmt.Println("2 in for i", i)
		fmt.Println("2 in for n", n)
		fmt.Println("2 in for start", start)
		if i >= 0 {
			start = 2 << i
		} else {
			start = 1
		}
	}

	fmt.Println("factorList", factorList)
	fmt.Println("diff", diff)
	// 根据n 的组成 因子判断 diff 需要几个
	res := 0
	for _, v := range factorList {
		if diff >= v {
			res++
			diff -= v
		}

		if diff == 0 {
			return res
		}
	}

	return -1
}
