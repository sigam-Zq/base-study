package leetcode

import (
	"container/heap"
	"reflect"
	"strconv"
	"testing"
)

// 3066. 超过阈值的最少操作数 II
/*

给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。

一次操作中，你将执行：

选择 nums 中最小的两个整数 x 和 y 。
将 x 和 y 从 nums 中删除。
将 min(x, y) * 2 + max(x, y) 添加到数组中的任意位置。
注意，只有当 nums 至少包含两个元素时，你才可以执行以上操作。

你需要使数组中的所有元素都大于或等于 k ，请你返回需要的 最少 操作次数。



示例 1：

输入：nums = [2,11,10,1,3], k = 10
输出：2
解释：第一次操作中，我们删除元素 1 和 2 ，然后添加 1 * 2 + 2 到 nums 中，nums 变为 [4, 11, 10, 3] 。
第二次操作中，我们删除元素 3 和 4 ，然后添加 3 * 2 + 4 到 nums 中，nums 变为 [10, 11, 10] 。
此时，数组中的所有元素都大于等于 10 ，所以我们停止操作。
使数组中所有元素都大于等于 10 需要的最少操作次数为 2 。
示例 2：

输入：nums = [1,1,2,4,9], k = 20
输出：4
解释：第一次操作后，nums 变为 [2, 4, 9, 3] 。
第二次操作后，nums 变为 [7, 4, 9] 。
第三次操作后，nums 变为 [15, 9] 。
第四次操作后，nums 变为 [33] 。
此时，数组中的所有元素都大于等于 20 ，所以我们停止操作。
使数组中所有元素都大于等于 20 需要的最少操作次数为 4 。


提示：

2 <= nums.length <= 2 * 105
1 <= nums[i] <= 109
1 <= k <= 109
输入保证答案一定存在，也就是说一定存在一个操作序列使数组中所有元素都大于等于 k 。
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f       func([]int, int) int
		nums    []int
		k       int
		want    int
		isDebug bool
	}{
		{
			f:       minOperations,
			nums:    []int{2, 11, 10, 1, 3},
			k:       10,
			want:    2,
			isDebug: false,
		},
		{
			f:       minOperations,
			nums:    []int{1, 1, 2, 4, 9},
			k:       20,
			want:    4,
			isDebug: false,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans := v.f(v.nums, v.k)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

type myHeap []int

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

func minOperations(nums []int, k int) int {

	list := &myHeap{}
	heap.Init(list)
	for _, v := range nums {
		heap.Push(list, v)
	}

	var ans int

	cur := heap.Pop(list).(int)
	cur2 := heap.Pop(list).(int)
	for cur < k {
		heap.Push(list, min(cur, cur2)*2+max(cur, cur2))
		ans++
		if list.Len() < 2 {
			break
		}
		cur = heap.Pop(list).(int)
		cur2 = heap.Pop(list).(int)
		// fmt.Printf(" len %d list %+v cur %d cur2 %d \n", list.Len(), list, cur, cur2)
	}

	return ans
}

func minOperationsLeetCode(nums []int, k int) int {
	res := 0
	pq := &MinHeap{}
	heap.Init(pq)
	for _, num := range nums {
		heap.Push(pq, num)
	}

	for (*pq)[0] < k {
		x := heap.Pop(pq).(int)
		y := heap.Pop(pq).(int)
		heap.Push(pq, x+x+y)
		res++
	}

	return res
}

// MinHeap
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/minimum-operations-to-exceed-threshold-value-ii/solutions/3040119/chao-guo-yu-zhi-de-zui-shao-cao-zuo-shu-y7tgx/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
