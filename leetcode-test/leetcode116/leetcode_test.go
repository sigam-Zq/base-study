package leetcode116

import (
	"reflect"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f    func([]int) []int
		arr  []int
		want []int
	}{
		{
			f:    replaceElements,
			arr:  []int{17, 18, 5, 4, 6, 1},
			want: []int{18, 6, 6, 6, 1, -1},
		},
		{
			f:    replaceElements,
			arr:  []int{400},
			want: []int{-1},
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

func replaceElements(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	ans[n-1] = -1

	for i := n - 2; i >= 0; i-- {
		maxNum := 0
		for j := i + 1; j < n; j++ {
			maxNum = max(maxNum, arr[j])
		}

		ans[i] = maxNum
	}

	return ans
}

func replaceElementsLeetCode(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	ans[n-1] = -1
	for i := n - 2; i >= 0; i-- {
		ans[i] = max(ans[i+1], arr[i+1])
	}
	return ans
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/replace-elements-with-greatest-element-on-right-side/solutions/101750/jiang-mei-ge-yuan-su-ti-huan-wei-you-ce-zui-da-y-5/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
