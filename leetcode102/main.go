package main

import (
	"log"
	"slices"
)

// 541. 反转字符串 II

func main() {
	log.Println(reverseStr("abcdefg", 2) == "bacdfeg")

	log.Println(reverseStr("abcd", 2) == "bacd")
}

func reverseStr(s string, k int) string {
	n := len(s)
	ans := ""
	tmp := make([]byte, 0)
	for i := 0; i < n; i++ {
		if i%(2*k) < k {
			tmp = append(tmp, s[i])
			if i > k && i%(2*k) == 0 {
				ans += string(s[i-k : i])
			}
		} else if i%(2*k) == k {
			slices.Reverse(tmp)
			ans += string(tmp)
			tmp = make([]byte, 0)
		}

		if i == (n - 1) {

			if i%(2*k) < k {
				slices.Reverse(tmp)
				ans += string(tmp)
				tmp = make([]byte, 0)
			} else {
				ans += string(s[i-i%k : n])
			}

		}

	}

	// log.Println(ans)
	return ans
}

func reverseStrLeetCode1(s string, k int) string {
	t := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		sub := t[i:min(i+k, len(s))]
		for j, n := 0, len(sub); j < n/2; j++ {
			sub[j], sub[n-1-j] = sub[n-1-j], sub[j]
		}
	}
	return string(t)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/reverse-string-ii/solutions/946553/fan-zhuan-zi-fu-chuan-ii-by-leetcode-sol-ua7s/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
