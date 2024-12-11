package leetcode58

import (
	"strconv"
	"testing"
)

/*
935. 骑士拨号器
象棋骑士有一个独特的移动方式，它可以垂直移动两个方格，水平移动一个方格，或者水平移动两个方格，垂直移动一个方格(两者都形成一个 L 的形状)。

象棋骑士可能的移动方式如下图所示:



我们有一个象棋骑士和一个电话垫，如下所示，骑士只能站在一个数字单元格上(即蓝色单元格)。



给定一个整数 n，返回我们可以拨多少个长度为 n 的不同电话号码。

你可以将骑士放置在任何数字单元格上，然后你应该执行 n - 1 次移动来获得长度为 n 的号码。所有的跳跃应该是有效的骑士跳跃。

因为答案可能很大，所以输出答案模 109 + 7.



示例 1：

输入：n = 1
输出：10
解释：我们需要拨一个长度为1的数字，所以把骑士放在10个单元格中的任何一个数字单元格上都能满足条件。
示例 2：

输入：n = 2
输出：20
解释：我们可以拨打的所有有效号码为[04, 06, 16, 18, 27, 29, 34, 38, 40, 43, 49, 60, 61, 67, 72, 76, 81, 83, 92, 94]
示例 3：

输入：n = 3131
输出：136006598
解释：注意取模


提示：

1 <= n <= 5000

*/

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f    func(int) int
		n    int
		want int
	}{
		{
			f:    knightDialer,
			n:    1,
			want: 10,
		},
		{
			f:    knightDialer,
			n:    2,
			want: 20,
		},
		{
			f:    knightDialer,
			n:    3,
			want: 46,
		},
		{
			f:    knightDialer,
			n:    3131,
			want: 136006598,
		},
	} {

		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.n); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})
	}
}

func knightDialer(n int) int {

	mod := int(1e9 + 7)

	numRoad := [][]int{{4, 6}, {6, 8}, {7, 9}, {4, 8}, {3, 9, 0}, {}, {1, 7, 0}, {2, 6}, {1, 3}, {2, 4}}

	// dp[i][j] 长度为 i+1 以  j结尾 位置上有dp[i][j]个可能
	dp := make([][10]int, n)

	for i := range numRoad {
		dp[0][i] = 1
	}
	// 状态转移方程 dp[i+1] =
	for i := 1; i < n; i++ {

		for j, v := range numRoad {
			// dp[i][j] = dp[i-1][j] * len(v) % mod
			for _, k := range v {
				dp[i][k] += dp[i-1][j] % mod
				// dp[i][k] = (dp[i][k] + dp[i-1][j]) % mod
			}
		}

	}

	var res int

	for i := range numRoad {
		res = (res + dp[n-1][i]) % mod
	}

	return res
}

/*
问
		go 语言中

		res = (res + dp[n-1][i]) % mod

		和
		res += dp[n-1][i] % mod
		这里有什么区别，会导致对 res返回的结果不同么

答

在 Go 语言中，以下两种表达式虽然看起来相似，但它们可能导致不同的结果，这取决于具体的 `mod` 值以及 `res` 的数据范围和初始值。

1. **表达式 1**
   ```go
   res = (res + dp[n-1][i]) % mod
   ```
   这里的计算步骤是：
   - 先将 `res + dp[n-1][i]` 求和。
   - 然后将结果对 `mod` 取模。

   这个表达式确保 `res` 始终保持在 `[0, mod)` 的范围内，不会出现溢出或负数。

2. **表达式 2**
   ```go
   res += dp[n-1][i] % mod
   ```
   这里的计算步骤是：
   - 先计算 `dp[n-1][i] % mod`。
   - 然后将结果加到 `res` 上。

   这里并没有直接对最终的 `res` 值取模，因此 `res` 可能会累积变得非常大，甚至溢出，尤其在循环多次后。

### 可能导致的不同
- 如果没有溢出问题，**理论上最终结果一致**，因为最终可以对 `res` 取模来保持在 `[0, mod)` 范围内。
- 如果可能存在溢出（例如，`res` 的类型为 `int`，`dp[n-1][i]` 累加后超出其表示范围），第二种写法可能导致错误结果甚至崩溃，而第一种写法通过每次取模，避免了这种风险。
- 如果 `res` 有负值的可能性，第二种写法不会处理负值取模问题，可能会导致意料之外的负数结果。

### 推荐做法
为了保证安全性和结果的一致性，推荐使用第一种写法：
```go
res = (res + dp[n-1][i]) % mod
```
这种写法不仅可以防止溢出，还能确保 `res` 始终在合法范围内。
*/

