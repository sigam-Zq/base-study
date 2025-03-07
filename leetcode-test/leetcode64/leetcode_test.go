package leetcode64

import (
	"log"
	"math"
	"strconv"
	"strings"
	"testing"
)

/*
3291. 形成目标字符串需要的最少字符串数 I
给你一个字符串数组 words 和一个字符串 target。

如果字符串 x 是 words 中 任意 字符串的
前缀
，则认为 x 是一个 有效 字符串。

现计划通过 连接 有效字符串形成 target ，请你计算并返回需要连接的 最少 字符串数量。如果无法通过这种方式形成 target，则返回 -1。



示例 1：

输入： words = ["abc","aaaaa","bcdef"], target = "aabcdabc"

输出： 3

解释：

target 字符串可以通过连接以下有效字符串形成：

words[1] 的长度为 2 的前缀，即 "aa"。
words[2] 的长度为 3 的前缀，即 "bcd"。
words[0] 的长度为 3 的前缀，即 "abc"。
示例 2：

输入： words = ["abababab","ab"], target = "ababaababa"

输出： 2

解释：

target 字符串可以通过连接以下有效字符串形成：

words[0] 的长度为 5 的前缀，即 "ababa"。
words[0] 的长度为 5 的前缀，即 "ababa"。
示例 3：

输入： words = ["abcdef"], target = "xyz"

输出： -1



提示：

1 <= words.length <= 100
1 <= words[i].length <= 5 * 103
输入确保 sum(words[i].length) <= 105。
words[i] 只包含小写英文字母。
1 <= target.length <= 5 * 103
target 只包含小写英文字母。

*/

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
			f:      minValidStringsFake,
			words:  []string{"b", "ccacc", "a"},
			target: "cccaaaacba",
			want:   8,
			// cc c a a a a c b a  9
			// c cca a a a c b a	8
		},
		{
			f:      minValidStringsFake,
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

	var dfs func(string) int
	// isFirst := true
	dfs = func(tar string) int {
		n := len(tar)
		if n == 0 {
			return 0
		}
		maxL := 0
		isLose := true
		for _, word := range words {
			for j := 1; j <= n; j++ {
				if strings.HasPrefix(word, tar[:j]) {
					isLose = false
					maxL = max(maxL, j)
				} else {
					break
				}
			}
		}

		if isLose {
			return -1
		}
		// isFirst = false
		res := dfs(tar[maxL:])
		if res == -1 {
			return res
		} else {
			return 1 + res
		}
	}

	return dfs(target)
}

func TestOther(t *testing.T) {
	t.Log(strings.HasPrefix("abaaa", "abc"))
}

func TestCut(t *testing.T) {
	t.Log("abcdefg"[0:])
	t.Log("abcdefg"[1:])
	t.Log("abcdefg"[2:])
	t.Log("abcdefg"[3:])
}

func minConcatenation(words []string, target string) int {
	// 获取 target 的长度
	n := len(target)
	// 初始化 dp 数组，dp[i] 表示形成 target[:i] 所需的最少字符串数量
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32 // 初始状态为无穷大
	}
	dp[0] = 0 // 空字符串不需要连接任何字符串

	// 遍历 target 的每个前缀位置
	for i := 1; i <= n; i++ {
		for _, word := range words { // 遍历 words 中的每个字符串
			length := len(word)
			// 检查当前 word 是否可以作为 target[i-length:i] 的前缀
			if i >= length && target[i-length:i] == word[:length] {
				dp[i] = min(dp[i], dp[i-length]+1)
			}
		}
	}

	// 如果 dp[n] 仍然是无穷大，说明无法形成 target
	if dp[n] == math.MaxInt32 {
		return -1
	}
	return dp[n]
}

