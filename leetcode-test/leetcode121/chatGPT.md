

# 问

对某题目有以下两种解法





func maximumBeautyOptimize(items [][]int, queries []int) []int {

	// fmt.Printf("%v \n", items)
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] == items[j][0] {
			return items[i][1] < items[j][1]
		}
		return items[i][0] < items[j][0]
	})

	compressItems := make([][]int, 0)
	haveMap := make(map[int]struct{})
	n := len(items)
	for i := n - 1; i >= 0; i-- {
		if _, ok := haveMap[items[i][0]]; !ok {
			compressItems = append(compressItems, items[i])
			haveMap[items[i][0]] = struct{}{}
		}
	}
	n = len(compressItems)
	slices.Reverse(compressItems)
	for i := 1; i < n; i++ {
		if compressItems[i][1] < compressItems[i-1][1] {
			compressItems[i][1] = compressItems[i-1][1]
		}
	}

	q := func(x int) (y int) {
		l, r := 0, len(compressItems)
		for l < r {
			mid := (l + r) / 2
			if compressItems[mid][0] > x {
				r = mid
			} else {
				l++
			}
		}

		if l == 0 {
			return 0
		} else {
			return compressItems[l-1][1]
		}

	}

	ans := make([]int, len(queries))
	for i, v := range queries {
		ans[i] = q(v)
	}

	return ans
}


和


func maximumBeautyLeetCode(items [][]int, queries []int) []int {
	// 将物品按价格升序排序
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	n := len(items)
	// 按定义修改美丽值
	// fmt.Printf("befor---%v \n", items)
	for i := 1; i < n; i++ {
		if items[i][1] < items[i-1][1] {
			items[i][1] = items[i-1][1]
		}
	}
	// fmt.Printf("after---%v \n", items)
	// 二分查找处理查询
	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = query(items, q)
	}
	return res
}

