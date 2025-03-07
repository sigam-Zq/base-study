package leetcode36

import (
	"log"
	"strconv"
	"testing"
)

/*
3249. 统计好节点的数目

https://leetcode.cn/problems/count-the-number-of-good-nodes/description/?envType=daily-question&envId=2024-11-14

现有一棵 无向 树，树中包含 n 个节点，按从 0 到 n - 1 标记。树的根节点是节点 0 。给你一个长度为 n - 1 的二维整数数组 edges，其中 edges[i] = [ai, bi] 表示树中节点 ai 与节点 bi 之间存在一条边。

如果一个节点的所有子节点为根的
子树
 包含的节点数相同，则认为该节点是一个 好节点。

返回给定树中 好节点 的数量。

子树 指的是一个节点以及它所有后代节点构成的一棵树。





示例 1：

输入：edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]

输出：7

说明：


树的所有节点都是好节点。

示例 2：

输入：edges = [[0,1],[1,2],[2,3],[3,4],[0,5],[1,6],[2,7],[3,8]]

输出：6

说明：


树中有 6 个好节点。上图中已将这些节点着色。

示例 3：

输入：edges = [[0,1],[1,2],[1,3],[1,4],[0,5],[5,6],[6,7],[7,8],[0,9],[9,10],[9,12],[10,11]]

输出：12

解释：


除了节点 9 以外其他所有节点都是好节点。



提示：

2 <= n <= 105
edges.length == n - 1
edges[i].length == 2
0 <= ai, bi < n
输入确保 edges 总表示一棵有效的树。
*/

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f     func([][]int) int
		edges [][]int
		want  int
	}{
		{
			f:     countGoodNodes,
			edges: [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
			want:  7,
		},
		{
			f:     countGoodNodes,
			edges: [][]int{{0, 1}, {1, 2}, {1, 3}, {1, 4}, {0, 5}, {5, 6}, {6, 7}, {7, 8}, {0, 9}, {9, 10}, {9, 12}, {10, 11}},
			want:  12,
		},
		{
			f:     countGoodNodes,
			edges: [][]int{{6, 0}, {1, 0}, {5, 1}, {2, 5}, {3, 1}, {4, 3}},
			want:  6,
		},
		{
			f:     countGoodNodes,
			edges: [][]int{{4, 0}, {6, 4}, {3, 6}, {5, 0}, {1, 5}, {2, 1}},
			want:  7,
		},
		{
			f:     countGoodNodes,
			edges: [][]int{{3, 0}, {1, 3}, {5, 0}, {6, 5}, {4, 6}, {2, 4}},
			want:  6,
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.edges); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})

	}
}

func countGoodNodes(edges [][]int) int {
	n := len(edges) + 1
	// index ---> 子阶段

	sortEdges := tidyEdge(edges)
	log.Printf("sortEdges -- %v --\n", sortEdges)
	treeNode := make([][]int, n)
	// treeNode 为该节点 可达节点
	for _, v := range sortEdges {
		treeNode[v[0]] = append(treeNode[v[0]], v[1])
	}
	log.Printf("treeNode -- %v --\n", treeNode)

	// 这里需要-递归才能拿到子节点的子节点

	// 倒序根据可达节点 计算节点权重
	treeHeavy := make([]int, n)
	var dfs func(idx int) int
	dfs = func(idx int) int {

		// 可以直接返回的情况
		// 叶子节点直接返回1
		if len(treeNode[idx]) == 0 {
			treeHeavy[idx] = 1
			return 1
		}

		// 寻找所有子节点
		heavy := 1
		for _, v := range treeNode[idx] {
			// heavy 维护过 直接加上自己返回
			if treeHeavy[v] != 0 {
				heavy += treeHeavy[v]
			} else {
				// 没有维护过
				heavy += dfs(v)
			}
		}
		return heavy
	}

	for i := n - 1; i >= 0; i-- {
		// treeHeavy[i] = len(treeNode[i]) + 1
		treeHeavy[i] = dfs(i)
	}

	log.Printf("treeHeavy -- %v --\n", treeHeavy)
	// 转换为 本节点 + 子节点重量
	treeHeavyNode := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		treeHeavyNode[i] = make([]int, len(treeNode[i]))
		for ii, v := range treeNode[i] {
			treeHeavyNode[i][ii] = treeHeavy[v]
		}
	}

	log.Printf("treeHeavyNode -- %v --\n", treeHeavyNode)

	res := 0
	for _, v := range treeHeavyNode {
		isAdd := true
		if len(v) > 0 {
			tmp := v[0]
			for _, vv := range v {
				if tmp != vv {
					isAdd = false
				}
			}
		}

		if isAdd {
			res++
		}

	}

	return res
}

