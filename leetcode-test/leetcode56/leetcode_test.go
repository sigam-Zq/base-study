package leetcode56

import (
	"math/bits"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f     func([][]int) int
		board [][]int
		want  int
	}{
		{
			f:     movesToChessboard,
			board: [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1}},
			want:  2,
		},
	} {

		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.board); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})
	}
}

func movesToChessboard(board [][]int) int {
	n := len(board)
	colMask, rowMask := 0, 0

	for i := 0; i < n; i++ {
		rowMask |= board[0][i] << i
		colMask |= board[i][0] << i
	}

	reversalRow := 1<<n - 1 ^ rowMask
	reversalCol := 1<<n - 1 ^ colMask

	rowCnt, colCnt := 0, 0
	for i := 0; i < n; i++ {
		curRow, curCol := 0, 0

		for j := 0; j < n; j++ {
			curRow |= board[i][j] << j
			curCol |= board[j][i] << j
		}

		// 取反 ^(  ( curCol == colMask || curCol || reversalCol  ) && (curRow == rowMask || curRow  == reversalRow) )
		if curCol != colMask && curCol != reversalCol || curRow != rowMask && curRow != reversalRow {
			return -1
		}

		if curCol == colMask {
			colCnt++
		}
		if curRow == rowMask {
			rowCnt++
		}

	}

	r := RowMove(uint(rowMask), rowCnt, n)
	if r == -1 {
		return r
	}

	c := RowMove(uint(colMask), colCnt, n)
	if c == -1 {
		return c
	}

	return r + c
}

func RowMove(mask uint, cnt int, n int) int {
	ones := bits.OnesCount(mask)

	// n 判断奇偶性
	if n&0 > 0 {
		// 奇数

		// 每一行的需要1的个数依次大于 总数的一半
		if abs(n-2*ones) != 1 || abs(n-2*cnt>>1) != 1 {
			return -1
		}

		// 第一位为 偶数
		if n>>1 == ones {
			return n/2 - bits.OnesCount(mask&0xAAAAAAAA)
		} else {
			return (n+1)/2 - bits.OnesCount(mask&0x55555555)
		}

	} else {
		// 偶数
		if n != ones>>1 || n != cnt>>1 {
			return -1
		}

		m := n/2 - bits.OnesCount(mask&0xAAAAAAAA)
		m2 := n/2 - bits.OnesCount(mask&0x55555555)

		return min(m, m2)
	}

}

// func abs(a int) int {
// 	return int(math.Abs(float64(a)))
// }

func getMoves(mask uint, count, n int) int {
	ones := bits.OnesCount(mask)
	if n&1 > 0 {
		// 如果 n 为奇数，则每一行中 1 与 0 的数目相差为 1，且满足相邻行交替
		if abs(n-2*ones) != 1 || abs(n-2*count) != 1 {
			return -1
		}
		if ones == n>>1 {
			// 偶数位变为 1 的最小交换次数
			return n/2 - bits.OnesCount(mask&0xAAAAAAAA)
		} else {
			// 奇数位变为 1 的最小交换次数
			return (n+1)/2 - bits.OnesCount(mask&0x55555555)
		}
	} else {
		// 如果 n 为偶数，则每一行中 1 与 0 的数目相等，且满足相邻行交替
		if ones != n>>1 || count != n>>1 {
			return -1
		}
		// 偶数位变为 1 的最小交换次数
		count0 := n/2 - bits.OnesCount(mask&0xAAAAAAAA)
		// 奇数位变为 1 的最小交换次数
		count1 := n/2 - bits.OnesCount(mask&0x55555555)
		return min(count0, count1)
	}
}

