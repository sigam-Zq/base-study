package leetcode88

import (
	"log"
	"math"
	"reflect"
	"strconv"
	"testing"
)

// 	3095. 或值至少 K 的最短子数组 I
/*
给你一个 非负 整数数组 nums 和一个整数 k 。

如果一个数组中所有元素的按位或运算 OR 的值 至少 为 k ，那么我们称这个数组是 特别的 。

请你返回 nums 中 最短特别非空
子数组
的长度，如果特别子数组不存在，那么返回 -1 。



示例 1：

输入：nums = [1,2,3], k = 2

输出：1

解释：

子数组 [3] 的按位 OR 值为 3 ，所以我们返回 1 。

注意，[2] 也是一个特别子数组。

示例 2：

输入：nums = [2,1,8], k = 10

输出：3

解释：

子数组 [2,1,8] 的按位 OR 值为 11 ，所以我们返回 3 。

示例 3：

输入：nums = [1,2], k = 0

输出：1

解释：

子数组 [1] 的按位 OR 值为 1 ，所以我们返回 1 。



提示：

1 <= nums.length <= 50
0 <= nums[i] <= 50
0 <= k < 64
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f       func([]int, int) int
		nums    []int
		k       int
		want    int
		isDebug bool
	}{
		{
			f:       minimumSubarrayLength,
			nums:    []int{1, 2, 3},
			k:       2,
			want:    1,
			isDebug: false,
		},
		{
			f:       minimumSubarrayLength,
			nums:    []int{2, 1, 8},
			k:       10,
			want:    3,
			isDebug: false,
		},
		{
			f:       minimumSubarrayLength,
			nums:    []int{1, 2},
			k:       0,
			want:    1,
			isDebug: false,
		},
		{
			f:       minimumSubarrayLength,
			nums:    []int{32, 2, 24, 1},
			k:       35,
			want:    3,
			isDebug: false,
		},
		{
			f:       minimumSubarrayLength,
			nums:    []int{32, 1, 25, 11, 2},
			k:       59,
			want:    4,
			isDebug: false,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans := v.f(v.nums, v.k)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

func minimumSubarrayLength(nums []int, k int) int {

	// sort.Ints(nums)
	n := len(nums)
	var res int
	res = math.MaxInt
	i, j := 0, 0
	for i < n {
		j++
		if OrOp(nums[i:j]) >= k {
			res = min(res, j-i)
		}
		if j == n {
			i++
			j = i
		}
	}
	if res == math.MaxInt {
		return -1
	}

	return res
}

func OrOp(num []int) int {
	log.Printf("\n\n")
	var ans int
	for _, v := range num {
		ans |= v
	}
	log.Printf("%v %d \n", num, ans)
	return ans
}