// 整理路径 确保 其中的因子 0-树 1-枝
func tidyEdge(edges [][]int) [][]int {
	newEdge := make([][]int, len(edges))
	n := len(edges) + 1

	// 靠近0的距离
	nearRoot := make([]int, n)

	idx := 0
	done := n
	for done > 0 {
		// 大于n进行取余
		if idx >= n-1 {
			idx %= (n - 1)
		}
		// 找到0时
		if edges[idx][0] == 0 {
			nearRoot[edges[idx][0]] = 0
			nearRoot[edges[idx][1]] = 1
			done--
		}
		if edges[idx][1] == 0 {
			nearRoot[edges[idx][1]] = 0
			nearRoot[edges[idx][0]] = 1
			done--
		}

		// 找到一个和0 相关联的点时候
		if edges[idx][0] != 0 && nearRoot[edges[idx][0]] == 0 && nearRoot[edges[idx][1]] != 0 {
			done--
			nearRoot[edges[idx][0]] = nearRoot[edges[idx][1]] + 1
		}
		if edges[idx][1] != 0 && nearRoot[edges[idx][1]] == 0 && nearRoot[edges[idx][0]] != 0 {
			done--
			nearRoot[edges[idx][1]] = nearRoot[edges[idx][0]] + 1
		}
		idx++
	}

	log.Printf("nearRoot -- %v --\n", nearRoot)
	// 0 必定为根 0 顺序必须在前面
	for i, v := range edges {
		newEdge[i] = make([]int, 2)

		// 根据哪个点靠近0 交换位置
		if nearRoot[v[0]] < nearRoot[v[1]] {
			newEdge[i][0], newEdge[i][1] = v[0], v[1]
		} else {
			newEdge[i][0], newEdge[i][1] = v[1], v[0]
		}

	}

	return newEdge
}

/*
方法一：深度优先搜索
思路

首先根据边数组 edges 构建邻接表 g。在树中，边的数量为节点数量减 1。因此，n 为 edges 的长度加 1。再构造深度优先搜索，输入为当前遍历的节点 node 和其父节点 parent，返回值为以 node 为根节点的树的节点数量。需要递归调用 node 的所有子节点。因为 g 中存的是邻接关系，所以要跳过节点 parent。在计算节点数量和的同时，需要判断 node 的所有子节点是否拥有相同的节点数。如果是的话，将结果加 1。最后调用 dfs(0,−1) 并返回结果。

代码

*/

func countGoodNodesLeetCode(edges [][]int) int {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	res := 0
	var dfs func(node, parent int) int
	dfs = func(node, parent int) int {
		valid := true
		treeSize := 0
		subTreeSize := 0

		for _, child := range g[node] {
			if child != parent {
				size := dfs(child, node)
				if subTreeSize == 0 {
					subTreeSize = size
				} else if size != subTreeSize {
					valid = false
				}
				treeSize += size
			}
		}
		if valid {
			res++
		}
		return treeSize + 1
	}

	dfs(0, -1)
	return res
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/count-the-number-of-good-nodes/solutions/2977809/tong-ji-hao-jie-dian-de-shu-mu-by-leetco-4q70/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
