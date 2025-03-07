package leetcode111

import (
	"fmt"
	"strconv"
	"testing"
)

// 80. 删除有序数组中的重复项 II

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func([]int) int
		nums []int
		want int
	}{
		{
			f:    removeDuplicatesLeetCode,
			nums: []int{1, 1, 1, 2, 2, 3},
			want: 5,
		},
		{
			f:    removeDuplicatesLeetCode,
			nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want: 7,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.nums)
			if got != v.want {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func removeDuplicates(nums []int) int {

	cnt, flag := 1, nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == flag {
			cnt++
			if cnt > 2 {
				nums = append(nums[:i], nums[i+1:]...)
			}
		} else {
			cnt = 1
			flag = nums[i]
		}
	}
	// fmt.Println(nums)
	return len(nums)
}

func removeDuplicatesLeetCode(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	fmt.Println(nums)
	return slow
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/solutions/702644/shan-chu-pai-xu-shu-zu-zhong-de-zhong-fu-yec2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
