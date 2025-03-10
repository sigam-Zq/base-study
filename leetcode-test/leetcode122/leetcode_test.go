package leetcode122

import (
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {
	divisorSubstrings(30003, 3)
}

func divisorSubstrings(num int, k int) int {
	numStr := strconv.Itoa(num)

	cnt := 0
	for i := 0; i < len(numStr)-k; i++ {
		numSig, _ := strconv.Atoi(numStr[i : i+k])
		if numSig != 0 && numSig%k == 0 {
			cnt++
		}
	}
	return cnt
}

func divisorSubstringsLeetcode(num int, k int) int {
	s := strconv.Itoa(num) // num 十进制表示字符串
	n := len(s)
	res := 0
	for i := 0; i <= n-k; i++ {
		// 枚举所有长度为 k 的子串
		tmp, _ := strconv.Atoi(s[i : i+k])
		if tmp != 0 && num%tmp == 0 {
			res++
		}
	}
	return res
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/find-the-k-beauty-of-a-number/solutions/1538673/zhao-dao-yi-ge-shu-zi-de-k-mei-li-zhi-by-jn5i/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
