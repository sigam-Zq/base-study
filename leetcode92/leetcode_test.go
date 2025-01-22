package leetcode92

import (
	"reflect"
	"strconv"
	"testing"
)

// 2218. 从栈中取出 K 个硬币的最大面值和

/*
一张桌子上总共有 n 个硬币 栈 。每个栈有 正整数 个带面值的硬币。

每一次操作中，你可以从任意一个栈的 顶部 取出 1 个硬币，从栈中移除它，并放入你的钱包里。

给你一个列表 piles ，其中 piles[i] 是一个整数数组，分别表示第 i 个栈里 从顶到底 的硬币面值。同时给你一个正整数 k ，请你返回在 恰好 进行 k 次操作的前提下，你钱包里硬币面值之和 最大为多少 。



示例 1：



输入：piles = [[1,100,3],[7,8,9]], k = 2
输出：101
解释：
上图展示了几种选择 k 个硬币的不同方法。
我们可以得到的最大面值为 101 。
示例 2：

输入：piles = [[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]], k = 7
输出：706
解释：
如果我们所有硬币都从最后一个栈中取，可以得到最大面值和。


提示：

n == piles.length
1 <= n <= 1000
1 <= piles[i][j] <= 105
1 <= k <= sum(piles[i].length) <= 2000


*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f       func([][]int, int) int
		piles   [][]int
		k       int
		want    int
		isDebug bool
	}{
		{
			f:       maxValueOfCoins,
			piles:   [][]int{{1, 100, 3}, {7, 8, 9}},
			k:       2,
			want:    101,
			isDebug: false,
		},
		{
			f:       maxValueOfCoins,
			piles:   [][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}},
			k:       7,
			want:    706,
			isDebug: false,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans := v.f(v.piles, v.k)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

func maxValueOfCoins(piles [][]int, k int) int {

	f := make([]int, k+1)
	f[0] = 0
	for i := 1; i <= k; i++ {
		f[i] = -1
	}

	for _, p := range piles {

		for i := k; i >= 0; i-- {
			v := 0
			for j := 1; j <= len(p) && i >= j; j++ {
				v += p[j-1]
				if i >= j && f[i-j] != -1 {
					f[i] = max(f[i], f[i-j]+v)
				}
			}

		}

	}
	return f[k]
}

// TODO
func maxValueOfCoinsDfs(piles [][]int, k int) int {

	var dfs func(curList []int, k int) int

	dfs = func(curList []int, k int) int {
		if k == 0 {
			return 0
		}

		for i := range piles {
			if curList[i] >= len(piles[i]) {
				continue
			}

			curList = append(curList, piles[i][0])
			dfs(curList, k-1)
		}
		return 0
	}

	var ans int
	// for i := range dp[k] {
	// 	ans = max(ans, dp[k][i])
	// }

	return ans
}

func maxValueOfCoinsLeetCode(piles [][]int, k int) int {
	f := make([]int, k+1)
	for i := range f {
		f[i] = -1
	}
	f[0] = 0
	for _, pile := range piles {
		for i := k; i > 0; i-- {
			value := 0
			for t := 1; t <= len(pile); t++ {
				value += pile[t-1]
				if i >= t && f[i-t] != -1 {
					f[i] = max(f[i], f[i-t]+value)
				}
			}
		}
	}
	return f[k]
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/maximum-value-of-k-coins-from-piles/solutions/3047590/cong-zhan-zhong-qu-chu-k-ge-ying-bi-de-z-4hua/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