func minValidStringsLeetCode(words []string, target string) int {
	prefixFunction := func(word, target string) []int {
		s := word + "#" + target
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

		pi := prefixFunction(word, target)
		log.Println(word, '-', target)
		log.Println(pi)
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
	log.Println("dp")
	log.Println(dp)
	return dp[n]
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/minimum-number-of-valid-strings-to-form-target-i/solutions/3014795/xing-cheng-mu-biao-zi-fu-chuan-xu-yao-de-3cop/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/*


我们来详细解释这道题目以及其题解的关键部分。题目目标是通过连接**有效字符串**来拼接出 `target`，并求出所需的最少字符串数量。如果无法拼接出 `target`，则返回 `-1`。

题解使用了字符串匹配的技巧（类似于 KMP 算法），动态规划，以及前缀处理。下面分步骤进行说明：

---

### 1. `prefixFunction` 函数的作用和原理

#### 作用：
`prefixFunction` 是计算字符串前缀函数（Prefix Function）的方法，类似于 KMP 字符串匹配算法中的部分匹配表（Partial Match Table）。
它的作用是：
- 计算字符串 `word` 在拼接后的字符串 `word + "#" + target` 中与 `target` 的**前缀匹配**情况。
- 通过前缀函数，我们可以知道 `target` 的前缀与 `word` 的后缀可以匹配多长。

#### 原理：
前缀函数 `pi[i]` 表示：在字符串 `s` 的前缀中，`s[0...i]` 的**最长**相等的前后缀的长度（前后缀不包括整个字符串本身）。

例如，对于字符串 `s = "ababc"`, 其前缀函数 `pi` 的计算如下：
- `pi[0] = 0` （第一个字符没有前后缀）
- `pi[1] = 0` （"ab" 的前后缀最大相等长度是 0）
- `pi[2] = 1` （"aba" 的最长前后缀是 "a"）
- `pi[3] = 2` （"abab" 的最长前后缀是 "ab"）
- `pi[4] = 0` （"ababc" 的前后缀没有匹配）

`prefixFunction` 的关键步骤：
1. 将 `word` 和 `target` 拼接成 `s = word + "#" + target`。
   - 这里 `#` 是分隔符，确保 `word` 和 `target` 之间不会相互干扰。
2. 遍历拼接后的字符串，计算其前缀函数。
3. 通过前缀函数，得出 `target` 的每个位置上，与 `word` 可以匹配的最长前缀长度。

#### 代码示例：
```go
prefixFunction := func(word, target string) []int {
    s := word + "#" + target
    n := len(s)
    pi := make([]int, n)
    for i := 1; i < n; i++ {
        j := pi[i - 1]
        for j > 0 && s[i] != s[j] {
            j = pi[j - 1]
        }
        if s[i] == s[j] {
            j++
        }
        pi[i] = j
    }
    return pi
}
```

问

某题目为

给你一个字符串数组 words 和一个字符串 target。

如果字符串 x 是 words 中 任意 字符串的
前缀
，则认为 x 是一个 有效 字符串。

现计划通过 连接 有效字符串形成 target ，请你计算并返回需要连接的 最少 字符串数量。如果无法通过这种方式形成 target，则返回 -1。


题解为
func minValidStrings(words []string, target string) int {
    prefixFunction := func(word, target string) []int {
        s := word + "#" + target
        n := len(s)
        pi := make([]int, n)
        for i := 1; i < n; i++ {
            j := pi[i - 1]
            for j > 0 && s[i] != s[j] {
                j = pi[j - 1]
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
        pi := prefixFunction(word, target)
        m := len(word)
        for i := 0; i < n; i++ {
            back[i] = int(math.Max(float64(back[i]), float64(pi[m + 1 + i])))
        }
    }

    dp := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        dp[i] = int(1e9)
    }
    for i := 0; i < n; i++ {
        dp[i + 1] = dp[i + 1 - back[i]] + 1
        if dp[i + 1] > n {
            return -1
        }
    }
    return dp[n]
}


请详细解释这个题解是怎么解决这个问题的

尤其是以下几个地方

1 prefixFunction  函数的作用和原理

2 back数组这里起到什么作用

3 关于状态转移方程   dp[i + 1] = dp[i + 1 - back[i]] + 1是怎么推理出来的


答


---

### 2. `back` 数组的作用

#### `back` 数组的含义：
`back[i]` 表示 `target` 中以位置 `i` 结尾的最长有效前缀的长度。

有效前缀是指：它是 `words` 中某个字符串的前缀。

#### 如何计算：
- 对于 `target` 的每个位置 `i`，我们遍历 `words` 中的所有字符串。
- 使用 `prefixFunction` 计算当前 `word` 与 `target` 的匹配长度。
- 更新 `back[i]`，使其等于当前已知的最长有效前缀长度。

#### 代码示例：
```go
for _, word := range words {
    pi := prefixFunction(word, target)
    m := len(word)
    for i := 0; i < n; i++ {
        back[i] = int(math.Max(float64(back[i]), float64(pi[m + 1 + i])))
    }
}
```

其中：
- `pi[m + 1 + i]` 是 `prefixFunction` 结果中对应 `target` 位置 `i` 处的匹配长度。
- `back[i]` 取当前最大值，确保我们得到最长的有效前缀。

**总结**：`back` 数组告诉我们，从 `target` 的某个位置往回看，能够匹配的最长有效字符串前缀的长度是多少。

---

### 3. 状态转移方程的推理

#### 状态定义：
`dp[i]` 表示拼接出 `target[0...i-1]`（即 `target` 的前 `i` 个字符）所需的最少字符串数量。

#### 初始化：
- `dp[0] = 0`，表示空字符串不需要拼接任何字符串。

#### 状态转移方程：
对于位置 `i`，我们知道 `back[i-1]` 是以 `target[i-1]` 结尾的最长有效前缀的长度。
这意味着我们可以用长度为 `back[i-1]` 的字符串拼接出 `target` 的后缀。

因此：
- 要拼接 `target` 的前 `i` 个字符，可以从位置 `i - back[i-1]` 转移过来。
- 在 `dp` 中，这相当于：
  \[
  dp[i] = \min(dp[i], dp[i - back[i-1]] + 1)
  \]
  其中 `+1` 表示我们拼接了一个有效字符串。

#### 转移方程代码：
```go
for i := 0; i < n; i++ {
    dp[i + 1] = dp[i + 1 - back[i]] + 1
    if dp[i + 1] > n {
        return -1
    }
}
```

#### 推理逻辑：
- `dp[i + 1 - back[i]]` 表示拼接出 `target[0...(i - back[i] - 1)]` 所需的最少字符串数量。
- `+1` 表示我们用一个有效字符串（长度为 `back[i]`）拼接出了 `target[i - back[i] ... i]`。
- 因此，`dp[i + 1]` 更新为 `dp[i + 1 - back[i]] + 1`。

#### 为什么需要判断 `dp[i + 1] > n`？
- 如果 `dp[i + 1]` 大于 `n`，说明无法拼接出 `target`，返回 `-1`。

---

### 综上总结

1. **`prefixFunction`**：计算 `word + "#" + target` 的前缀函数，找出 `target` 中的最长有效前缀。
2. **`back` 数组**：记录 `target` 每个位置结尾的最长有效前缀长度。
3. **动态规划**：`dp[i + 1] = dp[i + 1 - back[i]] + 1` 表示将当前有效前缀拼接到前一个位置，更新所需字符串数量。

通过以上步骤，我们可以高效求出拼接 `target` 所需的最少字符串数量。如果无法拼接，则返回 `-1`。
*/
