package leetcode29

import (
	"log"
	"math"
	"strconv"
	"testing"
)

/*
638. 大礼包
在 LeetCode 商店中， 有 n 件在售的物品。每件物品都有对应的价格。然而，也有一些大礼包，每个大礼包以优惠的价格捆绑销售一组物品。

给你一个整数数组 price 表示物品价格，其中 price[i] 是第 i 件物品的价格。另有一个整数数组 needs 表示购物清单，其中 needs[i] 是需要购买第 i 件物品的数量。

还有一个数组 special 表示大礼包，special[i] 的长度为 n + 1 ，其中 special[i][j] 表示第 i 个大礼包中内含第 j 件物品的数量，且 special[i][n] （也就是数组中的最后一个整数）为第 i 个大礼包的价格。

返回 确切 满足购物清单所需花费的最低价格，你可以充分利用大礼包的优惠活动。你不能购买超出购物清单指定数量的物品，即使那样会降低整体价格。任意大礼包可无限次购买。



示例 1：

输入：price = [2,5], special = [[3,0,5],[1,2,10]], needs = [3,2]
输出：14
解释：有 A 和 B 两种物品，价格分别为 ¥2 和 ¥5 。
大礼包 1 ，你可以以 ¥5 的价格购买 3A 和 0B 。
大礼包 2 ，你可以以 ¥10 的价格购买 1A 和 2B 。
需要购买 3 个 A 和 2 个 B ， 所以付 ¥10 购买 1A 和 2B（大礼包 2），以及 ¥4 购买 2A 。
示例 2：

输入：price = [2,3,4], special = [[1,1,0,4],[2,2,1,9]], needs = [1,2,1]
输出：11
解释：A ，B ，C 的价格分别为 ¥2 ，¥3 ，¥4 。
可以用 ¥4 购买 1A 和 1B ，也可以用 ¥9 购买 2A ，2B 和 1C 。
需要买 1A ，2B 和 1C ，所以付 ¥4 买 1A 和 1B（大礼包 1），以及 ¥3 购买 1B ， ¥4 购买 1C 。
不可以购买超出待购清单的物品，尽管购买大礼包 2 更加便宜。


提示：

n == price.length == needs.length
1 <= n <= 6
0 <= price[i], needs[i] <= 10
1 <= special.length <= 100
special[i].length == n + 1
0 <= special[i][j] <= 50
生成的输入对于 0 <= j <= n - 1 至少有一个 special[i][j] 非零。
*/

