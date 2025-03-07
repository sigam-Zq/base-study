package main

import (
	"container/heap"
	"log"
	"math"
	"sort"
	"time"
)

func main() {
	var res []int
	timeCost(func() {
		res = getFinalStatRaw([]int{66307295, 441787703, 589039035, 322281864}, 900900704, 641725)
	})
	log.Println(res)
}

func getFinalStatRaw(nums []int, k int, multiplier int) []int {
	mod := int64(1e9 + 7)
	numsEx := make([]int64, len(nums))

	for i, v := range nums {
		numsEx[i] = int64(v)
	}

	for i := 0; i < k; i++ {
		min, minIdx := int64(math.MaxInt64), 0
		for i, v := range numsEx {
			if min > v {
				min = v
				minIdx = i
			}
		}
		nums[minIdx] *= multiplier
	}

	for i, v := range numsEx {
		nums[i] = int(v % mod)
	}

	return nums

}

// 初始化自己的堆
type MyHeap [][]int

func getFinalState(nums []int, k int, multiplier int) []int {
	mod := int(1e9 + 7)

	// var h *MyHeap
	h := &MyHeap{}

	heap.Init(h)
	for i, v := range nums {
		u := []int{v, i}
		heap.Push(h, u)
	}
	// log.Println(h)
	// log.Println(heap.Pop(h).([]int))

	for i := 0; i < k; i++ {
		minIdx := heap.Pop(h).([]int)[1]
		nums[minIdx] = quickMul(nums[minIdx], multiplier, mod)
		u := []int{nums[minIdx], minIdx}
		heap.Push(h, u)
	}

	return nums
}

func quickMul(x, y, m int) int {
	res := int(1)
	for y > 0 {
		if (y & 1) == 1 {
			res = (res * x) % m
		}
		y >>= 1
		x = (x * x) % m
	}
	return res
}

func timeCost(f func()) {
	s := time.Now()
	f()

	tc := time.Since(s)

	log.Printf("time COST %v\n", tc)
}

func (m MyHeap) Len() int {
	return len(m)
}

func (m MyHeap) Less(i, j int) bool {
	if m[i][0] == m[j][0] {
		return m[i][1] < m[j][1]
	}
	return m[i][0] < m[j][0]
}

func (m MyHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MyHeap) Push(a any) {
	(*m) = append(*m, a.([]int))
}
func (m *MyHeap) Pop() any {
	// old := &m
	x := (*m)[m.Len()-1]
	(*m) = (*m)[:m.Len()-1]
	return x
}

/*
问:

这里有两个问题
1.
这里两个题解的 时间复杂度分别是 什么

2.
这里当k和 multiplier极大的时候
为什么 题解2 运行时间会长

参数为 []int{66307295, 441787703, 589039035, 322281864}, 900900704, 641725
时 题解1  5s  题解2 1m34s

答:

Let's analyze the two solutions in terms of time complexity and performance when `k` and `multiplier` are large.

### 1. Time Complexity

#### Solution 1: Brute Force Approach

```go
func getFinalState(nums []int, k int, multiplier int) []int {
    for i := 0; i < k; i++ {
        min, minIdx := math.MaxInt, 0
        for i, v := range nums {
            if min > v {
                min = v
                minIdx = i
            }
        }
        nums[minIdx] *= multiplier
    }
    return nums
}
```

- **Time Complexity**: The outer loop runs `k` times, and for each iteration, it scans the entire `nums` array to find the minimum value, which takes `O(n)` time. Therefore, the overall time complexity is `O(k * n)`.

#### Solution 2: Heap Approach

```go
type MyHeap [][]int

func getFinalState(nums []int, k int, multiplier int) []int {
    mod := int(1e9 + 7)
    h := &MyHeap{}
    heap.Init(h)
    for i, v := range nums {
        u := []int{v, i}
        heap.Push(h, u)
    }
    for i := 0; i < k; i++ {
        minIdx := heap.Pop(h).([]int)[1]
        nums[minIdx] = (nums[minIdx] * multiplier) % mod
        u := []int{nums[minIdx], minIdx}
        heap.Push(h, u)
    }
    return nums
}
```

- **Time Complexity**: The heap operations (push and pop) take `O(log n)` time. The loop runs `k` times, so the overall time complexity is `O(k * log n)`.

### 2. Performance with Large `k` and `multiplier`

- **Solution 1**: This solution will perform poorly with large `k` because it has to scan the entire array `k` times, leading to a time complexity of `O(k * n)`. However, it does not involve any additional operations that depend on the size of `multiplier`.

- **Solution 2**: While this solution is more efficient in terms of time complexity (`O(k * log n)`), it involves additional operations such as modulus operation with `multiplier` and heap operations. When `multiplier` is very large, the modulus operation can become computationally expensive, especially if `k` is also large, leading to longer execution times.

In summary, Solution 2 is generally more efficient due to its use of a heap, but the additional computational overhead from modulus operations and heap management can make it slower in practice when `k` and `multiplier` are extremely large.


*/

