package leetcode10

import (
	"testing"
)

func TestMain(t *testing.T) {
	args := []struct {
		name    string
		f       func([]int, int) int
		tickets []int
		k       int
		want    int
	}{
		{
			name:    "one",
			f:       timeRequiredToBuy,
			tickets: []int{5, 1, 1, 1},
			k:       0,
			want:    8,
		},
		{
			name:    "two",
			f:       timeRequiredToBuy,
			tickets: []int{2, 3, 2},
			k:       2,
			want:    6,
		},
	}

	for _, v := range args {
		t.Run(v.name, func(t *testing.T) {
			if got := v.f(v.tickets, v.k); v.want != got {
				t.Errorf(" got = %v ,want = %v \n", got, v.want)
			}
		})

	}
}

func timeRequiredToBuy(tickets []int, k int) int {
	res := 0
	for i, v := range tickets {
		// log.Printf("i ---- %d  v---%d \n", i, v)
		// 在 k 数值 和 之前的
		if i <= k {
			// 如果小于k 值 全部加上 +
			if v <= tickets[k] {
				res += v
				// log.Printf("add %d  \n", v)
			} else {
				// log.Printf("add %d  \n", tickets[k])
				res += tickets[k]
			}

		} else {
			// 在k 数值之后的
			// 如果小于k-1 值 全部加上 +
			if v <= tickets[k]-1 {
				res += v
			} else {
				res += tickets[k] - 1
			}
			// 在k 数值之后的
		}

	}

	return res
}
