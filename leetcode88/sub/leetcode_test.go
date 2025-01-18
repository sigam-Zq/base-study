package leetcode88

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

// 	3097. 或值至少为 K 的最短子数组 II
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

1 <= nums.length <= 2 * 105
0 <= nums[i] <= 109
0 <= k <= 109
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
		{
			f:       minimumSubarrayLength,
			nums:    []int{262144, 9575, 16013, 86, 10104, 8585, 1726, 10426, 5288, 7962, 5708, 7326, 16028, 11881, 12973, 609, 12557, 16372, 15991, 12024, 277, 7685, 1113, 10165, 12215, 15406, 6575, 4622, 3488, 9459, 293, 13893, 3556, 11003, 15779, 15036, 12832, 7502, 8327, 11065, 3783, 15305, 5311, 14704, 6087, 8442, 6783, 251, 15186, 15012, 3815, 8131, 3723, 11050, 14462, 12661, 6755, 4897, 9556, 6422, 5914, 546, 14238, 3547, 14378, 7508, 12745, 13571, 12427, 11653, 15607, 10486, 8423, 8491, 15640, 1173, 8718, 13097, 300, 8294, 14318, 10533, 4688, 8730, 8119, 14160, 7257, 14951, 7077, 11913, 1612, 7317, 13568, 9949, 9219, 13179, 13708, 14214, 11518, 4422, 12923, 1552, 14082, 2339, 15788, 12877, 7452, 2076, 2789, 10965, 9365, 4883, 3548, 11384, 15943, 11523, 8879, 15987, 8190, 15066, 2352, 1109, 10200, 9380, 14835, 13557, 16351, 15867, 2642, 13113, 3610, 5900, 3456, 3773, 11802, 12626, 839, 4933, 7717, 4156, 9294, 7973, 1456, 240, 6955, 2914, 2388, 7561, 10028, 3253, 7005, 1665, 3149, 7169, 1033, 13067, 2948, 4172, 9021, 1209, 12247, 13357, 9787, 8536, 4365, 8367, 2971, 2359, 2941, 13310, 10054, 4999, 9662, 6223, 6908, 12044, 13858, 6868, 4622, 9783, 7444, 3873, 10182, 5688, 13245, 15579, 12740, 12928, 6040, 2043, 6622, 4611, 15386, 6622, 10459, 13378, 14212, 12637, 4913, 4344, 8098},
			k:       285803,
			want:    -1,
			isDebug: false,
		},
		{
			f:       minimumSubarrayLength,
			nums:    []int{1, 2, 32, 21},
			k:       55,
			want:    3,
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

	maxV := 0
	c := 0
	for _, v := range nums {
		maxV = max(maxV, v)
		c |= v
	}
	// 满足最短情况
	if maxV >= k {
		return 1
	} else if c < k {
		return -1
	}

	// maxV 用二进制表示需要多长的数组
	maxL := 0
	for 1<<maxL < maxV {
		maxL++
	}

	bitsList := make([]int, maxL+2)
	n := len(nums)
	i, j := 0, 0
	res := math.MaxInt
	for ; j < n; j++ {
		addInBits(bitsList, nums[j])

		for i < j && bitsTransNum(bitsList) >= k {
			res = min(res, j-i+1)
			subtractBits(bitsList, nums[i])
			i++
		}

	}

	return res

}

func addInBits(bits []int, num int) {
	idx := 0
	for num>>idx > 0 {
		if (num>>idx)&1 == 1 {
			bits[idx]++
		}
		idx++
	}
}

func subtractBits(bits []int, num int) {
	idx := 0
	for num>>idx > 0 {
		if (num>>idx)&1 == 1 {
			bits[idx]--
		}
		idx++
	}

}

func bitsTransNum(bits []int) int {
	var ans int
	for i, v := range bits {
		if v > 0 {
			ans |= (1 << i)
		}
	}

	return ans
}

func TestOp(t *testing.T) {
	bitsList := make([]int, 10)
	addInBits(bitsList, 10)
	addInBits(bitsList, 1)
	t.Logf("%v \n", bitsList)

	t.Logf("%d \n", bitsTransNum(bitsList))
}

func minimumSubarrayLengthLeetCode(nums []int, k int) int {
	n := len(nums)
	bits := make([]int, 30)
	res := math.MaxInt32

	for left, right := 0, 0; right < n; right++ {
		for i := 0; i < 30; i++ {
			bits[i] += (nums[right] >> i) & 1
		}
		for left <= right && calc(bits) >= k {
			res = min(res, right-left+1)
			for i := 0; i < 30; i++ {
				bits[i] -= (nums[left] >> i) & 1
			}
			left++
		}
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func calc(bits []int) int {
	ans := 0
	for i := 0; i < len(bits); i++ {
		if bits[i] > 0 {
			ans |= 1 << i
		}
	}
	return ans
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/shortest-subarray-with-or-at-least-k-ii/solutions/3040101/huo-zhi-zhi-shao-wei-k-de-zui-duan-zi-sh-rzf8/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func Test1(t *testing.T) {

	t.Logf("%d \n", 21|32|2)
}
