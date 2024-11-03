package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	log.Println(smallestRangeII([]int{7, 8, 8, 5, 2}, 4))
	fmt.Println("--------------------")
	log.Println(smallestRangeIITwo([]int{7, 8, 8, 5, 2}, 4))
	fmt.Println("--------------------")
	log.Println(smallestRangeIIThree([]int{7, 8, 8, 5, 2}, 4))
}

func smallestRangeII(nums []int, k int) int {
	sum := 0
	minO, maxO := nums[0], nums[0]
	for _, v := range nums {
		sum += v

		if v > maxO {
			maxO = v
		}

		if v < minO {
			minO = v
		}
	}
	avg := float64(sum) / float64(len(nums))

	fistDiff := maxO - minO

	min, max := maxO, minO
	for _, v := range nums {
		tmp := 0
		if float64(v) > avg {
			tmp = v - k
		} else {
			tmp = v + k
		}

		if tmp > max {
			max = tmp
		}

		if tmp < min {
			min = tmp
		}

	}
	log.Printf("max  %d   min %d \n", max, min)
	twoDiff := max - min

	if fistDiff < twoDiff {
		return fistDiff
	}
	return twoDiff
}

func smallestRangeIITwo(nums []int, k int) int {

	sort.Ints(nums)
	min, max := nums[len(nums)-1], nums[0]
	for i, v := range nums {
		tmp := 0
		if i < len(nums)/2 {
			tmp = v + k
		} else {
			tmp = v - k
		}

		if tmp > max {
			max = tmp
		}

		if tmp < min {
			min = tmp
		}

	}

	return max - min
}

func smallestRangeIIThree(nums []int, k int) int {

	// 对数组排序
	sort.Ints(nums)
	n := len(nums)

	// 原始差值
	result := nums[n-1] - nums[0]

	// 遍历所有可能的分界点
	for i := 0; i < n-1; i++ {
		maxVal := max(nums[n-1]-k, nums[i]+k) // 最大值在nums[i]和nums[n-1]之间
		minVal := min(nums[0]+k, nums[i+1]-k) // 最小值在nums[0]和nums[i+1]之间
		result = min(result, maxVal-minVal)   // 更新最小差值
	}

	return result
	return 1
}

// 辅助函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
