package leetcode100

import (
	"reflect"
	"strconv"
	"testing"
)

// 219. 存在重复元素 II
/*

给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。



示例 1：

输入：nums = [1,2,3,1], k = 3
输出：true
示例 2：

输入：nums = [1,0,1,1], k = 1
输出：true
示例 3：

输入：nums = [1,2,3,1,2,3], k = 2
输出：false




提示：

1 <= nums.length <= 105
-109 <= nums[i] <= 109
0 <= k <= 105
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func([]int, int) bool
		nums []int
		k    int
		want bool
	}{
		{
			f:    containsNearbyDuplicate,
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			f:    containsNearbyDuplicate,
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			f:    containsNearbyDuplicate,
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.nums, v.k)

			if !reflect.DeepEqual(got, v.want) {
				t.Errorf(" got %v  want %v \n", got, v.want)
			}
		})

	}
}

func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)

	var res bool
Lab1:
	for i := range nums {
		for j := i + 1; j < n && j <= i+k; j++ {
			if nums[i] == nums[j] {
				res = true
				break Lab1
			}
		}
	}
	return res
}

func containsNearbyDuplicateLeetCode(nums []int, k int) bool {
	pos := map[int]int{}
	for i, num := range nums {
		if p, ok := pos[num]; ok && i-p <= k {
			return true
		}
		pos[num] = i
	}
	return false
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/contains-duplicate-ii/solutions/1218075/cun-zai-zhong-fu-yuan-su-ii-by-leetcode-kluvk/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func containsNearbyDuplicateLeetCode2(nums []int, k int) bool {
	set := map[int]struct{}{}
	for i, num := range nums {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/contains-duplicate-ii/solutions/1218075/cun-zai-zhong-fu-yuan-su-ii-by-leetcode-kluvk/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
