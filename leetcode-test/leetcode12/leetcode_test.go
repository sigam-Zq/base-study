package leetcode12

import "testing"

func TestMain(t *testing.T) {
	args := []struct {
		name  string
		f     func([]int, []int) int
		days  []int
		costs []int
		want  int
	}{
		{
			name:  "one",
			f:     mincostTickets,
			days:  []int{1, 4, 6, 7, 8, 20},
			costs: []int{2, 7, 15},
			want:  11,
		},
	}

	for _, v := range args {
		t.Run(v.name, func(t *testing.T) {
			if got := v.f(v.days, v.costs); v.want != got {
				t.Errorf(" got = %v ,want = %v \n", got, v.want)
			}
		})

	}

}

func mincostTickets(days []int, costs []int) int {

	return 0
}