func TestXxx(t *testing.T) {

	for idx, v := range []struct {
		f       func([]int, [][]int, []int) int
		price   []int
		special [][]int
		needs   []int
		want    int
	}{
		{
			f:     shoppingOffers,
			price: []int{2, 5},
			special: [][]int{
				{3, 0, 5},
				{1, 2, 10},
			},
			needs: []int{3, 2},
			want:  14,
		},
		{
			f:     shoppingOffers,
			price: []int{0, 0, 0},
			special: [][]int{
				{1, 1, 0, 4},
				{2, 2, 1, 9},
			},
			needs: []int{1, 1, 1},
			want:  0,
		},
		{
			f:     shoppingOffers,
			price: []int{9, 9},
			special: [][]int{
				{1, 1, 1},
			},
			needs: []int{2, 2},
			want:  2,
		},
		{
			f:     shoppingOffers,
			price: []int{2, 3, 4},
			special: [][]int{
				{1, 1, 0, 4},
				{2, 2, 1, 9},
			},
			needs: []int{1, 2, 1},
			want:  11,
		},
		{
			f:       shoppingOffers,
			price:   []int{9, 6, 1, 5, 3, 4},
			special: [][]int{{1, 2, 2, 1, 0, 4, 14}, {6, 3, 4, 0, 0, 1, 16}, {4, 5, 6, 6, 2, 4, 26}, {1, 1, 4, 3, 4, 3, 15}, {4, 2, 5, 4, 4, 5, 15}, {4, 0, 0, 2, 3, 5, 13}, {2, 4, 6, 4, 3, 5, 7}, {3, 3, 4, 2, 2, 6, 21}, {0, 3, 0, 2, 3, 3, 15}, {0, 2, 4, 2, 2, 5, 24}, {4, 1, 5, 4, 5, 4, 25}, {6, 0, 5, 0, 1, 1, 14}, {4, 0, 5, 2, 1, 5, 8}, {4, 1, 4, 4, 3, 1, 10}, {4, 4, 2, 1, 5, 0, 14}, {2, 4, 4, 1, 3, 1, 16}, {4, 2, 3, 1, 2, 1, 26}, {2, 4, 1, 6, 5, 3, 2}, {0, 2, 0, 4, 0, 0, 19}, {3, 1, 6, 3, 3, 1, 23}, {6, 2, 3, 2, 4, 4, 16}, {5, 3, 5, 5, 0, 4, 5}, {5, 0, 4, 3, 0, 2, 20}, {5, 3, 1, 2, 2, 5, 8}, {3, 0, 6, 1, 0, 2, 10}, {5, 6, 6, 1, 0, 4, 12}, {0, 6, 6, 4, 6, 4, 21}, {0, 4, 6, 5, 0, 0, 22}, {0, 4, 2, 4, 4, 6, 16}, {4, 2, 1, 0, 6, 5, 14}, {0, 1, 3, 5, 0, 3, 8}, {5, 5, 3, 3, 2, 0, 4}, {1, 0, 3, 6, 2, 3, 18}, {4, 2, 6, 2, 2, 5, 2}, {0, 2, 5, 5, 3, 6, 12}, {1, 0, 6, 6, 5, 0, 10}, {6, 0, 0, 5, 5, 1, 24}, {1, 4, 6, 5, 6, 3, 19}, {2, 2, 4, 2, 4, 2, 20}, {5, 6, 1, 4, 0, 5, 3}, {3, 3, 2, 2, 1, 0, 14}, {0, 1, 3, 6, 5, 0, 9}, {5, 3, 6, 5, 3, 3, 11}, {5, 3, 3, 1, 0, 2, 26}, {0, 1, 1, 4, 2, 1, 16}, {4, 2, 3, 2, 1, 4, 6}, {0, 2, 1, 3, 3, 5, 15}, {5, 6, 4, 1, 2, 5, 18}, {1, 0, 0, 1, 6, 1, 16}, {2, 0, 6, 6, 2, 2, 17}, {4, 4, 0, 2, 4, 6, 12}, {0, 5, 2, 5, 4, 6, 6}, {5, 2, 1, 6, 2, 1, 24}, {2, 0, 2, 2, 0, 1, 14}, {1, 1, 0, 5, 3, 5, 16}, {0, 2, 3, 5, 5, 5, 6}, {3, 2, 0, 6, 4, 6, 8}, {4, 0, 1, 4, 5, 1, 6}, {5, 0, 5, 6, 6, 3, 7}, {2, 6, 0, 0, 2, 1, 25}, {0, 4, 6, 1, 4, 4, 6}, {6, 3, 1, 4, 1, 1, 24}, {6, 2, 1, 2, 1, 4, 4}, {0, 1, 2, 3, 0, 1, 3}, {0, 2, 5, 6, 5, 2, 13}, {2, 6, 4, 2, 2, 3, 17}, {3, 4, 5, 0, 5, 4, 20}, {6, 2, 3, 4, 1, 3, 4}, {6, 4, 0, 0, 0, 5, 16}, {3, 1, 2, 5, 0, 6, 11}, {1, 3, 2, 2, 5, 6, 14}, {1, 3, 4, 5, 3, 5, 18}, {2, 1, 1, 2, 6, 1, 1}, {4, 0, 4, 0, 6, 6, 8}, {4, 6, 0, 5, 0, 2, 1}, {3, 1, 0, 5, 3, 2, 26}, {4, 0, 4, 0, 6, 6, 6}, {5, 0, 0, 0, 0, 4, 26}, {4, 3, 2, 2, 0, 2, 14}, {5, 2, 4, 0, 2, 2, 26}, {3, 4, 6, 0, 2, 4, 25}, {2, 1, 5, 5, 1, 3, 26}, {0, 5, 2, 4, 0, 2, 24}, {5, 2, 5, 4, 5, 0, 1}, {5, 3, 0, 1, 5, 4, 15}, {6, 1, 5, 1, 2, 1, 21}, {2, 5, 1, 2, 1, 4, 15}, {1, 4, 4, 0, 0, 0, 1}, {5, 0, 6, 1, 1, 4, 22}, {0, 1, 1, 6, 1, 4, 1}, {1, 6, 0, 3, 2, 2, 17}, {3, 4, 3, 3, 1, 5, 17}, {1, 5, 5, 4, 5, 2, 27}, {0, 6, 5, 5, 0, 0, 26}, {1, 4, 0, 3, 1, 0, 13}, {1, 0, 3, 5, 2, 4, 5}, {2, 2, 2, 3, 0, 0, 11}, {3, 2, 2, 1, 1, 1, 6}, {6, 6, 1, 1, 1, 6, 26}, {1, 5, 1, 2, 5, 2, 12}},
			needs:   []int{6, 6, 6, 1, 6, 6},
			want:    34,
		},
		{
			f:     shoppingOffers,
			price: []int{6, 3},
			special: [][]int{
				{3, 2, 9},
				{1, 2, 1},
				{0, 1, 5},
				{2, 5, 9},
			},
			needs: []int{6, 6},
			want:  17,
		},
		{
			f:     shoppingOffersLeetCode,
			price: []int{6, 3},
			special: [][]int{
				{3, 2, 9},
				{1, 2, 1},
				{0, 1, 5},
				{2, 5, 9},
			},
			needs: []int{6, 6},
			want:  17,
		},
	} {
		t.Run(strconv.Itoa(idx)+"-test", func(t *testing.T) {
			if got := v.f(v.price, v.special, v.needs); got != v.want {
				t.Errorf(" got %v  want %v \n", got, v.want)
			}
		})
	}
}

