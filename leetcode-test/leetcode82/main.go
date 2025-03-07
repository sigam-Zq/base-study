package main

import "log"

// 2264. 字符串中最大的 3 位相同数字
/*
给你一个字符串 num ，表示一个大整数。如果一个整数满足下述所有条件，则认为该整数是一个 优质整数 ：

该整数是 num 的一个长度为 3 的 子字符串 。
该整数由唯一一个数字重复 3 次组成。
以字符串形式返回 最大的优质整数 。如果不存在满足要求的整数，则返回一个空字符串 "" 。

注意：

子字符串 是字符串中的一个连续字符序列。
num 或优质整数中可能存在 前导零 。


示例 1：

输入：num = "6777133339"
输出："777"
解释：num 中存在两个优质整数："777" 和 "333" 。
"777" 是最大的那个，所以返回 "777" 。
示例 2：

输入：num = "2300019"
输出："000"
解释："000" 是唯一一个优质整数。
示例 3：

输入：num = "42352338"
输出：""
解释：不存在长度为 3 且仅由一个唯一数字组成的整数。因此，不存在优质整数。


提示：

3 <= num.length <= 1000
num 仅由数字（0 - 9）组成
*/

func main() {
	log.Println(largestGoodInteger("6777133339"))
}

func largestGoodInteger(num string) string {

	cnt := 1
	nowChar := num[0]
	resCharIdxList := []int{}
	for i := 1; i < len(num); i++ {
		if num[i] == nowChar {
			cnt++
		} else {
			nowChar = num[i]
			cnt = 1
		}

		if cnt == 3 {
			resCharIdxList = append(resCharIdxList, i-2)
		}
	}
	log.Println(resCharIdxList)

	switch len(resCharIdxList) {
	case 0:
		return ""
	case 1:
		return num[resCharIdxList[0] : resCharIdxList[0]+3]
	default:
		// 大于 两条 需要选择 最大的那个
		maxIdx := 0

		for i := 1; i < len(resCharIdxList); i++ {
			if num[resCharIdxList[maxIdx]] < num[resCharIdxList[i]] {
				maxIdx = i
			}
		}
		return num[resCharIdxList[maxIdx] : resCharIdxList[maxIdx]+3]
	}
}

func largestGoodIntegerLeetCode(num string) string {
	n := len(num)
	var res string
	for i := 0; i < n-2; i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			res = max(res, num[i:i+3])
		}
	}
	return res
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/largest-3-same-digit-number-in-string/solutions/1538493/zi-fu-chuan-zhong-zui-da-de-3-wei-xiang-isykz/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
