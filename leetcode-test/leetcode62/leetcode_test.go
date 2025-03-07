package leetcode62

import (
	"sort"
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
			f:    minSetSize,
			arr:  []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7},
			want: 2,
		},
	} {

		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.arr); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})
	}
}

func minSetSize(arr []int) int {
	n := len(arr)
	frequencyMap := make(map[int]int)

	for _, v := range arr {
		frequencyMap[v]++
		// 有超过半数的频次的数目直接返回
		if frequencyMap[v] >= n>>1 {
			return 1
		}
	}
	vList := make([]int, 0)
	for _, v := range frequencyMap {
		vList = append(vList, v)
	}
	sort.Ints(vList)

	// log.Println(vList)
	// log.Println("vList")
	res := 0
	s := 0
	for i := len(vList) - 1; i >= 0; i-- {
		s += vList[i]
		res++
		if s >= n>>1 {
			return res
		}
	}

	return -1
}
