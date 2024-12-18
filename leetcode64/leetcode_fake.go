package leetcode64

import "log"

func minValidStringsFake(words []string, target string) int {

	prefixFuncSearch := func(word, tar string) []int {
		s := word + "#" + tar
		n := len(s)
		pi := make([]int, n)
		for i := 1; i < n; i++ {
			j := pi[i-1]
			for j > 0 && s[i] != s[j] {
				j = pi[j-1]
			}

			if s[i] == s[j] {
				j++
			}
			pi[i] = j
		}
		return pi
	}
	n := len(target)

	back := make([]int, n)
	for _, word := range words {
		pi := prefixFuncSearch(word, target)
		m := len(word)
		for i := 0; i < n; i++ {
			back[i] = max(back[i], pi[m+1+i])
		}
	}
	log.Println("back")
	log.Println(back)

	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = int(1e9)
	}

	for i := 0; i < n; i++ {
		dp[i+1] = dp[i+1-back[i]] + 1
		if dp[i+1] > n {
			return -1
		}
	}
	log.Println(dp)
	return dp[n]
}
