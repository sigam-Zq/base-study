package main

import "log"

/*
3046. 分割数组
简单
相关标签
相关企业
提示
给你一个长度为 偶数 的整数数组 nums 。你需要将这个数组分割成 nums1 和 nums2 两部分，要求：

nums1.length == nums2.length == nums.length / 2 。
nums1 应包含 互不相同 的元素。
nums2也应包含 互不相同 的元素。
如果能够分割数组就返回 true ，否则返回 false 。



示例 1：

输入：nums = [1,1,2,2,3,4]
输出：true
解释：分割 nums 的可行方案之一是 nums1 = [1,2,3] 和 nums2 = [1,2,4] 。
示例 2：

输入：nums = [1,1,1,1]
输出：false
解释：分割 nums 的唯一可行方案是 nums1 = [1,1] 和 nums2 = [1,1] 。但 nums1 和 nums2 都不是由互不相同的元素构成。因此，返回 false 。


提示：

1 <= nums.length <= 100
nums.length % 2 == 0
1 <= nums[i] <= 100
*/

func main() {
	log.Println(isPossibleToSplitFix([]int{5, 9, 5, 5, 6, 8, 6, 1, 5, 7}))
}

/*
 互不相同的的元素是说 自身数组中包含互不相同元素， 而不是两个数组互相包含不相同元素

举例 输入 [5, 9, 5, 5, 6, 8, 6, 1, 5, 7]

nums1 = [9, 6, 1, 5, 7] nums2 = [5, 5, 6, 8, 5]

是指 num2 不满足 自身内部存在相同元素 5

而我错误理解成

nums1 包含对nums2 互不相同 的元素 1 9 nums2 包含对nums2 互不相同 的元素 8
*/

func isPossibleToSplitFix(nums []int) bool {

	nMap := make(map[int]int)
	for _, v := range nums {
		nMap[v]++
		if nMap[v] > 2 {
			return false
		}
	}
	return true
}

func isPossibleToSplit(nums []int) bool {

	nMap := make(map[int]int)

	for _, v := range nums {
		nMap[v]++
	}

	sigleLetter := 0

	for _, v := range nMap {
		if v == 1 {
			sigleLetter++
		}
	}
	log.Println(nMap)
	return (sigleLetter != 0 && sigleLetter&1 == 0)
}
