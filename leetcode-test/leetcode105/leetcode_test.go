package leetcode105

import (
	"strconv"
	"testing"
)

// 680. 验证回文串 II

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func(string) bool
		s    string
		want bool
	}{
		{
			f:    validPalindrome,
			s:    "aba",
			want: true,
		},
		{
			f:    validPalindrome,
			s:    "abca",
			want: true,
		},
		{
			f:    validPalindrome,
			s:    "abc",
			want: false,
		},
		{
			f:    validPalindrome,
			s:    "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
			want: true,
		},
		{
			f:    validPalindrome,
			s:    "eeccccbebaeeabebccceea",
			want: false,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := v.f(v.s); got != v.want {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			// 尝试减去左边
			if s[l] == s[r-1] || s[l+1] == s[r] {
				s1 := s[:l] + s[l+1:]
				s2 := s[:r] + s[r+1:]
				// log.Println(l, r)
				// log.Println(s)
				// log.Println(s1)
				// log.Println(s2)
				return validPalindromeSub(s1) || validPalindromeSub(s2)
			} else {
				return false
			}
		}
		l++
		r--
	}

	return true
}

func validPalindromeSub(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func validPalindromeLeetCode1(s string) bool {
	low, high := 0, len(s)-1
	for low < high {
		if s[low] == s[high] {
			low++
			high--
		} else {
			flag1, flag2 := true, true
			for i, j := low, high-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := low+1, high; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/valid-palindrome-ii/solutions/251842/yan-zheng-hui-wen-zi-fu-chuan-ii-by-leetcode-solut/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
