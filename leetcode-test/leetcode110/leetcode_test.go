package leetcode110

import (
	"strconv"
	"testing"
)

//63. 不同路径 II

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f            func([][]int) int
		obstacleGrid [][]int
		want         int
	}{
		{
			f:            uniquePathsWithObstaclesOptimize,
			obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			want:         2,
		},
		{
			f:            uniquePathsWithObstaclesOptimize,
			obstacleGrid: [][]int{{0, 1}, {0, 0}},
			want:         1,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.obstacleGrid)
			if got != v.want {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 0 {

				if i > 0 {
					dp[i][j] += dp[i-1][j]
				}

				if j > 0 {
					dp[i][j] += dp[i][j-1]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	return dp[m-1][n-1]
}

func uniquePathsWithObstaclesOptimize(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, m)

	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j > 0 && obstacleGrid[j-1][i] == 0 {
				dp[j] += dp[j-1]
			}
			if obstacleGrid[j][i] == 1 {
				dp[j] = 0
			}
		}
	}

	return dp[m-1]
}

func uniquePathsWithObstaclesLeetCode(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, m)
	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[len(f)-1]
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/unique-paths-ii/solutions/316968/bu-tong-lu-jing-ii-by-leetcode-solution-2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
