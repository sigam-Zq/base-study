package leetcode99

import (
	"reflect"
	"strconv"
	"testing"
)

// 119. 杨辉三角 II

/*
给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。





示例 1:

输入: rowIndex = 3
输出: [1,3,3,1]
示例 2:

输入: rowIndex = 0
输出: [1]
示例 3:

输入: rowIndex = 1
输出: [1,1]


提示:

0 <= rowIndex <= 33


进阶：

你可以优化你的算法到 O(rowIndex) 空间复杂度吗？
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f      func(int) []int
		rowIdx int
		want   []int
	}{
		{
			f:      getRow,
			rowIdx: 4,
			want:   []int{1, 4, 6, 4, 1},
		},
		{
			f:      getRow,
			rowIdx: 3,
			want:   []int{1, 3, 3, 1},
		},
		{
			f:      getRow,
			rowIdx: 0,
			want:   []int{1},
		},
		{
			f:      getRow,
			rowIdx: 1,
			want:   []int{1, 1},
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.rowIdx)

			if !reflect.DeepEqual(got, v.want) {
				t.Errorf(" got %d  want %d \n", got, v.want)
			}
		})

	}
}

func getRow(rowIndex int) []int {

	// for i := 0; i <= (rowIndex >> 1); i++ {
	// 	ans[i], ans[rowIndex-i] = (i + 1), (i + 1)
	// }
	PreInitDateSet()

	return ProDataSet[rowIndex]
}

var ProDataSet [][]int

func PreInitDateSet() {

	ans := make([][]int, 0)

	ans = append(ans, []int{1})

	ans = append(ans, []int{1, 1})
	for i := 2; i <= 33; i++ {
		ansSin := make([]int, i+1)
		ansSin[0], ansSin[i] = 1, 1
		for j := 1; j <= (i >> 1); j++ {
			ansSin[j], ansSin[i-j] = ans[i-1][j-1]+ans[i-1][j], ans[i-1][j-1]+ans[i-1][j]
		}
		ans = append(ans, ansSin)
	}
	ProDataSet = ans
}

func TestS(t *testing.T) {
	ans := make([][]int, 0)

	ans = append(ans, []int{1})

	ans = append(ans, []int{1, 1})
	for i := 2; i <= 33; i++ {
		ansSin := make([]int, i+1)
		ansSin[0], ansSin[i] = 1, 1
		for j := 1; j <= (i >> 1); j++ {
			ansSin[j], ansSin[i-j] = ans[i-1][j-1]+ans[i-1][j], ans[i-1][j-1]+ans[i-1][j]
		}
		ans = append(ans, ansSin)
	}

	t.Logf("%+v \n", ans)
}

func getRowLeetCode1(rowIndex int) []int {
	C := make([][]int, rowIndex+1)
	for i := range C {
		C[i] = make([]int, i+1)
		C[i][0], C[i][i] = 1, 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}
	return C[rowIndex]
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/pascals-triangle-ii/solutions/601082/yang-hui-san-jiao-ii-by-leetcode-solutio-shuk/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func getRowLeetCode2(rowIndex int) []int {
	var pre, cur []int
	for i := 0; i <= rowIndex; i++ {
		cur = make([]int, i+1)
		cur[0], cur[i] = 1, 1
		for j := 1; j < i; j++ {
			cur[j] = pre[j-1] + pre[j]
		}
		pre = cur
	}
	return pre
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/pascals-triangle-ii/solutions/601082/yang-hui-san-jiao-ii-by-leetcode-solutio-shuk/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func getRowLeetCode3(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/pascals-triangle-ii/solutions/601082/yang-hui-san-jiao-ii-by-leetcode-solutio-shuk/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func getRowLeetCode4(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}
	return row
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/pascals-triangle-ii/solutions/601082/yang-hui-san-jiao-ii-by-leetcode-solutio-shuk/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
