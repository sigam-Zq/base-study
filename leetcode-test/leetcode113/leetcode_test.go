package leetcode113

import (
	"strconv"
	"testing"
)

//1742. 盒子中小球的最大数量

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f         func(int, int) int
		lowLimit  int
		highLimit int
		want      int
	}{
		{
			f:         countBalls,
			lowLimit:  1,
			highLimit: 10,
			want:      2,
		},
		{
			f:         countBalls,
			lowLimit:  5,
			highLimit: 15,
			want:      2,
		},
		{
			f:         countBalls,
			lowLimit:  19,
			highLimit: 28,
			want:      2,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.lowLimit, v.highLimit)
			if got != v.want {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func countBalls(lowLimit int, highLimit int) int {
	resMap := make(map[int]int)
	var res int
	for i := lowLimit; i <= highLimit; i++ {
		var sum, t int
		for t = i; t > 0; t /= 10 {
			sum += (t % 10)
		}
		resMap[sum] += 1
		res = max(res, resMap[sum])
	}

	return res
}

func countBallsLeetCode(lowLimit, highLimit int) (ans int) {
	count := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
		for x := i; x > 0; x /= 10 {
			sum += x % 10
		}
		count[sum]++
		ans = max(ans, count[sum])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/maximum-number-of-balls-in-a-box/solutions/1984796/he-zi-zhong-xiao-qiu-de-zui-da-shu-liang-9sfh/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
