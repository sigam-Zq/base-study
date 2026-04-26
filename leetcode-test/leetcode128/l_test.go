package leetcode125

import (
	"log"
	"math/big"
	"reflect"
	"strconv"
	"testing"
)

// 1404.将二进制表示减到1的步骤数
// https://leetcode.cn/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/
func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func(s string) int
		s    string
		want int
	}{
		{
			f:    numSteps,
			s:    "1101",
			want: 6,
		},
		{
			f:    numSteps,
			s:    "1111110011101010110011100100101110010100101110111010111110110010",
			want: 89,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.s)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func numSteps(s string) int {
	var cnt int
	mockNum := mockNum(s)
	//  mockNum > 1
	for mockNum.Cmp(big.NewInt(1)) > 0 {
		// 奇数
		if mockNum.Bit(0) == 1 {
			cnt += 2
			mockNum = mockNum.Add(mockNum, big.NewInt(1))
		} else {
			cnt++
		}
		mockNum = mockNum.Rsh(mockNum, uint(1))
	}

	return cnt
}
func mockNum(s string) *big.Int {

	number := big.NewInt(0)
	for i, v := range s {
		if v == '1' {
			// big.NewInt(math.Pow(float64(2), float64(len(s)-i-1))) -> new(big.Int).Lsh(big.NewInt(1), uint(len(s)-i-1))
			number = number.Add(number, new(big.Int).Lsh(big.NewInt(1), uint(len(s)-i-1)))
		}
	}

	return number
}
func Test11Xxx(t *testing.T) {

	log.Println(mockNum("100"))
	log.Println(mockNum("110"))
	log.Println(mockNum("111"))
	log.Println(mockNum("1000"))
	log.Println(mockNum("1111110011101010110011100100101110010100101110111010111110110010"))
}

func numStepsLeetCode(s string) int {
	steps := 0
	bytes := []byte(s)

	for string(bytes) != "1" {
		steps++
		if bytes[len(bytes)-1] == '0' {
			// 偶数的情况
			bytes = bytes[:len(bytes)-1]
		} else {
			// 第一步：找出最低位的 0
			// 第二步：把这个 0 变成 1，并将后面所有的 1 变成 0，这样就实现了 +1
			// 特别地，如果 s 中全是 1，那么会有额外的进位
			for i := len(bytes) - 1; i >= 0; i-- {
				if bytes[i] == '1' {
					bytes[i] = '0'
					if i == 0 {
						bytes = append([]byte{'1'}, bytes...)
						break
					}
				} else {
					bytes[i] = '1'
					break
				}
			}
		}
	}
	return steps
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/solutions/202251/jiang-er-jin-zhi-biao-shi-jian-dao-1-de-bu-zou-shu/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
