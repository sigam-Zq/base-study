package leetcode

import (
	"sort"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f    func(int, []int) int
		n    int
		cuts []int
		want int
	}{
		{
			f:    minCost,
			n:    7,
			cuts: []int{1, 3, 4, 5},
			want: 16,
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.n, v.cuts); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}

		})

	}
}

func minCost(n int, cuts []int) int {
	sort.Ints(cuts)

	return 0
}
