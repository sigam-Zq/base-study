package soulationgo

type RangeFreqQuery struct {
	// 数值为键，出现下标数组为值的哈希表
	occurrence map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	occurrence := make(map[int][]int)
	// 顺序遍历数组初始化哈希表
	for i, v := range arr {
		occurrence[v] = append(occurrence[v], i)
	}
	return RangeFreqQuery{occurrence: occurrence}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	// 查找对应的出现下标数组，不存在则为空
	pos, exists := this.occurrence[value]
	if !exists {
		return 0
	}
	// 两次二分查找计算子数组内出现次数
	l := lowerBound(pos, left)
	r := upperBound(pos, right)
	return r - l
}

func lowerBound(pos []int, target int) int {
	low, high := 0, len(pos)-1
	for low <= high {
		mid := low + (high-low)/2
		if pos[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func upperBound(pos []int, target int) int {
	low, high := 0, len(pos)-1
	for low <= high {
		mid := low + (high-low)/2
		if pos[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/range-frequency-queries/solutions/1115337/qu-jian-nei-cha-xun-shu-zi-de-pin-lu-by-wh4ez/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
