package leetcode73

import (
	"reflect"
	"strconv"
	"testing"
)

/*

3159. 查询数组中元素的出现位置
中等

给你一个整数数组 nums ，一个整数数组 queries 和一个整数 x 。

对于每个查询 queries[i] ，你需要找到 nums 中第 queries[i] 个 x 的位置，并返回它的下标。如果数组中 x 的出现次数少于 queries[i] ，该查询的答案为 -1 。

请你返回一个整数数组 answer ，包含所有查询的答案。



示例 1：

输入：nums = [1,3,1,7], queries = [1,3,2,4], x = 1

输出：[0,-1,2,-1]

解释：

第 1 个查询，第一个 1 出现在下标 0 处。
第 2 个查询，nums 中只有两个 1 ，所以答案为 -1 。
第 3 个查询，第二个 1 出现在下标 2 处。
第 4 个查询，nums 中只有两个 1 ，所以答案为 -1 。
示例 2：

输入：nums = [1,2,3], queries = [10], x = 5

输出：[-1]

解释：

第 1 个查询，nums 中没有 5 ，所以答案为 -1 。


提示：

1 <= nums.length, queries.length <= 105
1 <= queries[i] <= 105
1 <= nums[i], x <= 104
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f       func([]int, []int, int) []int
		nums    []int
		queries []int
		x       int
		want    []int
	}{
		{
			f:       occurrencesOfElement,
			nums:    []int{1, 3, 1, 7},
			queries: []int{1, 3, 2, 4},
			x:       1,
			want:    []int{0, -1, 2, -1},
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.nums, v.queries, v.x)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}

func occurrencesOfElement(nums []int, queries []int, x int) []int {

	// xMap[i] = j x的第i个 出现在j位置
	xMap := make(map[int]int)
	xFrequency := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == x {
			xMap[xFrequency] = i
			xFrequency++
		}
	}

	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {

		if queries[i] >= xFrequency {
			ans[i] = -1
		} else {
			ans[i] = xMap[queries[i]]
		}
	}

	return ans
}
