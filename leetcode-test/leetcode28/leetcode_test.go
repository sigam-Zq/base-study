package leetcode28

import (
	"log"
	"strconv"
	"testing"
)

/*

来自未来的体育科学家给你两个整数数组 energyDrinkA 和 energyDrinkB，数组长度都等于 n。这两个数组分别代表 A、B 两种不同能量饮料每小时所能提供的强化能量。

你需要每小时饮用一种能量饮料来 最大化 你的总强化能量。然而，如果从一种能量饮料切换到另一种，你需要等待一小时来梳理身体的能量体系（在那个小时里你将不会获得任何强化能量）。

返回在接下来的 n 小时内你能获得的 最大 总强化能量。

注意 你可以选择从饮用任意一种能量饮料开始。



示例 1：

输入：energyDrinkA = [1,3,1], energyDrinkB = [3,1,1]

输出：5

解释：

要想获得 5 点强化能量，需要选择只饮用能量饮料 A（或者只饮用 B）。

示例 2：

输入：energyDrinkA = [4,1,1], energyDrinkB = [1,1,3]

输出：7

解释：

第一个小时饮用能量饮料 A。
切换到能量饮料 B ，在第二个小时无法获得强化能量。
第三个小时饮用能量饮料 B ，并获得强化能量。


提示：

n == energyDrinkA.length == energyDrinkB.length
3 <= n <= 105
1 <= energyDrinkA[i], energyDrinkB[i] <= 105
*/

