package leetcode60

import (
	"container/heap"
	"log"
	"math"
	"sort"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f      func([][]int) int64
		values [][]int
		want   int64
	}{
		{
			f:      maxSpendingRaw,
			values: [][]int{{8, 5, 2}, {6, 4, 1}, {9, 7, 3}},
			want:   285,
		},
		//[[10,8,6,4,2],[9,7,5,3,2]]
		{
			f:      maxSpendingRaw,
			values: [][]int{{10, 8, 6, 4, 2}, {9, 7, 5, 3, 2}},
			want:   386,
		},
	} {

		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.values); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})
	}
}

func maxSpendingRaw(values [][]int) int64 {
	m := len(values)
	n := len(values[0])

	// 索引
	idxList := make([]int, m)
	list := make([]int, m)
	for i := range idxList {
		idxList[i] = n - 1
	}

	for j := m - 1; j >= 0; j-- {
		// list.Push(values[j][idxList[j]])
		list[j] = values[j][idxList[j]]
	}
	var res int64
	for count := 1; count <= n*m; count++ {
		min := math.MaxInt
		minIdx := 0

		for i, v := range list {
			if v == 0 {
				continue
			}
			if min > v {
				min = v
				minIdx = i
			}
		}

		log.Printf("add %d count %d \n", min, count)
		res += int64(min * count)

		if idxList[minIdx] > 0 {
			idxList[minIdx]--
			list[minIdx] = values[minIdx][idxList[minIdx]]
		} else {
			// idx 这一列清空了 去找其他一个最小数的列
			list[minIdx] = 0
		}
	}

	return res
}

// 可以拿到值 排序。 但是找不到具体
type myHeap []int

func maxSpending(values [][]int) int64 {
	m := len(values)
	n := len(values[0])

	list := &myHeap{}

	// 索引
	idxList := make([]int, m)
	for i := range idxList {
		idxList[i] = n - 1
	}

	heap.Init(list)
	for j := m - 1; j >= 0; j-- {
		// list.Push(values[j][idxList[j]])
		heap.Push(list, values[j][idxList[j]])
		idxList[j]--
	}
	var res int64
	for count := 1; count <= n*m; count++ {
		// x := list.Pop().(int)
		x := heap.Pop(list).(int)
		log.Printf("add %d count %d \n", x, count)
		res += int64(x * count)
		// count++
		for i := 0; i < m; i++ {
			if x == values[i][idxList[i]+1] {
				// list.Push(values[i][idxList[i]])
				heap.Push(list, values[i][idxList[i]])
			}
		}
	}

	return res
}

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v any) {
	*h = append(*h, v.(int))
}

func maxSpendingViolenceSort(values [][]int) int64 {
	// m := len(values)
	// n := len(values[0])

	allList := make([]int, 0)

	for _, lis := range values {
		allList = append(allList, lis...)
	}
	sort.Ints(allList)

	var res int64
	for i, v := range allList {
		res += int64((i + 1) * v)
	}

	return res
}

/*
方法一：排序不等式 + 小根堆
思路与算法

由于每一个商店的物品都已经按照价值单调递减排好序了，那么当我们选择某个商店购买物品时，都可以买到该商店中价值最低的物品。由于我们可以任意选择商店，这就说，我们总是可以买到当前所有物品中价值最低的那个。

在开销的计算公式中，物品的价值会乘上购买它的天数。根据排序不等式，在理想状态下我们应该将所有商品按照价值从低到高排序，分别在第 1 到 m×n 天去购买。根据上一段的结论，我们一定是可以达到这个理想状态的。

因此，我们可以将 m×n 个商品按照价值进行排序，就可以得到答案，但这样做的时间复杂度是 O(mnlog(mn))，没有进一步用到「每一个商店的物品都已经按照价值单调递减排好序」这个性质。我们可以使用「23. 合并 K 个升序链表」中的方法，使用一个小根堆，存储每个商店当前价值最小的物品，那么小根堆的堆顶就是全局价值最小的物品。随后，我们将该物品在对应的商店中的下一个物品放入小根堆中，重复一共 m×n 次操作即可，时间复杂度降低至 O(mnlogm)。

作者：力扣官方题解
链接：https://leetcode.cn/problems/maximum-spending-after-buying-items/solutions/3004329/gou-mai-wu-pin-de-zui-da-kai-xiao-by-lee-zwv3/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func maxSpendingLeetCode(values [][]int) int64 {
	m, n := len(values), len(values[0])
	pq := &Heap{}
	for i := 0; i < m; i++ {
		heap.Push(pq, []int{values[i][n-1], i, n - 1})
	}
	ans := int64(0)
	for turn := 1; turn <= m*n; turn++ {
		top := heap.Pop(pq).([]int)
		val, i, j := top[0], top[1], top[2]
		ans += int64(val) * int64(turn)
		if j > 0 {
			heap.Push(pq, []int{values[i][j-1], i, j - 1})
		}
	}
	return ans
}

type Heap [][]int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] < h[j][0]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/maximum-spending-after-buying-items/solutions/3004329/gou-mai-wu-pin-de-zui-da-kai-xiao-by-lee-zwv3/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
