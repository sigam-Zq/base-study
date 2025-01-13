package leetcode85

import (
	"log"
	"reflect"
	"strconv"
	"testing"
)

// 2275. 按位与结果大于零的最长组合
/*
对数组 nums 执行 按位与 相当于对数组 nums 中的所有整数执行 按位与 。

例如，对 nums = [1, 5, 3] 来说，按位与等于 1 & 5 & 3 = 1 。
同样，对 nums = [7] 而言，按位与等于 7 。
给你一个正整数数组 candidates 。计算 candidates 中的数字每种组合下 按位与 的结果。

返回按位与结果大于 0 的 最长 组合的长度。

示例 1：

输入：candidates = [16,17,71,62,12,24,14]
输出：4
解释：组合 [16,17,62,24] 的按位与结果是 16 & 17 & 62 & 24 = 16 > 0 。
组合长度是 4 。
可以证明不存在按位与结果大于 0 且长度大于 4 的组合。
注意，符合长度最大的组合可能不止一种。
例如，组合 [62,12,24,14] 的按位与结果是 62 & 12 & 24 & 14 = 8 > 0 。
示例 2：

输入：candidates = [8,8]
输出：2
解释：最长组合是 [8,8] ，按位与结果 8 & 8 = 8 > 0 。
组合长度是 2 ，所以返回 2 。


提示：

1 <= candidates.length <= 105
1 <= candidates[i] <= 107

*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f          func([]int) int
		candidates []int
		want       int
		isDebug    bool
	}{
		{
			f:          largestCombination,
			candidates: []int{16, 17, 71, 62, 12, 24, 14},
			want:       4,
			isDebug:    false,
		},
		{
			f:          largestCombination,
			candidates: []int{8, 8},
			want:       2,
			isDebug:    false,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans := v.f(v.candidates)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

func largestCombination(candidates []int) int {
	numMax := 0
	for _, v := range candidates {
		numMax = max(numMax, v)
	}
	// 找到最存储的最大位数
	MaxLen := 1
	for 1<<MaxLen <= numMax {
		MaxLen++
	}

	listInts := make([]int, MaxLen)
	for _, v := range candidates {
		bitL := 0
		for (1 << bitL) <= v {
			if (v>>bitL)&1 == 1 {
				listInts[bitL]++
			}
			bitL++
		}
	}
	log.Println(listInts)
	ans := 0

	for _, v := range listInts {
		ans = max(ans, v)
	}

	return ans
}

func largestCombinationLeetCode(candidates []int) int {
	// 计算从低到高第 k 个二进制位数值为 1 的元素个数
	maxlen := func(k int) int {
		res := 0
		for _, num := range candidates {
			if num&(1<<k) != 0 {
				res++
			}
		}
		return res
	}

	res := 0
	for i := 0; i < 24; i++ {
		// 遍历二进制位
		res = max(res, maxlen(i))
	}
	return res
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/largest-combination-with-bitwise-and-greater-than-zero/solutions/1538671/an-wei-yu-jie-guo-da-yu-ling-de-zui-chan-hm7c/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