func TestXxx(t *testing.T) {

	for idx, v := range []struct {
		f            func([]int, []int) int64
		energyDrinkA []int
		energyDrinkB []int
		want         int64
	}{
		{
			f:            maxEnergyBoost,
			energyDrinkA: []int{1, 3, 1},
			energyDrinkB: []int{3, 1, 1},
			want:         int64(5),
		},
		{
			f:            maxEnergyBoost,
			energyDrinkA: []int{3, 3, 3, 5},
			energyDrinkB: []int{5, 3, 6, 5},
			want:         int64(19),
		},
		{
			f:            maxEnergyBoost,
			energyDrinkA: []int{5, 5, 6, 3, 4, 3, 3, 4},
			energyDrinkB: []int{5, 3, 3, 4, 4, 6, 6, 3},
			want:         int64(35),
		},
		{
			f:            maxEnergyBoostClaude,
			energyDrinkA: []int{1801, 60716, 37726, 68357, 81497, 93047, 51014, 72664, 25041, 63261, 82051, 44136, 24746, 49775, 39970, 20604, 31542, 40100, 55637, 96719, 56595, 31142, 77691, 82164, 53571, 15442, 95027, 49323, 76569, 47341, 20653, 68722, 49053, 29387, 11164, 13307, 23164, 46559, 50155, 85669, 57631, 31633, 30033, 45109, 21089, 96791, 29291, 73574, 48852, 67340, 83198, 2637, 34596, 87978, 82634, 32677, 18544, 26605, 10359, 50515, 90720, 14290, 5836, 63192, 86348, 51730, 14873, 62180, 70882, 54944, 13626, 96508, 48080, 47597, 53356, 66227, 14753, 47986, 89858, 30075, 40395, 98629, 17142, 73763, 41998, 51960, 14163, 12808, 33186, 72197, 62342, 64482, 39815, 25881, 69880, 12308, 25338, 69446, 95197, 70375},
			energyDrinkB: []int{87762, 48631, 29496, 98673, 82273, 70847, 91496, 47327, 27175, 66668, 93200, 97212, 71224, 6825, 66542, 46198, 62742, 98507, 19460, 60154, 89480, 72347, 97455, 49196, 91136, 30641, 52875, 10239, 11773, 20744, 20997, 75931, 42517, 25594, 87335, 86965, 87538, 15071, 10099, 21836, 89541, 83028, 84120, 47402, 52302, 50255, 95583, 10530, 65406, 14632, 30141, 60801, 97478, 51898, 69088, 95343, 51984, 47046, 52653, 74222, 76970, 51471, 48726, 20043, 12519, 71411, 31124, 89924, 6212, 224, 47824, 39655, 99606, 49848, 49694, 67176, 72858, 88032, 67789, 86604, 35969, 63301, 50995, 76069, 12816, 88380, 27420, 89934, 65464, 44451, 54694, 21096, 13210, 89821, 5536, 23074, 92064, 82735, 11068, 45899},
			want:         int64(5845167),
		},
	} {

		t.Run(strconv.Itoa(idx)+"-test", func(t *testing.T) {
			if got := v.f(v.energyDrinkA, v.energyDrinkB); got != v.want {
				t.Errorf(" got %v  want %v \n", got, v.want)
			}
		})

	}
}

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	// k - v (从 k索引开始后续所能量值的和)
	drinkARear := make([]int, len(energyDrinkA))
	drinkBRear := make([]int, len(energyDrinkB))

	drinkARear[len(energyDrinkA)-1] = energyDrinkA[len(energyDrinkA)-1]
	drinkBRear[len(energyDrinkA)-1] = energyDrinkB[len(energyDrinkA)-1]
	for i := len(energyDrinkA) - 2; i >= 0; i-- {
		drinkARear[i] = energyDrinkA[i] + drinkARear[i+1]
		drinkBRear[i] = energyDrinkB[i] + drinkBRear[i+1]
	}
	log.Printf("drinkARear %v, drinkBRear %v \n", drinkARear, drinkBRear)
	max := 0
	// 这里解法的瑕疵是  这里仅有一次换赛道的情况 排除了所有 AB 来回切换的情况
	// 前面的累加也没有 加进去
	// sumA, sumB := 0, 0
	// for i := 0; i < len(energyDrinkA)-2; i++ {
	// 	sumA += energyDrinkA[i]
	// 	sumB += energyDrinkB[i]
	// 	if max < (sumA + drinkARear[i+1]) {
	// 		max = sumA + drinkARear[i+1]
	// 	}

	// 	if max < (sumB + drinkBRear[i+1]) {
	// 		// log.Printf("energyDrinkB[i]  %d ,drinkBRear[i+1] %d \n ", energyDrinkB[i], drinkBRear[i+1])
	// 		max = sumB + drinkBRear[i+1]
	// 	}

	// 	if max < (sumA + drinkBRear[i+2]) {
	// 		max = sumA + drinkBRear[i+2]
	// 	}

	// 	if max < (sumB + drinkARear[i+2]) {
	// 		max = sumB + drinkARear[i+2]
	// 	}
	// }

	var dfs func(isA bool, nowSum, startPoint int)
	dfs = func(isA bool, nowSum, startPoint int) {
		// 小于数组长度就继续
		if startPoint >= len(energyDrinkA)-2 {
			return
		}

		// log.Printf("isA  %v ,nowSum %d ,startPoint %d \n ", isA, nowSum, startPoint)
		if isA {
			nowSum += energyDrinkA[startPoint]

			// 把后面全部选一种的情况列出来对比
			if max < nowSum+drinkARear[startPoint+1] {
				max = nowSum + drinkARear[startPoint+1]
			}

			if max < nowSum+drinkBRear[startPoint+2] {
				max = nowSum + drinkBRear[startPoint+2]
			}

		} else {
			nowSum += energyDrinkB[startPoint]

			if max < nowSum+drinkBRear[startPoint+1] {
				max = nowSum + drinkBRear[startPoint+1]
				// log.Printf("isA  %v ,nowSum %d ,startPoint %d ,max %d\n ", isA, nowSum, startPoint, max)
			}

			if max < nowSum+drinkARear[startPoint+2] {
				max = nowSum + drinkARear[startPoint+2]
			}

		}

		// 两种选择 走A ，或者切B
		dfs(!isA, nowSum, startPoint+2)
		dfs(isA, nowSum, startPoint+1)
	}

	dfs(true, 0, 0)
	dfs(false, 0, 0)

	return int64(max)
}

