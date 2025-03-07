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

	log.Println(isSubstringPresentFake("abcd"))

}

func isSubstringPresentFake(s string) bool {

	leterMap := make([]int, 26)
	for i := 0; i+1 < len(s); i++ {
		x, y := s[i]-'a', s[i+1]-'a'
		leterMap[x] |= (1 << y)
		if (leterMap[y]>>x)&1 != 0 {
			return true
		}
	}
	return false
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

/*
问：
针对题目


给你一个字符串 s ，请你判断字符串 s 是否存在一个长度为 2 的子字符串，在其反转后的字符串中也出现。

如果存在这样的子字符串，返回 true；如果不存在，返回 false 。

存在一个位运算解法

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

存在以下问题

1.   h[x] |= (1 << y)
这里的位或运算 和  1左移 y位 写入 h[x]在下面要起到什么作用，这一步有什么作用

2  (h[y]>>x)
上面的 h[y]>>x 把 h[y]  随后右移 x 位实现了什么效果，起到了什么作用

这里右移之后通过&1 != 0判断 该 数的个位数是否是奇数这里起到了什么作用

3. for 循环中重复 对  h[x] |= (1 << y) 进行 位移或运算，这里会达到什么效果

4.  通过上面的位运算怎么实现 反转后的字符串中也出现 的判断的



copilot 回答

1.
h[x] |= (1 << y)
：   这一步的作用是将字符
s[i+1]
 的信息记录在字符
s[i]
 对应的位置上。具体来说，`1 << y` 会生成一个只有第 y 位为 1 的整数，
 然后通过 `|=` 操作将这个整数与  h[x] 进行位或运算，从而在 h[x] 中记录下字符
s[i+1]
 的信息。
2. `(h[y] >> x)`：
   这一步的作用是检查字符
s[i+1]
 是否已经记录了字符
s[i]
 的信息。具体来说，
h[y] >> x
 会将
h[y]
 右移
x
 位，从而将第
x
 位移到最低位。然后通过 `& 1` 操作检查最低位是否为 1，从而判断字符
s[i+1]
 是否已经记录了字符
s[i]
 的信息。
3. `for` 循环中重复对
h[x] |= (1 << y)
 进行位移或运算：   这一步的作用是不断更新 h
 数组，记录每个字符后面出现的字符信息。通过这种方式，
 可以在遍历字符串的过程中逐步构建字符之间的关系。
4. 通过上面的位运算实现反转后的字符串中也出现的判断：   通过
h[x] |= (1 << y)
 操作，记录了每个字符后面出现的字符信息。
 然后通过 `(h[y] >> x) & 1 != 0` 操作，检查字符 s[i+1]
 是否已经记录了字符 s[i]
 的信息。如果存在这样的记录，说明存在一个长度为 2 的子字符串，
 在其反转后的字符串中也出现。


当然可以。让我们通过一个具体的例子来解释这个算法的工作原理。

假设字符串

s

 为 `"abca"`。

1. 初始化一个长度为 26 的数组

h

，每个元素初始值为 0。
   ```go
   h := make([]int, 26)
   ```

2. 遍历字符串

s

，对于每一对相邻字符

s[i]

 和

s[i+1]

，进行以下操作：

   - 第一次迭代 (

i = 0

):
     -

x = s[0] - 'a' = 'a' - 'a' = 0


     -

y = s[1] - 'a' = 'b' - 'a' = 1


     -

h[0] |= (1 << 1)

，即

h[0] = 0 | (1 << 1) = 0 | 2 = 2


     -

h

 数组更新为 `[2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]`
     - `(h[1] >> 0) & 1 = (0 >> 0) & 1 = 0 & 1 = 0`，不满足条件，继续下一次迭代。

   - 第二次迭代 (

i = 1

):
     -

x = s[1] - 'a' = 'b' - 'a' = 1


     -

y = s[2] - 'a' = 'c' - 'a' = 2


     -

h[1] |= (1 << 2)

，即

h[1] = 0 | (1 << 2) = 0 | 4 = 4


     -

h

 数组更新为 `[2, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]`
     - `(h[2] >> 1) & 1 = (0 >> 1) & 1 = 0 & 1 = 0`，不满足条件，继续下一次迭代。

   - 第三次迭代 (

i = 2

):
     -

x = s[2] - 'a' = 'c' - 'a' = 2


     -

y = s[3] - 'a' = 'a' - 'a' = 0


     -

h[2] |= (1 << 0)

，即

h[2] = 0 | (1 << 0) = 0 | 1 = 1


     -

h

 数组更新为 `[2, 4, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]`
     - `(h[0] >> 2) & 1 = (2 >> 2) & 1 = 0 & 1 = 0`，不满足条件，继续下一次迭代。

3. 遍历结束，没有找到满足条件的子字符串，返回 `false`。

通过这个例子，我们可以看到，算法通过位运算记录了每个字符后面出现的字符信息，并检查是否存在反转后的子字符串。
*/
