package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	TestFunc(countSpecialNumbers, 60020144)
	println("--------------------------")
	TestFunc(countSpecialNumbersOptimize, 60020144)
	println("--------------------------")
	TestFunc(countSpecialNumbersChatGPT, 60020144)
	println("--------------------------")
	TestFunc(countSpecialNumbersOptimizeGPTTwo, 60020144)
	println("--------------------------")
	TestFunc(countSpecialNumbersGptFake, 60020144)

}

func countSpecialNumbers(n int) int {
	onHave := 0
	// haveNum := []int{}
	for i1 := 1; i1 <= n; i1++ {
		atoms := make(map[int]struct{})
		for i2 := i1; i2 > 0; i2 /= 10 {
			if _, ok := atoms[i2%10]; ok {
				onHave++
				// haveNum = append(haveNum, i1)
				break
			} else {
				atoms[i2%10] = struct{}{}
			}
		}

	}
	// log.Printf("%v \n", haveNum)
	return n - onHave
}

// 转换为字符串处理
func countSpecialNumbersOptimize(n int) int {
	onHave := 0
	// haveNum := []int{}
	for i1 := 1; i1 <= n; i1++ {
		i1Str := strconv.Itoa(i1)
		atoms := make(map[byte]struct{})
		for i2 := 0; i2 < len(i1Str); i2++ {
			charO := i1Str[i2]
			if _, ok := atoms[charO]; ok {
				onHave++
				break
			} else {
				atoms[charO] = struct{}{}
			}
		}

	}
	// log.Printf("%v \n", haveNum)
	return n - onHave
}

func TestFunc(f func(n int) int, s int) {
	defer timeCost(time.Now())
	log.Println(f(s))
	// defer PrintMemUsage()

}

func timeCost(sT time.Time) {
	tc := time.Since(sT)
	fmt.Printf("run Cost %v \n", tc)
}

/*
你的算法基本思路是通过逐个遍历 1 到 n 之间的所有数字，然后检查每个数字是否有重复的数字位。这个方法的时间复杂度是 O(n)，当 n 较大时效率会很低。优化方向可以从减少遍历次数和跳过不必要的检查来进行改进。

以下是一些优化方向：

1. 数学组合数的优化
这个问题实际上可以通过排列组合来解决，避免逐个检查所有数字。思路是：

一个 d 位的特殊整数可以用排列组合的思想来计算。例如，计算 3 位数的特殊整数时，第一位可以是 1-9 的任意数字（9 种选择），第二位可以是 0-9 中除去第一位的数字（9 种选择），第三位则可以是剩下的 8 种选择。

总体思路是先计算小于 n 位数的所有特殊整数，再处理 n 本身的特殊性。

2. 递归 + 记忆化搜索
可以通过递归来解决该问题。具体步骤是：

使用 DFS（深度优先搜索）的方法遍历数字的每一位。
使用一个数组 used 来记录当前使用的数字，避免重复。
对于每一位，选择一个可用的数字，并递归处理后续位。
改进版代码
以下是基于排列组合和深度优先搜索的优化版本代码：
*/
func countSpecialNumbersChatGPT(n int) int {
	// 将数字 n 转换为字符串，便于逐位处理
	digits := []int{}
	for temp := n; temp > 0; temp /= 10 {
		digits = append([]int{temp % 10}, digits...)
	}

	// 计算小于给定位数的所有特殊数字个数
	count := func(d int) int {
		if d == 0 {
			return 0
		}
		result := 0
		used := 9
		for i := 0; i < d; i++ {
			if i == 0 {
				result += 9 // 第一位不能为 0
			} else {
				result *= used
				used--
			}
		}
		return result
	}
	// 处理位数小于 n 位的数字
	total := 0
	for i := 1; i < len(digits); i++ {
		total += count(i)
	}

	// DFS 递归计算与 n 同位数的情况
	var dfs func(pos int, isLimit bool, used []bool) int
	dfs = func(pos int, isLimit bool, used []bool) int {
		if pos == len(digits) {
			return 1 // 已经处理完所有位，找到一个特殊数字
		}
		limit := digits[pos]
		result := 0

		// 枚举当前位可以取的所有数字
		for d := 0; d <= 9; d++ {
			if d == 0 && pos == 0 { // 第一位不能为 0
				continue
			}
			if used[d] { // 不能重复使用数字
				continue
			}
			if isLimit && d > limit {
				break // 当前位超过限制
			}

			used[d] = true
			// 递归处理下一位
			result += dfs(pos+1, isLimit && d == limit, used)
			used[d] = false
		}

		return result
	}

	// 初始化 used 数组
	used := make([]bool, 10)

	// 递归计算与 n 同位数的特殊数字
	total += dfs(0, true, used)

	return total
}

