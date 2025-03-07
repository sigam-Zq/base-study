package leetcode40

import (
	"reflect"
	"strconv"
	"testing"
)

/*

3243. 新增道路查询后的最短距离 I
给你一个整数 n 和一个二维整数数组 queries。

有 n 个城市，编号从 0 到 n - 1。初始时，每个城市 i 都有一条单向道路通往城市 i + 1（ 0 <= i < n - 1）。

queries[i] = [ui, vi] 表示新建一条从城市 ui 到城市 vi 的单向道路。每次查询后，你需要找到从城市 0 到城市 n - 1 的最短路径的长度。

返回一个数组 answer，对于范围 [0, queries.length - 1] 中的每个 i，answer[i] 是处理完前 i + 1 个查询后，从城市 0 到城市 n - 1 的最短路径的长度。



示例 1：

输入： n = 5, queries = [[2, 4], [0, 2], [0, 4]]

输出： [3, 2, 1]

解释：



新增一条从 2 到 4 的道路后，从 0 到 4 的最短路径长度为 3。



新增一条从 0 到 2 的道路后，从 0 到 4 的最短路径长度为 2。



新增一条从 0 到 4 的道路后，从 0 到 4 的最短路径长度为 1。

示例 2：

输入： n = 4, queries = [[0, 3], [0, 2]]

输出： [1, 1]

解释：



新增一条从 0 到 3 的道路后，从 0 到 3 的最短路径长度为 1。



新增一条从 0 到 2 的道路后，从 0 到 3 的最短路径长度仍为 1。



提示：

3 <= n <= 500
1 <= queries.length <= 500
queries[i].length == 2
0 <= queries[i][0] < queries[i][1] < n
1 < queries[i][1] - queries[i][0]
查询中没有重复的道路。
*/

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f       func(int, [][]int) []int
		n       int
		queries [][]int
		want    []int
	}{
		{
			f:       shortestDistanceAfterQueriesFast,
			n:       5,
			queries: [][]int{{2, 4}, {0, 2}, {0, 4}},
			want:    []int{3, 2, 1},
		},
		{
			f:       shortestDistanceAfterQueriesFast,
			n:       38,
			queries: [][]int{{18, 31}, {8, 33}, {15, 32}, {17, 36}, {10, 14}, {10, 26}, {15, 27}, {6, 30}, {1, 7}, {9, 13}, {28, 35}, {1, 10}, {1, 28}, {22, 37}, {22, 30}, {17, 28}, {27, 35}, {14, 37}, {3, 13}, {2, 12}, {7, 30}, {8, 13}, {1, 20}, {16, 19}, {1, 37}, {10, 31}, {1, 12}, {10, 18}, {5, 17}, {13, 15}, {0, 15}, {0, 9}, {2, 11}, {11, 24}, {8, 36}, {3, 7}, {9, 17}, {29, 32}, {3, 32}, {4, 10}, {1, 23}, {21, 29}, {6, 37}, {6, 11}, {6, 16}, {10, 13}, {23, 31}, {0, 32}, {18, 25}, {2, 13}, {21, 31}, {19, 24}, {0, 27}, {1, 25}, {9, 25}, {0, 31}, {5, 36}, {8, 20}, {4, 25}, {16, 23}, {1, 19}, {30, 37}, {0, 3}, {13, 29}, {32, 35}, {8, 27}, {16, 36}, {0, 28}, {2, 32}, {11, 19}, {19, 23}, {7, 37}, {11, 30}, {27, 31}, {10, 12}, {7, 17}, {8, 17}, {12, 20}, {2, 14}, {1, 24}, {6, 31}, {3, 10}, {9, 19}, {14, 30}, {26, 34}, {9, 30}, {30, 35}, {4, 24}, {0, 36}, {6, 33}, {4, 8}, {6, 34}, {1, 5}, {17, 21}, {23, 26}, {12, 28}, {10, 32}, {19, 25}, {8, 10}, {24, 26}, {0, 16}, {18, 35}, {15, 26}, {12, 23}, {25, 37}, {3, 35}, {0, 8}, {23, 27}, {22, 28}, {11, 25}, {14, 34}, {5, 32}, {7, 29}, {15, 17}, {0, 13}, {12, 25}, {21, 27}, {16, 34}, {10, 28}, {0, 35}, {16, 30}, {11, 28}, {14, 32}, {5, 28}, {9, 16}, {8, 26}, {15, 20}, {11, 16}, {9, 28}, {3, 28}, {25, 33}, {2, 31}, {4, 34}, {16, 18}, {15, 34}, {6, 35}, {28, 37}, {5, 19}, {21, 24}, {12, 32}, {21, 32}, {6, 36}, {15, 35}, {15, 22}, {3, 14}, {1, 30}, {9, 11}, {6, 8}, {9, 29}, {13, 19}, {3, 34}, {16, 31}, {5, 15}, {22, 24}, {17, 26}, {3, 27}, {1, 36}, {0, 12}, {2, 23}, {9, 18}, {16, 35}, {3, 37}, {10, 35}, {2, 19}, {9, 34}, {7, 34}, {12, 33}, {10, 37}, {10, 34}, {25, 35}, {23, 30}, {2, 26}, {23, 25}, {27, 29}, {2, 16}, {5, 37}, {20, 30}, {31, 33}, {0, 37}, {2, 34}, {8, 19}, {34, 37}, {17, 30}, {3, 5}, {1, 17}, {2, 7}, {20, 37}, {4, 37}},
			want:    []int{3, 2, 1},
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.n, v.queries); !reflect.DeepEqual(got, v.want) {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})

	}
}

