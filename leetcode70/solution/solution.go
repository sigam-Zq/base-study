package solution

import "container/heap"

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/maximum-number-of-eaten-apples/solutions/1172766/chi-ping-guo-de-zui-da-shu-mu-by-leetcod-93ka/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/*
方法一：贪心 + 优先队列
为了将吃苹果的数目最大化，应该使用贪心策略，在尚未腐烂的苹果中优先选择腐烂日期最早的苹果。

为了得到腐烂日期最早的苹果，可以使用优先队列存储每个苹果的腐烂日期，优先队列中最小的元素（即最早的腐烂日期）会最先被取出。由于数组 apples 和 days 的长度 n 最大为 2×10
4
 ，两个数组中的每个元素最大为 2×10
4
 ，因此苹果的总数最大可达 (2×10
4
 )×(2×10
4
 )=4×10
8
 。如果直接使用优先队列存储每个苹果的腐烂日期，则会导致优先队列中的元素个数过多，时间复杂度和空间复杂度过高，因此需要使用优化的表示法。可以在优先队列中存储二元组，每个二元组表示苹果的腐烂日期和在该日期腐烂的苹果个数，则优先队列中的元素个数最多为 n 个（即数组长度），可以显著降低时间复杂度和空间复杂度。

计算吃苹果的最大数目分成两个阶段，第一阶段是第 0 天到第 n−1 天，即天数在数组下标范围内，第二阶段是第 n 天及以后，即天数在数组下标范围外。

在第一阶段，由于每天树上都可能长出苹果，因此需要对每一天分别计算是否能吃到苹果，并更新优先队列。具体做法如下：

将优先队列中的所有腐烂日期小于等于当前日期的元素取出，这些元素表示已经腐烂的苹果，无法食用；

根据 days 和 apples 的当前元素计算当天长出的苹果的腐烂日期和数量，如果数量大于 0，则将腐烂日期和数量加入优先队列；

如果优先队列不为空，则当天可以吃 1 个苹果，将优先队列的队首元素的数量减 1，如果队首元素的数量变成 0 则将队首元素取出。

在第二阶段，由于树上不会再长出苹果，因此只需要考虑能吃到的苹果数量。由于优先队列中的每个元素的数量可能很大，因此需要根据当前日期和优先队列的队首元素的腐烂日期和数量计算能吃到的苹果数量。

假设当前日期是第 i 天，首先将优先队列中的所有腐烂日期小于等于 i 的元素取出，如果优先队列不为空，则根据优先队列的队首元素计算能吃到的苹果数量。假设优先队列的队首元素的腐烂日期是 x，数量是 y，其中 x>i，则有 y 个苹果，距离腐烂还有 x−i 天，因此能吃到的苹果数量是 curr=min(x−i,y)。将优先队列的队首元素 (x,y) 取出并将日期增加 curr，重复上述操作直到优先队列为空，即可得到吃苹果的最大数目。

*/

func eatenApples(apples, days []int) (ans int) {
	h := hp{}
	i := 0
	for ; i < len(apples); i++ {
		for len(h) > 0 && h[0].end <= i {
			heap.Pop(&h)
		}
		if apples[i] > 0 {
			heap.Push(&h, pair{i + days[i], apples[i]})
		}
		if len(h) > 0 {
			h[0].left--
			if h[0].left == 0 {
				heap.Pop(&h)
			}
			ans++
		}
	}
	for len(h) > 0 {
		for len(h) > 0 && h[0].end <= i {
			heap.Pop(&h)
		}
		if len(h) == 0 {
			break
		}
		p := heap.Pop(&h).(pair)
		num := min(p.end-i, p.left)
		ans += num
		i += num
	}
	return
}

type pair struct{ end, left int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
复杂度分析

时间复杂度：O(nlogn)，其中 n 是数组 apples 和 days 的长度。优先队列中最多有 n 个元素，每个元素加入优先队列和从优先队列中取出各一次，每次操作的时间复杂度是 O(logn)，因此总时间复杂度是 O(nlogn)。

空间复杂度：O(n)，其中 n 是数组 apples 和 days 的长度。空间复杂度主要取决于优先队列，优先队列中的元素个数不会超过 n。

*/
