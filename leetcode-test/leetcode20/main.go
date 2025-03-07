package main

import "log"

/*
给你一个整数数组 hours，表示以 小时 为单位的时间，返回一个整数，表示满足 i < j 且 hours[i] + hours[j] 构成 整天 的下标对 i, j 的数目。

整天 定义为时间持续时间是 24 小时的 整数倍 。

例如，1 天是 24 小时，2 天是 48 小时，3 天是 72 小时，以此类推。



示例 1：

输入： hours = [12,12,30,24,24]

输出： 2

解释：

构成整天的下标对分别是 (0, 1) 和 (3, 4)。

示例 2：

输入： hours = [72,48,24,3]

输出： 3

解释：

构成整天的下标对分别是 (0, 1)、(0, 2) 和 (1, 2)。



提示：

1 <= hours.length <= 100
1 <= hours[i] <= 109

*/

func main() {
	log.Println(countCompleteDayPairs([]int{12, 12, 30, 24, 24}))
}

func countCompleteDayPairs(hours []int) int {
	count := 0
	for i, v := range hours {
		for ii, vv := range hours[i+1:] {

			log.Printf(" i %d  ii %d v %d  vv %d  \n", i, ii, v, vv)
			log.Printf(" (v+vv)%%24  %d \n", (v+vv)%24)
			if (v+vv)%24 == 0 {
				log.Println("i++")
				count++
			}

		}
	}

	return count
}

// 中等难度 下解
func countCompleteDayPairs2(hours []int) int64 {
	// 余数(0 ---23 ) - 个数
	remainderCount := make([]int, 24)

	for _, v := range hours {
		remainder := v % 24
		remainderCount[remainder]++
	}

	count := int64(remainderCount[0] * (remainderCount[0] - 1))

	for i := 1; i < 24; i++ {
		if i == 24-i {
			count += int64(remainderCount[i] * (remainderCount[i] - 1))
		} else {
			count += int64(remainderCount[i] * remainderCount[24-i])
		}
	}

	count /= int64(2)
	return count
}
