package leetcode45

import (
	"container/heap"
	"math"
)

var (
	numsC2 [][]int
	next2  []int
)

func Max2(x, y int) int {
	return max(x, y)
}

func smallestRangeCot(nums [][]int) []int {
	numsC2 = nums
	max := math.MinInt
	l := len(nums)
	next2 = make([]int, l)

	h := &IHead{}
	heap.Init(h)
	for i, v := range nums {
		heap.Push(h, i)
		max = Max2(max, v[0])
	}

	// for i := 0; i < l; i++ {
	// 	heap.Push(h, i)
	// 	max = Max2(max, nums[i][0])
	// }

	left, right := 0, math.MaxInt
	minRange := right - left
	for {
		minIdx := heap.Pop(h).(int)
		curRange := max - nums[minIdx][next2[minIdx]]
		if minRange > curRange {
			left, right = nums[minIdx][next2[minIdx]], max
			minRange = curRange
		}
		next2[minIdx]++
		if next2[minIdx] == len(nums[minIdx]) {
			break
		}

		heap.Push(h, minIdx)
		max = Max2(max, nums[minIdx][next2[minIdx]])
	}

	return []int{left, right}
}

type IHead []int

func (H IHead) Len() int {
	return len(H)
}

func (H IHead) Swap(i, j int) {
	H[i], H[j] = H[j], H[i]
	return
}

func (H IHead) Less(i, j int) bool {
	return numsC2[H[i]][next2[H[i]]] < numsC2[H[j]][next2[H[j]]]
}

func (H *IHead) Pop() any {
	old := *H
	n := len(old)
	x := old[n-1]
	*H = old[:n-1]
	return x
}

func (H *IHead) Push(v any) {
	*H = append(*H, v.(int))
}
