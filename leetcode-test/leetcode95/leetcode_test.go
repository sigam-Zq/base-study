package leetcode95

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

// 2944. 购买水果需要的最少金币数

/*
给你一个 下标从 1 开始的 整数数组 prices ，其中 prices[i] 表示你购买第 i 个水果需要花费的金币数目。

水果超市有如下促销活动：

如果你花费 prices[i] 购买了下标为 i 的水果，那么你可以免费获得下标范围在 [i + 1, i + i] 的水果。
注意 ，即使你 可以 免费获得水果 j ，你仍然可以花费 prices[j] 个金币去购买它以获得它的奖励。

请你返回获得所有水果所需要的 最少 金币数。



示例 1：

输入：prices = [3,1,2]

输出：4

解释：

用 prices[0] = 3 个金币购买第 1 个水果，你可以免费获得第 2 个水果。
用 prices[1] = 1 个金币购买第 2 个水果，你可以免费获得第 3 个水果。
免费获得第 3 个水果。
请注意，即使您可以免费获得第 2 个水果作为购买第 1 个水果的奖励，但您购买它是为了获得其奖励，这是更优化的。

示例 2：

输入：prices = [1,10,1,1]

输出：2

解释：

用 prices[0] = 1 个金币购买第 1 个水果，你可以免费获得第 2 个水果。
免费获得第 2 个水果。
用 prices[2] = 1 个金币购买第 3 个水果，你可以免费获得第 4 个水果。
免费获得第 4 个水果。
示例 3：

输入：prices = [26,18,6,12,49,7,45,45]

输出：39

解释：

用 prices[0] = 26 个金币购买第 1 个水果，你可以免费获得第 2 个水果。
免费获得第 2 个水果。
用 prices[2] = 6 个金币购买第 3 个水果，你可以免费获得第 4，5，6（接下来的三个）水果。
免费获得第 4 个水果。
免费获得第 5 个水果。
用 prices[5] = 7 个金币购买第 6 个水果，你可以免费获得第 7 和 第 8 个水果。
免费获得第 7 个水果。
免费获得第 8 个水果。
请注意，即使您可以免费获得第 6 个水果作为购买第 3 个水果的奖励，但您购买它是为了获得其奖励，这是更优化的。



提示：

1 <= prices.length <= 1000
1 <= prices[i] <= 105
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f       func([]int) int
		prices  []int
		want    int
		isDebug bool
	}{
		{
			f:       minimumCoins,
			prices:  []int{3, 1, 2},
			want:    4,
			isDebug: false,
		},
		{
			f:       minimumCoins,
			prices:  []int{1, 10, 1, 1},
			want:    2,
			isDebug: false,
		},
		{
			f:       minimumCoins,
			prices:  []int{26, 18, 6, 12, 49, 7, 45, 45},
			want:    39,
			isDebug: false,
		},
		{
			f:       minimumCoinsFake,
			prices:  []int{38, 23, 27, 32, 47, 45, 48, 24, 39, 26, 37, 42, 24, 45, 27, 26, 15, 16, 26, 6},
			want:    132,
			isDebug: false,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans := v.f(v.prices)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

// [i + 1, i + i]
// 在0 开始的坐标中 [i + 1, 2*i + 1]
func minimumCoins(prices []int) int {
	n := len(prices)

	dp := make([]int, n)
	dp[0] = prices[0]
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt
	}

	for i, _ := range prices {
		if i == 0 {
			continue
		}
		minPrice := math.MaxInt

		for j := i; j >= (i / 2); j-- {
			minPrice = min(minPrice, prices[j])
		}
		dp[i] = minPrice
	}
	// log.Println(dp)
	var ans int
	for i := 0; i < n; i = 2*i + 2 {
		ans += dp[i]
	}

	return ans
}

func minimumCoinsFake(prices []int) int {

	n := len(prices)

	memo := make(map[int]int)

	var f func(idx int) int

	f = func(idx int) int {
		if 2*idx+2 >= n {
			return prices[idx]
		}

		if val, ok := memo[idx]; ok {
			return val
		}

		maxV := math.MaxInt
		for i := idx + 1; i <= 2*idx+2; i++ {
			maxV = min(maxV, f(i))
		}
		memo[idx] = prices[idx] + maxV
		return memo[idx]

	}

	return f(0)
}

// [i + 1, i + i]
// 在0 开始的坐标中 [i + 1, 2*i + 1]
func minimumCoinsFack(prices []int) int {
	n := len(prices)

	dp := make([]int, n+1)
	dp[0] = prices[0]
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
	}

	for i := 0; i < n; i++ {
		for j := i / 2; j <= i; j++ {
			dp[i+1] = min(dp[i+1], dp[j]+prices[j])
		}
	}

	return dp[n]
}

func minimumCoinsLeetCode(prices []int) int {
	memo := make(map[int]int)

	var dp func(index int) int
	dp = func(index int) int {
		if 2*index+2 >= len(prices) {
			return prices[index]
		}
		if val, ok := memo[index]; ok {
			return val
		}
		minValue := math.MaxInt32
		for i := index + 1; i <= 2*index+2; i++ {
			minValue = min(minValue, dp(i))
		}
		memo[index] = prices[index] + minValue
		return memo[index]
	}

	return dp(0)
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/minimum-number-of-coins-for-fruits/solutions/3046625/gou-mai-shui-guo-xu-yao-de-zui-shao-jin-f6rsy/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func minimumCoinsLeetCode2(prices []int) int {
	n := len(prices)
	queue := [][2]int{{n, 0}}

	for i := n - 1; i >= 0; i-- {
		for len(queue) > 0 && queue[len(queue)-1][0] >= 2*i+3 {
			queue = queue[:len(queue)-1]
		}
		cur := queue[len(queue)-1][1] + prices[i]
		for len(queue) > 0 && queue[0][1] >= cur {
			queue = queue[1:]
		}
		queue = append([][2]int{{i, cur}}, queue...)
	}

	return queue[0][1]
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/minimum-number-of-coins-for-fruits/solutions/3046625/gou-mai-shui-guo-xu-yao-de-zui-shao-jin-f6rsy/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
