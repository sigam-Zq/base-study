package leetcode67

import (
	"log"
	"sort"
	"strconv"
	"testing"
)

// https://leetcode.cn/problems/sort-integers-by-the-power-value/solutions/168355/jiang-zheng-shu-an-quan-zhong-pai-xu-by-leetcode-s/?envType=daily-question&envId=2024-12-22

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func(int, int, int) int
		lo   int
		hi   int
		k    int
		want int
	}{
		{
			f:    getKth,
			lo:   12,
			hi:   15,
			k:    2,
			want: 13,
		},
		{
			f:    getKth,
			lo:   7,
			hi:   11,
			k:    4,
			want: 7,
		},
		{
			f:    getKthLeetCode1,
			lo:   1,
			hi:   1000,
			k:    777,
			want: 570,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.lo, v.hi, v.k)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}

func getKth(lo int, hi int, k int) int {

	dp := make([]int, 250505)
	dp[1] = 0

	// 递归计算权重
	var dfsWeightCalc func(x int) int
	dfsWeightCalc = func(x int) int {
		if x == 1 {
			return 0
		}

		var t int
		if x&1 == 0 {
			// 偶数
			t = x / 2

		} else {
			// 奇数
			t = x*3 + 1
		}

		if dp[t] != 0 {
			dp[x] = dp[t] + 1
			return dp[t] + 1
		}

		res := dfsWeightCalc(t) + 1
		dp[x] = res
		return res

	}

	p := make([][2]int, hi-lo+1)

	for i := lo; i <= hi; i++ {
		p[i-lo][0], p[i-lo][1] = i, dfsWeightCalc(i)
	}

	// 根据dp计算顺序
	sort.Slice(p, func(i, j int) bool {
		if p[i][1] == p[j][1] {
			return p[i][0] < p[j][0]
		} else {
			return p[i][1] < p[j][1]
		}
	})
	// log.Println(p)
	// log.Println(k)
	return p[k-1][0]
}

func TestClac(t *testing.T) {

	for i := 1; i < 1001; i++ {
		log.Println(Clac(i))
	}
	log.Println(maxI)
}

var dp = make([]int, 1e9)
var maxI int

func Clac(x int) int {
	if x == 1 {
		return 0
	}

	var t int
	if x&1 == 0 {
		// 偶数
		t = x / 2

	} else {
		// 奇数
		t = x*3 + 1
	}

	maxI = max(maxI, t)
	if dp[t] != 0 {
		dp[x] = dp[t] + 1
		return dp[t] + 1
	}

	res := Clac(t) + 1
	dp[x] = res
	return res

}

/*
题目分析
我们要按照权重为第一关键字，原值为第二关键字对区间 [lo, hi] 进行排序，关键在于我们怎么求权重。

方法一：递归
思路

记 x 的权重为 f(x)，按照题意很明显我们可以构造这样的递归式：

https://leetcode.cn/problems/sort-integers-by-the-power-value/solutions/168355/jiang-zheng-shu-an-quan-zhong-pai-xu-by-leetcode-s/?envType=daily-question&envId=2024-12-22

于是我们就可以递归求解每个数字的权重了。

代码

*/

func getKthLeetCode1(lo int, hi int, k int) int {
	v := []int{}
	for i := lo; i <= hi; i++ {
		v = append(v, i)
	}
	sort.Slice(v, func(i, j int) bool {
		if getFLeetCode1(v[i]) != getFLeetCode1(v[j]) {
			return getFLeetCode1(v[i]) < getFLeetCode1(v[j])
		}
		return v[i] < v[j]
	})
	return v[k-1]
}

func getFLeetCode1(x int) int {
	if x == 1 {
		return 0
	}
	if x&1 == 1 {
		return getFLeetCode1(x*3+1) + 1
	} else {
		return getFLeetCode1(x/2) + 1
	}
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/sort-integers-by-the-power-value/solutions/168355/jiang-zheng-shu-an-quan-zhong-pai-xu-by-leetcode-s/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/*
方法二：记忆化
思路

我们知道在求 f(3) 的时候会调用到 f(10)，在求 f(20) 的时候也会调用到 f(10)。同样的，如果单纯递归计算权重的话，
会存在很多重复计算，我们可以用记忆化的方式来加速这个过程，即「先查表，再计算」和「先记忆，再返回」。
我们可以用一个哈希映射作为这里的记忆化的「表」，这样保证每个元素的权值只被计算 1 次。在 [1,1000] 中所有 x 求 f(x) 的值的过程中，只可能出现 2228 种 x，于是效率就会大大提高。

*/

var f = make(map[int]int)

func getFLeetCode2(x int) int {
	if val, exists := f[x]; exists {
		return val
	}
	if x == 1 {
		f[x] = 0
		return 0
	}
	if x&1 == 1 {
		f[x] = getFLeetCode2(x*3+1) + 1
	} else {
		f[x] = getFLeetCode2(x/2) + 1
	}
	return f[x]
}

func getKthLeetCode2(lo int, hi int, k int) int {
	v := make([]int, 0)
	for i := lo; i <= hi; i++ {
		v = append(v, i)
	}
	sort.Slice(v, func(i, j int) bool {
		if getFLeetCode2(v[i]) != getFLeetCode2(v[j]) {
			return getFLeetCode2(v[i]) < getFLeetCode2(v[j])
		}
		return v[i] < v[j]
	})
	return v[k-1]
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/sort-integers-by-the-power-value/solutions/168355/jiang-zheng-shu-an-quan-zhong-pai-xu-by-leetcode-s/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