func shortestDistanceAfterQueries(n int, queries [][]int) []int {

	// 维护每个节点的可达数组
	toPoints := make([][]int, n)

	for i := 0; i < n; i++ {
		toPoints[i] = []int{i + 1}
	}

	var bfs func(int, int, int) int

	// 寻找从 startPoint 到 endPoint 的最短路径
	bfs = func(startPoint, endPoint, cost int) int {
		cost++
		lastConst := 0

		for _, v := range toPoints[startPoint] {
			if v == endPoint {
				return cost
			} else {
				if lastConst == 0 {
					lastConst = bfs(v, endPoint, cost)
				} else {
					lastConst = min(lastConst, bfs(v, endPoint, cost))
				}

			}
		}
		return lastConst
	}

	res := make([]int, len(queries))
	for i, v := range queries {
		toPoints[v[0]] = append(toPoints[v[0]], v[1])
		res[i] = bfs(0, n-1, 0)
	}

	return res
}

func shortestDistanceAfterQueriesFast(n int, queries [][]int) []int {
	neighbors := make([][]int, n)
	for i := 0; i < n-1; i++ {
		neighbors[i] = append(neighbors[i], i+1)
	}

	bfs := func(neigh [][]int) int {
		// 距离0点的距离
		dist := make([]int, n)
		for i := 1; i < n; i++ {
			dist[i] = -1
		}

		q := []int{0}
		for len(q) > 0 {
			x := q[0]
			q = q[1:]
			for _, y := range neigh[x] {
				if dist[y] >= 0 {
					continue
				} else {
					q = append(q, y)
					dist[y] = dist[x] + 1
				}
			}
		}

		return dist[n-1]
	}

	res := make([]int, len(queries))
	for i, v := range queries {
		neighbors[v[0]] = append(neighbors[v[0]], v[1])
		res[i] = bfs(neighbors)
	}

	return res
}

func shortestDistanceAfterQueriesLeetCodeFast(n int, queries [][]int) []int {
	neighbors := make([][]int, n)
	for i := 0; i < n-1; i++ {
		neighbors[i] = append(neighbors[i], i+1)
	}
	var res []int
	for _, query := range queries {
		neighbors[query[0]] = append(neighbors[query[0]], query[1])
		res = append(res, bfs(n, neighbors))
	}
	return res
}

func bfs(n int, neighbors [][]int) int {
	dist := make([]int, n)
	for i := 1; i < n; i++ {
		dist[i] = -1
	}
	q := []int{0}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range neighbors[x] {
			if dist[y] >= 0 {
				continue
			}
			q = append(q, y)
			dist[y] = dist[x] + 1
		}
	}
	return dist[n-1]

	// 作者：力扣官方题解
	// 链接：https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-i/solutions/2984418/xin-zeng-dao-lu-cha-xun-hou-de-zui-duan-9smce/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}

