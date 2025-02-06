package leetcode107

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func([]int) [][]int
		nums []int
		want [][]int
	}{
		{
			f:    subsetsWithDup,
			nums: []int{1, 2, 2},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		{
			f:    subsetsWithDup,
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			f:    subsetsWithDup,
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}},
		},
		{
			f:    subsetsWithDup,
			nums: []int{2, 1, 2, 1, 3},
			want: [][]int{{}, {1}, {1, 1}, {1, 1, 2}, {1, 1, 2, 2}, {1, 1, 2, 2, 3}, {1, 1, 2, 3}, {1, 1, 3}, {1, 2}, {1, 2, 2}, {1, 2, 2, 3}, {1, 2, 3}, {1, 3}, {2}, {2, 2}, {2, 2, 3}, {2, 3}, {3}},
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.nums)
			fmt.Println(got)
			fmt.Println(v.want)
			if !arraysEqualIgnoringOrderDup(got, v.want) {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	alreadyMap := make(map[string]struct{})
	sort.Ints(nums)
	var dfs func(layer int, res []int)
	dfs = func(layer int, res []int) {

		if _, ok := alreadyMap[fmt.Sprintf("%v", res)]; !ok {
			ans = append(ans, append([]int(nil), res...))
			alreadyMap[fmt.Sprintf("%v", res)] = struct{}{}
		}
		if layer == n {
			return
		}
		// copy1 := deepCopy(res)
		dfs(layer+1, append(res, nums[layer]))
		// copy2 := deepCopy(res)
		dfs(layer+1, res)
	}
	dfs(0, make([]int, 0))
	// 不满足跨元素组 子集
	// for i, j := 0, 0; i <= j && j <= n; {
	// 	nn := len(ans)
	// 	var isNotExist = true
	// 	for ii := 0; ii < nn; ii++ {
	// 		if arraysEqualIgnoringOrderMap(ans[ii], nums[i:j]) {
	// 			isNotExist = false
	// 			break
	// 		}
	// 	}

	// 	if isNotExist {
	// 		ans = append(ans, nums[i:j])
	// 	}

	// 	if j < n {
	// 		j++
	// 	} else {
	// 		i++
	// 		j = i
	// 	}
	// }

	return ans
}

func deepCopy(in []int) []int {
	out := make([]int, len(in))
	copy(out, in)
	return out
}

// 对每个子数组进行排序
func sortEachSubArray(arr [][]int) {
	for i := range arr {
		sort.Ints(arr[i])
	}
}

// 排序整个二维数组
func sort2DArray(arr [][]int) {
	sort.Slice(arr, func(i, j int) bool {
		for x := 0; x < len(arr[i]) && x < len(arr[j]); x++ {
			if arr[i][x] != arr[j][x] {
				return arr[i][x] < arr[j][x]
			}
		}
		return len(arr[i]) < len(arr[j])
	})
}

// 判断两个二维数组是否相等（忽略顺序）
func arraysEqualIgnoringOrderDup(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// 先对子数组排序
	sortEachSubArray(a)
	sortEachSubArray(b)

	// 再对整个二维数组排序
	sort2DArray(a)
	sort2DArray(b)

	// 逐个比较子数组
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func arraysEqualIgnoringOrderDual(a, b [][]int) bool {

	return true
}

func arraysEqualIgnoringOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	aCopy := make([]int, len(a))
	bCopy := make([]int, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	sort.Ints(aCopy)
	sort.Ints(bCopy)
	for i := range aCopy {
		if aCopy[i] != bCopy[i] {
			return false
		}
	}
	return true
}

func arraysEqualIgnoringOrderMap(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[int]int)

	// 统计数组 a 中元素的频次
	for _, num := range a {
		counts[num]++
	}

	// 统计数组 b，减少对应元素的频次
	for _, num := range b {
		if counts[num] == 0 {
			return false
		}
		counts[num]--
	}

	return true
}

// 将子数组转换为唯一字符串（先排序保证一致性）
func arrayToString(arr []int) string {
	sort.Ints(arr) // 确保 [1,2] 和 [2,1] 视为相同
	return strings.Trim(fmt.Sprint(arr), "[]")
}

// 判断两个二维数组是否相等（忽略顺序）
func arraysEqualIgnoringOrderMapDup(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[string]int)

	// 统计 a 的子数组频次
	for _, subArr := range a {
		key := arrayToString(subArr)
		counts[key]++
	}

	// 统计 b，减少对应元素的频次
	for _, subArr := range b {
		key := arrayToString(subArr)
		if counts[key] == 0 {
			return false
		}
		counts[key]--
	}

	return true
}
