package leetcode46

import (
	"container/heap"
	"log"
	"math"
	"strconv"
	"testing"
)

/*
743. 网络延迟时间

有 n 个网络节点，标记为 1 到 n。

给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， wi 是一个信号从源节点传递到目标节点的时间。

现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。



示例 1：



输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
输出：2
示例 2：

输入：times = [[1,2,1]], n = 2, k = 1
输出：1
示例 3：

输入：times = [[1,2,1]], n = 2, k = 2
输出：-1


提示：

1 <= k <= n <= 100
1 <= times.length <= 6000
times[i].length == 3
1 <= ui, vi <= n
ui != vi
0 <= wi <= 100
所有 (ui, vi) 对都 互不相同（即，不含重复边）有 n 个网络节点，标记为 1 到 n。

给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， wi 是一个信号从源节点传递到目标节点的时间。

现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。



示例 1：



输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
输出：2
示例 2：

输入：times = [[1,2,1]], n = 2, k = 1
输出：1
示例 3：

输入：times = [[1,2,1]], n = 2, k = 2
输出：-1


提示：

1 <= k <= n <= 100
1 <= times.length <= 6000
times[i].length == 3
1 <= ui, vi <= n
ui != vi
0 <= wi <= 100
所有 (ui, vi) 对都 互不相同（即，不含重复边）
*/

