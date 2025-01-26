package leetcode96

import (
	"log"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

// 2412. 完成所有交易的初始最少钱数

/*
给你一个下标从 0 开始的二维整数数组 transactions，其中transactions[i] = [costi, cashbacki] 。

数组描述了若干笔交易。其中每笔交易必须以 某种顺序 恰好完成一次。在任意一个时刻，你有一定数目的钱 money ，为了完成交易 i ，money >= costi 这个条件必须为真。执行交易后，你的钱数 money 变成 money - costi + cashbacki 。

请你返回 任意一种 交易顺序下，你都能完成所有交易的最少钱数 money 是多少。



示例 1：

输入：transactions = [[2,1],[5,0],[4,2]]
输出：10
解释：
刚开始 money = 10 ，交易可以以任意顺序进行。
可以证明如果 money < 10 ，那么某些交易无法进行。
示例 2：

输入：transactions = [[3,0],[0,3]]
输出：3
解释：
- 如果交易执行的顺序是 [[3,0],[0,3]] ，完成所有交易需要的最少钱数是 3 。
- 如果交易执行的顺序是 [[0,3],[3,0]] ，完成所有交易需要的最少钱数是 0 。
所以，刚开始钱数为 3 ，任意顺序下交易都可以全部完成。


提示：

1 <= transactions.length <= 105
transactions[i].length == 2
0 <= costi, cashbacki <= 109
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f            func([][]int) int64
		transactions [][]int
		want         int64
		isDebug      bool
	}{
		{
			f:            minimumMoney,
			transactions: [][]int{{2, 1}, {5, 0}, {4, 2}},
			want:         10,
			isDebug:      false,
		},
		{
			f:            minimumMoney,
			transactions: [][]int{{3, 0}, {0, 3}},
			want:         3,
			isDebug:      false,
		},
		{
			f:            minimumMoney,
			transactions: [][]int{{7, 2}, {0, 10}, {5, 0}, {4, 1}, {5, 8}, {5, 9}},
			want:         18,
			isDebug:      false,
		},
		{
			f:            minimumMoney,
			transactions: [][]int{{6, 5}, {0, 5}, {8, 5}, {3, 6}, {9, 0}, {10, 1}, {4, 10}},
			want:         27,
			isDebug:      false,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans := v.f(v.transactions)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

func minimumMoney(transactions [][]int) int64 {

	sort.Slice(transactions, func(i, j int) bool {
		// 如果 i[0] - i[1] >= 0 说明这个交易亏钱 (先做这个盈利)
		if transactions[i][0] <= transactions[i][1] {
			// 并且 j[0] - j[1] >= 0 说明另一个交易亏
			if transactions[j][0] <= transactions[j][1] {
				// 根据成本价排序
				return transactions[i][0] >= transactions[j][0]
			} else {
				// 另一个不盈利  i 放前面
				return false
			}
		} else {
			// 等价 if transactions[i][0] < transactions[i][1] {
			// j 和i 同样 亏钱 返现比成本多
			if transactions[j][0] > transactions[j][1] {
				// 根据返现价排序
				return transactions[i][1] <= transactions[j][1]
			} else {
				// 另一个不盈利  j 放前面
				return true
			}
		}

	})

	log.Println(transactions)

	var ans, res int64
	n := len(transactions)
	for i := 0; i < n; i++ {
		ans += int64(transactions[i][0])
		res = max(res, ans)
		if transactions[i][0] < transactions[i][1] {
			break
		}
		if i < n-1 {
			ans -= int64(transactions[i][1])
		}

	}

	return res
}

func minimumMoneyLeecode(transactions [][]int) int64 {
	var totalLose int64 = 0
	var res int = 0
	for _, t := range transactions {
		cost, cashback := t[0], t[1]
		totalLose += int64(max(cost-cashback, 0))
		res = max(res, min(cost, cashback))
	}
	return totalLose + int64(res)
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/minimum-money-required-before-transactions/solutions/3047584/wan-cheng-suo-you-jiao-yi-de-chu-shi-zui-cde1/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
