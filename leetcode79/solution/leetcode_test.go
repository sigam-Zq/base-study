package leetcode79

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f       func([]int, []int, int) []int
		act     []string
		params  [][]int
		want    [][]int
		isDebug bool
	}{
		{
			act:     []string{"ATM", "deposit", "withdraw", "deposit", "withdraw", "withdraw"},
			params:  [][]int{{}, {0, 0, 1, 2, 1}, {600}, {0, 1, 0, 1, 1}, {600}, {550}},
			want:    [][]int{{}, {}, {0, 0, 1, 0, 1}, {}, {-1}, {0, 1, 0, 0, 1}},
			isDebug: false,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var atm ATM
			var ans [][]int
			for ii, vAct := range v.act {
				switch vAct {
				case "ATM":
					atm = Constructor()
					ans = append(ans, []int{})
				case "deposit":
					atm.Deposit(v.params[ii])
					ans = append(ans, []int{})
				case "withdraw":
					res := atm.Withdraw(v.params[ii][0])
					ans = append(ans, res)
					if v.isDebug {
						fmt.Printf("ii %d res %v \n", ii, res)
						fmt.Printf("v.params %v \n", v.params[ii])
						fmt.Println("===================")

						fmt.Printf("%+v\n", atm)
						// for _, v := range atm.List {
						// 	fmt.Printf("%+v\n", v)
						// }
						fmt.Println("===================")
					}
				default:
					panic("not support act")
				}

			}
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

type ATM struct {
	cnt   []int64 // 每张钞票剩余数量
	value []int64 // 每张钞票面额
}

func Constructor() ATM {
	return ATM{
		cnt: make([]int64, 5),
		value: []int64{
			20, 50, 100, 200, 500,
		},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < 5; i++ {
		this.cnt[i] += int64(banknotesCount[i])
	}
}

func (this *ATM) Withdraw(amount int) []int {
	res := make([]int, 5)
	// 模拟尝试取出钞票的过程
	for i := 4; i >= 0; i-- {
		res[i] = int(min(this.cnt[i], int64(amount)/this.value[i]))
		amount -= res[i] * int(this.value[i])
	}
	if amount > 0 {
		// 无法完成该操作
		return []int{-1}
	}
	// 可以完成该操作
	for i := 0; i < 5; i++ {
		this.cnt[i] -= int64(res[i])
	}
	return res
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/design-an-atm-machine/solutions/1485677/she-ji-yi-ge-atm-ji-qi-by-leetcode-solut-etxe/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
