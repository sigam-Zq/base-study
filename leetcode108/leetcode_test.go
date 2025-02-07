package leetcode108

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func([]int) [][]int
		nums []int
		want [][]int
	}{
		{
			f:    permuteUnique,
			nums: []int{1, 1, 2},
			want: [][]int{{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1}},
		},
		{
			f:    permuteUnique,
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.nums)
			fmt.Println(got)
			fmt.Println(v.want)
			if !arraysEqualsDup(got, v.want) {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func permuteUniqueRaw(nums []int) [][]int {
	n := len(nums)
	alreadyMap := make(map[string]struct{})
	ans := make([][]int, 0)

	var dfs func(choice []int, noChoice []int)
	dfs = func(choice []int, noChoice []int) {

		if len(choice) == n {
			if _, ok := alreadyMap[fmt.Sprintf("%v", choice)]; !ok {
				ans = append(ans, append([]int(nil), choice...))
				alreadyMap[fmt.Sprintf("%v", choice)] = struct{}{}
			}
			return
		}
		l := len(noChoice)
		for i := 0; i < l; i++ {
			// TODO 这里传入 {1, 1, 2} 渐渐的 noChoice就变成 {2,2,2}了 找不到原因
			dfs(append(choice, noChoice[i]), append(noChoice[:i], noChoice[i+1:]...))
		}
	}
	dfs(make([]int, 0), nums)
	return ans
}

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	alreadyMap := make(map[string]struct{})
	ans := make([][]int, 0)

	var dfs func(choice []int, noChoice []int)
	dfs = func(choice []int, noChoice []int) {

		if len(choice) == n {
			if _, ok := alreadyMap[fmt.Sprintf("%v", choice)]; !ok {
				ans = append(ans, append([]int(nil), choice...))
				alreadyMap[fmt.Sprintf("%v", choice)] = struct{}{}
			}
			return
		}
		l := len(noChoice)
		for i := 0; i < l; i++ {
			x := noChoice[i]
			// 可以的代码
			// newNoChoice := make([]int, 0, len(noChoice)-1)
			// newNoChoice = append(newNoChoice, noChoice[:i]...)
			// newNoChoice = append(newNoChoice, noChoice[i+1:]...)
			// 可以的代码
			// newNoChoice := make([]int, 0)
			// newNoChoice = append(newNoChoice, noChoice[:i]...)
			// newNoChoice = append(newNoChoice, noChoice[i+1:]...)

			// 不可以的代码
			newNoChoice := make([]int, 0)
			newNoChoice = append(newNoChoice, append(noChoice[:i], noChoice[i+1:]...)...)
			// 不可以的代码
			// newNoChoice := deepCopy(append(noChoice[:i], noChoice[i+1:]...))
			// newNoChoice := deepCopy(append([]int(nil), append(noChoice[:i], noChoice[i+1:]...)...))

			dfs(append(choice, x), newNoChoice)
		}
	}
	dfs(make([]int, 0), nums)
	return ans
}

func arraysEqualsDup(v1 [][]int, v2 [][]int) bool {
	// 内部数组顺序需要判断，外部不需要判断顺序
	v1Map := make(map[string]int)
	for _, v := range v1 {
		v1Map[fmt.Sprintf("%v", v)]++
	}

	for _, v := range v2 {
		if x, ok := v1Map[fmt.Sprintf("%v", v)]; ok && x >= 0 {
			v1Map[fmt.Sprintf("%v", v)]--
		} else {
			return false
		}
	}

	return true
}

func deepCopy(in []int) []int {
	out := make([]int, len(in))
	copy(out, in)
	// fmt.Printf("out %p  in %p \n", &out, &in)
	return out
}

func permuteUniqueChatGpt1(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	visited := make(map[string]bool)

	var dfs func(choice []int, used []bool)
	dfs = func(choice []int, used []bool) {
		if len(choice) == n {
			key := fmt.Sprint(choice) // 仍然使用字符串，但减少 `fmt.Sprintf`
			if !visited[key] {
				visited[key] = true
				ans = append(ans, append([]int(nil), choice...))
			}
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			dfs(append(choice, nums[i]), used)
			used[i] = false
		}
	}
	dfs([]int{}, make([]bool, n))
	return ans
}

func permuteUniqueChatGpt2(nums []int) [][]int {
	sort.Ints(nums) // 先排序以方便剪枝
	n := len(nums)
	ans := [][]int{}
	used := make([]bool, n)

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for i := 0; i < n; i++ {
			// 剪枝，避免重复元素的重复使用
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}
			used[i] = true
			dfs(append(path, nums[i]))
			used[i] = false
		}
	}
	dfs([]int{})
	return ans
}

func TestArray(t *testing.T) {
	tA := []int{1, 2, 3, 4, 5, 6, 7}
	t.Log(tA[:0])
	t.Log(tA[:1])
	t.Log(append(tA[:1], tA[2:]...))
}

func permuteUniqueLeetCode1(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/permutations-ii/solutions/417937/quan-pai-lie-ii-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func TestDeepCopy(t *testing.T) {
	arr1 := []int{1, 1, 2, 2, 2, 3}

	t.Log(arr1)

	arr2 := deepCopy(arr1)
	arr3 := deepCopy(append(arr1[:2], arr1[3:]...))

	t.Log(arr2)
	arr1[1] = 3
	t.Log(arr1)
	t.Log(arr2)

	arr1[0] = 3
	t.Log(arr1)
	t.Log(arr3)

}

func TestDeepCopy2(t *testing.T) {
	arr1 := []int{1, 1, 2, 2, 2, 3}
	arr2 := append(arr1[:2], arr1[3:]...)

	t.Logf("arr1 %p  arr2 %p \n", arr1, arr2)

	arr3 := []int{1, 2, 3, 4, 5, 6}

	// 这里会更改arr3 的内容 1, 2, 3, 4, 5, 6  ==> 1 2 4 5 6 6
	arr4 := append(arr3[:2], arr3[3:]...)

	// leetcode_test.go:253: arr3 0xc0000ba0f0  arr4 0xc0000ba0f0
	// leetcode_test.go:254: arr3 [1 2 4 5 6 6]  arr4 [1 2 4 5 6]
	t.Logf("arr3 %p  arr4 %p \n", arr3, arr4)
	t.Logf("arr3 %v  arr4 %v \n", arr3, arr4)

	_ = append(arr3[:1], arr3[3:]...)
	t.Logf("arr3 %v   \n", arr3)
}