func query(items [][]int, q int) int {
	l, r := 0, len(items)
	for l < r {
		mid := l + (r-l)/2
		if items[mid][0] > q {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l == 0 {
		// 此时所有物品价格均大于查询价格
		return 0
	} else {
		// 返回小于等于查询价格的物品的最大美丽值
		return items[l-1][1]
	}
}

当传入参数 items 存在 十万条  query 有9万条的同样的输入情况下

第一种方法需要执行   需要十秒多，第二种方法只需要不到1秒相差十多倍

试分析上述两种解法的 时间复杂度

并解释为什么第一种方法要比第二种方法慢这么多


# 答


我们分别分析两种方法的时间复杂度，并解释为什么第一种方法比第二种方法慢这么多。

---

## **第一种方法 `maximumBeautyOptimize`**
### **时间复杂度分析**
1. **排序：**  
   ```go
   sort.Slice(items, func(i, j int) bool { ... })
   ```
   该操作的时间复杂度为 **O(n log n)**（假设 `n` 为 `items` 的长度）。

2. **数据去重和压缩 `compressItems`：**  
   ```go
   for i := n - 1; i >= 0; i-- {
       if _, ok := haveMap[items[i][0]]; !ok {
           compressItems = append(compressItems, items[i])
           haveMap[items[i][0]] = struct{}{}
       }
   }
   ```
   - 这部分的复杂度是 **O(n)**，因为每个 `items[i][0]` 只被访问一次，`map` 查询和插入的时间复杂度为 O(1)。

3. **前缀最大值计算：**  
   ```go
   for i := 1; i < n; i++ {
       if compressItems[i][1] < compressItems[i-1][1] {
           compressItems[i][1] = compressItems[i-1][1]
       }
   }
   ```
   - 这个遍历 `compressItems` 的过程是 **O(n)**。

4. **查询 `q(x)`：**
   ```go
   l, r := 0, len(compressItems)
   for l < r {
       mid := (l + r) / 2
       if compressItems[mid][0] > x {
           r = mid
       } else {
           l++
       }
   }
   ```
   - 这里 `l++` 使得查找方式变成了 **O(n)**（**而不是 O(log n)**），因为 `l` 逐渐递增，而不是正常的二分查找方式 `l = mid + 1`。
   - 正确的二分查找应该像第二种方法 `maximumBeautyLeetCode` 那样写，而这里 `l++` 导致最多需要 **O(n)** 次迭代，极端情况下变成线性扫描。

5. **对每个 `query` 执行查询：**
   ```go
   for i, v := range queries {
       ans[i] = q(v)
   }
   ```
   - 假设 `m` 为 `queries` 的长度，单个查询的 `q(x)` 复杂度是 **O(n)**，总共执行 `m` 次查询，所以查询的总复杂度是 **O(mn)**。

### **总体复杂度**
综合来看：
- **排序 O(n log n)**
- **数据压缩 O(n)**
- **前缀最大值计算 O(n)**
- **查询 O(mn)**（这里因为 `q(x)` 变成了 O(n)）

所以整体复杂度是：
\[
O(n \log n) + O(n) + O(n) + O(mn) = O(mn)
\]
在 `n=10^5, m=9 \times 10^4` 时，`mn` 级别的复杂度会导致运行时间极大，导致 **慢 10 倍以上** 的差距。

---

## **第二种方法 `maximumBeautyLeetCode`**
### **时间复杂度分析**
1. **排序：**  
   ```go
   sort.Slice(items, func(i, j int) bool { return items[i][0] < items[j][0] })
   ```
   - 复杂度为 **O(n log n)**。

2. **前缀最大值计算：**  
   ```go
   for i := 1; i < n; i++ {
       if items[i][1] < items[i-1][1] {
           items[i][1] = items[i-1][1]
       }
   }
   ```
   - 复杂度为 **O(n)**。

3. **二分查找查询 `query(x)`：**
   ```go
   l, r := 0, len(items)
   for l < r {
       mid := l + (r-l)/2
       if items[mid][0] > q {
           r = mid
       } else {
           l = mid + 1
       }
   }
   ```
   - 这是一个标准的 **二分查找**，复杂度是 **O(log n)**。

4. **对每个 `query` 执行查询：**
   ```go
   for i, q := range queries {
       res[i] = query(items, q)
   }
   ```
   - **单次查询 O(log n)，总共 `m` 次查询，复杂度 O(m log n)**。

### **总体复杂度**
\[
O(n \log n) + O(n) + O(m \log n) = O(n \log n + m \log n)
\]
在 `n=10^5, m=9 \times 10^4` 时，`m log n` 远远小于 `mn`，运行时间大大减少。

---

## **为什么第一种方法慢这么多？**
1. **错误的二分查找逻辑**  
   - 第一种方法 `q(x)` 使用 `l++`，导致查询退化成 O(n) 复杂度，而不是 O(log n)。
   - 第二种方法 `query(x)` 正确地使用了二分查找，确保了 O(log n) 复杂度。

2. **查询次数影响放大**  
   - 第一种方法：O(mn) ≈ **O(9 × 10^9)**，远超 10 秒。
   - 第二种方法：O(m log n) ≈ **O(9 × 10^5)**，在 1 秒内运行完毕。

3. **第一种方法额外的数据压缩步骤**  
   - `compressItems` 的构造和维护额外增加了 O(n) 的开销，尽管这个部分相对较小，但仍然影响整体效率。

---

## **结论**
- **第一种方法 `maximumBeautyOptimize` 的查询部分由于 `l++` 变成了 O(n) 复杂度，导致整体 O(mn)，在大规模输入时极其慢。**
- **第二种方法 `maximumBeautyLeetCode` 正确使用二分查找，查询部分 O(m log n)，大幅提升性能，运行时间大大减少。**
- **在 `m=9 \times 10^4`，`n=10^5` 的情况下，第一种方法可能需要 10 秒以上，而第二种方法只需要不到 1 秒。**