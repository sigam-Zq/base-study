package leetcode104

import (
	"strconv"
	"testing"
)

// 	598. 区间加法 II

/*
给你一个 m x n 的矩阵 M 和一个操作数组 op 。
矩阵初始化时所有的单元格都为 0 。ops[i] = [ai, bi] 意味着当所有的 0 <= x < ai 和 0 <= y < bi 时，
M[x][y] 应该加 1。

在 执行完所有操作后 ，计算并返回 矩阵中最大整数的个数 。



示例 1:



输入: m = 3, n = 3，ops = [[2,2],[3,3]]
输出: 4
解释: M 中最大的整数是 2, 而且 M 中有4个值为2的元素。因此返回 4。
示例 2:

输入: m = 3, n = 3, ops = [[2,2],[3,3],[3,3],[3,3],[2,2],[3,3],[3,3],[3,3],[2,2],[3,3],[3,3],[3,3]]
输出: 4
示例 3:

输入: m = 3, n = 3, ops = []
输出: 9


提示:


1 <= m, n <= 4 * 104
0 <= ops.length <= 104
ops[i].length == 2
1 <= ai <= m
1 <= bi <= n


*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func(int, int, [][]int) int
		m    int
		n    int
		ops  [][]int
		want int
	}{
		{
			f:    maxCount,
			m:    3,
			n:    3,
			ops:  [][]int{{2, 2}, {3, 3}},
			want: 4,
		},
		{
			f:    maxCount,
			m:    3,
			n:    3,
			ops:  [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}},
			want: 4,
		},
		{
			f:    maxCount,
			m:    3,
			n:    3,
			ops:  [][]int{},
			want: 9,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.m, v.n, v.ops)
			if got != v.want {
				t.Errorf(" got %d  want %d  \n", got, v.want)
			}
		})

	}

}

func maxCount(m int, n int, ops [][]int) int {
	ans := n * m

	for _, op := range ops {
		m = min(m, op[0])
		n = min(n, op[1])
		ans = min(ans, m*n)
	}

	return ans
}

func maxCountLeetCode1(m, n int, ops [][]int) int {
	mina, minb := m, n
	for _, op := range ops {
		mina = min(mina, op[0])
		minb = min(minb, op[1])
	}
	return mina * minb
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/range-addition-ii/solutions/1086781/fan-wei-qiu-he-ii-by-leetcode-solution-kcxq/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
