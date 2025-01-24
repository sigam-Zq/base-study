package leetcode94

import (
	"reflect"
	"strconv"
	"testing"
)

// 2920. 收集所有金币可获得的最大积分

/*
有一棵由 n 个节点组成的无向树，以 0  为根节点，节点编号从 0 到 n - 1 。给你一个长度为 n - 1 的二维 整数 数组 edges ，其中 edges[i] = [ai, bi] 表示在树上的节点 ai 和 bi 之间存在一条边。另给你一个下标从 0 开始、长度为 n 的数组 coins 和一个整数 k ，其中 coins[i] 表示节点 i 处的金币数量。

从根节点开始，你必须收集所有金币。要想收集节点上的金币，必须先收集该节点的祖先节点上的金币。

节点 i 上的金币可以用下述方法之一进行收集：

收集所有金币，得到共计 coins[i] - k 点积分。如果 coins[i] - k 是负数，你将会失去 abs(coins[i] - k) 点积分。
收集所有金币，得到共计 floor(coins[i] / 2) 点积分。如果采用这种方法，节点 i 子树中所有节点 j 的金币数 coins[j] 将会减少至 floor(coins[j] / 2) 。
返回收集 所有 树节点的金币之后可以获得的最大积分。



示例 1：


输入：edges = [[0,1],[1,2],[2,3]], coins = [10,10,3,3], k = 5
输出：11
解释：
使用第一种方法收集节点 0 上的所有金币。总积分 = 10 - 5 = 5 。
使用第一种方法收集节点 1 上的所有金币。总积分 = 5 + (10 - 5) = 10 。
使用第二种方法收集节点 2 上的所有金币。所以节点 3 上的金币将会变为 floor(3 / 2) = 1 ，总积分 = 10 + floor(3 / 2) = 11 。
使用第二种方法收集节点 3 上的所有金币。总积分 =  11 + floor(1 / 2) = 11.
可以证明收集所有节点上的金币能获得的最大积分是 11 。
示例 2：


输入：edges = [[0,1],[0,2]], coins = [8,4,4], k = 0
输出：16
解释：
使用第一种方法收集所有节点上的金币，因此，总积分 = (8 - 0) + (4 - 0) + (4 - 0) = 16 。


提示：

n == coins.length
2 <= n <= 105
0 <= coins[i] <= 104
edges.length == n - 1
0 <= edges[i][0], edges[i][1] < n
0 <= k <= 104
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f       func(edges [][]int, coins []int, k int) int
		edges   [][]int
		coins   []int
		k       int
		want    int
		isDebug bool
	}{
		{
			f:       maximumPoints,
			edges:   [][]int{{0, 1}, {1, 2}, {2, 3}},
			coins:   []int{10, 10, 3, 3},
			k:       5,
			want:    11,
			isDebug: false,
		},
		{
			f:       maximumPoints,
			edges:   [][]int{{0, 1}, {0, 2}},
			coins:   []int{8, 4, 4},
			k:       0,
			want:    16,
			isDebug: false,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans := v.f(v.edges, v.coins, v.k)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	roadList := make([][]int, n)

	for _, v := range edges {
		u, v := v[0], v[1]
		roadList[u] = append(roadList[u], v)
		roadList[v] = append(roadList[v], u)
	}

	memo := make([][14]int, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(node, parent, f int) int

	dfs = func(node, parent, f int) int {
		if memo[node][f] != -1 {
			return memo[node][f]
		}

		res1, res2 := (coins[node]>>f)-k, (coins[node] >> (f + 1))

		for _, targetIdx := range roadList[node] {
			if targetIdx == parent {
				continue
			}

			res1 += dfs(targetIdx, node, f)
			if f+1 < 14 {
				// 大于14 的值都为0
				res2 += dfs(targetIdx, node, f+1)
			}

		}

		memo[node][f] = max(res1, res2)
		return memo[node][f]
	}

	return dfs(0, -1, 0)
}

func maximumPointsLeetCode(edges [][]int, coins []int, k int) int {
	n := len(coins)
	children := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		children[u] = append(children[u], v)
		children[v] = append(children[v], u)
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 14)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(node, parent, f int) int
	dfs = func(node, parent, f int) int {
		if memo[node][f] >= 0 {
			return memo[node][f]
		}
		res0, res1 := (coins[node]>>f)-k, coins[node]>>(f+1)
		for _, child := range children[node] {
			if child == parent {
				continue
			}
			res0 += dfs(child, node, f)
			if f+1 < 14 {
				res1 += dfs(child, node, f+1)
			}
		}
		memo[node][f] = max(res0, res1)
		return memo[node][f]
	}

	return dfs(0, -1, 0)
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/maximum-points-after-collecting-coins-from-all-nodes/solutions/3047557/shou-ji-suo-you-jin-bi-ke-huo-de-de-zui-d6zuo/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
