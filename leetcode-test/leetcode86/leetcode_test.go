package leetcode86

import (
	"reflect"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f       func([]int) int
		nums    []int
		want    int
		isDebug bool
	}{
		{
			f:       waysToSplitArray,
			nums:    []int{10, 4, -8, 7},
			want:    2,
			isDebug: false,
		},
		{
			f:       waysToSplitArray,
			nums:    []int{2, 3, 1, 0},
			want:    2,
			isDebug: false,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans := v.f(v.nums)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

func waysToSplitArray(nums []int) int {
	n := len(nums)
	sum := 0

	for _, v := range nums {
		sum += v
	}
	v1 := 0
	v2 := sum
	ans := 0
	for i := 0; i < n-1; i++ {
		v1 += nums[i]
		v2 -= nums[i]
		if v1 >= v2 {
			ans++
		}
	}

	return ans
}

// 2270. 分割数组的方案数

func waysToSplitArrayLeetCode(nums []int) int {
	n := len(nums)
	left, right := int64(0), int64(0)
	for _, num := range nums {
		right += int64(num)
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		left += int64(nums[i])
		right -= int64(nums[i])
		if left >= right {
			ans++
		}
	}
	return ans
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/number-of-ways-to-split-array/solutions/1501536/fen-ge-shu-zu-de-fang-an-shu-by-leetcode-3ygv/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
