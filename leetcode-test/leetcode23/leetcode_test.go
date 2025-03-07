package leetcode23

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试执行第一步: main开始测试,是在测试之前执行的")
	m.Run()
}

/*
https://leetcode.cn/problems/redundant-connection/description/?envType=daily-question&envId=2024-10-27

树可以看成是一个连通且 无环 的 无向 图。

给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。

请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的那个。



示例 1：



输入: edges = [[1,2], [1,3], [2,3]]
输出: [2,3]
示例 2：



输入: edges = [[1,2], [2,3], [3,4], [1,4], [1,5]]
输出: [1,4]


提示:

n == edges.length
3 <= n <= 1000
edges[i].length == 2
1 <= ai < bi <= edges.length
ai != bi
edges 中无重复元素
给定的图是连通的
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
			f:     findRedundantConnection,
			edges: [][]int{{1, 2}, {1, 3}, {2, 3}},
			want:  []int{2, 3},
		},
		{
			name:  "TWOTest",
			f:     findRedundantConnection,
			edges: [][]int{{9, 10}, {5, 8}, {2, 6}, {1, 5}, {3, 8}, {4, 9}, {8, 10}, {4, 10}, {6, 8}, {7, 9}},
			want:  []int{4, 10},
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

// 关键在于找出可以形成回路的 若干点 然后返回 edges 包含他的最后的一个边
func findRedundantConnection(edges [][]int) []int {
	nodeNumLen := len(edges) + 1
	for {
		prefLen := len(edges)
		n := 0
		nodeNum := make([]int, nodeNumLen)

		for _, v := range edges {
			nodeNum[v[0]]++
			nodeNum[v[1]]++
		}

		for _, v := range edges {
			if nodeNum[v[0]] == 1 || nodeNum[v[1]] == 1 {
				n++
				continue
			}
			edges = append(edges, v)
		}

		log.Printf("check  %d , \n", n)
		if n == 0 {
			break
		}
		log.Printf("edge truncation before  %v \n", edges)
		edges = edges[prefLen:]
		log.Printf("edge truncation after  %v \n", edges)
	}

	return edges[len(edges)-1]
}

func findRedundantConnectionLeetCode(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/redundant-connection/solutions/557616/rong-yu-lian-jie-by-leetcode-solution-pks2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
/*
这个算法的目的是在给定的边列表中找到一条冗余边，使得删除这条边后，剩下的边仍然能够形成一棵树。树的特性是连通且无环，因此如果添加了一条边导致形成环，就需要找到并删除这条边。

### 解析 `find` 和 `union` 函数

1. **`find` 函数**：
   - 这个函数用于查找某个节点的根节点（或代表节点）。它使用了路径压缩的技术来优化查找过程。
   - 当调用 `find(x)` 时，它会返回节点 `x` 的根节点，并在过程中将 `x` 的父节点更新为其根节点，从而加速未来的查找。

2. **`union` 函数**：
   - 这个函数用于合并两个节点的集合。如果两个节点的根节点相同，说明它们已经在同一个集合中，返回 `false` 表示没有合并（即形成了环）。如果它们的根节点不同，则将它们合并，并返回 `true`。

### 举例说明

假设我们有以下边列表：

```go
edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
```

- **初始化**：
  - `parent` 数组初始化为 `[0, 1, 2, 3]`，表示每个节点的父节点是它自己。

- **处理边**：
  1. 处理边 `[1, 2]`：
     - 调用 `union(1, 2)`：
       - `find(1)` 返回 1，`find(2)` 返回 2。
       - 因为根节点不同，合并它们，`parent` 更新为 `[0, 2, 2, 3]`。
       - 返回 `true`，表示成功合并。

  2. 处理边 `[1, 3]`：
     - 调用 `union(1, 3)`：
       - `find(1)` 返回 2（经过路径压缩），`find(3)` 返回 3。
       - 根节点不同，合并它们，`parent` 更新为 `[0, 2, 3, 3]`。
       - 返回 `true`，表示成功合并。

  3. 处理边 `[2, 3]`：
     - 调用 `union(2, 3)`：
       - `find(2)` 返回 3（经过路径压缩），`find(3)` 返回 3。
       - 根节点相同，说明形成了环，返回 `false`。
       - 此时，返回边 `[2, 3]` 作为冗余边。

### 总结

- `find` 函数用于查找节点的根节点，并进行路径压缩以优化后续查找。
- `union` 函数用于合并两个节点的集合，并检查是否形成环。
- 通过遍历边列表，算法能够找到并返回最后一条冗余边。

*/
