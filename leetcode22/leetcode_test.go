package leetcode22

import (
	"math"
	"sort"
	"testing"
)

/*
3180. 执行操作可获得的最大总奖励 I

给你一个整数数组 rewardValues，长度为 n，代表奖励的值。

最初，你的总奖励 x 为 0，所有下标都是 未标记 的。你可以执行以下操作 任意次 ：

从区间 [0, n - 1] 中选择一个 未标记 的下标 i。
如果 rewardValues[i] 大于 你当前的总奖励 x，则将 rewardValues[i] 加到 x 上（即 x = x + rewardValues[i]），并 标记 下标 i。
以整数形式返回执行最优操作能够获得的 最大 总奖励。



示例 1：

输入：rewardValues = [1,1,3,3]

输出：4

解释：

依次标记下标 0 和 2，总奖励为 4，这是可获得的最大值。

示例 2：

输入：rewardValues = [1,6,4,3,2]

输出：11

解释：

依次标记下标 0、2 和 1。总奖励为 11，这是可获得的最大值。



提示：

1 <= rewardValues.length <= 2000
1 <= rewardValues[i] <= 2000

*/

func TestMain(t *testing.T) {
	args := []struct {
		name         string
		f            func(rewardValues []int) int
		rewardValues []int
		want         int
	}{
		{
			name:         "oneTest",
			f:            maxTotalReward,
			rewardValues: []int{1, 6, 4, 3, 2},
			want:         11,
		},
		{
			name:         "twoTest",
			f:            maxTotalReward,
			rewardValues: []int{3, 4, 17, 11},
			want:         32,
		},
		{
			name:         "fourTest",
			f:            maxTotalRewardClaude2,
			rewardValues: []int{3, 4, 17, 11},
			want:         32,
		},
		{
			name:         "threeTest",
			f:            maxTotalRewardClaude,
			rewardValues: []int{3, 4, 17, 11},
			want:         32,
		},
		{
			name:         "gptOneTest",
			f:            maxTotalRewardChatGpt,
			rewardValues: []int{3, 4, 17, 11},
			want:         32,
		},
	}

	for _, v := range args {
		t.Run(v.name, func(t *testing.T) {
			if got := v.f(v.rewardValues); v.want != got {
				t.Errorf(" got %d  want %d \n", got, v.want)
			}
		})
	}

}
func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	rewardValues = removeDuplicationMap(rewardValues)
	if len(rewardValues) == 0 {
		return 0
	} else if len(rewardValues) == 1 {
		return rewardValues[0]
	}

	diff := rewardValues[len(rewardValues)-1] - rewardValues[len(rewardValues)-2]
	sum := rewardValues[len(rewardValues)-1] + rewardValues[len(rewardValues)-2]

	if len(rewardValues) >= 3 {
		for i := len(rewardValues) - 3; i >= 0; i-- {
			if diff > rewardValues[i] {
				sum += rewardValues[i]
				diff -= rewardValues[i] // 更新diff为新的差值
			}
		}
	}

	return sum
}

func removeDuplicationMap(arr []int) []int {
	set := make(map[int]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}
	return arr[:j]
}

func maxTotalRewardClaude(rewardValues []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(rewardValues)))
	rewardValues = removeDuplicationMap(rewardValues)
	if len(rewardValues) == 0 {
		return 0
	} else if len(rewardValues) == 1 {
		return rewardValues[0]
	}

	sum := rewardValues[0]
	for i := 1; i < len(rewardValues); i++ {
		if rewardValues[i] <= sum {
			sum += rewardValues[i]
		} else {
			break
		}
	}

	return sum
}