/*
gpt

### 解1 事件复杂度分析

解1使用的是递归深度优先搜索（DFS）的方法。具体复杂度分析如下：

1. **初始化阶段**：
   - 构造 `toPoints` 数组的时间复杂度是 \(O(n)\)，因为需要遍历每个节点进行初始化。

2. **处理查询阶段**：
   - 对于每个查询，都会修改 `toPoints` 数组，同时调用 `bfs` 计算最短路径。
   - 在 `bfs` 方法中，最坏情况下会从起点遍历所有可能的路径到终点。因此，单次查询的复杂度为 \(O(V^V)\)，其中 \(V = n\)。这个复杂度非常高，因为每个节点可能有多条出边，导致递归调用的路径数指数增长。

3. **总复杂度**：
   - 如果有 \(q\) 个查询，总时间复杂度是 \(O(q \times V^V)\)。当节点数 \(n\) 较大时，这种复杂度会导致程序运行极其缓慢。

---

### 解2 事件复杂度分析

解2使用的是广度优先搜索（BFS）的方法。复杂度分析如下：

1. **初始化阶段**：
   - 构造 `neighbors` 数组的时间复杂度是 \(O(n)\)。

2. **处理查询阶段**：
   - 对于每个查询，添加一条边的复杂度是 \(O(1)\)，然后调用 `bfs` 方法。
   - 在 `bfs` 方法中，从起点出发的搜索会访问所有节点一次，并遍历所有出边一次。因此，单次调用 `bfs` 的复杂度是 \(O(V + E)\)，其中 \(V = n\) 是节点数，\(E\) 是当前的边数。
   - 每次查询更新后，最多增加一条边，所以 \(E\) 的最大值为初始边数加查询数 \(n-1 + q\)。

3. **总复杂度**：
   - 设查询数为 \(q\)，总时间复杂度为 \(O(q \times (V + E)) = O(q \times (n + q))\)。

---

### 解2 的实现详细解释

解2主要分为两部分：动态更新图结构和使用 BFS 计算最短路径。

#### 1. 动态更新图结构
- 每个节点的邻接节点存储在 `neighbors` 数组中。
- 初始时，节点 \(i\) 的邻接节点只有 \(i+1\)（形成一条从 0 到 \(n-1\) 的直线）。
- 每次查询 \( [u, v] \) 新增一条从 \(u\) 到 \(v\) 的边，直接将 \(v\) 添加到 \(u\) 的邻接节点列表中。

#### 2. 使用 BFS 计算最短路径
- **目的**：找到从 0 到 \(n-1\) 的最短路径。
- **步骤**：
  1. 初始化距离数组 `dist`，表示从起点到每个节点的最短距离。起点距离为 0，其余节点初始化为 -1。
  2. 使用队列 `q` 实现逐层搜索，初始时将起点 0 入队。
  3. 每次从队列中取出一个节点 \(x\)，遍历其所有邻接节点 \(y\)。如果 \(y\) 未被访问（即 `dist[y] == -1`），则将其加入队列，并更新其最短距离为 `dist[x] + 1`。
  4. 当队列为空时，结束搜索，此时 `dist[n-1]` 即为最短路径长度。
- BFS 的特性保证了首次访问到节点 \(n-1\) 时，一定是最短路径。

---

### 解1 和解2 的对比
- **复杂度**：解2的复杂度远优于解1，特别是当节点数和查询数较大时。
- **实现效率**：解2使用 BFS，在最短路径搜索中效率更高且更符合直觉，而解1的递归深搜方式会重复计算路径，导致冗余操作。
- **代码清晰度**：解2代码更简洁明了，逻辑易于理解。
*/

// 方法二：动态规划
// 根据题意，对于任一单向道路的起始点 u，终止点 v，都有 u<v，
//那么从城市 0 到任一城市的路径上，所经过的城市编号是单调递增的。//
//令 dp[i] 表示城市 0 到城市 i 的最短路径，同时使用 prev[i] 记录通往城市 i 的所有单向道路的起始城市集合，
//那么对于 i>0，有 dp[i]=min
// j∈prev[i]
// ​
//  dp[j]+1。

// 根据以上推论，我们可以遍历 queries，在每次查询时，
//更新 prev 数组，然后更新 dp 数组。注意到，每次新建一条从城市 u 到城市 v 的单向道路时，
//只有 i≥v 的 dp[i] 会发生变化，因此更新 dp 可以从 v 开始更新。

func shortestDistanceAfterQueriesLeetCodeDp(n int, queries [][]int) []int {
	prev := make([][]int, n)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		prev[i] = append(prev[i], i-1)
		dp[i] = i
	}
	var res []int
	for _, query := range queries {
		prev[query[1]] = append(prev[query[1]], query[0])
		for v := query[1]; v < n; v++ {
			for _, u := range prev[v] {
				dp[v] = min(dp[v], dp[u]+1)
			}
		}
		res = append(res, dp[n-1])
	}
	return res
}