func shoppingOffers(price []int, special [][]int, needs []int) int {

	// 先过滤去除大于总需求的special
	special = filterSpecials(special, needs)

	// log.Printf("before special %v ,cap %d ,len %d \n \n", special, cap(special), len(special))
	sortSpecials(price, &special)

	log.Printf("after special %v ,cap %d ,len %d \n \n", special, cap(special), len(special))
	// 如果有买 特惠礼包- 就使用特惠礼包

	// 初始值使用全部买单个
	cost := 0
	for i, v := range needs {
		cost += price[i] * v
	}
	// 依次去购买 特惠礼包 (目前未计算最优)--- 当前最外层便利寻找所有组合

	//TODO 当前 没有考虑组合之间相互 搭配的问题- (最省钱的套餐摆前面往死里薅)
	for startPoint := 0; startPoint < len(special); startPoint++ {
		needsTemp := make([]int, len(needs))
		copy(needsTemp, needs)
		costTemp := 0
		log.Printf("startPoint %d  \n", startPoint)
		for _, specialOne := range special[startPoint:] {
			part := math.MaxInt
			// 是否购买
			isBuy := true
			for i, v := range needsTemp {
				if specialOne[i] <= v {
					// 算出最小整除数 （最小买的份数）
					var tempPart int
					if specialOne[i] != 0 {
						tempPart = v / specialOne[i]
					} else {
						tempPart = math.MaxInt
					}
					if tempPart < part {
						part = tempPart
					}
				} else {
					isBuy = false
				}
			}
			log.Printf("isBuy %v, part %d \n ", isBuy, part)
			if isBuy {
				costTemp += part * specialOne[len(specialOne)-1]
				for i, v := range needsTemp {
					// 购买套餐后 去除所需的数目
					needsTemp[i] = v - (specialOne[i] * part)
				}
			}
		}

		log.Printf("residue needsTemp %v,  \n ", needsTemp)
		// 剩下的买单个
		for i, v := range needsTemp {
			costTemp += price[i] * v
		}

		log.Printf("costTemp %d,  \n ", costTemp)
		if cost > costTemp {
			cost = costTemp
		}
	}

	return cost
}

func filterSpecials(special [][]int, needs []int) [][]int {
	// 从后向前遍历，这样删除元素不会影响到还未遍历的元素
	for i := len(special) - 1; i >= 0; i-- {
		isExceed := false
		for ii, vv := range needs {
			if special[i][ii] > vv {
				isExceed = true
				break // 可以提前退出
			}
		}
		if isExceed {
			special = append(special[:i], special[i+1:]...)
		}
	}
	return special
}