func movesToChessboardLeetCode(board [][]int) int {
	n := len(board)
	// 棋盘的第一行与第一列
	rowMask, colMask := 0, 0
	for i := 0; i < n; i++ {
		rowMask |= board[0][i] << i
		colMask |= board[i][0] << i
	}
	reverseRowMask := 1<<n - 1 ^ rowMask
	reverseColMask := 1<<n - 1 ^ colMask
	rowCnt, colCnt := 0, 0
	for i := 0; i < n; i++ {
		currRowMask, currColMask := 0, 0
		for j := 0; j < n; j++ {
			currRowMask |= board[i][j] << j
			currColMask |= board[j][i] << j
		}
		if currRowMask != rowMask && currRowMask != reverseRowMask || // 检测每一行的状态是否合法
			currColMask != colMask && currColMask != reverseColMask { // 检测每一列的状态是否合法
			return -1
		}
		if currRowMask == rowMask {
			rowCnt++ // 记录与第一行相同的行数
		}
		if currColMask == colMask {
			colCnt++ // 记录与第一列相同的列数
		}
	}
	rowMoves := getMoves(uint(rowMask), rowCnt, n)
	colMoves := getMoves(uint(colMask), colCnt, n)
	if rowMoves == -1 || colMoves == -1 {
		return -1
	}
	return rowMoves + colMoves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*

问
关于上述题解 有以下问题

1  getMoves 函数在这里 入参的 count 在函数内 是什么作用

2.  其中 多次使用的  bits.OnesCount(mask) 是起到了什么作用

3.  currRowMask != rowMask && currRowMask != reverseRowMask ||    currColMask != colMask && currColMask != reverseColMask

这里 怎么 什么判断出 改矩阵不能转换出棋盘的


### 问题解答：

#### **1. `getMoves` 函数中 `count` 的作用是什么？**

`count` 是用来记录矩阵中，与参考模式（第一行或第一列的模式）**完全匹配的行或列的数量**。它的作用是确保：

- 矩阵中满足棋盘规则的行或列数量足够。
- 奇偶性检查是否一致（即行或列的分布规律是否符合棋盘模式）。

例如：
- 若 `n` 为奇数，符合模式的行或列数量应为 `⌈n / 2⌉` 或 `⌊n / 2⌋`。
- 若 `n` 为偶数，符合模式的行或列数量必须为 `n / 2`。

如果 `count` 的值和预期不符（即满足棋盘模式的行/列数量不对），则矩阵不能通过交换行或列变成棋盘。

---

#### **2. `bits.OnesCount(mask)` 的作用是什么？**

`bits.OnesCount(mask)` 用于计算 `mask` 的二进制表示中有多少个 `1`。

- 在这个问题中，`mask` 表示一行或一列的模式（例如 `1010`）。
- 通过 `bits.OnesCount`，可以得知一行或一列的 `1` 的数量，从而判断该行或列是否符合棋盘的基本规则：
  - 若 `n` 为奇数，`1` 的数量和 `0` 的数量相差最多为 1。
  - 若 `n` 为偶数，`1` 和 `0` 的数量必须相等。

此外，`bits.OnesCount` 还用来计算最小交换次数：
- `mask & 0xAAAAAAAA`：提取偶数位的 `1`。
- `mask & 0x55555555`：提取奇数位的 `1`。
通过这些计算，可以确定使 `1` 对齐到棋盘的最小交换次数。

---

#### **3. `currRowMask != rowMask && currRowMask != reverseRowMask || currColMask != colMask && currColMask != reverseColMask` 如何判断矩阵不能转换为棋盘？**

这段代码的逻辑是检查当前行/列是否满足棋盘的基本规则。

- `rowMask` 是第一行的模式（例如 `1010`）。
- `reverseRowMask` 是第一行模式的反转（例如 `0101`）。
- 类似地，`colMask` 和 `reverseColMask` 分别是第一列的模式及其反转。

对于一个矩阵能被转换为棋盘，所有的行和列必须要么完全匹配 `rowMask` / `colMask`，要么完全匹配 `reverseRowMask` / `reverseColMask`。

**若某一行或某一列既不符合模式，也不符合反模式**，说明该行或列无法通过交换转化为棋盘，直接返回 `-1`。

---

#### **总结**
1. `count` 确保行或列的数量分布满足棋盘规则。
2. `bits.OnesCount` 用于计算模式中的 `1` 数量，辅助判断是否符合棋盘规则，以及最小交换次数。
3. `currRowMask != rowMask && currRowMask != reverseRowMask || currColMask != colMask && currColMask != reverseColMask` 判断当前行或列是否符合棋盘的基本要求，不符合即返回 `-1`。
*/

/*

问

题解中
getMoves 函数中


当 n 为奇数时，下面的判断在全局中起到什么作用
ones == n>>1


答

在 `getMoves` 函数中，当矩阵的大小 `n` 为奇数时，`ones == n >> 1` 的判断在全局中起到了以下作用：

### **作用：**
1. **确定当前行/列的模式符合棋盘规则。**
   - 当 `n` 为奇数时，棋盘规则要求：每一行（或列）的 `1` 和 `0` 的数量必须相差最多为 1。这是由棋盘本身的交替结构决定的。
   - 在这种情况下：
     - 如果某一行或列的 `1` 的数量为 `n >> 1`（即 `n / 2` 的整数部分），那么 `0` 的数量就是 `(n >> 1) + 1`，正好满足相差为 1 的条件。
     - 如果 `ones != n >> 1`，说明该行或列的 `1` 和 `0` 的数量无法满足棋盘的要求。

2. **区分奇数位与偶数位的模式。**
   - 在棋盘规则下，行/列可以有两种排列方式：
     - 偶数位是 `1`，奇数位是 `0`（例如 `10101`）。
     - 奇数位是 `1`，偶数位是 `0`（例如 `01010`）。
   - 如果 `ones == n >> 1`，意味着**偶数位模式（如 `10101`）更适合匹配**。
   - 如果 `ones != n >> 1`，说明需要调整为奇数位模式（如 `01010`）。

---

### **全局逻辑分析**
#### 在 `getMoves` 函数中，当 `n` 为奇数时：
```go
if ones == n >> 1 {
    return n/2 - bits.OnesCount(mask & 0xAAAAAAAA)
} else {
    return (n+1)/2 - bits.OnesCount(mask & 0x55555555)
}
```

- 如果 `ones == n >> 1`：
  - 说明该行/列的 `1` 数量与偶数位模式匹配。
  - 接下来计算将偶数位对齐为 `1` 所需的最小交换次数。

- 如果 `ones != n >> 1`：
  - 则说明该行/列的 `1` 数量需要调整为奇数位模式。
  - 此时计算将奇数位对齐为 `1` 所需的最小交换次数。

#### **全局意义**
通过 `ones == n >> 1`，可以：
1. 确定当前的行或列是否满足棋盘的基本模式。
2. 根据满足的模式，计算需要的最小交换次数。

如果 `ones` 的值完全不符合条件（即两种模式都无法对齐），那意味着这行/列无法成为棋盘的一部分，后续代码可能返回 `-1`。

---

### **总结**
`ones == n >> 1` 在全局中起到了**验证当前行/列是否符合棋盘规则**和**选择最优交换策略（奇数位还是偶数位模式）**的作用。如果该条件不成立，说明当前行/列不符合棋盘模式要求，直接放弃这一行/列作为有效棋盘候选。

*/
