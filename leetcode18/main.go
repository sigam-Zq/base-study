package main

import "log"

func main() {
	log.Println(minOperations([]int{0, 1, 1, 1, 0, 0}))
}
func minOperations(nums []int) int {
	left := 0
	i := 0
	for left <= len(nums)-3 {
		if nums[left] == 0 {
			nums[left] = reversal(nums[left])
			nums[left+1] = reversal(nums[left+1])
			nums[left+2] = reversal(nums[left+2])
			i++
		}
		left++
		log.Printf(" i  %d 次 num %v \n ", i, nums)
	}
	log.Printf("最后  num %v \n ", nums)
	for _, v := range nums {
		if v == 0 {
			return -1
		}
	}

	return i
}

func reversal(source int) (target int) {
	if source == 0 {
		return 1
	} else {
		return 0
	}
}
