package main

import (
	"log"
	"strings"
)

/*
3083. 字符串及其反转中是否存在同一子字符串

给你一个字符串 s ，请你判断字符串 s 是否存在一个长度为 2 的子字符串，在其反转后的字符串中也出现。

如果存在这样的子字符串，返回 true；如果不存在，返回 false 。

示例 1：

输入：s = "leetcode"

输出：true

解释：子字符串 "ee" 的长度为 2，它也出现在 reverse(s) == "edocteel" 中。

示例 2：

输入：s = "abcba"

输出：true

解释：所有长度为 2 的子字符串 "ab"、"bc"、"cb"、"ba" 也都出现在 reverse(s) == "abcba" 中。

示例 3：

输入：s = "abcd"

输出：false

解释：字符串 s 中不存在满足「在其反转后的字符串中也出现」且长度为 2 的子字符串。



提示：

1 <= s.length <= 100
字符串 s 仅由小写英文字母组成。
*/

func main() {

	log.Println(isSubstringPresent("abcd"))

}

func isSubstringPresent(s string) bool {
	n := len(s)
	var dict [][2]byte
	dict = make([][2]byte, n-1)

	// for i := 0; i < n-1; i++ {
	// 	dict[i][0] = 0x00
	// }

	for i := 1; i < n; i++ {
		dict[i-1][0], dict[i-1][1] = s[i-1], s[i]

		for j := 0; j < i; j++ {

			if dict[j][0] == s[i] && dict[j][1] == s[i-1] {
				return true
			}
		}
	}

	return false
}

// 思路与算法

// 遍历字符串中每个长度为 2 的子串，将其翻转后判断是否在原串中出现即可。
func isSubstringPresentLeetCode1(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		substr := string([]byte{s[i+1], s[i]})
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/existence-of-a-substring-in-a-string-and-its-reverse/solutions/3016932/zi-fu-chuan-ji-qi-fan-zhuan-zhong-shi-fo-ra8p/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/*
方法二：哈希表 + 位运算优化
思路与算法

我们可以用哈希表提前存储字符串中的每个长度为 2 的子串，这样在判断翻转后的字符串是否出现时就避免了花费 O(n) 的时间查找。

由于字符仅包含小写字母，该哈希表可以用一个整数类型的二维数组实现，
形如 hash[26][26]。如果要进一步优化，还可以考虑将第二维使用二进制表示，
例如 hash[2] 二进制形式中，如果从低到高第 1 位为 1，则表示子串 ‘‘cb"
出现在字符串中（2 表示字符 c，1 表示字符 b）。

作者：力扣官方题解
链接：https://leetcode.cn/problems/existence-of-a-substring-in-a-string-and-its-reverse/solutions/3016932/zi-fu-chuan-ji-qi-fan-zhuan-zhong-shi-fo-ra8p/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func isSubstringPresentLeetCode2(s string) bool {
	h := make([]int, 26)
	for i := 0; i+1 < len(s); i++ {
		x, y := s[i]-'a', s[i+1]-'a'
		h[x] |= (1 << y)
		if (h[y]>>x)&1 != 0 {
			return true
		}
	}
	return false
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/existence-of-a-substring-in-a-string-and-its-reverse/solutions/3016932/zi-fu-chuan-ji-qi-fan-zhuan-zhong-shi-fo-ra8p/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
