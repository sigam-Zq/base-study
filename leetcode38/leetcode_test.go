package leetcode38

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"testing"
)

/*
3240. 最少翻转次数使二进制矩阵回文 II

给你一个 m x n 的二进制矩阵 grid 。

如果矩阵中一行或者一列从前往后与从后往前读是一样的，那么我们称这一行或者这一列是 回文 的。

你可以将 grid 中任意格子的值 翻转 ，也就是将格子里的值从 0 变成 1 ，或者从 1 变成 0 。

请你返回 最少 翻转次数，使得矩阵中 所有 行和列都是 回文的 ，且矩阵中 1 的数目可以被 4 整除 。

https://leetcode.cn/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-ii/description/?envType=daily-question&envId=2024-11-16


示例 1：

输入：grid = [[1,0,0],[0,1,0],[0,0,1]]

输出：3

解释：



示例 2：

输入：grid = [[0,1],[0,1],[0,0]]

输出：2

解释：



示例 3：

输入：grid = [[1],[1]]

输出：2

解释：





提示：

m == grid.length
n == grid[i].length
1 <= m * n <= 2 * 105
0 <= grid[i][j] <= 1
*/

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f    func([][]int) int
		grid [][]int
		want int
	}{
		{
			f:    minFlips,
			grid: [][]int{{0, 1}, {0, 1}, {0, 0}},
			want: 2,
		},
		{
			f:    minFlips,
			grid: [][]int{{1, 0, 1}, {0, 0, 0}, {0, 0, 0}, {0, 0, 1}},
			want: 1,
		},
		{
			f:    minFlips,
			grid: [][]int{{1}, {1}, {1}, {0}},
			want: 1,
		},
		{
			f:    minFlips,
			grid: [][]int{{0, 1}, {0, 1}, {0, 0}},
			want: 2,
		},
		{
			f:    minFlips,
			grid: [][]int{{0, 0, 1}, {1, 1, 0}, {1, 1, 1}, {0, 1, 1}},
			want: 4,
		},
		{
			f:    minFlips,
			grid: [][]int{{0, 0, 1}, {1, 1, 1}, {0, 1, 0}, {0, 1, 1}},
			want: 5,
		},
		{
			f:    minFlips,
			grid: [][]int{{1, 0, 0}, {0, 0, 1}, {0, 0, 1}},
			want: 3,
		},
		{
			f:    minFlips,
			grid: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			want: 3,
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.grid); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})

	}
}

