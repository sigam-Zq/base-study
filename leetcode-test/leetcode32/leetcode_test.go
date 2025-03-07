package leetcode32

import (
	"strconv"
	"testing"
)

/*

https://leetcode.cn/problems/check-if-the-rectangle-corner-is-reachable/description/?envType=daily-question&envId=2024-11-08


3235. 判断矩形的两个角落是否可达
给你两个正整数 xCorner 和 yCorner 和一个二维整数数组 circles ，其中 circles[i] = [xi, yi, ri] 表示一个圆心在 (xi, yi) 半径为 ri 的圆。

坐标平面内有一个左下角在原点，右上角在 (xCorner, yCorner) 的矩形。你需要判断是否存在一条从左下角到右上角的路径满足：路径 完全 在矩形内部，不会 触碰或者经过 任何 圆的内部和边界，同时 只 在起点和终点接触到矩形。

如果存在这样的路径，请你返回 true ，否则返回 false 。

示例 1：

输入：X = 3, Y = 4, circles = [[2,1,1]]

输出：true

解释：



黑色曲线表示一条从 (0, 0) 到 (3, 4) 的路径。

示例 2：

输入：X = 3, Y = 3, circles = [[1,1,2]]

输出：false

解释：



不存在从 (0, 0) 到 (3, 3) 的路径。

示例 3：

输入：X = 3, Y = 3, circles = [[2,1,1],[1,2,1]]

输出：false

解释：



不存在从 (0, 0) 到 (3, 3) 的路径。

示例 4：

输入：X = 4, Y = 4, circles = [[5,5,1]]

输出：true

解释：





提示：

3 <= xCorner, yCorner <= 109
1 <= circles.length <= 1000
circles[i].length == 3
1 <= xi, yi, ri <= 109

*/

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f       func(int, int, [][]int) bool
		xCorner int
		yCorner int
		circles [][]int
		want    bool
	}{
		{
			f:       canReachCorner,
			xCorner: 3,
			yCorner: 4,
			circles: [][]int{{2, 1, 1}},
			want:    true,
		},
		{
			f:       canReachCorner,
			xCorner: 5,
			yCorner: 8,
			circles: [][]int{{4, 7, 1}},
			want:    false,
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.xCorner, v.yCorner, v.circles); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}

		})

	}
}

func canReachCorner(xCorner int, yCorner int, circles [][]int) bool {
	res := true

	// // x余下的空间
	// xSpace := [][]int{{0,xCorner}}
	// // y余下的空间
	// ySpace := [][]int{{0,yCorner}}
	// circlesPointLeftUp := 0
	// circlesPointLeftDown := 0
	// circlesPointRightUp := 0
	// circlesPointRightDown := 0
	for _, circle := range circles {

		// 判断中心和 方块面积的关系
		// if (circle[0]+circle[2]) <= xCorner && (circle[1]+circle[2]) <= yCorner {
		// 	// 圆在方块内

		// } else if (circle[0]-circle[2]) <= xCorner && (circle[1]-circle[2]) <= yCorner {
		// 	// ⚪拟合的正方形 和矩形内容有相交

		// } else {
		// 	// 没有相交的情况
		// 	continue
		// }

		// 面积扩容

		// 拿到四个边界
		xAreaStart := circle[0] - circle[2]
		xAreaEnd := circle[0] + circle[2]
		yAreaStart := circle[1] - circle[2]
		yAreaEnd := circle[1] + circle[2]

		xCorner = min(xCorner, xAreaStart, xAreaEnd)
		yCorner = min(yCorner, yAreaStart, yAreaEnd)

	}

	// 正方形和⚪总和面积对比

	if xCorner <= 0 && yCorner <= 0 {
		res = false
	}

	return res
}

func min(list ...int) int {
	if len(list) == 0 {
		return 0
	}
	min := list[0]

	for _, v := range list {
		if min > v {
			min = v
		}
	}
	return min
}
