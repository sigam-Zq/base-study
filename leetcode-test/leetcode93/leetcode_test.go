package leetcode93

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f       func([]int) int
		piles   []int
		want    int
		isDebug bool
	}{
		{
			f:       maxCoins,
			piles:   []int{2, 4, 1, 2, 7, 8},
			want:    9,
			isDebug: false,
		},
		{
			f:       maxCoins,
			piles:   []int{2, 4, 5},
			want:    4,
			isDebug: false,
		},
		{
			f:       maxCoins,
			piles:   []int{9, 8, 7, 6, 5, 1, 2, 3, 4},
			want:    18,
			isDebug: false,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans := v.f(v.piles)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

func maxCoins(piles []int) int {
	var ans int
	sort.Ints(piles)
	// log.Println(piles)
	for i, j := 0, len(piles)-1; i < j; i, j = i+1, j-2 {
		ans += piles[j-1]
	}

	return ans
}

func maxCoinsLeetCode(piles []int) int {
	sort.Ints(piles)
	length := len(piles)
	rounds := length / 3
	coins := 0
	index := length - 2
	for i := 0; i < rounds; i++ {
		coins += piles[index]
		index -= 2
	}
	return coins
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/maximum-number-of-coins-you-can-get/solutions/409109/ni-ke-yi-huo-de-de-zui-da-ying-bi-shu-mu-by-leetco/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