func minFlips(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	// 分成三种情况  m n 都是偶数  m n 都是奇数 m n 存在一个奇数
	res := 0
	if n%2 == 0 && m%2 == 0 {
		// 回文即 为4 的倍数

		for i := 0; i < n/2; i++ {
			for j := 0; j < m/2; j++ {
				if grid[i][j] == grid[i][m-j-1] && grid[i][j] == grid[n-i-1][j] && grid[i][j] == grid[n-i-1][m-j-1] {
					continue
				} else {
					sum := grid[i][j] + grid[i][m-j-1] + grid[n-i-1][j] + grid[n-i-1][m-j-1]
					res += min(sum, 4-sum)
				}
			}
		}

		return res

	} else if n%2 == 1 && m%2 == 1 {
		// 回文 且中心位置必须为 0
		crossSpace := 0
		for i := 0; i <= n/2; i++ {
			for j := 0; j <= m/2; j++ {
				if grid[i][j] == grid[i][m-j-1] && grid[i][j] == grid[n-i-1][j] && grid[i][j] == grid[n-i-1][m-j-1] {
					continue
				} else {
					sum := grid[i][j] + grid[i][m-j-1] + grid[n-i-1][j] + grid[n-i-1][m-j-1]
					fmt.Printf("-------------sum %d\n", sum)
					if i == n/2 || j == m/2 {
						fmt.Printf("-i %d-  j %d--------\n", i, j)
						if i == n/2 && j == m/2 {
							continue
						}
						crossSpace++
					} else {
						res += min(sum, 4-sum)
					}
				}
			}
		}
		fmt.Printf("-------------res %d\n", res)
		fmt.Printf("-------------crossSpace %d\n", crossSpace)
		res += min(crossSpace, 4-crossSpace)

		// 判断中心位置
		nMid := n / 2
		mMid := m / 2
		if grid[nMid][mMid] == 1 {
			res++
		}

		return res

	} else {
		// m 或者 n 必然存在一个奇数 奇数的那一列或者一行 必须是 4的倍数（两个2）
		// 已经到达过中间一行或者一列了
		for i := 0; i < n/2; i++ {
			for j := 0; j < m/2; j++ {

				if grid[i][j] == grid[i][m-j-1] && grid[i][j] == grid[n-i-1][j] && grid[i][j] == grid[n-i-1][m-j-1] {
					continue
				} else {
					sum := grid[i][j] + grid[i][m-j-1] + grid[n-i-1][j] + grid[n-i-1][m-j-1]
					res += min(sum, 4-sum)
				}
			}
		}

		// 识别出 单列的情况

		if n%2 == 1 {

			// 中间一行 要为4 的倍数
			lineSum := 0
			for _, vv := range grid[n/2] {
				lineSum += vv
			}
			deviation := lineSum % 4
			res += min(deviation, 4-deviation)
		} else if m%2 == 1 {
			// 中间一列 要为4 的倍数
			rowSum := 0
			for k := 0; k < n; k++ {
				rowSum += grid[k][m/2]
			}
			deviation := rowSum % 4
			res += min(deviation, 4-deviation)

		}

		return res

	}

	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abandonFunc(grid [][]int) int {

	// 计算 1的数量
	allOne := 0
	for _, v := range grid {
		for _, vv := range v {
			allOne += vv
		}
	}
	//偏差值　该值需要需要控制在取余　为０
	// deviation := allOne % 4

	n := len(grid)
	m := len(grid[0])

	// 把 1--> 0 增加可替换的开销
	zeroSpace := 0
	// 把 0--> 1 增加可替换的开销
	oneSpace := 0

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 满足回文条件
			if grid[i][j] == grid[i][m-j-1] && grid[i][j] == grid[n-i-1][j] && grid[i][j] == grid[n-i-1][m-j-1] {
				continue
			} else {
				if 2 >= (grid[i][j]+grid[i][m-j-1]+grid[n-i-1][j]+grid[n-i-1][m-j-1]) && (allOne%4) < 3 {

					// 半数为 0 优先置零
					if grid[i][j] == 1 {
						grid[i][j] = 0
						allOne--
						res++
						zeroSpace++
					}

					if grid[i][m-j-1] == 1 {
						grid[i][m-j-1] = 0
						allOne--
						res++
						zeroSpace++
					}
					if grid[n-i-1][j] == 1 {
						grid[n-i-1][j] = 0
						allOne--
						res++
						zeroSpace++
					}

					if grid[n-i-1][m-j-1] == 1 {
						grid[n-i-1][m-j-1] = 0
						allOne--
						res++
						zeroSpace++
					}

				} else {
					// 超过半数为1
					// 优先置为一

					if grid[i][j] == 0 {
						grid[i][j] = 1
						allOne++
						res++
						oneSpace++
					}

					if grid[i][m-j-1] == 0 {
						grid[i][m-j-1] = 1
						allOne++
						res++
						oneSpace++
					}
					if grid[n-i-1][j] == 0 {
						grid[n-i-1][j] = 1
						allOne++
						res++
						oneSpace++
					}

					if grid[n-i-1][m-j-1] == 0 {
						grid[n-i-1][m-j-1] = 1
						allOne++
						res++
						oneSpace++
					}

				}
			}
		}
	}
	log.Printf("allOne %d  res %d \n", allOne, res)

	// 判断 allOne 与 4的关系
	deviation := allOne % 4

	log.Printf("deviation %d  zeroSpace %d oneSpace %d \n", deviation, zeroSpace, oneSpace)
	// 判断上述中是否有 可以被撤销的 变回去为  减少偏差值做出退让空间
	if deviation != (zeroSpace - oneSpace) {
		res += deviation
	}

	return res
}

