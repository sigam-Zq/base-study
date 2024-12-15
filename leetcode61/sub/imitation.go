package main

import (
	"container/heap"
	"sort"
)

func leetCodeImitation(nums []int, k int, multiplier int) []int {
	n, m := len(nums), int(1e9+7)

	var h = MyHeap{}
	heap.Init(&h)
	maxNum := 0
	for i, v := range nums {
		heap.Push(&h, []int{v, i})
		maxNum = max(maxNum, v)
	}

	for ; maxNum < h[0][0] && k > 0; k-- {
		x := heap.Pop(&h).([]int)
		x[0] *= multiplier
		heap.Push(&h, []int{x[0], x[1]})
	}
	sort.Slice(h, h.Less)

	for i := 0; i < n; i++ {
		t := k / n
		if i < k%n {
			t++
		}
		nums[h[i][1]] = (h[i][0] * quickMul(multiplier, t, m)) % m
	}

	return nums
}

/*
问：
int(((h[i][0]%m) * quickMul(int64(multiplier), int64(t), m)) % m)

int((h[i][0] * quickMul(int64(multiplier), int64(t), m)) % m)

在go 中上述两个表达式哪个是正确的 取余表达式


答：

在 Go 中，**第一个表达式**是正确的取余表达式：

```go
int(((h[i][0]%m) * quickMul(int64(multiplier), int64(t), m)) % m)
```

---

### **原因分析：**

#### **1. 取余运算的性质**
取模运算的核心性质是：
\[
(a \cdot b) \% m = [(a \% m) \cdot (b \% m)] \% m
\]

如果 `h[i][0]` 是一个可能较大的整数，直接参与乘法操作可能导致溢出（尽管 Go 中的 `int64` 类型可以表示相对大的数字）。通过提前对 `h[i][0]` 进行取模，可以避免乘法操作造成的溢出，并确保最终结果是正确的。

#### **2. 第二个表达式的问题**
```go
int((h[i][0] * quickMul(int64(multiplier), int64(t), m)) % m)
```

- 如果 `h[i][0]` 的值很大，那么 `h[i][0] * quickMul(...)` 的结果可能超出 `int64` 的表示范围（导致溢出问题）。
- 溢出后，即便对结果取模 `% m`，得到的值也可能不正确。

#### **3. 第一种表达式的优势**
```go
int(((h[i][0]%m) * quickMul(int64(multiplier), int64(t), m)) % m)
```

- **提前对 `h[i][0]` 取模**：通过 `(h[i][0] % m)`，将值限制在模 \( m \) 的范围内，避免后续乘法溢出。
- 整个表达式仍然符合取模运算的性质，结果正确。

---

### **举例对比**

假设：
- \( h[i][0] = 10^{18} \)
- \( m = 10^9 + 7 \)
- \( quickMul(...) = 10^6 \)

#### **第二种表达式：可能溢出**
```go
h[i][0] * quickMul(...) = 10^{18} \cdot 10^6 = 10^{24} // 超出 int64 范围
```

即使最后取模 `% m`，结果已经不可靠。

#### **第一种表达式：避免溢出**
```go
(h[i][0] % m) = 999999937   // 取模后，值限制在 [0, m-1]
(h[i][0] % m) * quickMul(...) = 999999937 \cdot 10^6  // 安全范围内
```

整个计算在安全范围内完成。

---

### **总结**
在 Go 中，选择第一个表达式是正确的方式：
```go
int(((h[i][0]%m) * quickMul(int64(multiplier), int64(t), m)) % m)
```
因为它避免了溢出，同时符合取模运算的性质。
*/