func maxEnergyBoostClaude(energyDrinkA []int, energyDrinkB []int) int64 {
	n := len(energyDrinkA)
	// dp[i][0]: 第i位置选A的最大值
	// dp[i][1]: 第i位置选B的最大值
	dp := make([][2]int, n)

	// 初始化第一个位置
	dp[0][0] = energyDrinkA[0]
	dp[0][1] = energyDrinkB[0]

	if n == 1 {
		return int64(max(dp[0][0], dp[0][1]))
	}

	// 初始化第二个位置
	dp[1][0] = max(energyDrinkA[1], dp[0][0]+energyDrinkA[1])
	dp[1][1] = max(energyDrinkB[1], dp[0][1]+energyDrinkB[1])

	for i := 2; i < n; i++ {
		// 选A的情况：要么直接选A，要么从i-2位置的B转换过来
		dp[i][0] = max(dp[i-1][0]+energyDrinkA[i], dp[i-2][1]+energyDrinkA[i])
		// 选B的情况：要么直接选B，要么从i-2位置的A转换过来
		dp[i][1] = max(dp[i-1][1]+energyDrinkB[i], dp[i-2][0]+energyDrinkB[i])
	}

	return int64(max(dp[n-1][0], dp[n-1][1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
让我分析一下这个解法：

**1. 算法正确性分析：**

这个算法存在几个问题：

1. **边界条件处理不当**：
```go
if startPoint >= len(energyDrinkA)-2 {
    return
}
```
这个终止条件可能会导致漏掉最后两个位置的处理。应该改为：
```go
if startPoint >= len(energyDrinkA) {
    return
}
```

2. **最大值初始化问题**：
```go
max := 0
```
如果输入数组都是负数，返回0是不正确的。应该初始化为数组中第一个元素的最大值。

3. **后缀和的使用方式**：虽然计算了后缀和，但在DFS中的使用可能会导致重复计算。

**2. 时间复杂度分析：**

当前DFS的时间复杂度是指数级的：O(2^n)
- 每个位置都有两个选择（继续当前饮料或切换）
- 每次可以前进1步或2步
- 没有记忆化处理，导致大量重复计算

**优化方向：**

1. **使用动态规划替代DFS**：
可以定义dp[i][0]表示在第i个位置选择A饮料的最大值，dp[i][1]表示选择B饮料的最大值。时间复杂度可以优化到O(n)。

2. **如果保留DFS，添加记忆化**：
```go
// 添加记忆化数组
memo := make([][2]int, len(energyDrinkA))
for i := range memo {
    memo[i] = [2]int{-1, -1}
}
```

这是一个优化后的动态规划解法示例：

```go
func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
    n := len(energyDrinkA)
    // dp[i][0]: 第i位置选A的最大值
    // dp[i][1]: 第i位置选B的最大值
    dp := make([][2]int, n)

    // 初始化第一个位置
    dp[0][0] = energyDrinkA[0]
    dp[0][1] = energyDrinkB[0]

    if n == 1 {
        return int64(max(dp[0][0], dp[0][1]))
    }

    // 初始化第二个位置
    dp[1][0] = max(energyDrinkA[1], dp[0][0] + energyDrinkA[1])
    dp[1][1] = max(energyDrinkB[1], dp[0][1] + energyDrinkB[1])

    for i := 2; i < n; i++ {
        // 选A的情况：要么直接选A，要么从i-2位置的B转换过来
        dp[i][0] = max(dp[i-1][0] + energyDrinkA[i], dp[i-2][1] + energyDrinkA[i])
        // 选B的情况：要么直接选B，要么从i-2位置的A转换过来
        dp[i][1] = max(dp[i-1][1] + energyDrinkB[i], dp[i-2][0] + energyDrinkB[i])
    }

    return int64(max(dp[n-1][0], dp[n-1][1]))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

这个优化后的解法：
1. 时间复杂度：O(n)
2. 空间复杂度：O(n)
3. 更容易理解和维护
4. 避免了递归调用的开销
5. 不会有栈溢出的风险
*/