/*
方法一：分类讨论
思路

题目要求所有行列都必须是回文的，即满足

grid[i][j]=grid[i][n−1−j]=grid[m−1−i][j]=grid[m−1−i][n−1−j]
其中，i 和 j 满足 0≤i≤⌊
2
m
​
 ⌋，0≤j≤⌊
2
n
​
 ⌋。

将这四个数都变为 0 需要的次数记作 cnt 次，那么将它们都变为 1 则需要 4−cnt 次。将这四个数变为相同所需要的次数就是 min(cnt,4−cnt) 次。当 m，n 都为偶数时，答案就是所有将四个数变为相同数字所需次数之和。

接下来讨论 m 或 n 为奇数时的情况。当 m 是奇数，矩阵正中间会多出一行；当 n 为奇数，矩阵正中间会多出一列。

当 m 和 n 都为奇数时，由于矩阵中 1 的数目可以被 4 整除，所以正中间的元素必须是 0。

当只有行数 n 为奇数时，需要满足对称性 grid[i][j]=grid[i][m−1−j]。除为了满足对称性所需要的操作次数外，我们可能还需要额外的操作来使该行中 1 的个数为 4 的整数倍。

将对称位置相同的 1 的个数记为 cnt
1
​
 ，对称位置的数不同的数对个数记作 diff。

当 cnt
1
​
  模 4 为 0 时，无需额外操作。
当 cnt
1
​
  模 4 为 2 时：
如果 diff>0，可以将其中一对数变为 1，其余数都变成 0，这样 cnt
1
​
  的数量就会增加 2，成为 4 的倍数。
如果 diff=0，可以把 cnt
1
​
  中的一对 1 变成 0，花费两次操作。
对于列数 m 为奇数时的讨论结果类似。

作者：力扣官方题解
链接：https://leetcode.cn/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-ii/solutions/2981502/zui-shao-fan-zhuan-ci-shu-shi-er-jin-zhi-hslz/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func minFlipsLeetCode(grid [][]int) int {
	m, n, ans := len(grid), len(grid[0]), 0
	for i := 0; i < m/2; i++ {
		for j := 0; j < n/2; j++ {
			cnt1 := grid[i][j] + grid[i][n-1-j] +
				grid[m-1-i][j] + grid[m-1-i][n-1-j]
			ans += min(cnt1, 4-cnt1)
		}
	}

	diff, cnt1 := 0, 0
	if m%2 == 1 {
		for j := 0; j < n/2; j++ {
			if grid[m/2][j]^grid[m/2][n-1-j] != 0 {
				diff++
			} else {
				cnt1 += grid[m/2][j] * 2
			}
		}
	}
	if n%2 == 1 {
		for i := 0; i < m/2; i++ {
			if grid[i][n/2]^grid[m-1-i][n/2] != 0 {
				diff++
			} else {
				cnt1 += grid[i][n/2] * 2
			}
		}
	}
	if m%2 == 1 && n%2 == 1 {
		ans += grid[m/2][n/2]
	}
	if diff > 0 {
		ans += diff
	} else {
		ans += cnt1 % 4
	}
	return ans
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-ii/solutions/2981502/zui-shao-fan-zhuan-ci-shu-shi-er-jin-zhi-hslz/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/*
方法二：动态规划
思路

将矩阵中对称的元素分组，每一组都必须相等。（当行数或列数为奇数时，不一定每一组都有 4 个元素）要么都为 1，要么都为 0。要使整个矩阵的 1 的数量模 4 等于 0，就是让每一组的 1 的个数的总和等于 4 的倍数。

定义 f[i][j] 表示在前 i 组中，1 的数量模 4 的余数为 j 时的最小操作数。那么对于第 i+1 组：

将这一组全都变为 0，那么余数不变，有 f[i+1][j]=f[i][j]。
将这一组全都变为 1，设组内有 cnt 个元素，那么有 f[i+1][(j+cnt)%4]=f[i][j]。
设总分组的数量为 group，f[group][0] 即所求答案。

作者：力扣官方题解
链接：https://leetcode.cn/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-ii/solutions/2981502/zui-shao-fan-zhuan-ci-shu-shi-er-jin-zhi-hslz/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func minFlipsLeetCodeDp(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([]int, 4)
	for i := range f {
		f[i] = math.MaxInt32 / 2
	}
	f[0] = 0
	for i := 0; i < (m+1)/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			ones := grid[i][j]
			cnt := 1
			if j != n-1-j {
				ones += grid[i][n-1-j]
				cnt++
			}
			if i != m-1-i {
				ones += grid[m-1-i][j]
				cnt++
			}
			if i != m-1-i && j != n-1-j {
				ones += grid[m-1-i][n-1-j]
				cnt++
			}
			// 计算将这一组全部变为 1 的代价
			cnt1 := cnt - ones
			// 计算将这一组全部变为 0 的代价
			cnt0 := ones
			tmp := make([]int, 4)
			for k := 0; k < 4; k++ {
				tmp[k] = f[k] + cnt0
			}
			for k := 0; k < 4; k++ {
				tmp[(k+cnt)%4] = min(tmp[(k+cnt)%4], f[k]+cnt1)
			}
			f = tmp
		}
	}
	return f[0]
}
