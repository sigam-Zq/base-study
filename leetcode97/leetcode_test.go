package leetcode97

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

// 40. 组合总和 II

/*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。



示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]


提示:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f          func([]int, int) [][]int
		candidates []int
		target     int
		want       [][]int
		isDebug    bool
	}{
		{
			f:          combinationSum2Fake,
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			want:       [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
			isDebug:    false,
		},
		{
			f:          combinationSum2Fake,
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			want:       [][]int{{1, 2, 2}, {5}},
			isDebug:    false,
		},
		{
			// fatal error: out of memory allocating heap arena metadata
			f:          combinationSum2Fake,
			candidates: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			target:     30,
			want:       [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
			isDebug:    false,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans := v.f(v.candidates, v.target)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)

	sort.Ints(candidates)

	res := make([][]int, 0)
	UniMap := make(map[int]int)
	for _, v := range candidates {
		if v < target {
			UniMap[v]++
		} else if v == target && len(res) == 0 {
			// 把单个的拿出来
			res = append(res, []int{v})
		}
	}
	var dfs func(pos, tar int, ans []int)
	dfs = func(pos, tar int, ans []int) {
		if pos >= n || candidates[pos] > tar || tar < 0 {
			return
		} else if candidates[pos] == tar {
			// 结束
			ans = append(ans, candidates[pos])
			res = append(res, ans)
			return
		}
		ans = append(ans, candidates[pos])
		backLen := len(ans)
		for i := range candidates[pos:] {
			ans = ans[:backLen]
			dfs(pos+i+1, tar-candidates[pos], ans)
		}
	}
	// dfs(0, target, make([]int, 0))
	for i := range candidates {
		if candidates[i] < target {
			dfs(i, target, make([]int, 0))
		}
	}

	// 去重
	deWeightRes := make([][]int, 0)
	haveUni := make([]map[int]int, 0)
	for i, v := range res {
		uni := make(map[int]int)
		for _, vv := range v {
			uni[vv]++
		}

		exist := false
		for _, uMap := range haveUni {
			if reflect.DeepEqual(uni, uMap) {
				exist = true
			}
		}
		if !exist {
			haveUni = append(haveUni, uni)
			deWeightRes = append(deWeightRes, res[i])
		}
	}

	return deWeightRes
}

func combinationSum2Fake(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	// freq := make([][2]int, 0)
	var freq [][2]int

	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var dfs func(pos, rest int)
	var seq []int
	var ans [][]int
	dfs = func(pos, rest int) {
		if rest == 0 {
			// ans = append(ans, seq)
			ans = append(ans, append([]int(nil), seq...))
			return
		}

		if pos >= len(freq) || rest < freq[pos][0] {
			return
		}

		// 不选择的情况
		dfs(pos+1, rest)

		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			seq = append(seq, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		seq = seq[:len(seq)-most]
	}
	dfs(0, target)
	return ans
}

func combinationSum2LeetCode(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var freq [][2]int
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var sequence []int
	var dfs func(pos, rest int)
	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}

		dfs(pos+1, rest)

		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0, target)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/combination-sum-ii/solutions/407850/zu-he-zong-he-ii-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
