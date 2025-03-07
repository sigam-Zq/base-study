package main

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SeatManager struct {
	availableSeats *IntHeap
}

func Constructor(n int) SeatManager {
	h := &IntHeap{}
	for i := 1; i <= n; i++ {
		heap.Push(h, i)
	}
	return SeatManager{availableSeats: h}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.availableSeats).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.availableSeats, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */

/*

//   方案一

 type SeatManager struct {
	seat []bool
}

func Constructor(n int) SeatManager {
	seat := make([]bool, n)
	return SeatManager{seat: seat}
}

func (this *SeatManager) Reserve() int {
	for i, v := range this.seat {
		if v {
			v = true
			return i + 1
		}
	}

	return 0
}

func (this *SeatManager) Unreserve(seatNumber int) {
	this.seat[seatNumber-1] = false
}


*/

/*

// 方案2 （优化第一版 --存在错误）

type SeatManager struct {
    seat []bool
    idx int
}


func Constructor(n int) SeatManager {
    seat := make([]bool,n)
    return SeatManager{seat:seat,idx:0}
}


func (this *SeatManager) Reserve() int {

    for _,v :=range this.seat[this.idx:] {
        if !v {
            this.seat[this.idx] = true
            if this.idx < len(this.seat) {
                this.idx++
            }
            return this.idx
        }
    }



    // for i,v :=range this.seat {
    //     if !v {
    //         // v = true
    //         this.seat[i] = true
    //         return i+1
    //     }
    // }

    return 0
}


func (this *SeatManager) Unreserve(seatNumber int)  {
    if seatNumber <= this.idx{
        this.idx = seatNumber-1
    }
    this.seat[seatNumber-1] = false
}


*/

/*

 您提供的题解确实可以实现所需功能，但正如您所说，它的时间复杂度不够优化。我们可以通过使用最小堆（优先队列）来显著提高性能。这里是一个优化后的解决方案：

```go:seatmanager.go
import (
    "container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type SeatManager struct {
    availableSeats *IntHeap
}

func Constructor(n int) SeatManager {
    h := &IntHeap{}
    for i := 1; i <= n; i++ {
        heap.Push(h, i)
    }
    return SeatManager{availableSeats: h}
}

func (this *SeatManager) Reserve() int {
    return heap.Pop(this.availableSeats).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
    heap.Push(this.availableSeats, seatNumber)
}
```

这个优化后的解决方案有以下几个主要改进：

1. 使用最小堆（优先队列）来存储可用的座位编号，而不是布尔数组。

2. `Reserve()` 方法现在的时间复杂度是 O(log n)，因为它只需要从堆中弹出最小的元素。

3. `Unreserve(seatNumber int)` 方法的时间复杂度也是 O(log n)，因为它只需要将座位号插入堆中。

4. 初始化时，我们将所有座位号添加到堆中，这个操作的时间复杂度是 O(n log n)。

这个优化后的解决方案在大多数操作上都比原来的解决方案快得多，特别是在频繁调用 `Reserve()` 和 `Unreserve()` 的情况下。原来的解决方案在 `Reserve()` 方法中最坏情况下需要 O(n) 的时间复杂度，而新的解决方案始终保持 O(log n) 的时间复杂度。
*/
