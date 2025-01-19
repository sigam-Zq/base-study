package leetcode88

import (
	"reflect"
	"strconv"
	"testing"
)

// 3287. 求出数组中最大序列值
/*


给你一个整数数组 nums 和一个 正 整数 k 。

定义长度为 2 * x 的序列 seq 的 值 为：

(seq[0] OR seq[1] OR ... OR seq[x - 1]) XOR (seq[x] OR seq[x + 1] OR ... OR seq[2 * x - 1]).
请你求出 nums 中所有长度为 2 * k 的
子序列
 的 最大值 。



示例 1：

输入：nums = [2,6,7], k = 1

输出：5

解释：

子序列 [2, 7] 的值最大，为 2 XOR 7 = 5 。

示例 2：

输入：nums = [4,2,5,6,7], k = 2

输出：2

解释：

子序列 [4, 5, 6, 7] 的值最大，为 (4 OR 5) XOR (6 OR 7) = 2 。



提示：

2 <= nums.length <= 400
1 <= nums[i] < 27
1 <= k <= nums.length / 2


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
			f:       maxValueFake,
			nums:    []int{2, 6, 7},
			k:       1,
			want:    5,
			isDebug: false,
		},
		{
			f:       maxValueFake,
			nums:    []int{4, 2, 5, 6, 7},
			k:       2,
			want:    2,
			isDebug: false,
		},
		{
			f:       maxValueFake,
			nums:    []int{1, 89, 11, 90},
			k:       2,
			want:    2,
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

func maxValueFake(nums []int, k int) int {
	n := len(nums)

	findOr := func(nums []int, k int) []map[int]bool {

		dp := make([]map[int]bool, 0)
		prev := make([]map[int]bool, k+1)

		for i := 0; i <= k; i++ {
			prev[i] = make(map[int]bool)
		}

		prev[0][0] = true

		for i := 0; i < n; i++ {

			for j := min(i+1, k-1); j >= 0; j-- {
				for x := range prev[j] {
					prev[j+1][nums[i]|x] = true
				}
			}

			current := make(map[int]bool)
			for key := range prev[k] {
				current[key] = true
			}
			dp = append(dp, current)
		}

		return dp
	}

	A := findOr(nums, k)
	revert(nums)
	B := findOr(nums, k)

	res := 0
	for i := k - 1; i < n-k; i++ {

		for a := range A[i] {
			for b := range B[n-i-2] {
				res = max(res, a^b)
			}
		}

	}

	return res
}

func revert(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func combination(n, k int) int {
	ans := 1
	for i := 0; i < k; i++ {
		ans *= (n - i)
		ans /= (i + 1)
	}
	return ans
}

func TestXor(t *testing.T) {
	t.Logf("%d\n", 1^8)
}

func TestCombination(t *testing.T) {
	t.Logf("%d\n", combination(5, 3))
}

func maxValueLeetCode(nums []int, k int) int {
	findORs := func(nums []int, k int) []map[int]bool {
		dp := make([]map[int]bool, 0)
		prev := make([]map[int]bool, k+1)
		for i := 0; i <= k; i++ {
			prev[i] = make(map[int]bool)
		}
		prev[0][0] = true

		for i := 0; i < len(nums); i++ {
			for j := min(k-1, i+1); j >= 0; j-- {
				for x := range prev[j] {
					prev[j+1][x|nums[i]] = true
				}
			}
			current := make(map[int]bool)
			for key := range prev[k] {
				current[key] = true
			}
			dp = append(dp, current)
		}
		return dp
	}

	A := findORs(nums, k)
	reverse(nums)
	B := findORs(nums, k)
	mx := 0

	for i := k - 1; i < len(nums)-k; i++ {
		for a := range A[i] {
			for b := range B[len(nums)-i-2] {
				if a^b > mx {
					mx = a ^ b
				}
			}
		}
	}
	return mx
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/find-the-maximum-sequence-value-of-array/solutions/3037275/qiu-chu-shu-zu-zhong-zui-da-xu-lie-zhi-b-bhnk/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
