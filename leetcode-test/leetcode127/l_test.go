package leetcode125

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

/*
1356.根据数字二进制下1的数目排序
https://leetcode.cn/problems/sort-integers-by-the-number-of-1-bits/
*/
func Test2Xxx(t *testing.T) {
	c := 1 & 1
	t.Log(c)
	c = 2 & 1
	t.Log(c)
	c = 3 & 1
	t.Log(c)
	c = 4 & 1
	t.Log(c)
	/*
	   l_test.go:15: 1
	   l_test.go:17: 0
	   l_test.go:19: 1
	   l_test.go:21: 0
	*/
}

func Test3Xxx(t *testing.T) {
	t.Log(CountOne(0))
	t.Log(CountOne(1))
	t.Log(CountOne(2))
	t.Log(CountOne(3))
	t.Log(CountOne(4))
	t.Log(CountOne(5))
	t.Log(CountOne(7))
	t.Log(CountOne(8))
	t.Log(CountOne(15))
}

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func(arr []int) []int
		arr  []int
		want []int
	}{
		{
			f:    sortByBits,
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			want: []int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			f:    sortByBits,
			arr:  []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
			want: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
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
func sortByBits(arr []int) []int {
	customArr := CustomInts(arr)
	sort.Sort(&customArr)

	return customArr
}

type CustomInts []int

func (c *CustomInts) Len() int {
	return len(*c)
}
func (c *CustomInts) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}
func (c *CustomInts) Less(i, j int) bool {
	if CountOne((*c)[i]) == CountOne((*c)[j]) {
		return (*c)[i] < (*c)[j]
	}
	return CountOne((*c)[i]) < CountOne((*c)[j])
}

func CountOne(t int) int {
	var count int
	if t == 0 {
		return count
	}
	for i := t; i > 0; i = i >> 1 {
		if (i & 1) == 1 {
			count++
		}
	}
	return count
}

func onesCount(x int) (c int) {
	for ; x > 0; x /= 2 {
		c += x % 2
	}
	return
}

func sortByBitsLeetCode(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		cx, cy := onesCount(x), onesCount(y)
		return cx < cy || cx == cy && x < y
	})
	return a
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/sort-integers-by-the-number-of-1-bits/solutions/109168/gen-ju-shu-zi-er-jin-zhi-xia-1-de-shu-mu-pai-xu-by/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
