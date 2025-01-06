package leetcode79

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

// 2241. 设计一个 ATM 机器

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
	twenty, fifty, oneHundred, twoHundred, fiveHundred int
}

func Constructor() ATM {
	return ATM{}
}

func (this *ATM) Deposit(banknotesCount []int) {
	this.twenty += banknotesCount[0]
	this.fifty += banknotesCount[1]
	this.oneHundred += banknotesCount[2]
	this.twoHundred += banknotesCount[3]
	this.fiveHundred += banknotesCount[4]
}

func (this *ATM) Withdraw(amount int) []int {
	total := 20*this.twenty + 50*this.fifty + 100*this.oneHundred + 200*this.twoHundred + 500*this.fiveHundred
	// 存的钱不够取的
	if total < amount {
		return []int{-1}
	}
	var fiveHundredPay, twoHundredPay, oneHundredPay, fiftyPay, twentyPay int
	if amount >= 500 && this.fiveHundred != 0 {
		fiveHundredPay = min(amount/500, this.fiveHundred)
		amount -= fiveHundredPay * 500
	}
	if amount >= 200 && this.twoHundred != 0 {
		twoHundredPay = min(amount/200, this.twoHundred)
		amount -= twoHundredPay * 200
	}
	if amount >= 100 && this.oneHundred != 0 {
		oneHundredPay = min(amount/100, this.oneHundred)
		amount -= oneHundredPay * 100
	}
	if amount >= 50 && this.fifty != 0 {
		fiftyPay = min(amount/50, this.fifty)
		amount -= fiftyPay * 50
	}
	if amount >= 20 && this.twenty != 0 {
		twentyPay = min(amount/20, this.fifty)
		amount -= twentyPay * 20
	}
	if amount != 0 {
		return []int{-1}
	}

	this.fiveHundred -= fiveHundredPay
	this.twoHundred -= twoHundredPay
	this.oneHundred -= oneHundredPay
	this.fifty -= fiftyPay
	this.twenty -= twentyPay

	return []int{twentyPay, fiftyPay, oneHundredPay, twoHundredPay, fiveHundredPay}
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
