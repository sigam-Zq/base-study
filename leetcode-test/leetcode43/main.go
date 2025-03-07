package main

import (
	"log"
	"math"
	"time"
)

/*
3233. 统计不是特殊数字的数字数量
给你两个 正整数 l 和 r。对于任何数字 x，x 的所有正因数（除了 x 本身）被称为 x 的 真因数。

如果一个数字恰好仅有两个 真因数，则称该数字为 特殊数字。例如：

数字 4 是 特殊数字，因为它的真因数为 1 和 2。
数字 6 不是 特殊数字，因为它的真因数为 1、2 和 3。
返回区间 [l, r] 内 不是 特殊数字 的数字数量。



示例 1：

输入： l = 5, r = 7

输出： 3

解释：

区间 [5, 7] 内不存在特殊数字。

示例 2：

输入： l = 4, r = 16

输出： 11

解释：

区间 [4, 16] 内的特殊数字为 4 和 9。



提示：

1 <= l <= r <= 109
*/

func main() {
	timeCost(func() {
		log.Println(nonSpecialCount(2879, 52145))
	})

	// timeCost(func() {
	// 	log.Println(nonSpecialCountTrans(2879, 52145))
	// })

	timeCost(func() {
		log.Println(nonSpecialCountTrans(1, 2))
	})

}

func nonSpecialCount(l int, r int) int {
	res := r - l + 1

	// 因素中 有大于等于两次的 缓存
	verified := make(map[int]bool)

	for i := l; i <= r; i++ {
		cnt := 0

		// 偶数 除了 4 之外的都不满住
		if i > 4 && i%2 == 0 {
			// verified[i] = true
			continue
		}

		// 3的倍数 除了 9 之外的都不满住
		if i > 9 && i%3 == 0 {
			// verified[i] = true
			continue
		}

		for j := 1; j < i; j++ {
			if i%j == 0 {
				if _, ok := verified[j]; ok {
					cnt = 3
					break
				}
				cnt++
				// log.Printf(" cnt++ i %d  j %d\n", i, j)
			}
		}

		if cnt >= 2 {
			verified[i] = true
		}

		if cnt == 2 {
			// log.Printf("res-- i %d \n", i)
			res--
		}
	}

	return res
}

func timeCost(f func()) {
	s := time.Now()
	f()

	tc := time.Since(s)

	log.Printf("time COST %v\n", tc)
}

// 转换为 求素数 -然后 转换为素数的平方
func nonSpecialCountTrans(l int, r int) int {

	primeList := make([]int, 0)

	for i := int(math.Floor(math.Sqrt(float64(l)))); i <= int(math.Ceil(math.Sqrt(float64(r)))); i++ {
		// 不算 1算
		cnt := 0
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				cnt++
			}
		}
		if i != 1 && cnt == 0 {
			primeList = append(primeList, i)
		}
	}

	log.Printf("prime %v \n", primeList)

	res := r - l + 1
	for _, v := range primeList {
		if l <= v*v && r >= v*v {
			res--
		}
	}

	return res
}

/*
方法一：质数筛
思路与算法

特殊数字首先是一个平方数，并且除去自身和 1 之后的另一个因子一定是一个质数。这是因为：

因子一般是成双成对的，若一个数字有奇数个因子，那么该数一定是平方数。
该数除去自身和 1 仅有一个因子，因此该因子一定是质数。
因此，我们可以在 [1,
r
​
 ] 的范围内遍历所有质数（使用质数筛，具体方法可以参考题解 204. 计数质数），然后将它们的平方从 [l,r] 的范围中去除即可。

由于 r 的范围不超过 10
9
 ，因此质数的遍历范围不超过 31622，而使用很简单的埃氏筛（复杂度为 O(nlognlogn)，其中 n 为质数遍历范围）就可以轻松通过本题。

作者：力扣官方题解
链接：https://leetcode.cn/problems/find-the-count-of-numbers-which-are-not-special/solutions/2984412/tong-ji-bu-shi-te-shu-shu-zi-de-shu-zi-s-kq6j/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

func nonSpecialCountLeetCode(l int, r int) int {
	n := int(math.Sqrt(float64(r)))
	v := make([]int, n+1)
	res := r - l + 1
	for i := 2; i <= n; i++ {
		if v[i] == 0 {
			if i*i >= l && i*i <= r {
				res--
			}
			for j := i * 2; j <= n; j += i {
				v[j] = 1
			}
		}
	}
	return res
}
