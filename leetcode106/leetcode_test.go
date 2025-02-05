package leetcode106

import (
	"reflect"
	"strconv"
	"testing"
)

//922. 按奇偶排序数组 II

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func([]int) []int
		nums []int
		want []int
	}{
		{
			f:    sortArrayByParityII,
			nums: []int{4, 2, 5, 7},
			want: []int{4, 5, 2, 7},
		},
		{
			f:    sortArrayByParityII,
			nums: []int{2, 3},
			want: []int{2, 3},
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.nums)

			if reflect.DeepEqual(got, v.want) {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func sortArrayByParityII(nums []int) []int {
	n := len(nums)
	ansSingular := make([]int, n/2)
	ansEven := make([]int, n/2)
	sI, eI := 0, 0
	for i := 0; i < n; i++ {
		if nums[i]&1 == 0 {
			// 偶数
			ansEven[eI] = nums[i]
			eI++
		} else {
			// 奇数
			ansSingular[sI] = nums[i]
			sI++
		}
	}

	ans := make([]int, n)

	sI, eI = 0, 0
	for i := 0; i < n; i++ {
		if i&1 == 1 {
			// 奇数
			ans[i] = ansSingular[sI]
			sI++
		} else {
			// 奇数
			ans[i] = ansEven[eI]
			eI++
		}
	}

	return ans
}

func sortArrayByParityIILeetCode1(nums []int) []int {
	ans := make([]int, len(nums))
	i := 0
	for _, v := range nums {
		if v%2 == 0 {
			ans[i] = v
			i += 2
		}
	}
	i = 1
	for _, v := range nums {
		if v%2 == 1 {
			ans[i] = v
			i += 2
		}
	}
	return ans
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/sort-array-by-parity-ii/solutions/481450/an-qi-ou-pai-xu-shu-zu-ii-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func sortArrayByParityIILeetCode2(nums []int) []int {
	for i, j := 0, 1; i < len(nums); i += 2 {
		if nums[i]%2 == 1 {
			for nums[j]%2 == 1 {
				j += 2
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/sort-array-by-parity-ii/solutions/481450/an-qi-ou-pai-xu-shu-zu-ii-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