func TestXxx(t *testing.T) {
	for i, v := range []struct {
		f     func(times [][]int, n int, k int) int
		times [][]int
		n     int
		k     int
		want  int
	}{
		// {
		// 	f:     networkDelayTimeMockTime,
		// 	times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
		// 	n:     4,
		// 	k:     2,
		// 	want:  2,
		// },
		// {
		// 	f:     networkDelayTimeMockTime,
		// 	times: [][]int{{1, 2, 1}},
		// 	n:     2,
		// 	k:     2,
		// 	want:  -1,
		// },
		// {
		// 	f:     networkDelayTimeMockTime,
		// 	times: [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 2}},
		// 	n:     3,
		// 	k:     1,
		// 	want:  2,
		// },
		// {
		// 	f:     networkDelayTimeMockTime,
		// 	times: [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}},
		// 	n:     3,
		// 	k:     1,
		// 	want:  3,
		// },
		{
			f:     networkDelayTimeMockTime,
			times: [][]int{{3, 5, 78}, {2, 1, 1}, {1, 3, 0}, {4, 3, 59}, {5, 3, 85}, {5, 2, 22}, {2, 4, 23}, {1, 4, 43}, {4, 5, 75}, {5, 1, 15}, {1, 5, 91}, {4, 1, 16}, {3, 2, 98}, {3, 4, 22}, {5, 4, 31}, {1, 2, 0}, {2, 5, 4}, {4, 2, 51}, {3, 1, 36}, {2, 3, 59}},
			n:     5,
			k:     5,
			want:  31,
		},
		{
			f:     networkDelayTimeMockTime,
			times: [][]int{{14, 1, 8}, {11, 2, 25}, {14, 15, 37}, {3, 7, 70}, {11, 7, 60}, {13, 11, 87}, {15, 10, 67}, {13, 10, 58}, {5, 4, 56}, {9, 3, 26}, {5, 11, 51}, {11, 4, 92}, {7, 6, 8}, {7, 10, 95}, {14, 9, 0}, {4, 13, 1}, {7, 9, 89}, {3, 14, 24}, {11, 15, 30}, {13, 2, 91}, {15, 8, 60}, {1, 4, 96}, {8, 2, 71}, {6, 8, 38}, {14, 13, 46}, {2, 12, 48}, {10, 11, 92}, {8, 12, 28}, {8, 7, 12}, {9, 13, 82}, {8, 6, 27}, {3, 2, 65}, {4, 10, 62}, {11, 13, 55}, {1, 2, 52}, {8, 3, 98}, {7, 12, 85}, {6, 12, 97}, {9, 4, 90}, {2, 4, 23}, {9, 11, 20}, {1, 14, 61}, {8, 9, 77}, {6, 5, 80}, {14, 11, 33}, {9, 8, 54}, {13, 1, 42}, {13, 8, 13}, {10, 14, 40}, {9, 7, 18}, {14, 3, 50}, {14, 6, 83}, {14, 8, 14}, {2, 1, 86}, {9, 5, 54}, {11, 5, 29}, {9, 12, 43}, {9, 2, 74}, {14, 4, 87}, {12, 7, 98}, {7, 14, 13}, {4, 12, 33}, {5, 2, 60}, {15, 11, 33}, {8, 4, 99}, {9, 6, 98}, {4, 6, 57}, {6, 11, 5}, {9, 15, 37}, {1, 3, 30}, {9, 10, 60}, {13, 12, 73}, {13, 14, 56}, {1, 11, 13}, {14, 2, 8}, {4, 15, 60}, {11, 3, 90}, {2, 5, 86}, {11, 1, 1}, {13, 4, 2}, {15, 7, 91}, {15, 4, 51}, {11, 6, 70}, {2, 7, 51}, {11, 9, 37}, {4, 2, 92}, {10, 4, 4}, {7, 2, 30}, {13, 9, 79}, {8, 15, 41}, {11, 8, 18}, {15, 2, 4}, {12, 14, 88}, {12, 6, 9}, {12, 9, 44}, {1, 6, 87}, {15, 14, 42}, {4, 9, 41}, {7, 15, 90}, {4, 1, 84}, {7, 11, 9}, {3, 11, 75}, {5, 9, 2}, {2, 11, 96}, {12, 5, 89}, {6, 15, 25}, {5, 13, 7}, {15, 5, 32}, {13, 5, 84}, {7, 5, 9}, {15, 3, 14}, {12, 13, 4}, {5, 3, 73}, {6, 9, 85}, {6, 10, 29}, {1, 8, 24}, {12, 3, 85}, {4, 3, 60}, {1, 13, 6}, {1, 5, 58}, {2, 3, 29}, {14, 5, 67}, {13, 15, 70}, {5, 14, 94}, {15, 1, 95}, {3, 1, 17}, {10, 2, 6}, {11, 10, 44}, {9, 14, 62}, {4, 11, 32}, {15, 13, 48}, {2, 10, 77}, {3, 13, 90}, {5, 7, 68}, {10, 6, 78}, {3, 6, 95}, {10, 12, 68}, {13, 6, 73}, {10, 1, 8}, {10, 7, 18}, {10, 5, 64}, {5, 1, 55}, {13, 7, 90}, {1, 9, 67}, {3, 12, 76}, {14, 10, 22}, {12, 8, 83}, {4, 7, 76}, {8, 13, 25}, {5, 6, 57}, {13, 3, 90}, {6, 2, 96}, {11, 14, 61}, {12, 1, 94}, {12, 15, 12}, {4, 8, 88}, {4, 14, 27}, {7, 4, 25}, {3, 9, 57}, {2, 15, 90}, {1, 12, 85}, {12, 11, 44}, {5, 10, 13}, {5, 12, 96}, {14, 7, 24}, {14, 12, 98}, {10, 9, 36}, {15, 6, 17}, {8, 10, 11}, {2, 13, 5}, {10, 3, 78}, {6, 13, 11}, {5, 15, 34}, {12, 10, 12}, {9, 1, 68}, {10, 13, 1}, {7, 13, 86}, {1, 7, 62}, {2, 14, 53}, {8, 14, 75}, {2, 6, 49}, {10, 15, 83}, {7, 8, 88}, {6, 1, 87}, {8, 1, 38}, {8, 11, 73}, {3, 15, 1}, {3, 8, 93}, {2, 8, 26}, {4, 5, 26}, {3, 4, 58}, {7, 1, 55}, {7, 3, 84}, {5, 8, 97}, {12, 4, 42}, {6, 3, 71}, {6, 7, 48}, {15, 12, 3}, {1, 15, 30}, {10, 8, 11}, {2, 9, 49}, {6, 14, 95}, {3, 10, 68}, {6, 4, 14}, {11, 12, 29}, {1, 10, 93}, {8, 5, 55}, {12, 2, 86}, {3, 5, 26}, {15, 9, 12}},
			n:     15,
			k:     11,
			want:  38,
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.times, v.n, v.k); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})

	}
}

