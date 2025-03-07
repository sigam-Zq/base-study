package leetcode31

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"testing"
)

/*
3254. 长度为 K 的子数组的能量值
给你一个长度为 n 的整数数组 nums 和一个正整数 k 。

一个数组的 能量值 定义为：

如果 所有 元素都是依次 连续 且 上升 的，那么能量值为 最大 的元素。
否则为 -1 。
你需要求出 nums 中所有长度为 k 的
子数组
 的能量值。

请你返回一个长度为 n - k + 1 的整数数组 results ，其中 results[i] 是子数组 nums[i..(i + k - 1)] 的能量值。



示例 1：

输入：nums = [1,2,3,4,3,2,5], k = 3

输出：[3,4,-1,-1,-1]

解释：

nums 中总共有 5 个长度为 3 的子数组：

[1, 2, 3] 中最大元素为 3 。
[2, 3, 4] 中最大元素为 4 。
[3, 4, 3] 中元素 不是 连续的。
[4, 3, 2] 中元素 不是 上升的。
[3, 2, 5] 中元素 不是 连续的。
示例 2：

输入：nums = [2,2,2,2,2], k = 4

输出：[-1,-1]

示例 3：

输入：nums = [3,2,3,2,3,2], k = 2

输出：[-1,3,-1,3,-1]
*/

func TestXxx(t *testing.T) {

	for idx, v := range []struct {
		f    func([]int, int) []int
		nums []int
		k    int
		want []int
	}{
		{
			f:    resultsArray,
			nums: []int{1, 2, 3, 4, 3, 2, 5},
			k:    3,
			want: []int{3, 4, -1, -1, -1},
		},
		{
			f:    resultsArray,
			nums: []int{2, 2, 2, 2, 2},
			k:    4,
			want: []int{-1, -1},
		},
		{
			f:    resultsArray,
			nums: []int{1, 3, 4},
			k:    2,
			want: []int{3, 4},
		},
		{
			f:    resultsArrayOptimize,
			nums: []int{3, 2, 3, 2, 3, 2},
			k:    2,
			want: []int{-1, 3, -1, 3, -1},
		},
	} {
		t.Run(strconv.Itoa(idx)+"-test", func(t *testing.T) {
			if got := v.f(v.nums, v.k); !reflect.DeepEqual(v.want, got) {
				t.Errorf(" got %v  want %v \n", got, v.want)
			} else {
				log.Printf("succeed %v \n", got)
			}
		})
	}

}

func resultsArray(nums []int, k int) []int {

	i := 0
	result := make([]int, 0)
	for i <= len(nums)-k {
		subList := nums[i : i+k]

		isInc := true
		for ii, v := range subList[1:] {
			log.Printf("subList %v  in   for -- ii %d\n", subList, ii)
			if v-subList[ii] != 1 {
				isInc = false
			}
		}
		if isInc {
			result = append(result, subList[k-1])
		} else {
			result = append(result, -1)
		}

		i++
	}
	return result
}

func resultsArrayOptimize(nums []int, k int) []int {
	// 加入数组的连续递增缓存 true 表示满足条件
	isIncCache := make([]bool, len(nums))
	isIncCache[0] = true
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			isIncCache[i] = true
		} else {
			isIncCache[i] = false
		}
	}
	fmt.Printf("isIncCache %v \n", isIncCache)
	res := make([]int, 0)
	for i := 0; i <= len(nums)-k; i++ {

		// 当前如果最后一位 小于k 直接加 -1 跳转下一次
		if nums[i+k-1] < k {
			res = append(res, -1)
			continue
		}

		if i == 0 {
			// 第一次要全量判断
			if AndBool(isIncCache[i+1 : i+k]...) {
				res = append(res, nums[i+k-1])
			} else {
				res = append(res, -1)
			}
		} else {
			// 第二次往上判断前面的res 是否符合条件
			fmt.Printf("isIncCache[i+k-1] %v i %d k %d isIncCache[i : i+k] %v  \n", isIncCache[i+k-1], i, k, isIncCache[i:i+k])

			if isIncCache[i+k-1] && res[len(res)-1] != -1 {
				// 上次为递增 且 下一个位置也是递增 总体就是递增
				res = append(res, nums[i+k-1])
			} else if !isIncCache[i+k-1] {
				// 当前位数不满足 不用判断前面 也是不满足
				res = append(res, -1)
			} else {
				fmt.Printf("--全量判断--\n")
				fmt.Printf("--isIncCache[i : i+k] %v--\n", isIncCache[i:i+k])
				//全量判断
				if AndBool(isIncCache[i+1 : i+k]...) {
					res = append(res, nums[i+k-1])
				} else {
					res = append(res, -1)
				}
			}
		}
	}

	return res
}

func AndBool(flagValue ...bool) bool {
	var res bool
	res = flagValue[0]
	for _, v := range flagValue {
		res = res && v
	}

	return res
}
