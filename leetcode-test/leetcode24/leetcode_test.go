package leetcode24

import (
	"log"
	"reflect"
	"testing"
)

/*
685. 冗余连接 II

https://leetcode.cn/problems/redundant-connection-ii/description/

在本问题中，有根树指满足以下条件的 有向 图。该树只有一个根节点，所有其他节点都是该根节点的后继。该树除了根节点之外的每一个节点都有且只有一个父节点，而根节点没有父节点。

输入一个有向图，该图由一个有着 n 个节点（节点值不重复，从 1 到 n）的树及一条附加的有向边构成。附加的边包含在 1 到 n 中的两个不同顶点间，这条附加的边不属于树中已存在的边。

结果图是一个以边组成的二维数组 edges 。 每个元素是一对 [ui, vi]，用以表示 有向 图中连接顶点 ui 和顶点 vi 的边，其中 ui 是 vi 的一个父节点。

返回一条能删除的边，使得剩下的图是有 n 个节点的有根树。若有多个答案，返回最后出现在给定二维数组的答案。



示例 1：


输入：edges = [[1,2],[1,3],[2,3]]
输出：[2,3]
示例 2：


输入：edges = [[1,2],[2,3],[3,4],[4,1],[1,5]]
输出：[4,1]


提示：

n == edges.length
3 <= n <= 1000
edges[i].length == 2
1 <= ui, vi <= n

*/

func TestXxx(t *testing.T) {
	args := []struct {
		name  string
		f     func(edges [][]int) []int
		edges [][]int
		want  []int
	}{
		{
			name:  "oneTest",
			f:     findRedundantDirectedConnection,
			edges: [][]int{{1, 2}, {1, 3}, {2, 3}},
			want:  []int{2, 3},
		},
		{
			name:  "twoTest",
			f:     findRedundantDirectedConnection,
			edges: [][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}},
			want:  []int{2, 1},
		},
		{
			name:  "case2",
			f:     findRedundantDirectedConnection,
			edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}},
			want:  []int{4, 1},
		},
		{
			name:  "threeTest",
			f:     findRedundantDirectedConnection,
			edges: [][]int{{4, 2}, {1, 5}, {5, 2}, {5, 3}, {2, 4}},
			want:  []int{4, 2},
		},
		{
			name:  "LeetCodeThreeTest",
			f:     findRedundantDirectedConnectionLeetCode,
			edges: [][]int{{4, 2}, {1, 5}, {5, 2}, {5, 3}, {2, 4}},
			want:  []int{4, 2},
		},
		{
			name:  "fix",
			f:     findRedundantDirectedConnectionFix,
			edges: [][]int{{4, 2}, {1, 5}, {5, 2}, {5, 3}, {2, 4}},
			want:  []int{4, 2},
		},
		{
			name:  "fourCase",
			f:     findRedundantDirectedConnection,
			edges: [][]int{{5, 2}, {5, 1}, {3, 1}, {3, 4}, {3, 5}},
			want:  []int{3, 1},
		},
		{
			name:  "fiveCase",
			f:     findRedundantDirectedConnection,
			edges: [][]int{{3, 1}, {1, 4}, {3, 5}, {1, 2}, {1, 5}},
			want:  []int{1, 5},
		},
	}

	for _, v := range args {
		t.Run(v.name, func(t *testing.T) {
			if got := v.f(v.edges); !reflect.DeepEqual(got, v.want) {
				t.Errorf(" got %v  want %v \n", got, v.want)
			}
		})
	}

}

func findRedundantDirectedConnectionFix(edges [][]int) []int {

	// // idx -> vlue
	// graph := make([][]int, len(edges)+1)
	// // 入度
	// nodeInNum := make([]int, len(edges)+1)

	// for _, road := range edges {
	// 	graph[road[0]] = append(graph[road[0]], road[1])
	// 	nodeInNum[road[1]]++
	// }
	// log.Printf("fix----%d \n", graph)
	// isKeyRoad := make([]bool, len(edges)+1)

	// // 选择一个 没有入度的（根） 或者 有两个出度的 非叶子 为起点便利
	// startNode := 0
	// // 异常节点
	// abnormalIdx := 0
	// for i, v := range nodeInNum {
	// 	if v == 0 {
	// 		startNode = i
	// 	}

	// 	if v > 1 {
	// 		abnormalIdx = i
	// 	}
	// }
	// if startNode == 0

	// 1-n 是否有未出现的值

	return nil
}

func findRedundantDirectedConnection(edges [][]int) []int {
	nodeInNum := make([]int, len(edges)+1)
	nodeOutNum := make([]int, len(edges)+1)

	// idx -> vlue
	graphToLow := make([][]int, len(edges)+1)
	graphToHigh := make([][]int, len(edges)+1)
	for _, road := range edges {
		graphToLow[road[0]] = append(graphToLow[road[0]], road[1])
		graphToHigh[road[1]] = append(graphToLow[road[1]], road[0])
		nodeOutNum[road[0]]++
		nodeInNum[road[1]]++
	}

	abnormalIdx := 0
	// 入度为 0 的节点 或者出度为0的节点的 延伸节点的路径起始节点（不可删除节点 起点）
	isFromKeyRoad := make([]bool, len(edges)+1)
	// 不可删除的 到达节点
	isToKeyRoad := make([]bool, len(edges)+1)
	// 图中仅存在一个 多点出 X 不对 非叶子节点都是多出
	// 正常树 只有一个节点 无入有出 其余节点均只有一个入
	for i, v := range nodeInNum {
		if v > 1 {
			abnormalIdx = i
		}
		if v == 0 && nodeOutNum[i] == 1 {
			isFromKeyRoad[i] = true
			isToKeyRoad[graphToLow[i][0]] = true
		}
	}

	for i, v := range nodeOutNum {
		if v == 0 && nodeInNum[i] == 1 {
			isToKeyRoad[i] = true
		}
	}
	log.Printf("--isKeyRoad %v  \n", isFromKeyRoad)
	log.Printf("--graphToLow %v  \n", graphToLow)
	log.Printf("--graphToHigh %v  \n", graphToHigh)
	// 如果是进入根节点的路径idx
	for i := len(edges) - 1; i >= 0; i-- {
		if abnormalIdx == edges[i][1] && !isFromKeyRoad[edges[i][0]] && !isToKeyRoad[edges[i][1]] {
			return edges[i]
		}
	}

	for i := len(edges) - 1; i >= 0; i-- {
		if !isFromKeyRoad[edges[i][0]] {
			return edges[i]
		}
	}

	// 当使用出发节点定义关键节点都没有出现返回值时 返回任意一个带有异常节点的出度

	for i := len(edges) - 1; i >= 0; i-- {
		if abnormalIdx == edges[i][1] {
			return edges[i]
		}
	}
	return nil

}

func findRedundantDirectedConnectionLeetCode(edges [][]int) []int {
	parent := make([]int, len(edges)+1)

	for i := range parent {
		parent[i] = i
	}

	var find func(int) int

	find = func(x int) int {
		if x != parent[x] {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return true
		}
		parent[y] = x

		return false
	}

	for _, road := range edges {
		if union(road[0], road[1]) {
			return road
		}
	}

	return nil
}
