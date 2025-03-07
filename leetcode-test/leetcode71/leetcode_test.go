package leetcode71

import (
	"reflect"
	"strconv"
	"testing"
)

//  切蛋糕的最小总开销 I

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f             func(int, int, []int, []int) int
		m             int
		n             int
		horizontalCut []int
		verticalCut   []int
		want          int
	}{
		{
			f:             minimumCost,
			m:             3,
			n:             2,
			horizontalCut: []int{1, 3},
			verticalCut:   []int{5},
			want:          13,
		},
		{
			f:             minimumCost,
			m:             2,
			n:             2,
			horizontalCut: []int{7},
			verticalCut:   []int{4},
			want:          15,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.m, v.n, v.horizontalCut, v.verticalCut)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}

}

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	ans := 0
	isHorizontal := false
	rawM, rawN := m, n

	for len(verticalCut)+len(horizontalCut) > 0 {
		maxFee := 0
		maxFeeIdx := -1
		for i := 0; i < len(horizontalCut); i++ {
			if maxFee < horizontalCut[i] {
				maxFee = horizontalCut[i]
				isHorizontal = true
				maxFeeIdx = i
			}
		}

		for i := 0; i < len(verticalCut); i++ {
			if maxFee < verticalCut[i] {
				maxFee = verticalCut[i]
				isHorizontal = false
				maxFeeIdx = i
			}
		}
		// if maxFeeIdx == -1 {
		// 	break
		// }

		if isHorizontal {
			n--
			ans += maxFee * (rawM - m + 1)
			horizontalCut = append(horizontalCut[:maxFeeIdx], horizontalCut[maxFeeIdx+1:]...)

		} else {
			m--
			ans += maxFee * (rawN - n + 1)
			verticalCut = append(verticalCut[:maxFeeIdx], verticalCut[maxFeeIdx+1:]...)
		}

	}

	return ans
}