func quickMulLeetCode(x, y, m int64) int64 {
	res := int64(1)
	for y > 0 {
		if (y & 1) == 1 {
			res = (res * x) % m
		}
		y >>= 1
		x = (x * x) % m
	}
	return res
}

func getFinalStateLeetCode(nums []int, k int, multiplier int) []int {
	if multiplier == 1 {
		return nums
	}
	n, m := len(nums), int64(1e9+7)
	mx := 0
	var v minHeap
	for i, num := range nums {
		mx = max(mx, num)
		v = append(v, pair{int64(num), i})
	}
	heap.Init(&v)
	for ; v[0].first < int64(mx) && k > 0; k-- {
		x := heap.Pop(&v).(pair)
		x.first *= int64(multiplier)
		heap.Push(&v, x)
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i].first < v[j].first || v[i].first == v[j].first && v[i].second < v[j].second
	})
	for i := 0; i < n; i++ {
		t := k / n
		if i < k%n {
			t++
		}
		nums[v[i].second] = int((v[i].first % m) * quickMulLeetCode(int64(multiplier), int64(t), m) % m)
	}
	return nums
}

type pair struct {
	first  int64
	second int
}

type minHeap []pair

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i].first < h[j].first || h[i].first == h[j].first && h[i].second < h[j].second
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *minHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return res
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/final-array-state-after-k-multiplication-operations-ii/solutions/3014793/k-ci-cheng-yun-suan-hou-de-zui-zhong-shu-74yw/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/*
问
给你一个整数数组 nums ，一个整数 k  和一个整数 multiplier 。

你需要对 nums 执行 k 次操作，每次操作中：

找到 nums 中的 最小 值 x ，如果存在多个最小值，选择最 前面 的一个。
将 x 替换为 x * multiplier 。
k 次操作以后，你需要将 nums 中每一个数值对 109 + 7 取余。

请你返回执行完 k 次乘运算以及取余运算之后，最终的 nums 数组。

题目如上

题解如下

func quickMul(x, y, m int64) int64 {
    res := int64(1)
    for y > 0 {
        if (y & 1) == 1 {
            res = (res * x) % m
        }
        y >>= 1
        x = (x * x) % m
    }
    return res
}

func getFinalState(nums []int, k int, multiplier int) []int {
    if multiplier == 1 {
        return nums
    }
    n, m := len(nums), int64(1e9+7)
    mx := 0
    var v minHeap
    for i, num := range nums {
        mx = max(mx, num)
        v = append(v, pair{int64(num), i})
    }
    heap.Init(&v)
    for ; v[0].first < int64(mx) && k > 0; k-- {
        x := heap.Pop(&v).(pair)
        x.first *= int64(multiplier)
        heap.Push(&v, x)
    }
    sort.Slice(v, func(i, j int) bool {
        return v[i].first < v[j].first || v[i].first == v[j].first && v[i].second < v[j].second
    })
    for i := 0; i < n; i++ {
        t := k / n
        if i < k % n {
            t++
        }
        nums[v[i].second] = int((v[i].first % m) * quickMul(int64(multiplier), int64(t), m) % m)
    }
    return nums
}

type pair struct {
    first  int64
    second int
}

type minHeap []pair

func (h minHeap) Len() int{
    return len(h)
}
func (h minHeap) Less(i, j int) bool {
    return h[i].first < h[j].first || h[i].first == h[j].first && h[i].second < h[j].second
}
func (h minHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
    *h = append(*h, x.(pair))
}

func (h *minHeap) Pop() interface{} {
    n := len(*h)
    res := (*h)[n - 1]
    *h = (*h)[0 : n - 1]
    return res
}
有以下问题
1
为什么循环判断中有
v[0].first < int64(mx)
为什么不满足后要跳出循环 不进行Pop Push 操作了呢
2
heap 最小堆的性质 不应该是每次 heap.Push 后自动排序了么
为什么在 循环 对变量名 为v 的heap Push 后执行了
sort.Slice  再次进行了排序

3
第二次循环
for i := 0; i < n; i++ {
        t := k / n
        if i < k % n {
            t++
        }
        nums[v[i].second] = int((v[i].first % m) * quickMul(int64(multiplier), int64(t), m) % m)
    }
这里是其中 维护的t是什么，在这里起到什么作用 ，这里再题目中和乘数 相乘的操作是在上面没做完么，为什么这里变为 求幂指数取余了

答：

我们逐一解答你的问题：

---

### **问题 1**
**为什么循环判断中有 `v[0].first < int64(mx)`，不满足后要跳出循环不进行 `Pop` `Push` 操作了呢？**

**解答：**
- 这个条件是为了优化操作次数。当堆顶的最小值 `v[0].first` 已经不小于数组中的最大值 `mx` 时，继续将这个数乘以 `multiplier` 已经没有意义，因为其他所有数都比它大，无法通过这些操作改变最终结果的顺序。
- 如果没有这个条件，即使继续进行 `k` 次操作，堆顶的数会越来越大，甚至可能超出数组中其他元素的值范围。这会导致无谓的计算，因此提前退出循环。

---

### **问题 2**
**堆最小堆的性质不应该是每次 `heap.Push` 后自动排序了么？为什么在循环对 `v`（一个堆） `Push` 后，还执行了 `sort.Slice` 再次排序？**

**解答：**
- `heap.Push` 确实会维护堆的**最小堆性质**，也就是说，堆顶元素始终是最小值。
- 但是堆本质上是一个**部分排序的数据结构**，并不是完全排序的。虽然堆顶是最小值，但其他元素的顺序未必满足全局排序规则。
- 在后续计算中，代码需要 `v` 按照值从小到大进行完整排序（比如按值升序排列，相同值按索引升序）。因此需要额外调用 `sort.Slice` 进行完整排序。

总结：堆操作维护的是堆的局部最小性质，`sort.Slice` 是为了全局排序以满足后续逻辑。

---

### **问题 3**
**第二次循环中 `t` 的作用是什么？为什么这里用幂指数和取模的形式？**

**解答：**
1. **`t` 的定义与作用：**
   - `t` 表示每个数在剩余的 `k` 次操作中将被乘以 `multiplier` 的次数。
   - `k / n` 表示每个数被完整循环中操作的次数；
   - `k % n` 是剩余的操作数，优先作用于前面的数，因此 `i < k % n` 用来判断当前数是否需要额外操作一次。

2. **为什么这里用幂指数和取模的形式？**
   - 由于乘法的幂次性质 \( x \times m^t = x \times (m \times m \times \ldots \times m) \)，连续的 \( t \) 次乘法可以表示为一次幂运算 \( x \times m^t \)。
   - 为了避免直接计算大数（可能溢出），这里使用快速幂算法（`quickMul`）计算 \( m^t \mod (10^9+7) \)。
   - 这样既减少了乘法操作的次数，也保持结果在有效范围内。

3. **与之前堆操作的关系：**
   - 在前面的循环中，仅对堆顶的最小值进行了 `multiplier` 操作，其他元素未被处理。
   - 在这里，通过幂指数的方式补充完成了剩余的 `k` 次操作，使得每个数都被乘以 `multiplier` 恰当的次数。

---

### **总结：**
1. **循环中退出的条件 `v[0].first < int64(mx)` 是为了优化操作，避免无效计算。**
2. **堆仅维护局部有序性，因此需要额外的排序操作以确保全局顺序。**
3. **通过计算幂指数 (`t`) 和快速幂取模，完成了剩余的 `k` 次乘法操作，且保证计算高效并避免溢出。**

问：
问题1的回答中

关于
当堆顶的最小值 v[0].first 已经不小于数组中的最大值 mx 时，继续将这个数乘以 multiplier 已经没有意义，因为其他所有数都比它大，无法通过这些操作改变最终结果的顺序。

这里  题目中说的是  找到 nums 中的 最小 值 x ，如果存在多个最小值，选择最 前面 的一个。

题目仅关心的是其中nums 中最小的值  ，其中 原来数组中的最大的值变为最小的值之后这里继续计算为什么没有意义了，
为什么 无法通过这些操作改变最终结果的顺序。

答:
你的问题非常好，让我们深入分析一下逻辑。

---

### 1. **题目重点：最小值操作的作用**
题目要求每次找到 `nums` 中的最小值 `x` 并将其替换为 `x * multiplier`，如果有多个最小值，取第一个。执行 `k` 次操作后，最终只需要关心数组中的数值以及这些数值的顺序。

- 操作的核心是将最小值逐渐放大。
- 堆用于高效地获取当前数组中的最小值，并对其进行操作。

---

### 2. **为什么堆顶不小于 `mx` 时，继续计算没有意义？**

#### **初始状态和操作分析：**
- 假设数组初始的最小值为 `min`，最大值为 `mx`。
- 每次将当前最小值 `min` 替换为 `min * multiplier` 后，它会变大，可能超过其他元素。

#### **关键问题：当堆顶大于等于 `mx` 后，是否还有意义？**

1. **堆顶变成全局最大值后：**
   - 假设当前堆顶的最小值 `v[0].first` 已经变成大于或等于初始的最大值 `mx`。
   - 由于 `v[0].first` 现在是整个数组中的最大值，继续对它执行乘法操作只会让它变得更大。
   - 而这与题目「操作最小值」的初衷背道而驰，因为本次操作的结果（最小值的变化）不会再影响最终排序和数值。

2. **堆顶的值增长速度远快于其他元素：**
   - 当堆顶大于等于 `mx` 时，后续操作会使它的值指数级增长。而其他元素在后续循环中，仍然是以较慢的速度增长。
   - 因此，继续对它操作只会拉大它与其他元素的差距，而不改变数组的最小值顺序。

#### **举个例子：**
初始数组：`nums = [1, 2, 3]`，`multiplier = 10`，`k = 3`。

1. 初始堆：`[1, 2, 3]`（最小值 `1`）。
2. 第一次操作：将 `1` 替换为 `1 * 10 = 10`，堆变为 `[2, 3, 10]`。
3. 第二次操作：将 `2` 替换为 `2 * 10 = 20`，堆变为 `[3, 10, 20]`。
4. 第三次操作：将 `3` 替换为 `3 * 10 = 30`，堆变为 `[10, 20, 30]`。

到此为止，堆的最小值 `v[0].first` 已经等于初始的最大值 `mx = 3`。如果还有剩余的操作：
- 继续操作堆顶（`10`）只会让它变得更大，例如从 `10` 变为 `100`。
- 但数组的顺序 `[10, 20, 30]` 已经不会因为堆顶的变化而受到影响。

---

### 3. **为什么说无法改变最终结果的顺序？**

题目要求最终结果是所有数值都模 \( 10^9 + 7 \) 后的数组顺序。而当堆顶 `v[0].first` 超过 `mx` 后：
- **无论怎么乘法操作，堆顶永远是最大的，最终的模运算结果也只会让它更大。**
- 剩余的操作只会让这个数增长得更快，而不会改变其他数的值或位置。

---

### **总结：**
当 `v[0].first >= mx` 时，继续操作堆顶不会改变数组的排序和数值关系，因此停止无意义的操作是一个优化策略。

问：
延续上面讨论中举出的例子

继续操作 [10, 20, 30] 的 堆顶元素 10 变为一百

堆中元素 不应该变为  [20,30,100] 么 ，后续然后继续操作 20 进行乘

这里是因为堆的性质原因导致  10  变为100 后仍在 堆顶，还是考虑到溢出的场景呢

答：
在你的例子中，[10, 20, 30] 是一个**最小堆**。当堆顶元素 `10` 被操作变成 `100` 后，堆需要重新调整以满足**最小堆的性质**，最终会变成 `[20, 30, 100]`。堆顶元素更新后仍然是最小值，堆的结构也会自动调整。

所以，**你的理解是正确的**，操作完 `10` 后堆中应该是 `[20, 30, 100]`，然后继续操作 `20`。

---

### **问题：为什么代码中提到“继续操作堆顶没有意义”？**
这与堆的性质无关，完全是代码中的优化逻辑——**不再操作堆顶的条件并不是因为堆顶的调整问题，而是因为溢出或优化操作的策略**。以下是详细原因：

---

### **1. 不再操作堆顶与堆的性质无关**
堆本身会自动调整，比如例子中：
1. 初始堆是 `[10, 20, 30]`。
2. 操作 `10` 后，堆更新为 `[20, 30, 100]`，最小堆性质依然满足。
3. 再次操作堆顶 `20`，更新为 `[30, 100, 200]`。

堆的调整是正常的，问题在于优化逻辑何时停止操作。

---

### **2. 溢出问题的影响**
由于题目要求的结果是对 \( 10^9 + 7 \) 取模的，数值可能会非常大：
- 每次操作 `min = min * multiplier` 会指数级增长。
- 即便取模后保存了有效值，但多余的操作可能不会影响最终结果顺序，而是带来额外的计算开销。

---

### **3. 停止操作是逻辑优化**
当 `v[0].first >= mx`（堆顶元素已经超过初始最大值）时，代码判断后续操作不再会影响结果，停止是基于以下理由：

1. **后续操作的数值顺序不会改变：**
   - 操作最小值后，其值会快速超过其他元素的最大值。
   - 剩余的操作对排序结果无影响，因此可以直接退出循环。

2. **减少不必要的计算：**
   - 假如后续操作堆顶的次数 \( t \) 很大，直接用幂运算 \( x \cdot m^t \mod (10^9 + 7) \) 就可以高效计算最终结果，避免逐次操作。

---

### **4. 举例进一步说明**
假设 `nums = [1, 2, 3]`，`multiplier = 10`，`k = 5`：

1. 初始堆为 `[1, 2, 3]`。
2. 第一次操作堆顶 `1`，堆变为 `[2, 3, 10]`。
3. 第二次操作堆顶 `2`，堆变为 `[3, 10, 20]`。
4. 第三次操作堆顶 `3`，堆变为 `[10, 20, 30]`。
5. **此时堆顶 `10 >= mx (3)`，后续逻辑判断：继续操作堆顶无法改变数值关系，停止循环。**

实际上：
- 停止循环后，剩余的两次操作会用幂运算高效完成。
- 对于堆中的每个数，用 `quickMul` 计算乘以 `multiplier` 的次数 \( t \) 后的最终结果，避免逐次操作浪费时间。

---

### **总结**
1. 堆的性质会自动维护最小堆，堆顶元素操作后会调整，问题与堆性质无关。
2. **停止操作是基于逻辑优化**，当堆顶的值已经超过初始的最大值后，直接用幂运算完成剩余计算更高效。
3. 停止循环不是因为堆顶未调整，而是后续的逐次操作已经没有实际意义。



问：
回到上述问题和题解的部分


    for ; v[0].first < int64(mx) && k > 0; k-- {
        x := heap.Pop(&v).(pair)
        x.first *= int64(multiplier)
        heap.Push(&v, x)
    }
这一次循环是否有什么意义，是否可以去掉，
去掉是否有什么影响，请举例说明

答：
让我们仔细分析这段代码的作用以及是否可以去掉，并通过举例说明去掉后的影响。

---

### **代码分析**

循环代码如下：

```go
for ; v[0].first < int64(mx) && k > 0; k-- {
    x := heap.Pop(&v).(pair)       // 取出当前堆顶最小值
    x.first *= int64(multiplier)   // 将最小值乘以 multiplier
    heap.Push(&v, x)               // 将更新后的值重新放回堆中
}
```

#### **作用：**
1. **优先处理小于 `mx` 的元素：**
   - 这段代码的主要目的是通过堆的特性（快速获取最小值），优先处理当前值小于 `mx` 的元素，逐步缩小最小值与其他值之间的差距。
   - 这样做可以尽量避免某些较小的值被遗漏，从而导致最终结果不正确。

2. **减少无意义计算：**
   - 如果当前最小值已经大于等于 `mx`，它的值在经过乘法操作后会快速变大，但其他较小值仍未被操作。
   - 假如直接停止操作这些较小值，可能会导致结果中某些值没有被充分操作，无法满足题目要求。

#### **去掉这段代码的影响：**
- 如果直接跳过这段逻辑，即不再区分小于 `mx` 和大于等于 `mx` 的元素，所有剩余的操作将被均匀分配给数组中的每个元素（通过后续的幂次计算）。
- 这可能导致：
  1. **最终数组顺序错误：** 如果某些较小的值未被优先操作，其值可能会比预期的大。
  2. **数值不一致：** 去掉这段逻辑后，最小值的更新过程无法有效完成，可能会与题目要求的逻辑偏差。

---

### **举例说明**

#### **1. 保留循环的情况**
初始 `nums = [1, 3, 5]`，`multiplier = 10`，`k = 3`，堆中最小值为 `1`，最大值 `mx = 5`。

- 第1次操作：堆顶 `1` 变为 `10`，堆变为 `[3, 5, 10]`。
- 第2次操作：堆顶 `3` 变为 `30`，堆变为 `[5, 10, 30]`。
- 第3次操作：堆顶 `5` 变为 `50`，堆变为 `[10, 30, 50]`。

最终结果经过取模处理后排序为 `[10, 30, 50]`，顺序正确。

---

#### **2. 去掉循环的情况**
假如直接跳过这段循环（不处理堆中小于 `mx` 的元素），进入后续计算：
- 所有剩余操作直接均匀分配给每个元素。
- 对于每个元素 `nums[i]`，计算乘以 `multiplier` 的幂次，结果可能如下：
  - `1 * 10^1 = 10`
  - `3 * 10^1 = 30`
  - `5 * 10^1 = 50`

最终结果也可能是 `[10, 30, 50]`，但这是**偶然正确**，因为所有操作正好均匀分配。如果 `k` 和 `nums` 的大小不成比例（比如 `k > n` 或 `k < n`），最终顺序会出错。

---

#### **影响场景**
假如 `nums = [1, 1, 10]`，`multiplier = 2`，`k = 4`，最大值 `mx = 10`：

- **保留循环：**
  1. 第1次操作：`1` -> `2`，堆变为 `[1, 10, 2]`。
  2. 第2次操作：`1` -> `2`，堆变为 `[2, 10, 2]`。
  3. 第3次操作：`2` -> `4`，堆变为 `[2, 10, 4]`。
  4. 第4次操作：`2` -> `4`，堆变为 `[4, 10, 4]`。
  最终结果 `[4, 4, 10]`。

- **去掉循环：**
  - 所有元素均匀分配 \( k = 4 \) 次操作，结果可能是：
    - `1 * 2^2 = 4`
    - `1 * 2^2 = 4`
    - `10 * 2^2 = 40`
  最终结果 `[4, 4, 40]`。

去掉循环导致最小值 `1` 没有被优先操作，`10` 的增长过快，最终结果与题目要求不符。

---

### **结论**
这段代码的意义在于优先处理较小的值，确保操作顺序与题目要求一致，避免后续幂次计算导致结果错误。如果去掉：
1. **可能影响最终的数值顺序或大小。**
2. **当 `k` 与数组长度不成比例时，问题更明显。**

因此，这段循环不能随意去掉，它是实现题目逻辑的重要部分。

*/
