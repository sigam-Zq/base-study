package leetcode118

import (
	"reflect"
	"strconv"
	"testing"
)

//2080. 区间内查询数字的频率

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		Oper []string
		Args [][]int
		want []int
	}{
		{
			Oper: []string{"RangeFreqQuery", "query", "query"},
			Args: [][]int{{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56}, {1, 2, 4}, {0, 11, 33}},
			want: []int{0, 1, 2},
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			var obj RangeFreqQuery

			ans := make([]int, len(v.Oper))
			for ii, op := range v.Oper {

				switch op {
				case "RangeFreqQuery":
					obj = Constructor(v.Args[ii])
					ans[ii] = 0
				case "query":
					x, y, z := v.Args[ii][0], v.Args[ii][1], v.Args[ii][2]
					ans[ii] = obj.Query(x, y, z)
				}

			}

			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf(" got %v want %v \n", ans, v.want)
			}

		})

	}

}

type RangeFreqQuery struct {
	list []int
	// map[1][2][3]= 2  1  在2-3 有两个 ,
	FreqMap map[int]map[int]map[int]int
}

func Constructor(arr []int) RangeFreqQuery {
	return RangeFreqQuery{
		list: arr,
	}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {

	if freqOneMap, ok := this.FreqMap[value]; ok {
		if freqTwoMap, ok := freqOneMap[left]; ok {
			if res, ok := freqTwoMap[right]; ok {
				return res
			}
		}
	}

	cnt := 0
	for i := left; i <= right; i++ {
		if this.list[i] == value {
			cnt++
		}
	}

	if freqOneMap, ok := this.FreqMap[value]; !ok {
		this.FreqMap = make(map[int]map[int]map[int]int)
		this.FreqMap[value] = make(map[int]map[int]int)
		this.FreqMap[value][left] = make(map[int]int)
		this.FreqMap[value][left][right] = cnt
	} else {
		if freqTwoMap, ok := freqOneMap[left]; !ok {
			this.FreqMap[value] = make(map[int]map[int]int)
			this.FreqMap[value][left] = make(map[int]int)
			this.FreqMap[value][left][right] = cnt
		} else {
			if _, ok := freqTwoMap[right]; !ok {
				this.FreqMap[value][left] = make(map[int]int)
				this.FreqMap[value][left][right] = cnt
			}
		}
	}

	return cnt
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
