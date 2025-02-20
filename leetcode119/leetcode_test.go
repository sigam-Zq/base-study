package leetcode119

import (
	"reflect"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f      func([][]int) int
		arrays [][]int
		want   int
	}{
		{
			f:      maxDistance,
			arrays: [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}},
			want:   4,
		},
		{
			f:      maxDistance,
			arrays: [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}},
			want:   4,
		},
		{
			f:      maxDistance,
			arrays: [][]int{{1, 4}, {0, 5}},
			want:   4,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.arrays)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func maxDistance(arrays [][]int) int {
	dis := 0

	for i, arr := range arrays {
		baseArrMax := arr[len(arr)-1]
		baseArrMin := arr[0]

		for j, arr2 := range arrays {
			if i == j {
				continue
			}
			dis = max(dis, abs(baseArrMax-arr2[0]))
			dis = max(dis, abs(baseArrMin-arr2[len(arr2)-1]))
		}
	}
	return dis
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func maxDistanceLeetCode1(arrays [][]int) int {
	res := 0
	n := len(arrays[0])
	minVal := arrays[0][0]
	maxVal := arrays[0][n-1]
	for i := 1; i < len(arrays); i++ {
		n = len(arrays[i])
		res = max(res, max(abs(arrays[i][n-1]-minVal),
			abs(maxVal-arrays[i][0])))
		minVal = min(minVal, arrays[i][0])
		maxVal = max(maxVal, arrays[i][n-1])
	}
	return res
}

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/maximum-distance-in-arrays/solutions/2393060/shu-zu-lie-biao-zhong-de-zui-da-ju-chi-b-f9x4/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
