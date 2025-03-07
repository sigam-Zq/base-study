package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

/*
3138. 同位字符串连接的最小长度

给你一个字符串 s ，它由某个字符串 t 和若干 t  的 同位字符串 连接而成。

请你返回字符串 t 的 最小 可能长度。

同位字符串 指的是重新排列一个单词得到的另外一个字符串，原来字符串中的每个字符在新字符串中都恰好只使用一次。


示例 1：

输入：s = "abba"

输出：2

解释：

一个可能的字符串 t 为 "ba" 。

示例 2：

输入：s = "cdef"

输出：4

解释：

一个可能的字符串 t 为 "cdef" ，注意 t 可能等于 s 。



提示：

1 <= s.length <= 105
s 只包含小写英文字母。

*/

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f    func(string) int
		s    string
		want int
	}{
		{
			f:    minAnagramLength,
			s:    "abba",
			want: 2,
		},
		{
			f:    minAnagramLength,
			s:    "cdef",
			want: 4,
		},
		{
			f:    minAnagramLength,
			s:    "aabb",
			want: 4,
		},
		{
			f:    minAnagramLength,
			s:    "jjj",
			want: 1,
		},
		{
			f:    minAnagramLength,
			s:    "pqqppqpqpq",
			want: 2,
		},
	} {

		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.s); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})
	}
}

func minAnagramLength(s string) int {
	sLen := len(s)
	unitMap := make(map[byte]int)
	for i := 0; i < sLen; i++ {
		unitMap[s[i]]++
	}
	if len(unitMap) == 1 {
		return 1
	}
	mLen := len(unitMap)

	var res int

	for i := mLen; i < sLen; i++ {
		// i需要是 sLen的因数
		if sLen%i != 0 {
			continue
		}
		subStr := s[:i]
		pass := true
		for j := i; j < sLen; j += i {
			if !isParityStr(subStr, s[j:j+i]) {
				pass = false
				break
			}
		}
		if pass {
			res = i
			break
		}

	}
	if res == 0 {
		res = sLen
	}

	return res
}

func isParityStr(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	n := len(a)
	aMap := make(map[byte]int)
	bMap := make(map[byte]int)

	for i := 0; i < n; i++ {
		aMap[a[i]]++
		bMap[b[i]]++
	}

	return reflect.DeepEqual(aMap, bMap) // 比较两个map是否相等
}

func TestIsParityStr(t *testing.T) {
	t.Log(isParityStr("aab", "bab"))
}

/*
方法一：枚举
由题意可知，字符串 t 的长度一定为字符串 s 的长度 n 的因数，因此我们可以从小到大枚举 n 的因数作为 t 的长度。令当前枚举的因数为 i，我们将字符串 s 切分为若干个长度为 i 的子字符串，用 count
0
​
  统计前一子字符串的字符出现次数，用 count
1
​
  统计后一子字符串的出现次数，如果 count
0
​
  不等于 count
1
​
 ，那么说明 i 不符合题意；否则说明所有子字符串的字符出现次数都相等，那么返回 i 作为 t 的最小可能长度。

作者：力扣官方题解
链接：https://leetcode.cn/problems/minimum-length-of-anagram-concatenation/solutions/3014840/tong-wei-zi-fu-chuan-lian-jie-de-zui-xia-74z1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func minAnagramLengthLeetCode(s string) int {
	n := len(s)
	check := func(m int) bool {
		var count0 [26]int
		for j := 0; j < n; j += m {
			var count1 [26]int
			for k := j; k < j+m; k++ {
				count1[s[k]-'a']++
			}
			if j > 0 && count0 != count1 {
				return false
			}
			count0 = count1
		}
		return true
	}
	for i := 1; i < n; i++ {
		if n%i != 0 {
			continue
		}
		if check(i) {
			return i
		}
	}
	return n
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/minimum-length-of-anagram-concatenation/solutions/3014840/tong-wei-zi-fu-chuan-lian-jie-de-zui-xia-74z1/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func TestEqSlice(t *testing.T) {
	a := []int{1, 2, 3}
	// b := []int{1, 2, 3}
	t.Logf("type %T\n", a)
	t.Logf("type %s\n", reflect.TypeOf(a).Kind())
	//invalid operation: a == b (slice can only be compared to nil)
	// t.Log(a == b)
	// 数组对象可以 对比  但是 切片对象不行
	var a2 [26]int
	var b2 [26]int
	t.Logf("type %T\n", a2)
	t.Logf("type %s\n", reflect.TypeOf(a2).Kind())
	a2[1] = 2
	b2[1] = 1
	t.Log(a2 == b2)

}
