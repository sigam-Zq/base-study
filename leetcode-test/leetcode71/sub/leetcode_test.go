package leetcode71

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

//  3219. 切蛋糕的最小总开销 II

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f             func(int, int, []int, []int) int64
		m             int
		n             int
		horizontalCut []int
		verticalCut   []int
		want          int64
	}{
		{
			f:             minimumCost,
			m:             3,
			n:             2,
			horizontalCut: []int{1, 3},
			verticalCut:   []int{5},
			want:          13,
		},
		{
			f:             minimumCost,
			m:             2,
			n:             2,
			horizontalCut: []int{7},
			verticalCut:   []int{4},
			want:          15,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.m, v.n, v.horizontalCut, v.verticalCut)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}

}

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
	ans := int64(0)

	sort.Ints(horizontalCut)
	sort.Ints(verticalCut)
	hIdx := len(horizontalCut) - 1
	vIdx := len(verticalCut) - 1
	hSeg, vSeg := 1, 1

	for hIdx >= 0 && vIdx >= 0 {
		if horizontalCut[hIdx] > verticalCut[vIdx] {
			ans += int64(horizontalCut[hIdx] * hSeg)
			vSeg++
			hIdx--
		} else {
			ans += int64(verticalCut[vIdx] * vSeg)
			hSeg++
			vIdx--
		}
	}

	for hIdx >= 0 {
		ans += int64(horizontalCut[hIdx] * hSeg)
		hIdx--
	}

	for vIdx >= 0 {
		ans += int64(verticalCut[vIdx] * vSeg)
		vIdx--
	}

	return ans
}

func minCost(m int, n int, horizontalcut []int, verticalcut []int) int {
	// 排序水平和垂直切割数组（降序）
	sort.Sort(sort.Reverse(sort.IntSlice(horizontalcut)))
	sort.Sort(sort.Reverse(sort.IntSlice(verticalcut)))

	// 初始化段数
	hSegments := 1
	vSegments := 1

	// 指针用于遍历切割数组
	hIndex, vIndex := 0, 0

	// 最小总开销
	totalCost := 0

	// 贪心选择开销较大的切割线
	for hIndex < len(horizontalcut) && vIndex < len(verticalcut) {
		if horizontalcut[hIndex] >= verticalcut[vIndex] {
			// 水平切割
			totalCost += horizontalcut[hIndex] * vSegments
			hSegments++
			hIndex++
		} else {
			// 垂直切割
			totalCost += verticalcut[vIndex] * hSegments
			vSegments++
			vIndex++
		}
	}

	// 剩余的水平切割线
	for hIndex < len(horizontalcut) {
		totalCost += horizontalcut[hIndex] * vSegments
		hIndex++
	}

	// 剩余的垂直切割线
	for vIndex < len(verticalcut) {
		totalCost += verticalcut[vIndex] * hSegments
		vIndex++
	}

	return totalCost
}

// func main() {
// 	// 示例输入
// 	m := 3
// 	n := 3
// 	horizontalcut := []int{2, 1}
// 	verticalcut := []int{3, 1}

// 	// 计算最小开销
// 	result := minCost(m, n, horizontalcut, verticalcut)
// 	fmt.Println(result) // 输出 9
// }
