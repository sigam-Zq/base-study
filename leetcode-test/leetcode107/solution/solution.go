package solution

import "sort"

func subsetsWithDupLeetCode1(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
outer:
	for mask := 0; mask < 1<<n; mask++ {
		t := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
					continue outer
				}
				t = append(t, v)
			}
		}
		ans = append(ans, append([]int(nil), t...))
	}
	return
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/subsets-ii/solutions/690549/zi-ji-ii-by-leetcode-solution-7inq/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	t := []int{}
	var dfs func(bool, int)
	dfs = func(choosePre bool, cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		dfs(false, cur+1)
		if !choosePre && cur > 0 && nums[cur-1] == nums[cur] {
			return
		}
		t = append(t, nums[cur])
		dfs(true, cur+1)
		t = t[:len(t)-1]
	}
	dfs(false, 0)
	return
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/subsets-ii/solutions/690549/zi-ji-ii-by-leetcode-solution-7inq/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
