package leetcode45

import (
	"container/heap"
	"log"
	"math"
	"reflect"
	"strconv"
	"testing"
)

/*

https://leetcode.cn/problems/smallest-range-covering-elements-from-k-lists/solutions/355881/zui-xiao-qu-jian-by-leetcode-solution/?envType=daily-question&envId=2024-11-24

632. 最小区间
你有 k 个 非递减排列 的整数列表。找到一个 最小 区间，使得 k 个列表中的每个列表至少有一个数包含在其中。

我们定义如果 b-a < d-c 或者在 b-a == d-c 时 a < c，则区间 [a,b] 比 [c,d] 小。



示例 1：

输入：nums = [[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
输出：[20,24]
解释：
列表 1：[4, 10, 15, 24, 26]，24 在区间 [20,24] 中。
列表 2：[0, 9, 12, 20]，20 在区间 [20,24] 中。
列表 3：[5, 18, 22, 30]，22 在区间 [20,24] 中。
示例 2：

输入：nums = [[1,2,3],[1,2,3],[1,2,3]]
输出：[1,1]


提示：

nums.length == k
1 <= k <= 3500
1 <= nums[i].length <= 50
-10^5 <= nums[i][j] <= 10^5
nums[i] 按非递减顺序排列

*/

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f    func([][]int) []int
		nums [][]int
		want []int
	}{
		{
			f:    smallestRangeCot,
			nums: [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}},
			want: []int{20, 24},
		},
		{
			f:    smallestRange,
			nums: [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}},
			want: []int{20, 24},
		},
		{
			f:    smallestRange,
			nums: [][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}},
			want: []int{1, 7},
		},
		{
			f:    smallestRange,
			nums: [][]int{{1, 5}},
			want: []int{1, 1},
		},
		{
			f: smallestRange,
			want: []int{24463, 47806},
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.nums); !reflect.DeepEqual(got, v.want) {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})

	}
}

func smallestRange(nums [][]int) []int {

	// 维护一个 多个数组扁平化处理 flatList 以上数组压平为一个数组
	flatList := make([]int, 0)
	// 维护一个 flatIndex --> 来自于 nums的第几个数组
	flatIdxNums := make([]int, 0)

	k := len(nums)
	done := 0
	for done < k {
		// minIdx 还需要调整到一个不为0的数组上
		minIdx := 0
		for i, v := range nums {
			if len(v) > 0 {
				minIdx = i
			}
		}
		// 找到最小的那个索引
		for j := 0; j < k; j++ {
			if len(nums[j]) == 0 {
				continue
			}
			if nums[minIdx][0] >= nums[j][0] {
				minIdx = j
			}
		}

		flatList = append(flatList, nums[minIdx][0])
		flatIdxNums = append(flatIdxNums, minIdx)
		// 取值后截短被取走的数字
		nums[minIdx] = nums[minIdx][1:]

		if len(nums[minIdx]) == 0 {
			// 如果其中一个数组归零就 结束标志加1
			done++
		}
	}
	// log.Printf("flatList %v  \n", flatList)
	// log.Printf("flatIdxNums %v  \n", flatIdxNums)
	// 这里可以转化 使用动态规划了
	// 求得 rightIdx -- leftIdx [flatList[rightIdx],flatList[leftIdx]] 区间最小
	// 且  [flatIdxNums[rightIdx],flatIdxNums[leftIdx]] 满足 0->k 均存在

	cursor := initCursorList(k)

	// done = 0
	leftIdx, rightIdx, idx := 0, 0, 0
	resIdx := make([]int, 2)
	resIdx[0], resIdx[1] = -1, -1
	// res 索引对应的 差值
	resDeviation := 0
	for idx < len(flatList) {
		// cursor.cursors[flatIdxNums[idx]] = idx
		// log.Printf("before outer cursor %+v \n", cursor)
		cursor.PushUnit(flatIdxNums[idx], idx)
		// log.Printf("after outer cursor %+v \n", cursor)
		// log.Printf("outer cursor.Sum %d \n", cursor.Sum())

		// if cursor.Sum() >= 0 {
		if cursor.sum >= 0 {
			// 游标已满的情况（全部都有）
			// log.Printf("cursor %v \n", cursor)
			leftIdx, rightIdx = cursor.First(), cursor.Last()
			// leftIdx, rightIdx = cursor.first, cursor.last
			// log.Printf("leftIdx %d rightIdx %d \n", leftIdx, rightIdx)
			// 初次初始化
			if resIdx[0] == -1 && resIdx[1] == -1 {
				resIdx[0], resIdx[1] = leftIdx, rightIdx
				resDeviation = flatList[rightIdx] - flatList[leftIdx]
				// log.Printf("resDeviation %d leftIdx %d rightIdx %d \n", resDeviation, leftIdx, rightIdx)
			} else {
				// 对比取出最小的偏差值
				if resDeviation > flatList[rightIdx]-flatList[leftIdx] {
					resIdx[0], resIdx[1] = leftIdx, rightIdx
					resDeviation = flatList[rightIdx] - flatList[leftIdx]
				}
			}

		}
		// log.Printf("now i %d cursor %v resIdx %v resDeviation %d \n", idx, cursor, resIdx, resDeviation)
		idx++
	}

	// log.Printf("resIdx %v \n", resIdx)
	return []int{flatList[resIdx[0]], flatList[resIdx[1]]}
}

