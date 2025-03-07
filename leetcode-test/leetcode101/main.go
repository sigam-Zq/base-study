package main

import (
	"log"
	"sort"
)

// 350. 两个数组的交集 II

/*
给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。

示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]

提示：

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000

进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小，哪种方法更优？
如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
*/
func main() {
	log.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersect(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)
	nums1Maps := make(map[int]int)

	for i := range nums1 {
		nums1Maps[nums1[i]]++
	}

	for i := range nums2 {
		if fre, ok := nums1Maps[nums2[i]]; ok && fre != 0 {
			ans = append(ans, nums2[i])
			nums1Maps[nums2[i]]--
		}
	}

	return ans
}

func intersectLeetCode1(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}

	intersection := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			intersection = append(intersection, num)
			m[num]--
		}
	}
	return intersection
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/intersection-of-two-arrays-ii/solutions/327356/liang-ge-shu-zu-de-jiao-ji-ii-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func intersectLeetCode2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	index1, index2 := 0, 0

	intersection := []int{}
	for index1 < length1 && index2 < length2 {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			intersection = append(intersection, nums1[index1])
			index1++
			index2++
		}
	}
	return intersection
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/intersection-of-two-arrays-ii/solutions/327356/liang-ge-shu-zu-de-jiao-ji-ii-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
