package leetcode109

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

// 59. 螺旋矩阵 II

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func(int) [][]int
		n    int
		want [][]int
	}{
		{
			f: generateMatrix,
			n: 3,
			want: [][]int{{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5}},
		},
		{
			f:    generateMatrix,
			n:    1,
			want: [][]int{{1}},
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.n)
			fmt.Println(got)
			fmt.Println(v.want)
			if !arraysEqualsDup(got, v.want) {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)

	for i := range ans {
		ans[i] = make([]int, n)
	}

	x, y, mode, modeLayer := 0, 0, 0, 0
	for i := 1; i <= n*n; i++ {
		fmt.Printf("x %d, y %d, i %d ,  mode %d, modeLayer %d \n", x, y, i, mode, modeLayer)
		ans[x][y] = i
		switch mode {
		case 0:
			if y < n-modeLayer-1 {
				y++
			} else {
				// 到头换模式
				mode = 1
				x++
			}
		case 1:

			if x < n-modeLayer-1 {
				x++
			} else {
				// 到头换模式
				mode = 2
				y--
			}
		case 2:
			if y > modeLayer {
				y--
			} else {
				// 到头换模式
				mode = 3
				x--
			}
		case 3:
			if x > modeLayer+1 {
				x--
			} else {
				// 到头换模式
				mode = 0
				y++
				modeLayer++
			}

		}
	}

	return ans
}

func arraysEqualsDup(v1 [][]int, v2 [][]int) bool {
	return reflect.DeepEqual(v1, v2)
}

type pair struct{ x, y int }

var dirs = []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上

func generateMatrixLeetCode1(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	row, col, dirIdx := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := dirs[dirIdx]
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			dirIdx = (dirIdx + 1) % 4 // 顺时针旋转至下一个方向
			dir = dirs[dirIdx]
		}
		row += dir.x
		col += dir.y
	}
	return matrix
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/spiral-matrix-ii/solutions/658676/luo-xuan-ju-zhen-ii-by-leetcode-solution-f7fp/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func generateMatrixLeetCode2(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 1
	left, right, top, bottom := 0, n-1, 0, n-1
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			matrix[top][column] = num
			num++
		}
		for row := top + 1; row <= bottom; row++ {
			matrix[row][right] = num
			num++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				matrix[bottom][column] = num
				num++
			}
			for row := bottom; row > top; row-- {
				matrix[row][left] = num
				num++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return matrix
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/spiral-matrix-ii/solutions/658676/luo-xuan-ju-zhen-ii-by-leetcode-solution-f7fp/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
