package sub

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

/*
52. N 皇后 II
n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。

示例 1：


输入：n = 4
输出：2
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：

输入：n = 1
输出：1


提示：

1 <= n <= 9

*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func(int) int
		n    int
		want int
	}{
		{
			f:    solveNQueensBit,
			n:    4,
			want: 2,
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.n); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})
	}

}

// 回溯法
func solveNQueens(n int) int {
	cnt := 0

	queens := make([]int, n)
	for i := range queens {
		queens[i] = -1
	}
	column, diagonal1, diagonal2 := map[int]bool{}, map[int]bool{}, map[int]bool{}

	var backTrack func(row int, col, dia1, dia2 map[int]bool)
	backTrack = func(row int, col, dia1, dia2 map[int]bool) {
		if n == row {
			cnt++
			return
		}

		for i := 0; i < n; i++ {
			if col[i] {
				continue
			}
			d1 := row - i
			if dia1[d1] {
				continue
			}
			d2 := row + i
			if dia2[d2] {
				continue
			}

			dia1[d1] = true
			dia2[d2] = true
			col[i] = true
			backTrack(row+1, col, dia1, dia2)
			delete(dia1, d1)
			delete(dia2, d2)
			delete(col, i)

		}

	}

	backTrack(0, column, diagonal1, diagonal2)

	return cnt
}

func TestBitOp(t *testing.T) {
	fmt.Printf(" %b \n", (1<<3 - 1))
	fmt.Printf(" %cn \n", (1<<3 - 1))
	fmt.Printf(" %8b \n", 8)
	fmt.Printf(" %b \n", (1<<3-1)&^8)
	fmt.Printf(" %b \n", 3)
	fmt.Printf(" %b \n", ^3)
	fmt.Printf("5 = %b \n", 5)
	fmt.Printf("^5 = %b %d %d %d \n", ^5, ^5, 1^5, 0^5)
	fmt.Printf("-5 = %b \n", -5)
	fmt.Printf(" %b   %b \n", 4, -4)
	fmt.Printf(" %b  \n", 4&-4)
	fmt.Printf(" %b   %b x&y %b  \n", 10, (1 << 2), 10&(1<<1))
}

// 位 回溯
func solveNQueensBit(n int) int {
	ans := 0

	var resolve func(row, col, dia1, dia2 int)
	resolve = func(row, col, dia1, dia2 int) {
		if row == n {
			ans++
			return
		}

		activePosition := (1<<n - 1) &^ (col | dia1 | dia2)
		log.Printf("row %d activePosition %b \n", row, activePosition)
		for activePosition > 0 {
			position := activePosition & -activePosition
			log.Printf("row %d position %b \n", row, position)
			resolve(row+1, (col | position), (dia1|position)<<1, (dia2|position)>>1)
			activePosition = activePosition &^ position
		}

	}
	resolve(0, 0, 0, 0)
	return ans
}