func maxTotalRewardClaude2(rewardValues []int) int {
	sort.Ints(rewardValues)
	n := len(rewardValues)
	dp := make([]int, 2001) // 因为 1 <= rewardValues[i] <= 2000

	for i := 0; i < n; i++ {
		for j := 2000; j >= 0; j-- {
			if rewardValues[i] > j {
				dp[rewardValues[i]] = max(dp[rewardValues[i]], j+rewardValues[i])
			}
		}
	}

	maxReward := 0
	for _, v := range dp {
		maxReward = max(maxReward, v)
	}

	return maxReward
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*


您说得对，我为之前的错误再次道歉。您指出了一个非常重要的点：我们需要考虑所有可能的组合，而不仅仅是贪心地选择大于当前和的值。这个问题实际上需要使用动态规划来解决。让我们来看一个正确的解决方案：

```go:path/to/file
func maxTotalReward(rewardValues []int) int {
    sort.Ints(rewardValues)
    n := len(rewardValues)
    dp := make([]int, 2001) // 因为 1 <= rewardValues[i] <= 2000

    for i := 0; i < n; i++ {
        for j := 2000; j >= 0; j-- {
            if rewardValues[i] > j {
                dp[rewardValues[i]] = max(dp[rewardValues[i]], j+rewardValues[i])
            }
        }
    }

    maxReward := 0
    for _, v := range dp {
        maxReward = max(maxReward, v)
    }

    return maxReward
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

这个解决方案的工作原理如下：

1. 我们首先对奖励值进行升序排序。

2. 我们创建一个 dp 数组，其中 dp[i] 表示当前总和为 i 时可以获得的最大奖励。

3. 我们遍历每个奖励值，对于每个奖励值，我们从大到小更新 dp 数组。

4. 如果当前奖励值大于当前的总和 j，我们就可以选择这个奖励值，更新 dp[rewardValues[i]]。

5. 最后，我们遍历 dp 数组找出最大值，这就是我们能获得的最大总奖励。

让我们用 [3,4,17,11] 这个例子来验证：

- 首先添加 3：dp[3] = 3
- 添加 4：dp[4] = 4
- 添加 11：dp[11] = 11, dp[15] = 15 (3+11)
- 添加 17：dp[17] = 17, dp[21] = 21 (4+17), dp[28] = 28 (11+17), dp[32] = 32 (4+11+17)

最终结果是 32，这是正确的答案。

这个解决方案的时间复杂度是 O(n * M)，其中 n 是奖励值的数量，M 是奖励值的最大可能值（在这个问题中是 2000）。空间复杂度是 O(M)。

非常感谢您的耐心和指出问题。这个问题确实比初看起来要复杂得多，需要考虑所有可能的组合。如果您还有任何疑问或需要进一步的解释，请随时告诉我。

*/

func maxTotalRewardChatGpt(rewardValues []int) int {
	n := len(rewardValues)
	// 初始化dp数组，长度为n，初始值均为0
	dp := make([]int, n)

	// 排序数组
	sort.Ints(rewardValues)

	maxReward := 0 // 用于存储最大总奖励

	for i := 0; i < n; i++ {
		dp[i] = rewardValues[i] // 每个位置先初始化为自己本身的值
		for j := 0; j < i; j++ {
			// 如果前面的dp[j]可以累加当前的rewardValues[i]
			if rewardValues[i] > dp[j] {
				dp[i] = max(dp[i], dp[j]+rewardValues[i])
			}
		}
		// 更新最大总奖励
		maxReward = max(maxReward, dp[i])
	}

	return maxReward
}

// 返回两个整数中的最大值
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

/*
使用动态规划的思路确实可以更好地解决这个问题。我们可以用动态规划的方式来追踪在不同状态下所能获得的最大奖励值。

### 动态规划思路
定义一个动态规划数组 `dp`，其中 `dp[i]` 表示当处理到 `rewardValues` 中的第 `i` 个元素时，能够获得的最大总奖励。

#### 状态转移方程
1. 如果我们选择 `rewardValues[i]` 加到当前总奖励 `x` 中，则需要满足条件 `rewardValues[i] > x`。
2. 我们遍历每个元素 `rewardValues[i]`，更新 `dp[i]`，并根据上一次累加的结果 `dp[j]` 来判断是否可以选择当前值 `rewardValues[i]`。
3. 最后，我们需要遍历所有 `dp[i]` 来找到最大的累加值。

#### 实现代码
代码实现如下：

```go
func maxTotalReward(rewardValues []int) int {
    n := len(rewardValues)
    // 初始化dp数组，长度为n，初始值均为0
    dp := make([]int, n)

    // 排序数组
    sort.Ints(rewardValues)

    maxReward := 0 // 用于存储最大总奖励

    for i := 0; i < n; i++ {
        dp[i] = rewardValues[i] // 每个位置先初始化为自己本身的值
        for j := 0; j < i; j++ {
            // 如果前面的dp[j]可以累加当前的rewardValues[i]
            if rewardValues[i] > dp[j] {
                dp[i] = max(dp[i], dp[j]+rewardValues[i])
            }
        }
        // 更新最大总奖励
        maxReward = max(maxReward, dp[i])
    }

    return maxReward
}

// 返回两个整数中的最大值
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### 示例分析

#### 示例 `[1,1,3,3]`
- 排序后为 `[1,1,3,3]`
- `dp` 数组的更新过程为：
    - `dp[0] = 1`
    - `dp[1] = 1`
    - `dp[2] = 4` （`1 + 3`）
    - `dp[3] = 4` （因为 `dp[2]` 已经是 4，无法增加新的值）

最终 `maxReward = 4`

#### 示例 `[1,6,4,3,2]`
- 排序后为 `[1,2,3,4,6]`
- `dp` 数组的更新过程为：
    - `dp[0] = 1`
    - `dp[1] = 3` （`1 + 2`）
    - `dp[2] = 6` （`3 + 3`）
    - `dp[3] = 10` （`6 + 4`）
    - `dp[4] = 11` （因为 `10 + 6` 不满足条件）

最终 `maxReward = 11`

### 复杂度分析
- **时间复杂度**：`O(n^2)`，因为我们遍历每个 `i` 时，还要遍历所有小于 `i` 的元素 `j`。
- **空间复杂度**：`O(n)`，用于存储 `dp` 数组。

通过这种动态规划的方式，我们可以确保获得的总奖励是最大值。
*/

// depth first search
func maxTotalRewardDFS(rewardValues []int) int {
	sort.Ints(rewardValues) // 对 rewardValues 排序，使得小值优先被考虑
	return dfs(rewardValues, 0, 0)
}

// dfs 从当前索引 start 开始，尝试获取最大总奖励
func dfs(rewardValues []int, start int, x int) int {
	if start == len(rewardValues) {
		return x
	}

	// 不选择当前元素，直接递归到下一个元素
	maxReward := dfs(rewardValues, start+1, x)

	// 如果选择当前元素
	if rewardValues[start] > x {
		maxReward = max(maxReward, dfs(rewardValues, start+1, x+rewardValues[start]))
	}

	return maxReward
}

// 获取最大值
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

/*

可以使用广度优先搜索（BFS）或深度优先搜索（DFS）来解决此问题，但由于这是一个优化问题，而不是单纯的路径搜索问题，广度或深度遍历不一定是最优解。贪心算法在这里更有效率，因为每次选择对当前情况最有利的奖励值，以便尽可能快速地满足条件，获得最大奖励。

不过，假设我们要用 DFS 或 BFS，则需要把每种选择视为一棵状态树的节点，递归或迭代地搜索所有可能的选择路径。以下是如何使用 DFS 来解决这个问题。

### DFS 解法

1. **状态定义**：`dfs(i, x)` 表示当前总奖励 `x`，从下标 `i` 开始选择未标记的奖励。
2. **状态转移**：对于每个元素，我们有两个选择：
   - 不选当前元素，直接跳到下一个。
   - 如果当前元素的值 `rewardValues[i]` 大于当前 `x`，则选择当前元素并递归。
3. **返回最大值**：递归到达数组末尾时返回 `x` 的最大值。

### 代码实现（DFS）

```go
package main

import (
	"fmt"
	"sort"
)

func maxRewardDFS(rewardValues []int) int {
	sort.Ints(rewardValues) // 对 rewardValues 排序，使得小值优先被考虑
	return dfs(rewardValues, 0, 0)
}

// dfs 从当前索引 start 开始，尝试获取最大总奖励
func dfs(rewardValues []int, start int, x int) int {
	if start == len(rewardValues) {
		return x
	}

	// 不选择当前元素，直接递归到下一个元素
	maxReward := dfs(rewardValues, start+1, x)

	// 如果选择当前元素
	if rewardValues[start] > x {
		maxReward = max(maxReward, dfs(rewardValues, start+1, x+rewardValues[start]))
	}

	return maxReward
}

// 获取最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 示例 1
	rewardValues1 := []int{1, 1, 3, 3}
	fmt.Println("最大总奖励:", maxRewardDFS(rewardValues1)) // 输出: 4

	// 示例 2
	rewardValues2 := []int{1, 6, 4, 3, 2}
	fmt.Println("最大总奖励:", maxRewardDFS(rewardValues2)) // 输出: 11
}
```

### 复杂度分析

- **时间复杂度**：这种 DFS 方法的时间复杂度是指数级 \(O(2^n)\)，因为每个元素都有两种选择（选择或不选择），这在输入较大时会导致较高的计算开销。
- **空间复杂度**：递归栈的空间复杂度为 \(O(n)\)。

这种解法虽然可以得到正确答案，但相比贪心算法效率低，因此在这个问题上并不推荐。

*/

func maxRewardDP(rewardValues []int) int {
	// 获取 rewardValues 的总和来定义 dp 数组的大小
	sumRewards := 0
	for _, reward := range rewardValues {
		sumRewards += reward
	}

	// 初始化 dp 数组，设 dp[0] 为 0，其他为无穷小的负数（表示无法达到的状态）
	dp := make([]int, sumRewards+1)
	for i := 1; i <= sumRewards; i++ {
		dp[i] = math.MinInt32
	}

	// 遍历每个 reward 值，更新 dp 数组
	for _, reward := range rewardValues {
		// 为了避免重复更新，倒序更新 dp
		for x := sumRewards; x >= 0; x-- {
			if dp[x] != math.MinInt32 && x < reward {
				dp[x+reward] = max(dp[x+reward], dp[x]+reward)
			}
		}
	}

	// 找出 dp 数组中的最大值
	maxReward := 0
	for _, val := range dp {
		if val > maxReward {
			maxReward = val
		}
	}
	return maxReward
}

/*
确实，贪心算法在这种情况下不能保证最优结果。针对每次「从最小的可选元素开始加」的策略，在一些情况下无法得到最大总奖励，因为这种策略会错过那些稍大但可以增加更多奖励的值。

### 问题的改进方向

为了优化，我们可以考虑不同的选择顺序，动态规划（DP）是一种更适合的方法。使用 DP 可以记录每种总奖励 `x` 状态下的最优结果，从而避免因一次错误选择导致的全局最优损失。

### 动态规划思路

1. **定义状态**：用一个 DP 数组 `dp[x]` 表示在某个总奖励为 `x` 时的可达状态。
2. **状态转移**：对于每个 `reward` 值，我们检查是否可以将其加入到当前的总奖励上。可以通过将 `reward` 值逐步累加到 `dp` 中每个当前有效的总奖励上，逐步扩大可能的总奖励。
3. **更新 DP 状态**：对于每个 `reward`，我们从当前状态中所有符合 `reward > x` 的 `x` 状态更新 `dp[x + reward]`，并确保每次选择路径都能达到最大奖励。
4. **最终答案**：`dp` 数组中的最大值就是结果。

### 动态规划代码实现

以下是具体代码实现：

```go
package main

import (
	"fmt"
	"math"
)

func maxRewardDP(rewardValues []int) int {
	// 获取 rewardValues 的总和来定义 dp 数组的大小
	sumRewards := 0
	for _, reward := range rewardValues {
		sumRewards += reward
	}

	// 初始化 dp 数组，设 dp[0] 为 0，其他为无穷小的负数（表示无法达到的状态）
	dp := make([]int, sumRewards+1)
	for i := 1; i <= sumRewards; i++ {
		dp[i] = math.MinInt32
	}

	// 遍历每个 reward 值，更新 dp 数组
	for _, reward := range rewardValues {
		// 为了避免重复更新，倒序更新 dp
		for x := sumRewards; x >= 0; x-- {
			if dp[x] != math.MinInt32 && x < reward {
				dp[x+reward] = max(dp[x+reward], dp[x] + reward)
			}
		}
	}

	// 找出 dp 数组中的最大值
	maxReward := 0
	for _, val := range dp {
		if val > maxReward {
			maxReward = val
		}
	}
	return maxReward
}

// 获取最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 示例 1
	rewardValues1 := []int{1, 1, 3, 3}
	fmt.Println("最大总奖励:", maxRewardDP(rewardValues1)) // 输出: 4

	// 示例 2
	rewardValues2 := []int{1, 6, 4, 3, 2}
	fmt.Println("最大总奖励:", maxRewardDP(rewardValues2)) // 输出: 11
}
```

### 复杂度分析

- **时间复杂度**：由于我们需要遍历每个 `reward` 值，并对 `dp` 数组进行更新，总时间复杂度为 \(O(n \cdot \text{sumRewards})\)，其中 `n` 是 `rewardValues` 的长度，`sumRewards` 是 `rewardValues` 的和。
- **空间复杂度**：空间复杂度为 \(O(\text{sumRewards})\)，用于保存 `dp` 数组。

### 解释

动态规划解法确保我们不会错过可能的最佳路径，因为每次更新 `dp[x + reward]`，都考虑到了所有可行的奖励路径。
*/