/*

(以下代码报错)
panic: runtime error: slice bounds out of range [60:58] [recovered]
        panic: runtime error: slice bounds out of range [60:58]


	问

	// 先过滤去除大于总需求的special
	for i, specialOne := range special {
		isExceed := false
		for ii, vv := range needs {
			if specialOne[ii] > vv {
				isExceed = true
			}
		}
		// 去除
		if isExceed {
			special = append(special[:i], special[i+1:]...)
		}
	}


这里在 special  的 for 循环  中改变 special   会存在什么问题么。存在问题要怎么解决

答


是的，在遍历切片的同时修改切片是很危险的操作，可能会导致以下问题：
1. 跳过元素：因为删除元素后，后面的元素会前移，但索引 `i` 仍会增加
2. 索引越界：因为切片长度在变化
3. 不可预期的遍历行为

这里有两种常见的解决方案：

1. 从后向前遍历（推荐）：

```go
func filterSpecials(special [][]int, needs []int) [][]int {
    // 从后向前遍历，这样删除元素不会影响到还未遍历的元素
    for i := len(special) - 1; i >= 0; i-- {
        isExceed := false
        for ii, vv := range needs {
            if special[i][ii] > vv {
                isExceed = true
                break  // 可以提前退出
            }
        }
        if isExceed {
            special = append(special[:i], special[i+1:]...)
        }
    }
    return special
}
```

2. 创建新切片存储结果：

```go
func filterSpecials(special [][]int, needs []int) [][]int {
    result := make([][]int, 0, len(special))
    // 遍历原切片，将符合条件的元素添加到新切片
    for _, specialOne := range special {
        isExceed := false
        for ii, vv := range needs {
            if specialOne[ii] > vv {
                isExceed = true
                break
            }
        }
        if !isExceed {
            result = append(result, specialOne)
        }
    }
    return result
}
```

两种方案的比较：
- 方案1（从后向前）：
  - 优点：原地修改，不需要额外内存
  - 缺点：代码不够直观

- 方案2（新切片）：
  - 优点：代码清晰，易于理解
  - 缺点：需要额外内存空间

使用示例：
```go
// 使用方案1或方案2
special = filterSpecials(special, needs)
```

我个人更推荐方案2，因为：
1. 代码更清晰，更容易维护
2. 避免了在原切片上的复杂操作
3. 现代计算机的内存通常不是瓶颈
4. 如果需要保留原始数据，方案2更合适

另外，无论使用哪种方案，都建议：
1. 添加提前退出条件（使用 `break`）
2. 将过滤逻辑封装成独立函数
3. 考虑返回新切片而不是修改原切片

*/

// 排序 - 根据能省下钱大小排序
func sortSpecials(price []int, special *[][]int) {
	// k-(省下的钱) v-special的Idx们
	spareToIdxsMap := make(map[int][]int)

	// 其中省下钱的最大金额
	max := 0
	for i, specialOne := range *special {
		total := 0
		for ii, v := range price {
			total += specialOne[ii] * v
		}
		spare := total - specialOne[len(specialOne)-1]
		// if _, ok := spareToIdxsMap[spare]; ok {
		// }
		// 使用append 不需要make
		spareToIdxsMap[spare] = append(spareToIdxsMap[spare], i)
		if spare > max {
			max = spare
		}
	}
	validCount := 0
	newSpecial := make([][]int, len(*special))
	for i, idx := max, 0; i > 0; i-- {
		if v, ok := spareToIdxsMap[i]; ok {
			for _, vv := range v {
				newSpecial[idx] = (*special)[vv]
				idx++
				validCount++
			}
		} else {
			continue
		}
	}

	// log.Printf("newSpecial %v ,cap %d ,len %d \n", newSpecial, cap(newSpecial), len(newSpecial))

	// 这样并不能 改变函数外部的值
	// special = &newSpecial
	// 这样可以改变函数外部的参数的值
	*special = newSpecial[:validCount]

	// log.Printf("validCount %d   \n", validCount)
	// 舍去为负数或者0 的套餐
	// special = special[:validCount]
	// for i := range special {
	// 	special[i] = newSpecial[i]
	// }
	// log.Printf("special %v ,cap %d ,len %d \n \n", special, cap(special), len(special))
}