// WALK 走不全所有路径 --
func networkDelayTimeMostLong(times [][]int, n int, k int) int {

	// 距离K的最短距离
	leaveK := make([]int, n+1)
	for i := 1; i < len(leaveK); i++ {
		if i == k {
			leaveK[i] = 0
		} else {
			leaveK[i] = -1
		}
	}

	// i(第i个元素) -> []int （times的索引）
	toPoint := make([][]int, n+1)
	for idx, road := range times {
		toPoint[road[0]] = append(toPoint[road[0]], idx)
	}

	alreadyDone := make(map[int]bool)
	// var walk func(startPoint int)
	walk := func(startPoint int) int {
		idxList := toPoint[startPoint]
		alreadyDone[startPoint] = true

		distanceMinIdx := 0
		distanceMin := math.MaxInt

		for _, v := range idxList {

			// log.Printf("road--- %v \n", times[v])
			if _, ok := alreadyDone[times[v][1]]; !ok {
				// 没有来过的点
				if leaveK[startPoint] != -1 {
					leaveK[times[v][1]] = leaveK[startPoint] + times[v][2]
				}
			} else {
				// 如果来过的点
				// 对比取代最小的值
				leaveK[times[v][1]] = min(leaveK[times[v][1]], leaveK[startPoint]+times[v][2])
			}
			distanceMin = min(distanceMin, leaveK[times[v][1]])
		}
		return distanceMinIdx
	}
	walk(k)
	for i := 1; i < n+1; i++ {
	}

	log.Printf("leaveK %v \n", leaveK)
	if len(alreadyDone) < n {
		return -1
	}
	maxCost := 0
	for _, v := range leaveK {
		maxCost = max(maxCost, v)
	}

	return maxCost
}

func networkDelayTimeMockTime(times [][]int, n int, k int) int {

	alreadyDone := make(map[int]bool)

	// i(第i个元素) -> []int （times的索引）
	toPoint := make([][]int, n+1)
	roadList := make([][]int, 0)
	for idx, road := range times {
		toPoint[road[0]] = append(toPoint[road[0]], idx)
		if road[0] == k {
			roadList = append(roadList, road)
		}
	}
	alreadyDone[k] = true

	// 对时间进行模拟
	cost := 0
	for {
		// 便利完了
		if len(alreadyDone) == n {
			break
		}

		// 取出一个最小值
		minSpend := math.MaxInt
		for i := range roadList {
			minSpend = min(minSpend, roadList[i][2])
		}
		log.Printf("minSpend %d", minSpend)

		//操作路
		for i := range roadList {
			roadList[i][2] -= minSpend
		}

		tmpRoad := make([][]int, 0)
		for i := len(roadList) - 1; i >= 0; i-- {
			// 路走完了
			if roadList[i][2] == 0 {
				// 标记到达
				alreadyDone[roadList[i][1]] = true
				// 加入待走的路
				for _, roadIdx := range toPoint[roadList[i][1]] {
					// 没到达的地儿再去加
					if _, ok := alreadyDone[times[roadIdx][1]]; !ok {
						tmpRoad = append(tmpRoad, times[roadIdx])
					}
				}
				roadList = append(roadList[:i], roadList[i+1:]...)
			}
		}
		for _, road := range tmpRoad {
			newRoad := make([]int, len(road))
			copy(newRoad, road)
			roadList = append(roadList, newRoad)
		}
		// 这里  roadList 数组 再 append  下面 tmpRoad 的时 会引用同一个片的地址  导致上面   roadList[i][2] -= minSpend 会执行两次 ，这里要怎么处理这里引用同一个地址减两次的情况
		// roadList = append(roadList, tmpRoad...)

		cost += minSpend

		// log.Printf("minSpend %d \n", minSpend)
		// 没有路走了
		if len(roadList) == 0 {
			break
		}

	}

	if len(alreadyDone) < n {
		return -1
	}

	return cost
}

