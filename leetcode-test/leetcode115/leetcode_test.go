package leetcode115

import (
	"reflect"
	"strconv"
	"testing"
)

// 1706. 球会落何处

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f    func([][]int) []int
		grid [][]int
		want []int
	}{
		{
			f:    findBallFix,
			grid: [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}},
			want: []int{1, -1, -1, -1, -1},
		},
		{
			f:    findBallFix,
			grid: [][]int{{-1}},
			want: []int{-1},
		},
		{
			f:    findBallFix,
			grid: [][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}},
			want: []int{0, 1, 2, 3, 4, -1},
		},
		{
			f:    findBallFix,
			grid: [][]int{{-1, 1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, 1, -1, -1, -1, 1, 1, 1, -1, -1, 1, 1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, 1, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, 1, -1, 1, -1, -1, 1, 1, -1, 1, -1, -1, -1, -1, 1, 1, 1, 1, 1, 1, -1, 1, 1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, 1, -1, -1, 1, 1, -1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, -1}},
			want: []int{-1, -1, -1, 2, 3, 4, 5, 6, -1, -1, 9, 10, 11, 14, -1, -1, 15, 16, 19, 20, -1, -1, 21, 24, -1, -1, 25, -1, -1, 28, 29, 30, 31, 32, 33, 34, 35, -1, -1, -1, -1, 40, 41, 42, 43, 44, 45, -1, -1, 48, -1, -1, -1, -1, 53, 56, -1, -1, -1, -1, 59, 60, 61, 64, 65, 66, 67, 68, -1, -1, 71, 72, -1, -1, 75, 76, -1, -1, 77, 78, -1, -1, -1, -1, 83, 86, -1, -1, 87, -1, -1, -1, -1, 94, 95, -1, -1, 96, 97, 98},
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.grid)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		ans[i] = i
	}

	for i := 0; i < m; i++ {

		for j := 0; j < n; j++ {
			if ans[j] != -1 {
				ans[j] += grid[i][j]
				// 撞墙情况
				if ans[j] < 0 || ans[j] >= n {
					ans[j] = -1
				}
			}
		}

		// 识别v 形情况
		pre, preIdx := 0, 0
		for j := 0; j < n; j++ {
			if ans[j] == -1 {
				continue
			}
			if pre == 0 {
				// 初次直接赋值
				pre = grid[i][ans[j]]
				preIdx = j
			} else {
				if (pre + grid[i][ans[j]]) == 0 {
					ans[j] = -1
					ans[preIdx] = -1
					pre, preIdx = 0, 0
				} else {
					pre = grid[i][j]
					preIdx = j
				}
			}

		}
	}

	return ans
}

func findBallFix(grid [][]int) []int {
	m, n := len(grid), len(grid[0])

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		col := i

		for j := 0; j < m; j++ {
			nextCol := col + grid[j][col]

			if nextCol < 0 || nextCol == n || grid[j][col] != grid[j][nextCol] {
				col = -1
				break
			}

			col = nextCol
		}

		ans[i] = col
	}

	return ans
}

func findBallLeetCode(grid [][]int) []int {
	n := len(grid[0])
	ans := make([]int, n)
	for j := range ans {
		col := j // 球的初始列
		for _, row := range grid {
			dir := row[col]
			col += dir                                  // 移动球
			if col < 0 || col == n || row[col] != dir { // 到达侧边或 V 形
				col = -1
				break
			}
		}
		ans[j] = col // col >= 0 为成功到达底部
	}
	return ans
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/where-will-the-ball-fall/solutions/1285377/qiu-hui-luo-he-chu-by-leetcode-solution-xqop/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func findBallChatGPT(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		col := i // 记录当前列
		for j := 0; j < m; j++ {
			nextCol := col + grid[j][col] // 计算下一步的列位置
			if nextCol < 0 || nextCol >= n || grid[j][col] != grid[j][nextCol] {
				col = -1 // 卡住
				break
			}
			col = nextCol
		}
		ans[i] = col // 记录最终位置
	}

	return ans
}
