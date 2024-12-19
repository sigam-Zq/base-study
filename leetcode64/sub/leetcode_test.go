package test

import (
	"log"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f      func([]string, string) int
		words  []string
		target string
		want   int
	}{
		{
			f:      minValidStrings,
			words:  []string{"abc", "aaaaa", "bcdef"},
			target: "aabcdabc",
			want:   3,
		},
		{
			f:      minValidStrings,
			words:  []string{"abababab", "ab"},
			target: "ababaababa",
			want:   2,
		},
		{
			f:      minValidStrings,
			words:  []string{"abcdef"},
			target: "xyz",
			want:   -1,
		},
		{
			f:      minValidStrings,
			words:  []string{"aacbabbbabacacbbcbbb"},
			target: "a",
			want:   1,
		},
		{
			f:      minValidStrings,
			words:  []string{"b", "abacaacabbb"},
			target: "abcca",
			want:   -1,
		},
		{
			f:      minValidStrings,
			words:  []string{"b", "ccacc", "a"},
			target: "cccaaaacba",
			want:   8,
			// cc c a a a a c b a  9
			// c cca a a a c b a	8
		},
		{
			f:      minValidStrings,
			words:  []string{"aaaaabbaccbcbaaaacb"},
			target: "b",
			want:   -1,
		},
	} {

		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.words, v.target); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})
	}
}

func minValidStrings(words []string, target string) int {

	prefixSearchFunc := func(word, tar string) []int {
		s := word + "#" + tar
		n := len(s)
		pi := make([]int, n)
		for i := 1; i < n; i++ {
			j := pi[i-1]
			for j > 0 && s[i] != s[j] {
				// j--
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
		pi := prefixSearchFunc(word, target)
		m := len(word)
		log.Println("dp", word, target)
		log.Println(pi)
		for i := 0; i < n; i++ {
			back[i] = max(back[i], pi[i+1+m])
		}
	}
	log.Println("back")
	log.Println(back)

	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = n + 1
	}

	for i := 0; i < n; i++ {
		dp[i+1] = dp[i+1-back[i]] + 1
		if dp[i+1] > n {
			return -1
		}
	}
	log.Println("dp")
	log.Println(dp)

	return dp[n]
}
