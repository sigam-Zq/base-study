package main

import "container/heap"

func getFinalStateChatGPT(nums []int, k int, multiplier int) []int {
	mod := int(1e9 + 7)

	// var h *MyHeap2
	h := &MyHeap2{}

	heap.Init(h)
	for i, v := range nums {
		u := []int{v, i}
		heap.Push(h, u)
	}
	// log.Println(h)
	// log.Println(heap.Pop(h).([]int))

	for i := 0; i < k; i++ {
		minIdx := heap.Pop(h).([]int)[1]
		nums[minIdx] = (nums[minIdx] * multiplier) % mod
		u := []int{nums[minIdx], minIdx}
		heap.Push(h, u)
	}

	return nums
}

// 初始化自己的堆
type MyHeap2 [][]int

func (m MyHeap2) Len() int {
	return len(m)
}

func (m MyHeap2) Less(i, j int) bool {
	if m[i][0] == m[j][0] {
		return m[i][1] < m[j][1]
	}
	return m[i][0] < m[j][0]
}

func (m MyHeap2) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MyHeap2) Push(a any) {
	(*m) = append(*m, a.([]int))
}
func (m *MyHeap2) Pop() any {
	// old := &m
	x := (*m)[m.Len()-1]
	(*m) = (*m)[:m.Len()-1]
	return x
}
