package leetcode50

import (
	"log"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

/*

51. N 皇后
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。



示例 1：


输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：

输入：n = 1
输出：[["Q"]]


提示：

1 <= n <= 9
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func(int) [][]string
		n    int
		want [][]string
	}{
		{
			f:    solveNQueens,
			n:    4,
			want: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
		{
			f:    solveNQueensLeetCode1,
			n:    4,
			want: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.n); !reflect.DeepEqual(got, v.want) {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})
	}

}

func solveNQueens(n int) [][]string {

	res := make([][]string, 0)

	// Q表示皇后 - 表示皇后影响不到（皇后可落子） . 表示已经被皇后影响（不能再有皇后落子）
	// 初始化棋盘
	checkerboard := make([][]rune, n)
	for i := 0; i < n; i++ {
		checkerboard[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			checkerboard[i][j] = '-'
		}
	}

	// dfs 在 done行找到一共可下的位置
	// 表示在 done 行 下到 第 i 个位置\
	var dfs func(i, done int, board [][]rune)
	dfs = func(i, done int, board [][]rune) {

		// done 行 下到 第 i 个位置 - 这里要求进入这个函数前判断这个位置必须为 -（没有被皇后波及）

		// 棋子下去 -标记其他位置
		board = draw(i, done, board)
		log.Printf("line %d row %d draw after %v \n", i, done, board)

		// 已经下满
		if done == n-1 {
			// 计入 res
			tmp := make([]string, n)
			for i, v := range board {
				tmp[i] = string(v)
			}
			res = append(res, tmp)
			return
		}
		// 已经没有位置可下
		pass := false
		for _, b := range board {
			if strings.Contains(string(b), "-") {
				pass = true
			}
		}
		if !pass {
			return
		}

		// 查看下一行还有那些可以下棋
		for i, v := range board[done+1] {
			if v == '-' {
				nextBoard := make([][]rune, n)
				for i, v := range board {
					nextBoard[i] = make([]rune, n)
					copy(nextBoard[i], v)
				}
				dfs(i, done+1, nextBoard)
			}
		}

	}

	for i := 0; i < n; i++ {
		// 第一行下棋 的位置
		// 复制棋盘
		nowBoard := make([][]rune, n)
		for i, v := range checkerboard {
			nowBoard[i] = make([]rune, n)
			copy(nowBoard[i], v)
		}
		dfs(i, 0, nowBoard)
	}

	return res
}

// 下棋 - row,line 下 Q 行列斜边为 .
func draw(line, row int, board [][]rune) [][]rune {
	n := len(board)

	// // 画皇后一列的 .
	for i := range board {
		board[i][line] = '.'
	}
	// 画皇后一行的 .
	for j := range board[row] {
		board[row][j] = '.'
	}

	board[row][line] = 'Q'

	// 画左上方斜边
	for i, j := row-1, line-1; i >= 0 && j >= 0; {
		log.Printf("i %d,j %d \n", i, j)
		board[i][j] = '.'
		// // 斜边左对边（line 为中心轴）
		// if 2*line-j < n {
		// 	board[i][2*line-j] = '.'
		// }
		// 自减
		i--
		j--
	}
	// 画右上方斜边
	for i, j := row-1, line+1; i >= 0 && j < n; {
		board[i][j] = '.'
		i--
		j++
	}

	// 画右下方斜边
	for i, j := row+1, line+1; i < n && j < n; {
		board[i][j] = '.'
		// // 斜边右对边（line 为中心轴）
		// if 2*line-j > 0 {
		// 	board[i][2*line-j] = '.'
		// }
		// 自增
		i++
		j++
	}

	// 画左下方斜边
	for i, j := row+1, line-1; i < n && j >= 0; {
		board[i][j] = '.'
		i++
		j--
	}

	return board
}

func TestXxx2(t *testing.T) {
	log.Printf(" - num %d  \n", '-')
	log.Printf(" . num %d  \n", '.')
	log.Printf(" Q num %d  \n", 'Q')

	n := 4

	checkerboard := make([][]rune, n)
	for i := 0; i < n; i++ {
		checkerboard[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			checkerboard[i][j] = '-'
		}
	}

	b := draw(3, 1, checkerboard)

	log.Printf(" %v \n", b)

}

/*

前言
「N 皇后问题」研究的是如何将 N 个皇后放置在 N×N 的棋盘上，并且使皇后彼此之间不能相互攻击。

皇后的走法是：可以横直斜走，格数不限。因此要求皇后彼此之间不能相互攻击，等价于要求任何两个皇后都不能在同一行、同一列以及同一条斜线上。

直观的做法是暴力枚举将 N 个皇后放置在 N×N 的棋盘上的所有可能的情况，并对每一种情况判断是否满足皇后彼此之间不相互攻击。暴力枚举的时间复杂度是非常高的，因此必须利用限制条件加以优化。

显然，每个皇后必须位于不同行和不同列，因此将 N 个皇后放置在 N×N 的棋盘上，一定是每一行有且仅有一个皇后，每一列有且仅有一个皇后，且任何两个皇后都不能在同一条斜线上。基于上述发现，可以通过回溯的方式寻找可能的解。

回溯的具体做法是：使用一个数组记录每行放置的皇后的列下标，依次在每一行放置一个皇后。每次新放置的皇后都不能和已经放置的皇后之间有攻击：即新放置的皇后不能和任何一个已经放置的皇后在同一列以及同一条斜线上，并更新数组中的当前行的皇后列下标。当 N 个皇后都放置完毕，则找到一个可能的解。当找到一个可能的解之后，将数组转换成表示棋盘状态的列表，并将该棋盘状态的列表加入返回列表。

由于每个皇后必须位于不同列，因此已经放置的皇后所在的列不能放置别的皇后。第一个皇后有 N 列可以选择，第二个皇后最多有 N−1 列可以选择，第三个皇后最多有 N−2 列可以选择（如果考虑到不能在同一条斜线上，可能的选择数量更少），因此所有可能的情况不会超过 N! 种，遍历这些情况的时间复杂度是 O(N!)。

为了降低总时间复杂度，每次放置皇后时需要快速判断每个位置是否可以放置皇后，显然，最理想的情况是在 O(1) 的时间内判断该位置所在的列和两条斜线上是否已经有皇后。

以下两种方法分别使用集合和位运算对皇后的放置位置进行判断，都可以在 O(1) 的时间内判断一个位置是否可以放置皇后，算法的总时间复杂度都是 O(N!)。

方法一：基于集合的回溯
为了判断一个位置所在的列和两条斜线上是否已经有皇后，使用三个集合 columns、diagonals
1
​
  和 diagonals
2
​
  分别记录每一列以及两个方向的每条斜线上是否有皇后。

列的表示法很直观，一共有 N 列，每一列的下标范围从 0 到 N−1，使用列的下标即可明确表示每一列。

如何表示两个方向的斜线呢？对于每个方向的斜线，需要找到斜线上的每个位置的行下标与列下标之间的关系。

方向一的斜线为从左上到右下方向，同一条斜线上的每个位置满足行下标与列下标之差相等，例如 (0,0) 和 (3,3) 在同一条方向一的斜线上。因此使用行下标与列下标之差即可明确表示每一条方向一的斜线。



方向二的斜线为从右上到左下方向，同一条斜线上的每个位置满足行下标与列下标之和相等，例如 (3,0) 和 (1,2) 在同一条方向二的斜线上。因此使用行下标与列下标之和即可明确表示每一条方向二的斜线。



每次放置皇后时，对于每个位置判断其是否在三个集合中，如果三个集合都不包含当前位置，则当前位置是可以放置皇后的位置。

作者：力扣官方题解
链接：https://leetcode.cn/problems/n-queens/solutions/398929/nhuang-hou-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

CODE 如下


复杂度分析

时间复杂度：O(N!)，其中 N 是皇后数量。

空间复杂度：O(N)，其中 N 是皇后数量。空间复杂度主要取决于递归调用层数、记录每行放置的皇后的列下标的数组以及三个集合，递归调用层数不会超过 N，数组的长度为 N，每个集合的元素个数都不会超过 N。


*/

// 回溯法
func solveNQueensLeetCode1(n int) [][]string {

	res := make([][]string, 0)

	queens := make([]int, n)
	for i := range queens {
		queens[i] = -1
	}

	// 当前列 ，左上到右下对角线 ，右上到左下对角线 是否有皇后影响
	columns, diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}, map[int]bool{}

	var backTrack func(queens []int, row, n int, col, dig1, dig2 map[int]bool)
	backTrack = func(queens []int, row, n int, col, dig1, dig2 map[int]bool) {
		// 完成便利后进行加入结果集
		if row == n {
			tmp := generateBoard(queens, n)
			res = append(res, tmp)
			return
		}

		for i := 0; i < n; i++ {

			if col[i] {
				continue
			}

			d1 := row - i
			if dig1[d1] {
				continue
			}

			d2 := row + i
			if dig2[d2] {
				continue
			}

			queens[row] = i
			col[i] = true
			dig1[d1] = true
			dig2[d2] = true
			backTrack(queens, row+1, n, col, dig1, dig2)
			queens[row] = -1
			delete(col, i)
			delete(dig1, d1)
			delete(dig2, d2)
		}

	}

	backTrack(queens, 0, n, columns, diagonals1, diagonals2)

	return res
}

