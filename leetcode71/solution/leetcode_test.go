package leetcode71

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f             func(int, int, []int, []int) int
		m             int
		n             int
		horizontalCut []int
		verticalCut   []int
		want          int
	}{
		{
			f:             minimumCost,
			m:             3,
			n:             2,
			horizontalCut: []int{1, 3},
			verticalCut:   []int{5},
			want:          13,
		},
		{
			f:             minimumCost,
			m:             2,
			n:             2,
			horizontalCut: []int{7},
			verticalCut:   []int{4},
			want:          15,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.m, v.n, v.horizontalCut, v.verticalCut)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}

}

/*

方法一：记忆化搜索
思路

将原问题分解成子问题，设计函数 dp[row1,col1​,row2,col2]
表示切割完整的子矩形至最小单元的代价，其中子矩形的两个对角顶点的坐标分别为 (row1,col1 ) 和 (row2,col2)。
  我们可以任意水平或者垂直切一刀，然后将问题分解成更小的子问题，
  直到分到最小单元。遍历所有可能的切法，取出最小值作为子问题的返回值。因为递归过程中有许多重复状态，
  我们利用记忆话搜索的方式来降低时间复杂度。最后返回 dp(0,0,m−1,n−1) 即可。

作者：力扣官方题解
链接：https://leetcode.cn/problems/minimum-cost-for-cutting-cake-i/solutions/3016927/qie-dan-gao-de-zui-xiao-zong-kai-xiao-i-7kpj5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	cache := make([]int, m*m*n*n)
	for i := range cache {
		cache[i] = -1
	}

	index := func(row1, col1, row2, col2 int) int {
		return (row1*n+col1)*m*n + row2*n + col2
	}

	var dp func(row1, col1, row2, col2 int) int
	dp = func(row1, col1, row2, col2 int) int {
		if row1 == row2 && col1 == col2 {
			return 0
		}
		ind := index(row1, col1, row2, col2)
		if cache[ind] >= 0 {
			return cache[ind]
		}
		cache[ind] = math.MaxInt32
		for i := row1; i < row2; i++ {
			cache[ind] = min(cache[ind], dp(row1, col1, i, col2)+dp(i+1, col1, row2, col2)+horizontalCut[i])
		}
		for i := col1; i < col2; i++ {
			cache[ind] = min(cache[ind], dp(row1, col1, row2, i)+dp(row1, i+1, row2, col2)+verticalCut[i])
		}
		return cache[ind]
	}

	return dp(0, 0, m-1, n-1)
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/minimum-cost-for-cutting-cake-i/solutions/3016927/qie-dan-gao-de-zui-xiao-zong-kai-xiao-i-7kpj5/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
