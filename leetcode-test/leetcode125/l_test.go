package leetcode125

import (
	"math"
	"strconv"
	"testing"
)

/*
1461. 检查一个字符串是否包含所有长度为 K 的二进制子串
https://leetcode.cn/problems/check-if-a-string-contains-all-binary-codes-of-size-k/
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func(string, int) bool
		s    string
		k    int
		want bool
	}{
		{
			f:    hasAllCodes,
			s:    "00110110",
			k:    2,
			want: true,
		},
		{
			f:    hasAllCodes,
			s:    "0110",
			k:    1,
			want: true,
		},
		{
			f:    hasAllCodes,
			s:    "0110",
			k:    2,
			want: false,
		},
		{
			f:    hasAllCodes,
			s:    "00110",
			k:    2,
			want: true,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.s, v.k)
			if got != v.want {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func hasAllCodes(s string, k int) bool {
	//  应该存在的组合次数
	targetHava := int(math.Pow(2, float64(k)))
	record := make(map[string]struct{})

	for i := 0; i <= len(s)-k; i++ {
		if _, ok := record[s[i:i+k]]; !ok {
			record[s[i:i+k]] = struct{}{}
			if len(record) == targetHava {
				return true
			}
		}
	}

	return false
}
func hasAllCodesLeetCode(s string, k int) bool {
	if len(s) < (1<<k)+k-1 {
		return false
	}

	exists := make(map[string]bool)
	for i := 0; i+k <= len(s); i++ {
		substring := s[i : i+k]
		exists[substring] = true
	}
	return len(exists) == (1 << k)
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/check-if-a-string-contains-all-binary-codes-of-size-k/solutions/519007/jian-cha-yi-ge-zi-fu-chuan-shi-fou-bao-h-1no1/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
