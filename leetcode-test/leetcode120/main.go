package main

import "log"

func main() {
	log.Println(evenOddBit(50))
}

func evenOddBit(n int) []int {
	ans := make([]int, 2)
	even, odd, isEven := 0, 0, true
	for n >= 1 {
		if isEven {
			if n&1 == 1 {
				even++
			}
			// even |= n & 1
		} else {
			if n&1 == 1 {
				odd++
			}
			// odd |= n & 1
		}
		n >>= 1
		isEven = !isEven
	}
	ans[0] = even
	ans[1] = odd
	return ans
}

func evenOddBitLeetCode(n int) []int {
	res := []int{0, 0}
	i := 0
	for n > 0 {
		res[i] += n & 1
		n >>= 1
		i ^= 1
	}
	return res
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/number-of-even-and-odd-bits/solutions/3065366/qi-ou-wei-shu-by-leetcode-solution-v0cc/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
