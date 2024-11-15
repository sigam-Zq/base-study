package leetcode37

import (
	"strconv"
	"testing"
)

/*

3239. 最少翻转次数使二进制矩阵回文 I

给你一个 m x n 的二进制矩阵 grid 。

如果矩阵中一行或者一列从前往后与从后往前读是一样的，那么我们称这一行或者这一列是 回文 的。

你可以将 grid 中任意格子的值 翻转 ，也就是将格子里的值从 0 变成 1 ，或者从 1 变成 0 。

请你返回 最少 翻转次数，使得矩阵 要么 所有行是 回文的 ，要么所有列是 回文的 。

https://leetcode.cn/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-i/description/?envType=daily-question&envId=2024-11-15

示例 1：

输入：grid = [[1,0,0],[0,0,0],[0,0,1]]

输出：2

解释：



将高亮的格子翻转，得到所有行都是回文的。

示例 2：

输入：grid = [[0,1],[0,1],[0,0]]

输出：1

解释：



将高亮的格子翻转，得到所有列都是回文的。

示例 3：

输入：grid = [[1],[0]]

输出：0

解释：

所有行已经是回文的。

*/

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f    func([][]int) int
		grid [][]int
		want int
	}{
		{
			f:    minFlips,
			grid: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}},
			want: 2,
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
	// 列长度 - 行数
	m := len(grid)
	// 行长度 - 列数
	n := len(grid[0])

	rowOperaNumber := 0
	for i := 0; i < m; i++ {
		rowOperaNumber += needOperaNumberToPalindrome(grid[i]...)
	}

	columnOperaNumber := 0
	for i := 0; i < n; i++ {
		column := make([]int, m)
		for ii, v := range grid {
			// log.Printf("ii %d i %d v %v \n", ii, i, v)
			column[ii] = v[i]
		}
		// log.Printf("i %d column %v \n", i, column)
		columnOperaNumber += needOperaNumberToPalindrome(column...)
	}
	// log.Printf("columnOperaNumber %d   rowOperaNumber %d \n", columnOperaNumber, rowOperaNumber)
	var min int

	if columnOperaNumber > rowOperaNumber {
		min = rowOperaNumber
	} else {
		min = columnOperaNumber
	}

	return min
}

func needOperaNumberToPalindrome(line ...int) int {

	lineLen := len(line)
	if lineLen == 1 {
		return 0
	}
	needNumber := 0
	for i := lineLen - 1; i >= (lineLen / 2); i-- {
		if line[i] != line[lineLen-i-1] {
			needNumber++
		}
	}

	return needNumber
}