// 广度便利BFS  -  这里好像没法找到最优的
func networkDelayTime(times [][]int, n int, k int) int {

	alreadyDone := make(map[int]bool)

	// i(第i个元素) -> []int （times的索引）
	toPoint := make([][]int, n+1)
	for idx, road := range times {
		toPoint[road[0]] = append(toPoint[road[0]], idx)
	}

	log.Printf("--toPoint- %v\n", toPoint)
	q := make([]int, 0)
	alreadyDone[k] = true
	q = append(q, toPoint[k]...)

	cost := 0
	for {

		if len(alreadyDone) == n {
			break
		}

		maxCost, sub := 0, 0
		for _, v := range q {
			subCost := 0
			if _, ok := alreadyDone[times[v][1]]; !ok {
				subCost += times[v][2]

				alreadyDone[times[v][1]] = true
				q = append(q, toPoint[times[v][1]]...)
			}
			maxCost = max(maxCost, subCost)
			sub++
		}
		cost += maxCost
		log.Printf("--q- %v\n", q)
		q = q[sub:]

		if len(q) == 0 {
			break
		}

	}

	if len(alreadyDone) < n {
		return -1
	}

	return cost
}

/*
前言
本题需要用到单源最短路径算法 Dijkstra，现在让我们回顾该算法，其主要思想是贪心。

将所有节点分成两类：已确定从起点到当前点的最短路长度的节点，以及未确定从起点到当前点的最短路长度的节点（下面简称「未确定节点」和「已确定节点」）。

每次从「未确定节点」中取一个与起点距离最短的点，将它归类为「已确定节点」，并用它「更新」从起点到其他所有「未确定节点」的距离。直到所有点都被归类为「已确定节点」。

用节点 A「更新」节点 B 的意思是，用起点到节点 A 的最短路长度加上从节点 A 到节点 B 的边的长度，去比较起点到节点 B 的最短路长度，如果前者小于后者，就用前者更新后者。这种操作也被叫做「松弛」。

这里暗含的信息是：每次选择「未确定节点」时，起点到它的最短路径的长度可以被确定。

可以这样理解，因为我们已经用了每一个「已确定节点」更新过了当前节点，无需再次更新（因为一个点不能多次到达）。而当前节点已经是所有「未确定节点」中与起点距离最短的点，不可能被其它「未确定节点」更新。所以当前节点可以被归类为「已确定节点」。

方法一：Dijkstra 算法
根据题意，从节点 k 发出的信号，到达节点 x 的时间就是节点 k 到节点 x 的最短路的长度。因此我们需要求出节点 k 到其余所有点的最短路，其中的最大值就是答案。若存在从 k 出发无法到达的点，则返回 −1。

下面的代码将节点编号减小了 1，从而使节点编号位于 [0,n−1] 范围。

作者：力扣官方题解
链接：https://leetcode.cn/problems/network-delay-time/solutions/909575/wang-luo-yan-chi-shi-jian-by-leetcode-so-6phc/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func networkDelayTimeLeetCode1(times [][]int, n, k int) (ans int) {
	const inf = math.MaxInt64 / 2
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, t := range times {
		x, y := t[0]-1, t[1]-1
		g[x][y] = t[2]
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[k-1] = 0
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		x := -1
		for y, u := range used {
			if !u && (x == -1 || dist[y] < dist[x]) {
				x = y
			}
		}
		used[x] = true
		for y, time := range g[x] {
			dist[y] = min(dist[y], dist[x]+time)
		}
	}

	for _, d := range dist {
		if d == inf {
			return -1
		}
		ans = max(ans, d)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func networkDelayTimeLeetCode2(times [][]int, n, k int) (ans int) {
	type edge struct{ to, time int }
	g := make([][]edge, n)
	for _, t := range times {
		x, y := t[0]-1, t[1]-1
		g[x] = append(g[x], edge{y, t[2]})
	}

	const inf int = math.MaxInt64 / 2
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[k-1] = 0
	h := &hp{{0, k - 1}}
	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		x := p.x
		if dist[x] < p.d {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			if d := dist[x] + e.time; d < dist[y] {
				dist[y] = d
				heap.Push(h, pair{d, y})
			}
		}
	}

	for _, d := range dist {
		if d == inf {
			return -1
		}
		ans = max(ans, d)
	}
	return
}

type pair struct{ d, x int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
