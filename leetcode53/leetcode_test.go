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

/*
方法一：分类讨论
思路与算法

根据棋盘中 车 的移动规则可以知道，对于棋盘中任意的位置 (x,y)，位置 (i,j) 的 车 至多需要 2 步即可移动到指定位置 (x,y)，移动方法如下：

对于 车 来说，可以从起始位置 (i,j) 先垂直移动到 (x,j)，再水平移动到 (x,y)，或者先水平移动到 (i,y)，再垂直移动到 (x,y)；
如果 白色象、黑皇后 处在同一条对角线或者 白色车、黑皇后 处于同一条线上，此时可能只需要 1 次移动即可捕获 黑皇后，根据题意分析如下：

如果 白色象 与 黑皇后 处在同一条对角线时，且此时该路径上无 白色车 阻挡时，此时只需移动 1 次即可捕获 黑皇后；

如果 白色车 与 黑皇后 处在同一行或者同一列时，且此时二者移动路线之间无 白色象 阻挡时，此时只需移动 1 次即可捕获 黑皇后；

其余情况下，白色车 最多需要 2 次移动即可捕获 黑皇后；

作者：力扣官方题解
链接：https://leetcode.cn/problems/minimum-moves-to-capture-the-queen/solutions/2995567/bu-huo-hei-huang-hou-xu-yao-de-zui-shao-vxmt1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func minMovesToCaptureTheQueenLeetCode(a int, b int, c int, d int, e int, f int) int {
	// 车与皇后处在同一行，且中间没有象
	if a == e && (c != a || d <= min(b, f) || d >= max(b, f)) {
		return 1
	}
	// 车与皇后处在同一列，且中间没有象
	if b == f && (d != b || c <= min(a, e) || c >= max(a, e)) {
		return 1
	}
	// 象、皇后处在同一条对角线，且中间没有车
	if abs(c-e) == abs(d-f) && ((c-e)*(b-f) != (a-e)*(d-f) || a < min(c, e) || a > max(c, e)) {
		return 1
	}
	return 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