func knightDialerLeetCode1(n int) int {
	const mod = 1_000_000_007
	moves := [][]int{
		{4, 6},
		{6, 8},
		{7, 9},
		{4, 8},
		{3, 9, 0},
		{},
		{1, 7, 0},
		{2, 6},
		{1, 3},
		{2, 4},
	}
	d := [2][10]int{}
	for i := 0; i < 10; i++ {
		d[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		x := i & 1
		for j := 0; j < 10; j++ {
			d[x][j] = 0
			for _, k := range moves[j] {
				d[x][j] = (d[x][j] + d[x^1][k]) % mod
			}
		}
	}
	res := 0
	for _, x := range d[n%2] {
		res = (res + x) % mod
	}
	return res
}

/*
方法二：矩阵快速幂

思路与算法

本题还有另外一种时间复杂度更低的做法，是使用矩阵快速幂。我们把每次 10 个一组的递推视作矩阵相乘，假设现在有一个形状为 1×10 的矩阵 s 表示十个位置的答案，那么它的一次递推大概是做这样一次运算：

[
s
0

	,s

1

	,s

2

	,s

3

	,s

4

	,s

5

	,s

6

	,s

7

	,s

8

	,s

9

	]×

⎣
⎡
​
0,0,0,0,1,0,1,0,0,0
0,0,0,0,0,0,1,0,1,0
0,0,0,0,0,0,0,1,0,1
0,0,0,0,1,0,0,0,1,0
1,0,0,1,0,0,0,0,0,1
0,0,0,0,0,0,0,0,0,0
1,1,0,0,0,0,0,1,0,0
0,0,1,0,0,0,1,0,0,0
0,1,0,1,0,0,0,0,0,0
0,0,1,0,1,0,0,0,0,0
​

⎦
⎤
​

我们把右侧那个矩阵称作 base，因为它是我们后面每次转移的基座，我们注意观察它第四列的元素，在第 0,3,9 行处为 1，其他地方为 0，那么相乘结果就是 s
0

	+s

2

	+s

9

	，这正是 s

4
′

	的答案。

因此，我们其实只需要求出 base
n−1

	， 就可以快速的求出答案了！而求解 base

n−1

	可以使用矩阵快速幂加速，因此复杂度将降低至 O(10

3

	logn)，由于这里的 n 最大为 5000，所以实际优化不明显，但它的优势将在 n 更大时显现。

作者：力扣官方题解
链接：https://leetcode.cn/problems/knight-dialer/solutions/3002022/qi-shi-bo-hao-qi-by-leetcode-solution-5x9m/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
const mod = 1_000_000_007

func knightDialerLeetCode2(n int) int {
	base := [][]int{
		{0, 0, 0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 1, 0},
		{1, 0, 0, 1, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0, 0, 0},
	}
	res := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	base2 := make([][]int, 10)
	for i := 0; i < 10; i++ {
		base2[i] = make([]int, 10)
		base2[i][i] = 1
	}
	n--
	for n > 0 {
		if n&1 == 1 {
			base2 = mul(base2, base)
		}
		base = mul(base, base)
		n >>= 1
	}
	res = mul(res, base2)
	ret := 0
	for _, x := range res[0] {
		ret = (ret + x) % mod
	}
	return ret
}

func mul(lth, rth [][]int) [][]int {
	rows, cols, inner := len(lth), len(rth[0]), len(lth[0])
	res := make([][]int, rows)
	for i := range res {
		res[i] = make([]int, cols)
	}
	for k := 0; k < inner; k++ {
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				res[i][j] = (res[i][j] + lth[i][k]*rth[k][j]%mod) % mod
			}
		}
	}
	return res
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/knight-dialer/solutions/3002022/qi-shi-bo-hao-qi-by-leetcode-solution-5x9m/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