// 20240921
/*

详细解释：
count 函数：

用于计算小于 n 位数的所有特殊数字个数，计算的是排列组合。
递归 dfs 函数：

pos 表示当前处理的数字位。
isLimit 表示当前位是否受 n 的限制。如果当前位是受限制的，那么只能选择不大于 n 当前位的数字；如果不受限制，意味着我们可以选择任意数字。
used 数组记录当前已经使用的数字，防止重复。
每一位数字从 0 到 9 依次尝试：
如果是第一位，则跳过 0。
如果当前数字已经被使用，则跳过。
如果当前位受限且选择的数字大于 n 的该位数字，则跳过。
如果选择的数字符合条件，则递归到下一位。
最后将该位的数字从 used 数组中移除，继续尝试其他数字。
递归处理与 n 同位数的情况：

使用 DFS 来处理每一位，递归处理数字的每一位。每当选择了与当前位相等的数字时，我们继续处理后面的数字。
如果处理完所有位，找到一个合法的特殊数字，则返回 1。
累计所有可能的结果，并将结果加到 total 中。
总计：

total 包含所有小于 n 位数的特殊数字，加上与 n 同位数的特殊数字。
例子：611
第一位选择 1-5，我们通过排列组合计算所有满足条件的三位数：5 * 9 * 8。
当第一位选择 6 时，继续递归处理第二位 1，此时我们可以选择 0，跳过重复的 1。
最终递归处理所有位，得到所有符合条件的特殊数字。
*/

// 辅助函数：检查数字数组是否是特殊数字
// func isValid(digits []int) bool {
// 	used := make([]bool, 10)
// 	for _, d := range digits {
// 		if used[d] {
// 			return false
// 		}
// 		used[d] = true
// 	}
// 	return true
// }

/*
优化分析：
1. 避免逐个遍历：通过计算小于 n 位数的所有特殊整数，减少了逐个遍历 1 到 n 的开销。
2. 深度优先搜索和记忆化：对于每一个可能的数字位，递归地构建特殊整数，减少了重复的计算。
3. 时间复杂度：由于我们避免了逐个检查每一个数字，时间复杂度可以近似认为是 O(d!)，其中 d 是 n 的位数。
*/

func countSpecialNumbersOptimizeGPTTwo(n int) int {
	onHave := 0
	for i1 := 1; i1 <= n; i1++ {
		used := [10]bool{} // 用于标记 0-9 的数字是否已经出现
		// isSpecial := true                // 标记是否为特殊数字
		for i2 := i1; i2 > 0; i2 /= 10 { // 直接操作数字的每一位
			digit := i2 % 10 // 取最后一位数字
			if used[digit] { // 如果该数字已经出现过，则不是特殊数字
				onHave++
				// isSpecial = false
				break
			}
			used[digit] = true // 记录该数字已经出现
		}
		// 这里是否是多余的
		// if !isSpecial {
		// 	continue
		// }
	}
	return n - onHave
}

func countSpecialNumbersGptFake(n int) int {

	digits := []int{}
	for i := n % 10; i > 0; i /= 10 {
		digits = append([]int{i}, digits...)
	}

	count := func(len int) int {
		result := 0
		used := 9
		for i := 0; i < len; i++ {
			if i == 0 {
				result += 9
			} else {
				result *= used
				used--
			}
		}
		return result
	}

	total := 0
	// 小于位数的特殊整数先加上
	for i := 0; i < len(digits); i++ {
		total += count(i)
	}

	// 处理相同位数

	// 把最顶位小一位的 加入总数
	head := digits[0] - 1
	for i := 9; i > len(digits); i-- {
		head *= i
	}
	total += head

	return total
}
