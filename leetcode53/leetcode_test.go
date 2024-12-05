package leetcode53

import (
	"strconv"
	"testing"
)

/*
3001. 捕获黑皇后需要的最少移动次数

现有一个下标从 1 开始的 8 x 8 棋盘，上面有 3 枚棋子。

给你 6 个整数 a 、b 、c 、d 、e 和 f ，其中：

(a, b) 表示白色车的位置。
(c, d) 表示白色象的位置。
(e, f) 表示黑皇后的位置。
假定你只能移动白色棋子，返回捕获黑皇后所需的最少移动次数。

请注意：

车可以向垂直或水平方向移动任意数量的格子，但不能跳过其他棋子。
象可以沿对角线方向移动任意数量的格子，但不能跳过其他棋子。
如果车或象能移向皇后所在的格子，则认为它们可以捕获皇后。
皇后不能移动。

*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f      func(int, int, int, int, int, int) int
		pieces []int
		want   int
	}{
		{
			f:      minMovesToCaptureTheQueen,
			pieces: []int{1, 1, 8, 8, 2, 3},
			want:   2,
		},
		{
			f:      minMovesToCaptureTheQueen,
			pieces: []int{4, 3, 3, 4, 5, 2},
			want:   2,
		},
		{
			f:      minMovesToCaptureTheQueen,
			pieces: []int{1, 1, 1, 4, 1, 8},
			want:   2,
		},
		{
			f:      minMovesToCaptureTheQueen,
			pieces: []int{2, 3, 1, 4, 3, 3},
			want:   1,
		},
		{
			f:      minMovesToCaptureTheQueen,
			pieces: []int{5, 3, 5, 1, 4, 3},
			want:   1,
		},
		{
			f:      minMovesToCaptureTheQueen,
			pieces: []int{1, 3, 1, 2, 1, 1},
			want:   2,
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.pieces[0], v.pieces[1], v.pieces[2], v.pieces[3], v.pieces[4], v.pieces[5]); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})
	}

}

/*
 q o b r o o o o
 o o o o o o o o
 o o o o o o o o
 o o o o o o o o
 o o o o o o o o
 o o o o o o o o
 o o o o o o o o
 o o o o o o o o

*/

/*
(a, b) 表示白色车的位置。
(c, d) 表示白色象的位置。
(e, f) 表示黑皇后的位置。
*/
func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {

	if c+d == e+f {
		// 一步吃的过程中有其他棋子挡住的情况
		// 1.1 .象被挡住 左上到右下对角线 【在一条对角线上 && 在 象和皇后中间】
		if a+b == c+d && ((a > c && a < e) || (a < c && a > e)) {
			return 2
		}

		return 1
	}

	if c-d == e-f {
		// 1.2 .象被挡住 右上到左下对角线 【在一条对角线上 && 在 象和皇后中间】
		if a-b == c-d && ((a > c && a < e) || (a < c && a > e)) {
			return 2
		}
		return 1
	}
	if a == e {
		// 2.1 .车被挡住  行  【在一行上 && 在 车和皇后中间】
		if a == c && ((d > b && d < f) || (d < b && d > f)) {
			return 2
		}
		return 1
	}
	if b == f {
		// 2.1 .车被挡住  列  【在一列上 && 在 车和皇后中间】
		if b == d && ((c > a && c < e) || (c < a && c > e)) {
			return 2
		}

		return 1
	}

	return 2
}
