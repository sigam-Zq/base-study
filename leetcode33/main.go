package main

import "sort"

/*

540. 有序数组中的单一元素

给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。

请你找出并返回只出现一次的那个数。

你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。



示例 1:

输入: nums = [1,1,2,3,3,4,4,8,8]
输出: 2
示例 2:

输入: nums =  [3,3,7,7,10,11,11]
输出: 10



提示:

1 <= nums.length <= 105
0 <= nums[i] <= 105

*/

func main() {

}

func singleNonDuplicate(nums []int) int {

	// 其实这里并不需要排序 - 题目中给到的例子就是有顺序的
	sort.Ints(nums)

	// 这里复杂度为 O(n) 本不满足条件的
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

/*
方法一：全数组的二分查找
思路和算法

假设只出现一次的元素位于下标 x，由于其余每个元素都出现两次，因此下标 x 的左边和右边都有偶数个元素，数组的长度是奇数。

由于数组是有序的，因此数组中相同的元素一定相邻。对于下标 x 左边的下标 y，如果 nums[y]=nums[y+1]，则 y 一定是偶数；对于下标 x 右边的下标 z，如果 nums[z]=nums[z+1]，则 z 一定是奇数。由于下标 x 是相同元素的开始下标的奇偶性的分界，因此可以使用二分查找的方法寻找下标 x。

初始时，二分查找的左边界是 0，右边界是数组的最大下标。每次取左右边界的平均值 mid 作为待判断的下标，根据 mid 的奇偶性决定和左边或右边的相邻元素比较：

如果 mid 是偶数，则比较 nums[mid] 和 nums[mid+1] 是否相等；

如果 mid 是奇数，则比较 nums[mid−1] 和 nums[mid] 是否相等。

如果上述比较相邻元素的结果是相等，则 mid<x，调整左边界，否则 mid≥x，调整右边界。调整边界之后继续二分查找，直到确定下标 x 的值。

得到下标 x 的值之后，nums[x] 即为只出现一次的元素。

细节

利用按位异或的性质，可以得到 mid 和相邻的数之间的如下关系，其中 ⊕ 是按位异或运算符：

当 mid 是偶数时，mid+1=mid⊕1；

当 mid 是奇数时，mid−1=mid⊕1。

因此在二分查找的过程中，不需要判断 mid 的奇偶性，mid 和 mid⊕1 即为每次需要比较元素的两个下标。

代码




*/

func singleNonDuplicateLeetCode(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == nums[mid^1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}
