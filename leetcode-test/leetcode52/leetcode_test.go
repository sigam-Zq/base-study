package leetcode52

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

//TODO 未完成
/*
2056. 棋盘上有效移动组合的数目

https://leetcode.cn/problems/number-of-valid-move-combinations-on-chessboard/description/?envType=daily-question&envId=2024-12-04
有一个 8 x 8 的棋盘，它包含 n 个棋子（棋子包括车，后和象三种）。给你一个长度为 n 的字符串数组 pieces ，其中 pieces[i] 表示第 i 个棋子的类型（车，后或象）。除此以外，还给你一个长度为 n 的二维整数数组 positions ，其中 positions[i] = [ri, ci] 表示第 i 个棋子现在在棋盘上的位置为 (ri, ci) ，棋盘下标从 1 开始。

棋盘上每个棋子都可以移动 至多一次 。每个棋子的移动中，首先选择移动的 方向 ，然后选择 移动的步数 ，同时你要确保移动过程中棋子不能移到棋盘以外的地方。棋子需按照以下规则移动：

车可以 水平或者竖直 从 (r, c) 沿着方向 (r+1, c)，(r-1, c)，(r, c+1) 或者 (r, c-1) 移动。
后可以 水平竖直或者斜对角 从 (r, c) 沿着方向 (r+1, c)，(r-1, c)，(r, c+1)，(r, c-1)，(r+1, c+1)，(r+1, c-1)，(r-1, c+1)，(r-1, c-1) 移动。
象可以 斜对角 从 (r, c) 沿着方向 (r+1, c+1)，(r+1, c-1)，(r-1, c+1)，(r-1, c-1) 移动。
移动组合 包含所有棋子的 移动 。每一秒，每个棋子都沿着它们选择的方向往前移动 一步 ，直到它们到达目标位置。所有棋子从时刻 0 开始移动。如果在某个时刻，两个或者更多棋子占据了同一个格子，那么这个移动组合 不有效 。

请你返回 有效 移动组合的数目。

注意：

初始时，不会有两个棋子 在 同一个位置 。
有可能在一个移动组合中，有棋子不移动。
如果两个棋子 直接相邻 且两个棋子下一秒要互相占据对方的位置，可以将它们在同一秒内 交换位置 。


示例 1:



输入：pieces = ["rook"], positions = [[1,1]]
输出：15
解释：上图展示了棋子所有可能的移动。
示例 2：



输入：pieces = ["queen"], positions = [[1,1]]
输出：22
解释：上图展示了棋子所有可能的移动。
示例 3:



输入：pieces = ["bishop"], positions = [[4,3]]
输出：12
解释：上图展示了棋子所有可能的移动。
示例 4:



输入：pieces = ["rook","rook"], positions = [[1,1],[8,8]]
输出：223
解释：每个车有 15 种移动，所以总共有 15 * 15 = 225 种移动组合。
但是，有两个是不有效的移动组合：
- 将两个车都移动到 (8, 1) ，会导致它们在同一个格子相遇。
- 将两个车都移动到 (1, 8) ，会导致它们在同一个格子相遇。
所以，总共有 225 - 2 = 223 种有效移动组合。
注意，有两种有效的移动组合，分别是一个车在 (1, 8) ，另一个车在 (8, 1) 。
即使棋盘状态是相同的，这两个移动组合被视为不同的，因为每个棋子移动操作是不相同的。
示例 5：



输入：pieces = ["queen","bishop"], positions = [[5,7],[3,4]]
输出：281
解释：总共有 12 * 24 = 288 种移动组合。
但是，有一些不有效的移动组合：
- 如果后停在 (6, 7) ，它会阻挡象到达 (6, 7) 或者 (7, 8) 。
- 如果后停在 (5, 6) ，它会阻挡象到达 (5, 6) ，(6, 7) 或者 (7, 8) 。
- 如果象停在 (5, 2) ，它会阻挡后到达 (5, 2) 或者 (5, 1) 。
在 288 个移动组合当中，281 个是有效的。


提示：

n == pieces.length
n == positions.length
1 <= n <= 4
pieces 只包含字符串 "rook" ，"queen" 和 "bishop" 。
棋盘上最多只有一个后。
1 <= ri, ci <= 8
每一个 positions[i] 互不相同。

*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f         func([]string, [][]int) int
		pieces    []string
		positions [][]int
		want      int
	}{
		{
			f:         countCombinations,
			pieces:    []string{"rook"},
			positions: [][]int{{1, 1}},
			want:      15,
		},
		{
			f:         countCombinations,
			pieces:    []string{"bishop"},
			positions: [][]int{{4, 3}},
			want:      12,
		},
		{
			f:         countCombinations,
			pieces:    []string{"queen", "bishop"},
			positions: [][]int{{5, 7}, {3, 4}},
			want:      281,
		},
		{
			f:         countCombinations,
			pieces:    []string{"bishop", "rook"},
			positions: [][]int{{8, 5}, {7, 7}},
			want:      96,
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.pieces, v.positions); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})
	}

}

func countCombinations(pieces []string, positions [][]int) int {
	n := len(pieces)
	checkerboard := make([][]int, 8)
	piecesPossibility := make([]int, n)
	for i := range checkerboard {
		checkerboard[i] = make([]int, 8)
	}

	for i := 0; i < n; i++ {
		pieceLoc := positions[i]
		// 去除 棋盘上 从1开始的偏差
		rowNum := pieceLoc[0] - 1
		lineNum := pieceLoc[1] - 1
		cnt := 0
		switch pieces[i] {
		case "rook":
			// 车走直线
			for j := 0; j < 8; j++ {
				// 同一行
				if rowNum != j {
					checkerboard[j][lineNum]++
					cnt++
				}
				if lineNum != j {
					checkerboard[rowNum][j]++
					cnt++
				}
			}
			checkerboard[rowNum][lineNum]++
			cnt++

		case "bishop":
			// 主教走斜线
			dig1Base := rowNum - lineNum
			dig2Base := rowNum + lineNum
			for j := 0; j < 8; j++ {
				for k := 0; k < 8; k++ {
					if j == rowNum || k == lineNum {
						continue
					}

					if j-k == dig1Base {
						checkerboard[j][k]++
						cnt++
					}
					if j+k == dig2Base {
						checkerboard[j][k]++
						cnt++
					}

				}
			}
			checkerboard[rowNum][lineNum]++
			cnt++

		case "queen":

			// 走直线
			for j := 0; j < 8; j++ {
				// 同一行
				if rowNum != j {
					checkerboard[j][lineNum]++
					cnt++
				}
				if lineNum != j {
					checkerboard[rowNum][j]++
					cnt++
				}
			}
			// 走斜线
			dig1Base := rowNum - lineNum
			dig2Base := rowNum + lineNum
			for j := 0; j < 8; j++ {
				for k := 0; k < 8; k++ {
					if j == rowNum || k == lineNum {
						continue
					}
					if j-k == dig1Base {
						checkerboard[j][k]++
						cnt++
					}
					if j+k == dig2Base {
						checkerboard[j][k]++
						cnt++
					}

				}
			}

			checkerboard[rowNum][lineNum]++
			cnt++

		default:
			panic("unexpected piece")
		}

		piecesPossibility[i] = cnt

	}
	log.Printf("\n")
	for _, v := range checkerboard {
		log.Printf(" %v \n", v)
	}
	log.Printf("\n")
	log.Printf("piecesPossibility %v \n", piecesPossibility)
	var res int
	res = 1
	for _, v := range piecesPossibility {
		res *= v
	}

	// 找出所有大于1的交叉点 并配对距离最远的棋子
	// 第i 枚棋子被遮蔽的点为 [x,y]
	deviation := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if checkerboard[i][j] > 1 {
				// 交叉点 求出交叉点距离边的距离 加到偏差值
				a := min(i+1, j+1)
				b := min(8-i, 8-j)
				deviation += min(a, b)
				fmt.Printf("deviation add i %d j %d  sum %d \n", i, j, min(a, b))
			}
		}
	}
	fmt.Println(deviation)

	return res - deviation
}

/*
方法一：枚举
思路与算法

题目描述看似复杂，实则非常好理解，给出一个 8×8 的国际象棋棋盘，只有三种类型棋子：

车，用 “rook" 表示，只能水平或者竖直移动，共 4 个方向；
后，用 “queen" 表示，可以水平、竖直或者斜对角移动，共 8 个方向；
象，用 “bishop" 表示，可以斜对角移动，共 4 个方向。
现在每个棋子最多可移动一个回合，即选择一个方向和步数进行移动一次，或者选择停留在原地。棋子的移动不是瞬时的，每一秒只能移动一步，并且它们的移动方向总是直线（也就是直接从起点到终点）。定义有效组合为所有棋子在移动时不会存在某个时刻两个棋子出现在同一个位置上，现在请你计算有多少个这样的有效组合。

我们首先看一下数据范围，棋盘的大小是固定的，棋子最多只有 4 个，并且最多只有一个后（后的移动可能最多）。而对于某个棋子来说，在同一个方向及其相反方向上（例如向上或者向下），最多只有 7 种可能，因此车有 15 种可能（算上停在原地的一种），象有 14 种可能（因为 8 是偶数，无论如何无法同时取到两条最长的对角线），后有 28 种可能。因此，最多也不过 94500 种可能，并且这里面还有很多初始状态就不符合的组合。判断某个组合是否合法的时间复杂度约为 O(8n
2
 )=O(n
2
 )。

因此，我们可以一一枚举这些组合，并判断它们是否是有效组合。在枚举时，我们可以使用深度优先搜索。在具体实现时，我们可以等某个组合内的移动方案都确定后再判断它的合法性，也可以在填充移动方案时提前判定，后者的运行时间更小，因为提前做了剪枝。

作者：力扣官方题解
链接：https://leetcode.cn/problems/number-of-valid-move-combinations-on-chessboard/solutions/2995574/qi-pan-shang-you-xiao-yi-dong-zu-he-de-s-83cm/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
// 定义车、象、后棋子的方向
var rookDirections = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var bishopDirections = [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
var queenDirections = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func countCombinationsLeetCode(pieces []string, positions [][]int) int {
	n, res := len(pieces), 0
	stack := []Movement{}

	// Check 判断第 u 个棋子是否与之前的棋子发生相交
	check := func(u int) bool {
		for v := 0; v < u; v++ {
			if stack[u].cross(&stack[v]) {
				return false
			}
		}
		return true
	}

	// DFS 深度优先搜索
	var dfs func(u int)
	dfs = func(u int) {
		if u == n {
			res++
			return
		}
		var directions [][2]int
		switch pieces[u] {
		case "rook":
			directions = rookDirections
		case "queen":
			directions = queenDirections
		default:
			directions = bishopDirections
		}

		// 处理第 u 个棋子原地不动的情况
		stack = append(stack, Movement{startX: positions[u][0], startY: positions[u][1], endX: positions[u][0], endY: positions[u][1]})
		if check(u) {
			dfs(u + 1)
		}
		stack = stack[:len(stack)-1]

		// 枚举第 u 个棋子在所有方向、所有步数的情况
		for _, dir := range directions {
			for j := 1; j < 8; j++ {
				x := positions[u][0] + dir[0]*j
				y := positions[u][1] + dir[1]*j
				if x < 1 || x > 8 || y < 1 || y > 8 {
					break
				}
				stack = append(stack, Movement{startX: positions[u][0], startY: positions[u][1], endX: x, endY: y, dx: dir[0], dy: dir[1]})
				if check(u) {
					dfs(u + 1)
				}
				stack = stack[:len(stack)-1]
			}
		}
	}

	dfs(0)
	return res
}

// Movement 结构体表示棋子的一个移动
type Movement struct {
	startX, startY, endX, endY, dx, dy, curX, curY int
}

// Reset 重置棋子的当前位置
func (m *Movement) reset() {
	m.curX = m.startX
	m.curY = m.startY
}

// Stopped 判断棋子是否停止
func (m *Movement) stopped() bool {
	return m.curX == m.endX && m.curY == m.endY
}

// Advance 让棋子按照步长移动
func (m *Movement) advance() {
	if !m.stopped() {
		m.curX += m.dx
		m.curY += m.dy
	}
}

// Cross 判断两个棋子是否相遇
func (m *Movement) cross(oth *Movement) bool {
	m.reset()
	oth.reset()
	for !m.stopped() || !oth.stopped() {
		m.advance()
		oth.advance()
		if m.curX == oth.curX && m.curY == oth.curY {
			return true
		}
	}

	return false
}
