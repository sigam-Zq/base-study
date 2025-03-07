package leetcode88

import (
	"reflect"
	"strconv"
	"testing"
)

// 2266. 统计打字方案数
/*
Alice 在给 Bob 用手机打字。数字到字母的 对应 如下图所示。



为了 打出 一个字母，Alice 需要 按 对应字母 i 次，i 是该字母在这个按键上所处的位置。

比方说，为了按出字母 's' ，Alice 需要按 '7' 四次。类似的， Alice 需要按 '5' 两次得到字母  'k' 。
注意，数字 '0' 和 '1' 不映射到任何字母，所以 Alice 不 使用它们。
但是，由于传输的错误，Bob 没有收到 Alice 打字的字母信息，反而收到了 按键的字符串信息 。

比方说，Alice 发出的信息为 "bob" ，Bob 将收到字符串 "2266622" 。
给你一个字符串 pressedKeys ，表示 Bob 收到的字符串，请你返回 Alice 总共可能发出多少种文字信息 。

由于答案可能很大，将它对 109 + 7 取余 后返回。



示例 1：

输入：pressedKeys = "22233"
输出：8
解释：
Alice 可能发出的文字信息包括：
"aaadd", "abdd", "badd", "cdd", "aaae", "abe", "bae" 和 "ce" 。
由于总共有 8 种可能的信息，所以我们返回 8 。
示例 2：

输入：pressedKeys = "222222222222222222222222222222222222"
输出：82876089
解释：
总共有 2082876103 种 Alice 可能发出的文字信息。
由于我们需要将答案对 109 + 7 取余，所以我们返回 2082876103 % (109 + 7) = 82876089 。


提示：

1 <= pressedKeys.length <= 105
pressedKeys 只包含数字 '2' 到 '9' 。
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f           func(string) int
		pressedKeys string
		want        int
		isDebug     bool
	}{
		{
			f:           countTexts,
			pressedKeys: "22233",
			want:        8,
			isDebug:     false,
		},
		{
			f:           countTexts,
			pressedKeys: "222222222222222222222222222222222222",
			want:        82876089,
			isDebug:     false,
		},
		{
			f:           countTexts,
			pressedKeys: "777777799995",
			want:        448,
			isDebug:     false,
		},
		{
			f:           countTexts,
			pressedKeys: "444444444444444444444444444444448888888888888888999999999999333333333333333366666666666666662222222222222222666666666666666633333333333333338888888888888888222222222222222244444444444444448888888888888222222222222222288888888888889999999999999999333333333444444664",
			want:        537551452,
			isDebug:     false,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans := v.f(v.pressedKeys)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

// 不同数字对应的变化
var NumberLimit = map[byte]int{
	'2': 1,
	'3': 1,
	'4': 1,
	'5': 1,
	'6': 1,
	'8': 1,
	'7': 2,
	'9': 2,
}

type piar struct {
	t int // 对应类型  1 三字母案件 2 四字母按键
	l int // 长度
}

const mod = int(1e9 + 7)

var calc3DP = []int{0, 1, 2, 4}
var calc4DP = []int{0, 1, 2, 4, 8}

func (p piar) ClacNum() int {

	if p.t == 1 {
		for len(calc3DP) <= p.l {
			calc3DP = append(calc3DP, (calc3DP[len(calc3DP)-1]+calc3DP[len(calc3DP)-2]+calc3DP[len(calc3DP)-3])%mod)
		}
		return calc3DP[p.l]
	} else {
		for len(calc4DP) <= p.l {
			calc4DP = append(calc4DP, (calc4DP[len(calc4DP)-1]+calc4DP[len(calc4DP)-2]+calc4DP[len(calc4DP)-3]+calc4DP[len(calc4DP)-4])%mod)
		}
		return calc4DP[p.l]
	}
}

func countTexts(pressedKeys string) int {

	mod := int(1e9 + 7)

	pressedList := make([]piar, 0)

	cur, curNum := pressedKeys[0], 1
	for i := 1; i < len(pressedKeys); i++ {

		if cur == pressedKeys[i] {
			curNum++
		} else if cur != pressedKeys[i] {

			pressedList = append(pressedList, piar{t: NumberLimit[cur], l: curNum})
			cur = pressedKeys[i]
			curNum = 1
		}
	}
	pressedList = append(pressedList, piar{t: NumberLimit[cur], l: curNum})

	// log.Printf("pressedList %+v \n", pressedList)
	ans := 1

	for _, v := range pressedList {
		ans = (ans * v.ClacNum() % mod)
	}

	// log.Printf("calc3DP %+v \n", calc3DP)
	// log.Printf("calc4DP %+v \n", calc4DP)

	return ans

}

func countTextsLeetCode(pressedKeys string) int {
	m := 1000000007
	n := len(pressedKeys)
	dp3 := []int{1, 1, 2, 4} // 连续按多次 3 个字母按键对应的方案数
	dp4 := []int{1, 1, 2, 4} // 连续按多次 4 个字母按键对应的方案数
	for i := 4; i <= n; i++ {
		dp3 = append(dp3, (dp3[i-1]+dp3[i-2]+dp3[i-3])%m)
		dp4 = append(dp4, (dp4[i-1]+dp4[i-2]+dp4[i-3]+dp4[i-4])%m)
	}
	res := 1 // 总方案数
	cnt := 1 // 当前字符连续出现的次数
	for i := 1; i < n; i++ {
		if pressedKeys[i] == pressedKeys[i-1] {
			cnt++
		} else {
			// 对按键对应字符数量讨论并更新总方案数
			if pressedKeys[i-1] == '7' || pressedKeys[i-1] == '9' {
				res = (res * dp4[cnt]) % m
			} else {
				res = (res * dp3[cnt]) % m
			}
			cnt = 1
		}
	}
	// 更新最后一段连续字符子串对应的方案数
	if pressedKeys[n-1] == '7' || pressedKeys[n-1] == '9' {
		res = (res * dp4[cnt]) % m
	} else {
		res = (res * dp3[cnt]) % m
	}
	return res
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/count-number-of-texts/solutions/1538660/tong-ji-da-zi-fang-an-shu-by-leetcode-so-714a/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
