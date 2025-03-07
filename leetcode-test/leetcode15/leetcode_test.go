package leetcode15

import (
	"log"
	"sort"
	"testing"
)

func TestMain(t *testing.T) {
	args := []struct {
		name string
		f    func(nums1 []int, nums2 []int, k int) int64
		num1 []int
		num2 []int
		k    int
		want int64
	}{
		{
			name: "one",
			f:    numberOfPairs,
			num1: []int{1, 3, 4},
			num2: []int{1, 3, 4},
			k:    1,
			want: int64(5),
		},
		{
			name: "one-claude2",
			f:    numberOfPairsClaude2,
			num1: []int{1, 3, 4},
			num2: []int{1, 3, 4},
			k:    1,
			want: int64(5),
		},
		{
			name: "one-claude",
			f:    numberOfPairsClaude,
			num1: []int{1, 3, 4},
			num2: []int{1, 3, 4},
			k:    1,
			want: int64(5),
		},
	}

	for _, v := range args {
		t.Run(v.name, func(t *testing.T) {
			if got := v.f(v.num1, v.num2, v.k); v.want != got {
				t.Errorf(" got %d  want %d \n", got, v.want)
			}
		})
	}

}

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {

	count := 0
	alreadyPass := make(map[int]int)
	for _, v1 := range nums1 {
		if v1%k == 0 {
			newV1 := v1 / k
			if alreadCount, ok := alreadyPass[v1]; ok {
				count += alreadCount
				continue
			}

			v1Count := 0
			for _, v2 := range nums2 {
				if newV1%v2 == 0 {
					count++
					v1Count++
				}
			}
			alreadyPass[v1] = v1Count
		}
	}

	return int64(count)
}

func numberOfPairsClaude(nums1 []int, nums2 []int, k int) int64 {
	count := 0
	m := len(nums2)

	// 预处理 nums2
	kNums2 := make([]int, m)
	for i, v := range nums2 {
		kNums2[i] = v * k
	}
	sort.Ints(kNums2)

	for _, v1 := range nums1 {
		// 使用二分查找找到可能的整除范围
		left := lowerBound(kNums2, v1)
		right := upperBound(kNums2, v1)

		// 在范围内检查每个元素是否真的满足整除条件
		log.Printf("--v1 %d  l %d ->  r %d  \n", v1, left, right)
		for i := left; i < right; i++ {
			if v1%(kNums2[i]) == 0 {
				count++
			}
		}
	}

	return int64(count)
}

// 二分查找，返回大于 target 的第一个元素的索引
func upperBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// 二分查找，返回大于等于 target 的第一个元素的索引
func lowerBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func numberOfPairsClaude2(nums1 []int, nums2 []int, k int) int64 {
	const MAX = 100001 // 根据题目约束，数组元素的最大值
	count := make([]int, MAX)
	var result int64 = 0

	// 第一次遍历：统计 nums2 中元素的出现次数
	for _, num := range nums2 {
		count[num]++
	}

	// 第二次遍历：处理 nums1 中的每个元素
	for _, num := range nums1 {
		if num%k != 0 {
			continue
		}
		target := num / k
		for j := 1; j*j <= target; j++ {
			if target%j == 0 {
				if j < MAX {
					result += int64(count[j])
				}
				if j*j != target && target/j < MAX {
					result += int64(count[target/j])
				}
			}
		}
	}

	return result
}

