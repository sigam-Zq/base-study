package leetcode114

import (
	"sort"
	"strconv"
	"testing"
)

// 1552. 两球之间的磁力

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f        func([]int, int) int
		position []int
		m        int
		want     int
	}{
		{
			f:        maxDistance,
			position: []int{1, 2, 3, 4, 7},
			m:        3,
			want:     3,
		},
		{
			f:        maxDistance,
			position: []int{5, 4, 3, 2, 1, 1000000000},
			m:        2,
			want:     999999999,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.position, v.m)
			if got != v.want {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func maxDistance(position []int, m int) int {

	sort.Ints(position)
	l, r := 1, position[len(position)-1]-position[0]

	var ans = -1
	for l <= r {
		mid := (l + r) / 2

		if check2(mid, position, m) {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}

	}

	return ans
}

func check2(x int, pos []int, m int) bool {
	pre, cnt := pos[0], 1
	n := len(pos)
	for i := 1; i < n; i++ {
		if x <= pos[i]-pre {
			pre = pos[i]
			cnt++
		}
	}

	return m <= cnt
}

func maxDistanceLeetCode(position []int, m int) int {
	sort.Ints(position)
	return sort.Search(position[len(position)-1], func(f int) bool {
		prev := position[0]
		cnt := 1
		for _, curr := range position {
			if curr-prev >= f {
				cnt++
				prev = curr
			}
		}
		return cnt < m
	}) - 1
}

// 作者：ylb
// 链接：https://leetcode.cn/problems/magnetic-force-between-two-balls/solutions/3074160/python3javacgotypescript-yi-ti-yi-jie-er-czva/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func maxDistanceLeetCode2(position []int, m int) int {
	sort.Ints(position)
	left, right := 1, position[len(position)-1]-position[0]
	ans := -1
	for left <= right {
		mid := (left + right) / 2
		if check(mid, position, m) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func check(x int, position []int, m int) bool {
	pre, cnt := position[0], 1
	for i := 1; i < len(position); i++ {
		if position[i]-pre >= x {
			pre = position[i]
			cnt++
		}
	}
	return cnt >= m
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/magnetic-force-between-two-balls/solutions/403701/liang-qiu-zhi-jian-de-ci-li-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