type CursorList struct {
	cursors []int
	sum     int
	first   int
	last    int
}

// cursor长度 初始内容为 -2 * n 确保区分只有游标满的时候 和为正数
func initCursorList(n int) *CursorList {
	cursorList := make([]int, n)
	sum := 0
	for i, _ := range cursorList {
		cursorList[i] = -math.MaxInt / n
		sum += cursorList[i]
	}
	return &CursorList{
		cursors: cursorList,
		sum:     sum,
		first:   math.MaxInt,
		last:    0,
	}
}

// 求和-判断 当前游标是否有效
func (c CursorList) Sum() int {
	sum := 0
	for _, v := range c.cursors {
		sum += v
	}
	return sum
}

// 获取当前游标中的最大值
func (c *CursorList) PushUnit(idx, v int) {

	// log.Printf("PushUnit idx %d v %d \n", idx, v)
	// 更新sum
	c.sum -= c.cursors[idx]
	c.sum += v

	c.cursors[idx] = v

	// // 有效之后 再去更新这两个值
	// if c.sum >= 0 {
	// 更新 最大和最小值
	// if c.first > v {
	// 	c.first = v
	// }
	// if c.last < v {
	// 	c.last = v
	// }
	// }

	// log.Printf("inner cursor %+v \n", c)
}

// 获取当前游标中的最大值
func (c CursorList) Last() int {
	max := 0
	for _, v := range c.cursors {
		if max < v {
			max = v
		}
	}
	return max
}

// 获取当前游标中的最小值
func (c CursorList) First() int {
	min := math.MaxInt
	for _, v := range c.cursors {
		if min > v {
			min = v
		}
	}
	return min
}

// ---------------------------leetCode------------------------------------------------------

