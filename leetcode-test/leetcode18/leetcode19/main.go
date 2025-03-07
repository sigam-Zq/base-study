package main

import "log"

func main() {
	log.Println(minOperations([]int{0, 1, 1, 0, 1}))
}

func minOperations(nums []int) int {

	left := 0
	i := 0
	for left < len(nums) {

		if i%2 == 0 {
			// 双数
			if nums[left] == 0 {
				nums[left] = reversal(nums[left])
				i++
			}
		} else {
			// 单数
			if reversal(nums[left]) == 0 {
				i++
			}
		}
		left++
	}
	return i
}

func reversal(in int) (out int) {
	if in == 0 {
		return 1
	} else {
		return 0
	}
}