func generateBoard(q []int, n int) []string {
	res := make([]string, n)
	for i := range res {
		for j := 0; j < n; j++ {
			if j == q[i] {
				res[i] += "Q"
			} else {
				res[i] += "."
			}
		}
	}
	return res
}

/*
方法二：基于位运算的回溯
方法一使用三个集合记录分别记录每一列以及两个方向的每条斜线上是否有皇后，每个集合最多包含 N 个元素，因此集合的空间复杂度是 O(N)。如果利用位运算记录皇后的信息，就可以将记录皇后信息的空间复杂度从 O(N) 降到 O(1)。

具体做法是，使用三个整数 columns、diagonals1和 diagonals2分别记录每一列以及两个方向的每条斜线上是否有皇后，每个整数有 N 个二进制位。棋盘的每一列对应每个整数的二进制表示中的一个数位，其中棋盘的最左列对应每个整数的最低二进制位，最右列对应每个整数的最高二进制位。

那么如何根据每次放置的皇后更新三个整数的值呢？在说具体的计算方法之前，首先说一个例子。

棋盘的边长和皇后的数量 N=8。如果棋盘的前两行分别在第 2 列和第 4 列放置了皇后（下标从 0 开始），则棋盘的前两行如下图所示。

作者：力扣官方题解
链接：https://leetcode.cn/problems/n-queens/solutions/398929/nhuang-hou-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
