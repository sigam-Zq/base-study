package leetcode22

import (
	"log"
	"math/big"
	"sort"
	"testing"
)

func TestMain2(t *testing.T) {
	args := []struct {
		name         string
		f            func(rewardValues []int) int
		rewardValues []int
		want         int
	}{
		{
			name:         "oneTest",
			f:            maxTotalReward2,
			rewardValues: []int{1, 6, 4, 3, 2},
			want:         11,
		},
		{
			name:         "twoTest",
			f:            maxTotalReward2,
			rewardValues: []int{1},
			want:         1,
		},
		{
			name:         "threeTest",
			f:            maxTotalReward2,
			rewardValues: []int{1, 10, 9},
			want:         19,
		},
		{
			name: "fourTest",
			// f:            maxTotalReward2,
			f:            maxTotalRewardLeetcode2BitOp,
			rewardValues: []int{84, 27, 99, 32, 12, 37, 39, 47, 45, 73, 52, 20, 47, 47, 50, 85, 35, 9, 10, 47, 45, 54},
			want:         197,
		},
		{
			name:         "fourLeetcodeTest",
			f:            maxTotalRewardLeetcode,
			rewardValues: []int{84, 27, 99, 32, 12, 37, 39, 47, 45, 73, 52, 20, 47, 47, 50, 85, 35, 9, 10, 47, 45, 54},
			want:         197,
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

func maxTotalReward2(rewardValues []int) int {
	// max := 0

	var dfs func(rewardValues []int, idx, sum int, alreadyDone map[int]bool) int

	dfs = func(rewardValues []int, idx, sum int, alreadyDone map[int]bool) int {
		log.Printf(" -->, idx %d, sum %d, alreadyDone %v \n", idx, sum, alreadyDone)

		if sum >= rewardValues[idx] {
			// log.Printf(" sum %d rewardValues[idx] %d \n ", sum, rewardValues[idx])
			return sum
		} else {
			sum += rewardValues[idx]
			alreadyDone[idx] = true
			maxSum := sum
			originSum := sum
			for i, _ := range rewardValues {
				if alreadyDone[i] {
					continue
				}

				// 创建 alreadyDone 的副本（避免引用传递的问题）
				newAlreadyDone := make(map[int]bool)
				for k, v := range alreadyDone {
					newAlreadyDone[k] = v
				}
				log.Printf("max  %d \n", maxSum)
				maxSum = max(dfs(rewardValues, i, originSum, newAlreadyDone), maxSum)
			}
			return maxSum
		}

	}

	// for _, v := range rewardValues {

	// }
	sumAll := make([]int, len(rewardValues))
	for i, _ := range rewardValues {

		alreadyDone := make(map[int]bool)
		for i, _ := range rewardValues {
			alreadyDone[i] = false
		}

		sumAll = append(sumAll, dfs(rewardValues, i, 0, alreadyDone))
	}
	max := 0
	for _, v := range sumAll {
		if max < v {
			max = v
		}
	}

	return max
}

func maxTotalRewardLeetcode(rewardValues []int) int {
	sort.Ints(rewardValues)
	m := rewardValues[len(rewardValues)-1]
	dp := make([]int, 2*m)
	dp[0] = 1
	for _, x := range rewardValues {
		for k := 2*x - 1; k >= x; k-- {
			if dp[k-x] == 1 {
				dp[k] = 1
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] == 1 {
			res = i
		}
	}
	return res
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/maximum-total-reward-using-operations-i/solutions/2959293/zhi-xing-cao-zuo-ke-huo-de-de-zui-da-zon-rd1t/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
/*
方法一：动态规划
记 rewardValues 的最大值为 m，因为最后一次操作前的总奖励一定小于等于 m−1，所以可获得的最大总奖励小于等于 2m−1。假设上一次操作选择的奖励值为 x
1
​
 ，那么执行操作后的总奖励 x≥x
1
​
 ，根据题意，后面任一操作选择的奖励值 x
2
​
  一定都大于 x，从而有 x
2
​
 >x
1
​
 ，因此执行的操作是按照奖励值单调递增的。

根据以上推断，首先将 rewardValues 从小到大进行排序，使用 dp[k] 表示总奖励 k 是否可获得，初始时 dp[0]=1，表示不执行任何操作获得总奖励 0。然后我们对 rewardValues 进行遍历，令当前值为 x，那么对于 k∈[x,2x−1]（将 k 倒序枚举），将 dp[k] 更新为 dp[k−x] ∣ dp[k]（符号 ∣ 表示或操作），表示先前的操作可以获得总奖励 k−x，那么加上 x 后，就可以获取总奖励 k。最后返回 dp 中可以获得的最大总奖励。

*/

func maxTotalRewardLeetcode2BitOp(rewardValues []int) int {
	n := len(rewardValues)
	sort.Ints(rewardValues)
	if n >= 2 && rewardValues[n-2] == rewardValues[n-1]-1 {
		return 2*rewardValues[n-1] - 1
	}
	f0, f1 := big.NewInt(1), big.NewInt(0)
	for _, x := range rewardValues {
		mask, one := big.NewInt(0), big.NewInt(1)
		mask.Sub(mask.Lsh(one, uint(x)), one)
		f1.Lsh(f1.And(f0, mask), uint(x))
		f0.Or(f0, f1)
	}
	return f0.BitLen() - 1
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/maximum-total-reward-using-operations-ii/solutions/2959298/zhi-xing-cao-zuo-ke-huo-de-de-zui-da-zon-8r4o/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/*
记 rewardValues 的最大值为 m，因为最后一次操作前的总奖励一定小于等于 m−1，所以可获得的最大总奖励小于等于 2m−1。

如果 rewardValues 中存在 m−1，那么结果必定为 2m−1，可以直接返回结果。

假设上一次操作选择的奖励值为 x
1
​
 ，那么执行操作后的总奖励 x≥x
1
​
 ，根据题意，后面任一操作选择的奖励值 x
2
​
  一定都大于 x，从而有 x
2
​
 >x
1
​
 ，因此执行的操作是按照奖励值单调递增的。

首先将 rewardValues 从小到大进行排序，使用 dp[k] 表示总奖励 k 是否可获得，初始时 dp[0]=1，表示不执行任何操作获得总奖励 0。然后我们对 rewardValues 进行遍历，令当前值为 x，那么对于 k∈[x,2x−1]（将 k 倒序枚举），将 dp[k] 更新为 dp[k−x] ∣ dp[k]（符号 ∣ 表示或操作），表示先前的操作可以获得总奖励 k−x，那么加上 x 后，就可以获取总奖励 k。

要根据以上推断，一次操作前后的 dp 数组（分别记为 f
1和 f0）之前的递推关系为：f0[k]=f1 [k−x] ∣ f0[k],k∈[x,2x−1]可以用位运算表示为 f0=(maskx (f1 )≪x) ∣ f0 ，其中 maskx
 (f1) 函数表示取 f1 的低 x 位。因此我们可以使用位运算对动态规划进行优化，降低时间复杂度。

作者：力扣官方题解
链接：https://leetcode.cn/problems/maximum-total-reward-using-operations-ii/solutions/2959298/zhi-xing-cao-zuo-ke-huo-de-de-zui-da-zon-8r4o/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/*
让我们逐步解析这段代码中关于 `f0` 和 `f1` 的操作，以及这些操作如何与动态规划关联起来。

### 1. 方法解释

#### `Sub`, `Lsh`, `And`, `Or`

这些方法是 Go 语言的 `math/big` 包中的操作，用于处理大整数。这里的具体作用如下：

- **`Lsh(x, n)`**：将 `x` 左移 `n` 位，相当于 `x * 2^n`。例如，`Lsh(big.NewInt(1), 3)` 会返回 `8`（即 `1 * 2^3`）。

- **`Sub(x, y)`**：计算 `x - y`。用于从 `x` 中减去 `y`。

- **`And(x, y)`**：计算 `x` 和 `y` 的按位与操作。这在此用来找到可以标记的奖励的集合。

- **`Or(x, y)`**：计算 `x` 和 `y` 的按位或操作。用于更新当前的状态集合。

### 2. 函数操作与动态规划的关联

这段代码通过位运算模拟了动态规划的状态转移过程。以下是对代码的详细解释：

- **初始化**：
  ```go
  f0, f1 := big.NewInt(1), big.NewInt(0)
  ```
  - `f0` 表示当前可以获得的奖励的状态，初始化为 `1`，表示没有奖励时的状态。
  - `f1` 用于计算在当前奖励值 `x` 下可以获得的新的状态，初始化为 `0`。

- **循环处理奖励值**：
  ```go
  for _, x := range rewardValues {
      mask, one := big.NewInt(0), big.NewInt(1)
      mask.Sub(mask.Lsh(one, uint(x)), one)
      f1.Lsh(f1.And(f0, mask), uint(x))
      f0.Or(f0, f1)
  }
  ```
  - **创建掩码**：
    ```go
    mask.Sub(mask.Lsh(one, uint(x)), one)
    ```
    这行代码的作用是生成一个掩码，该掩码的作用是限制只能选择 `rewardValues[i]` 大于当前总奖励 `x` 的元素。`mask` 中的位被设置为 `1`，表示可以选择的状态。

  - **更新 `f1`**：
    ```go
    f1.Lsh(f1.And(f0, mask), uint(x))
    ```
    在此行中，`f1` 计算出当前状态下可以添加 `rewardValues[i]` 的新状态。通过 `f0` 和 `mask` 进行按位与操作，确保只有在状态有效的情况下才会进行奖励值的累加。

  - **更新 `f0`**：
    ```go
    f0.Or(f0, f1)
    ```
    将新的状态 `f1` 合并到 `f0` 中，更新可达到的总奖励状态。

### 总结

通过使用位运算，这段代码以高效的方式跟踪可以获得的总奖励。与动态规划的关联在于，它通过维护状态集合 `f0` 和 `f1`，在每一步中通过条件约束来更新可选的状态，从而最终计算出最大总奖励。这种方法可以处理较大的输入，并能有效地跟踪状态变化。
*/