/*

问

以下为执行的方法
// 排序 - 根据能省下钱大小排序
func sortSpecials(price []int, special [][]int) {
	// k-(省下的钱) v-special的Idx们
	spareToIdxsMap := make(map[int][]int)
	// 其中省下钱的最大金额
	max := 0
	for i, specialOne := range special {
		total := 0
		for ii, v := range price {
			total += specialOne[ii] * v
		}
		spare := total - specialOne[len(specialOne)-1]
		// if _, ok := spareToIdxsMap[spare]; ok {
		// }
		// 使用append 不需要make
		spareToIdxsMap[spare] = append(spareToIdxsMap[spare], i)
		if spare > max {
			max = spare
		}
	}
	newSpecial := make([][]int, len(special))
	for i, idx := max, 0; i >= 0; i-- {
		if v, ok := spareToIdxsMap[i]; ok {
			for _, vv := range v {
				newSpecial[idx] = special[vv]
				idx++
			}
		} else {
			continue
		}
	}
	special = newSpecial

	log.Printf("newSpecial %v \n", newSpecial)
}


主函数为


	log.Printf("sort before %v \n", special)
	sortSpecials(price, special[:])
	log.Printf("sort after %v \n", special)



这里 special 未切片类型 应该是引用传递  但是在方法sortSpecials 改变了其中 special 但是主函数没有发生顺序改变呢



答

我明白问题所在了。这是 Go 语言中切片参数传递的一个常见误解。让我解释一下问题并提供解决方案：

1. 虽然 `special[:]` 是切片，并且切片确实是引用类型，但是在 `sortSpecials` 函数中的赋值 `special = newSpecial` 只是改变了局部变量 `special` 的引用，而不是修改底层数组的内容。

2. 要修改原切片的内容，我们需要直接修改切片的元素，而不是重新赋值切片变量。

这是修改后的代码：

```go:path/to/file
func sortSpecials(price []int, special [][]int) {
    // ... existing code ...

    // 修改这里：不要重新赋值special，而是复制元素
    for i := range special {
        special[i] = newSpecial[i]
    }

    log.Printf("newSpecial %v \n", newSpecial)
}
```

或者另一种方案是返回排序后的切片：

```go:path/to/file
func sortSpecials(price []int, special [][]int) [][]int {
    // ... existing code ...

    return newSpecial
}

// 主函数中这样调用：
special = sortSpecials(price, special[:])
```

解释：
1. 在 Go 中，当你传递切片时，你实际上是传递了一个包含指针、长度和容量的结构体的副本
2. 当你在函数内部对切片变量重新赋值时，只是修改了这个副本，而不是原始切片
3. 要修改原始切片的内容，你需要：
   - 要么直接修改切片的元素（方案1）
   - 要么返回新切片并在调用处重新赋值（方案2）

这就是为什么原代码中主函数里的 `special` 没有改变顺序的原因。
*/

func shoppingOffersLeetCode(price []int, special [][]int, needs []int) int {
	n := len(price)

	// 过滤不需要计算的大礼包，只保留需要计算的大礼包
	filterSpecial := [][]int{}
	for _, s := range special {
		totalCount, totalPrice := 0, 0
		for i, c := range s[:n] {
			totalCount += c
			totalPrice += c * price[i]
		}
		if totalCount > 0 && totalPrice > s[n] {
			filterSpecial = append(filterSpecial, s)
		}
	}
	log.Printf("filterSpecial %v \n", filterSpecial)
	// 记忆化搜索计算满足购物清单所需花费的最低价格
	dp := map[string]int{}
	var dfs func([]byte) int
	dfs = func(curNeeds []byte) (minPrice int) {
		if res, has := dp[string(curNeeds)]; has {
			return res
		}
		for i, p := range price {
			minPrice += int(curNeeds[i]) * p // 不购买任何大礼包，原价购买购物清单中的所有物品
		}
		nextNeeds := make([]byte, n)
	outer:
		for _, s := range filterSpecial {
			for i, need := range curNeeds {
				if need < byte(s[i]) { // 不能购买超出购物清单指定数量的物品
					continue outer
				}
				nextNeeds[i] = need - byte(s[i])
			}
			minPrice = min(minPrice, dfs(nextNeeds)+s[n])
		}
		dp[string(curNeeds)] = minPrice
		return
	}

	curNeeds := make([]byte, n)
	for i, need := range needs {
		curNeeds[i] = byte(need)
	}
	return dfs(curNeeds)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/shopping-offers/solutions/1062534/da-li-bao-by-leetcode-solution-p1ww/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
