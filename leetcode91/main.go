package main

import "log"

func main() {
	log.Println(findClosestNumber([]int{-4, -2, 1, 4, 8}))
}

func findClosestNumber(nums []int) int {
	var ans int
	ans = nums[0]
	for i := 1; i < len(nums); i++ {

		if abs(nums[i]) < abs(ans) {
			ans = nums[i]
		} else if abs(nums[i]) == abs(ans) && nums[i] > ans {
			ans = nums[i]
		}

	}

	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