/*
方法一：贪心 + 最小堆
给定 k 个列表，需要找到最小区间，使得每个列表都至少有一个数在该区间中。该问题可以转化为，从 k 个列表中各取一个数，使得这 k 个数中的最大值与最小值的差最小。

假设这 k 个数中的最小值是第 i 个列表中的 x，对于任意 j

=i，设第 j 个列表中被选为 k 个数之一的数是 y，则为了找到最小区间，y 应该取第 j 个列表中大于等于 x 的最小的数，这是一个贪心的策略。贪心策略的正确性简单证明如下：假设 z 也是第 j 个列表中的数，且 z>y，则有 z−x>y−x，同时包含 x 和 z 的区间一定不会小于同时包含 x 和 y 的区间。因此，其余 k−1 个列表中应该取大于等于 x 的最小的数。

由于 k 个列表都是升序排列的，因此对每个列表维护一个指针，通过指针得到列表中的元素，指针右移之后指向的元素一定大于或等于之前的元素。

使用最小堆维护 k 个指针指向的元素中的最小值，同时维护堆中元素的最大值。初始时，k 个指针都指向下标 0，最大元素即为所有列表的下标 0 位置的元素中的最大值。每次从堆中取出最小值，根据最大值和最小值计算当前区间，如果当前区间小于最小区间则用当前区间更新最小区间，然后将对应列表的指针右移，将新元素加入堆中，并更新堆中元素的最大值。

如果一个列表的指针超出该列表的下标范围，则说明该列表中的所有元素都被遍历过，堆中不会再有该列表中的元素，因此退出循环。

作者：力扣官方题解
链接：https://leetcode.cn/problems/smallest-range-covering-elements-from-k-lists/solutions/355881/zui-xiao-qu-jian-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

var (
	next  []int
	numsC [][]int
)

func smallestRangeLeetcode1(nums [][]int) []int {
	numsC = nums
	rangeLeft, rangeRight := 0, math.MaxInt32
	minRange := rangeRight - rangeLeft
	max := math.MinInt32
	size := len(nums)
	next = make([]int, size)
	h := &IHeap{}
	heap.Init(h)

	for i := 0; i < size; i++ {
		heap.Push(h, i)
		max = Max(max, nums[i][0])
	}
	log.Printf("head %+v \n", h)
	for {

		log.Printf("next %+v   \n", next)
		log.Printf("Pop head %+v  \n", h)
		minIndex := heap.Pop(h).(int)
		curRange := max - nums[minIndex][next[minIndex]]

		log.Printf("max %d  min %d \n", max, nums[minIndex][next[minIndex]])
		if curRange < minRange {
			minRange = curRange
			rangeLeft, rangeRight = nums[minIndex][next[minIndex]], max
		}
		next[minIndex]++
		if next[minIndex] == len(nums[minIndex]) {
			break
		}
		heap.Push(h, minIndex)
		log.Printf("Push head %+v \n", h)
		max = Max(max, nums[minIndex][next[minIndex]])
	}
	return []int{rangeLeft, rangeRight}
}

type IHeap []int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return numsC[h[i]][next[h[i]]] < numsC[h[j]][next[h[j]]] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
方法二：哈希表 + 滑动窗口
思路

在讲这个方法之前我们先思考这样一个问题：有一个序列 A={a
1
​
 ,a
2
​
 ,⋯,a
n
​
 } 和一个序列 B={b
1
​
 ,b
2
​
 ,⋯,b
m
​
 }，请找出一个 B 中的一个最小的区间，使得在这个区间中 A 序列的每个数字至少出现一次，请注意 A 中的元素可能重复，也就是说如果 A 中有 p 个 u，那么你选择的这个区间中 u 的个数一定不少于 p。没错，这就是我们五月份的一道打卡题：「76. 最小覆盖子串」。官方题解使用了一种滑动窗口的方法，遍历整个 B 序列并用一个哈希表表示当前窗口中的元素：

右边界在每次遍历到新元素的时候右移，同时将拓展到的新元素加入哈希表；
左边界右移当且仅当当前区间为一个合法的答案区间，即当前窗口内的元素包含 A 中所有的元素，同时将原来左边界指向的元素从哈希表中移除；
答案更新当且仅当当前窗口内的元素包含 A 中所有的元素。
如果这个地方不理解，可以参考「76. 最小覆盖子串的官方题解」。

回到这道题，我们发现这两道题的相似之处在于都要求我们找到某个符合条件的最小区间，我们可以借鉴「76. 最小覆盖子串」的做法：这里序列 {0,1,⋯,k−1} 就是上面描述的 A 序列，即 k 个列表，我们需要在一个 B 序列当中找到一个区间，可以覆盖 A 序列。这里的 B 序列是什么？我们可以用一个哈希映射来表示 B 序列—— B[i] 表示 i 在哪些列表当中出现过，这里哈希映射的键是一个整数，表示列表中的某个数值，哈希映射的值是一个数组，这个数组里的元素代表当前的键出现在哪些列表里。也许文字表述比较抽象，大家可以结合下面这个例子来理解。

如果列表集合为：
0: [-1, 2, 3]
1: [1]
2: [1, 2]
3: [1, 1, 3]
那么可以得到这样一个哈希映射
-1: [0]
 1: [1, 2, 3, 3]
 2: [0, 2]
 3: [0, 3]
我们得到的这个哈希映射就是这里的 B 序列。我们要做的就是在 B 序列上使用两个指针维护一个滑动窗口，并用一个哈希表维护当前窗口中已经包含了哪些列表中的元素，记录它们的索引。遍历 B 序列的每一个元素：

指向窗口右边界的指针右移当且仅当每次遍历到新的元素，并将这个新的元素对应的值数组中的每一个数加入到哈希表中；
指向窗口左边界的指针右移当且仅当当前区间内的元素包含 A 中所有的元素，同时将原来左边界对应的值数组的元素们从哈希表中移除；
答案更新当且仅当当前窗口内的元素包含 A 中所有的元素。
大家可以参考代码理解这个过程。

*/

func smallestRangeLeetCode2(nums [][]int) []int {
	size := len(nums)
	indices := map[int][]int{}
	xMin, xMax := math.MaxInt32, math.MinInt32
	for i := 0; i < size; i++ {
		for _, x := range nums[i] {
			indices[x] = append(indices[x], i)
			xMin = min(xMin, x)
			xMax = max(xMax, x)
		}
	}
	freq := make([]int, size)
	inside := 0
	left, right := xMin, xMin-1
	bestLeft, bestRight := xMin, xMax
	for right < xMax {
		right++
		if len(indices[right]) > 0 {
			for _, x := range indices[right] {
				freq[x]++
				if freq[x] == 1 {
					inside++
				}
			}
			for inside == size {
				if right-left < bestRight-bestLeft {
					bestLeft, bestRight = left, right
				}
				for _, x := range indices[left] {
					freq[x]--
					if freq[x] == 0 {
						inside--
					}
				}
				left++
			}
		}
	}
	return []int{bestLeft, bestRight}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}