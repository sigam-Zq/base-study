package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f    func([]int) int
		arr  []int
		want int
	}{
		{
			f:    findSpecialInteger,
			arr:  []int{1, 2, 2, 6, 6, 6, 6, 7, 10},
			want: 6,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.arr)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func findSpecialInteger(arr []int) int {
	n := len(arr)

	if n == 1 {
		return arr[0]
	}

	cnt, maxNum := 1, arr[0]

	for i := 1; i < n; i++ {
		if arr[i] == maxNum {
			cnt++
			if cnt > n/4 {
				return arr[i]
			}
		} else {
			cnt, maxNum = 1, arr[i]
		}
	}

	return 0
}

func findSpecialIntegerLeetCode(arr []int) int {
	n := len(arr)
	cur := arr[0]
	cnt := 0
	for i := 0; i < n; i++ {
		if arr[i] == cur {
			cnt++
			if cnt*4 > n {
				return cur
			}
		} else {
			cur = arr[i]
			cnt = 1
		}
	}
	return -1
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/element-appearing-more-than-25-in-sorted-array/solutions/101725/you-xu-shu-zu-zhong-chu-xian-ci-shu-chao-guo-25d-3/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
