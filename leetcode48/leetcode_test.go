package leetcode48

import (
	"log"
	"strconv"
	"testing"
)

/*
3250. 单调数组对的数目 I
给你一个长度为 n 的 正 整数数组 nums 。

如果两个 非负 整数数组 (arr1, arr2) 满足以下条件，我们称它们是 单调 数组对：

两个数组的长度都是 n 。
arr1 是单调 非递减 的，换句话说 arr1[0] <= arr1[1] <= ... <= arr1[n - 1] 。
arr2 是单调 非递增 的，换句话说 arr2[0] >= arr2[1] >= ... >= arr2[n - 1] 。
对于所有的 0 <= i <= n - 1 都有 arr1[i] + arr2[i] == nums[i] 。
请你返回所有 单调 数组对的数目。

由于答案可能很大，请你将它对 109 + 7 取余 后返回。



示例 1：

输入：nums = [2,3,2]

输出：4

解释：

单调数组对包括：

([0, 1, 1], [2, 2, 1])
([0, 1, 2], [2, 2, 0])
([0, 2, 2], [2, 1, 0])
([1, 2, 2], [1, 1, 0])
示例 2：

输入：nums = [5,5,5,5]

输出：126



提示：

1 <= n == nums.length <= 2000
1 <= nums[i] <= 50


*/

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f    func([]int) int
		nums []int
		want int
	}{
		{
			f:    countOfPairs,
			nums: []int{2, 3, 2},
			want: 4,
		},
		{
			f:    countOfPairs,
			nums: []int{40, 40, 40, 40, 41, 42, 43, 44, 45, 45},
			want: 272278100,
		},
		{
			f:    countOfPairs,
			nums: []int{40, 40, 40, 40, 40, 40, 41, 41, 41, 42, 42, 43, 43, 43, 44, 44, 44, 45, 45, 45, 45, 45, 46, 46, 46, 46, 46, 46, 46, 47, 47, 47, 47, 48, 48, 49, 49, 49, 49, 49, 49, 50, 50, 50, 50},
			want: 221670461,
		},
		{
			f:    countOfPairsLeetCode,
			nums: []int{40, 40, 40, 40, 40, 40, 41, 41, 41, 42, 42, 43, 43, 43, 44, 44, 44, 45, 45, 45, 45, 45, 46, 46, 46, 46, 46, 46, 46, 47, 47, 47, 47, 48, 48, 49, 49, 49, 49, 49, 49, 50, 50, 50, 50},
			want: 221670461,
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.nums); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})

	}
}

func countOfPairs(nums []int) int {

	mod := int(1e9 + 7)
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, nums[i]+1)
		// dp[i][0] = 1
	}
	for i := 0; i <= nums[0]; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for v2 := 0; v2 <= nums[i]; v2++ {
			for v1 := 0; v1 <= v2; v1++ {
				if nums[i-1]-v1 >= nums[i]-v2 {
					dp[i][v2] = (dp[i][v2] + dp[i-1][v1]) % mod
				}
			}
		}
	}
	log.Printf("dp %v  \n", dp)

	res := 0
	for _, v := range dp[n-1] {
		res = (res + v) % mod
	}

	return res
}

/*
方法一：动态规划
思路与算法

我们使用动态规划来解决本题目，定义 dp[i][j] 表示当 arr1[i]=j 时，前 i+1 个元素组成的单调数组的数目。
因为 arr1[0] 可以为 0 到 nums[0] 之间的任意数，初始化 dp[0][j]=1，其中 j 小于等于 nums[0]，其它初始化为零。
我们遍历数据，并且枚举 arr1中之前和现在的值，按照题目要求的检查单调性，可得到转移方程 dp[i][v2]=v1∑dp[i−1][1]。
其中满足 v
2和 nums[i−1]−v
1≥nums[i]−v2≥0。

最后，我们返回 dp[n−1] 的和即为结果。

代码

作者：力扣官方题解
链接：https://leetcode.cn/problems/find-the-count-of-monotonic-pairs-i/solutions/2992035/dan-diao-shu-zu-dui-de-shu-mu-i-by-leetc-7x5r/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func countOfPairsLeetCode(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 51)
	}
	mod := int(1e9 + 7)

	for v := 0; v <= nums[0]; v++ {
		dp[0][v] = 1
	}

	for i := 1; i < n; i++ {
		for v2 := 0; v2 <= nums[i]; v2++ {
			for v1 := 0; v1 <= v2; v1++ {
				if nums[i-1]-v1 >= nums[i]-v2 {
					dp[i][v2] = (dp[i][v2] + dp[i-1][v1]) % mod
				}
			}
		}
	}

	res := 0
	for _, v := range dp[n-1] {
		res = (res + v) % mod
	}
	return res
}

/*
问


	for v := 0; v <= nums[0]; v++ {
		dp[0][v] = 1
	}
这里的初始化 给dp 中的所有的位置为v的地方写入1 是代表什么

回答


在你的代码中，初始化部分：

```go
for v := 0; v <= nums[0]; v++ {
    dp[0][v] = 1
}
```

这段代码的作用是初始化动态规划数组 `dp` 的第一行。具体来说，它表示在第一个位置（`i = 0`）时，所有可能的 `arr1[0]` 的值（从 `0` 到 `nums[0]`）都可以形成一个有效的单调数组对 `(arr1, arr2)`。

### 解释：

- **`arr1[0]` 的可能值**: `arr1[0]` 可以是从 `0` 到 `nums[0]` 的任何值。
- **`arr2[0]` 的值**: 对于每个 `arr1[0] = v`，`arr2[0]` 必须是 `nums[0] - v`，以满足 `arr1[0] + arr2[0] = nums[0]`。
- **单调性**: 因为这是第一个元素，单调性条件自动满足（没有前一个元素进行比较）。
- **初始化 `dp`**: `dp[0][v] = 1` 表示在第一个位置，`arr1[0]` 取值为 `v` 的情况下，有一种方法可以形成一个有效的单调数组对。

通过这种初始化，后续的动态规划过程可以基于这些初始值，逐步构建出满足条件的单调数组对的数量。每个 `dp[i][v2]` 表示在位置 `i`，`arr1[i]` 取值为 `v2` 的情况下，形成有效单调数组对的数量。

*/