/**

问: 这里出现 两次的   j * j   是 什么作用

答:
非常好的问题！`j * j` 在这个算法中扮演了重要的角色。让我解释一下它的作用：

```go
for j := 1; j * j <= target; j++ {
    if target % j == 0 {
        if j < MAX {
            result += int64(count[j])
        }
        if j * j != target && target / j < MAX {
            result += int64(count[target/j])
        }
    }
}
```

这段代码的目的是找出 `target` 的所有因子，但它使用了一个优化技巧来减少迭代次数。这里的 `j * j` 有两个主要作用：

1. 限制循环范围：
   - 条件 `j * j <= target` 限制了循环只需要进行到 `target` 的平方根。
   - 这是因为如果 `j` 是 `target` 的因子，那么 `target / j` 也是 `target` 的因子。
   - 通过只迭代到平方根，我们可以在找到一个因子 `j` 的同时也考虑到它的配对因子 `target / j`。

2. 避免重复计算：
   - 在循环体内，我们检查 `j * j != target`。这是为了避免在 `target` 是完全平方数时重复计算中间的因子。
   - 例如，如果 `target = 16`，我们不希望在 `j = 4` 时重复计算 `4`。

让我们通过一个例子来说明这个过程：

假设 `target = 12`：

1. 当 `j = 1`：
   - 1 是因子，我们计入 `count[1]`
   - 12/1 = 12 也是因子，我们计入 `count[12]`

2. 当 `j = 2`：
   - 2 是因子，我们计入 `count[2]`
   - 12/2 = 6 也是因子，我们计入 `count[6]`

3. 当 `j = 3`：
   - 3 是因子，我们计入 `count[3]`
   - 12/3 = 4 也是因子，我们计入 `count[4]`

4. 当 `j = 4`：
   - 4 * 4 > 12，循环结束

通过这种方式，我们只需要迭代到 `sqrt(12)` ≈ 3.46，就可以找到所有的因子对。这大大减少了需要检查的数字数量，特别是对于大的 `target` 值。

这个优化将找因子的时间复杂度从 O(target) 降低到了 O(sqrt(target))，这在处理大数时特别有效。
**/

/*
您好！我很乐意帮您分析当前题解的时间复杂度，并提供一个优化后的解决方案。让我们逐步来看：

1. 当前题解的时间复杂度分析：

当前解法的时间复杂度是 O(n * m)，其中 n 是 nums1 的长度，m 是 nums2 的长度。这是因为对于 nums1 中的每个元素，我们都可能需要遍历整个 nums2 数组。虽然使用了 map 来缓存结果，但在最坏情况下（所有元素都不同），仍然需要进行 n * m 次比较。

2. 优化方案：

我们可以通过预处理 nums2 并使用二分查找来优化算法。主要思路如下：

a) 将 nums2 中的每个元素乘以 k，并排序。
b) 对于 nums1 中的每个元素，使用二分查找在处理后的 nums2 中找到可以整除它的元素数量。

这样可以将时间复杂度降低到 O(n * log(m) + m * log(m))。

下面是优化后的代码：

```go:solution.go
import (
    "sort"
)

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
    count := 0
    n, m := len(nums1), len(nums2)

    // 预处理 nums2
    kNums2 := make([]int, m)
    for i, v := range nums2 {
        kNums2[i] = v * k
    }
    sort.Ints(kNums2)

    for _, v1 := range nums1 {
        // 使用二分查找找到可以整除 v1 的元素数量
        count += upperBound(kNums2, v1) - lowerBound(kNums2, v1)
    }

    return int64(count)
}

// 二分查找，返回大于 target 的第一个元素的索引
func upperBound(arr []int, target int) int {
    left, right := 0, len(arr)
    for left < right {
        mid := left + (right - left) / 2
        if arr[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

// 二分查找，返回大于等于 target 的第一个元素的索引
func lowerBound(arr []int, target int) int {
    left, right := 0, len(arr)
    for left < right {
        mid := left + (right - left) / 2
        if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}
```

优化后算法的详细解释：

1. 首先，我们创建一个新的切片 `kNums2`，将 nums2 中的每个元素乘以 k。这一步的时间复杂度是 O(m)。

2. 然后，我们对 `kNums2` 进行排序。这一步的时间复杂度是 O(m * log(m))。

3. 接下来，我们遍历 nums1 中的每个元素 v1。对于每个 v1，我们需要找到在 kNums2 中能够整除 v1 的元素数量。

4. 为了高效地找到这个数量，我们使用两次二分查找：
   - `upperBound(kNums2, v1)` 找到第一个大于 v1 的元素的索引。
   - `lowerBound(kNums2, v1)` 找到第一个大于等于 v1 的元素的索引。
   这两个索引之间的元素就是能够整除 v1 的元素。

5. 每次二分查找的时间复杂度是 O(log(m))，我们对 nums1 中的每个元素都进行这样的操作，所以这部分的总时间复杂度是 O(n * log(m))。

总的时间复杂度是 O(m * log(m) + n * log(m))，这比原来的 O(n * m) 要好得多，特别是当 m 很大时。

空间复杂度方面，我们使用了一个额外的 `kNums2` 切片，所以空间复杂度是 O(m)。

这个优化利用了排序和二分查找来加速查找过程，大大提高了算法的效率，特别是在处理大规模数据时。
*/
