package leetcode55

import (
	"log"
	"math"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {

	// 设置内存限制为100MB   这个没效果   最后使用的 ulimit -v 1000000 限制 1G 内存使用
	// limit := 100 * 1024 * 1024
	// if retu := debug.SetMemoryLimit(int64(limit)); retu != 0 {
	// 	fmt.Println("设置前内存 :", retu)
	// 	// return
	// }
	// fmt.Println("内存限制设置为100MB")
	for i, v := range []struct {
		f      func(int, int, int, int) float64
		n      int
		k      int
		column int
		row    int
		want   float64
	}{
		{
			f:      knightProbabilityRaw,
			n:      3,
			k:      2,
			column: 0,
			row:    0,
			want:   0.0625,
		},
		{
			f:      knightProbabilityRaw,
			n:      1,
			k:      0,
			column: 0,
			row:    0,
			want:   1,
		},
		{
			f:      knightProbabilityProgress,
			n:      8,
			k:      30,
			column: 6,
			row:    4,
			want:   1,
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.n, v.k, v.column, v.row); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})
	}
}

// 代表移动的八个方向
var direction = [][]int{{1, 2}, {-1, 2}, {-1, -2}, {1, -2}, {2, 1}, {2, -1}, {-2, -1}, {-2, 1}}

// 有内存爆炸风险方法
func knightProbabilityRaw(n int, k int, row int, column int) float64 {
	// 最后到达的终点列表
	goalList := [][]int{}

	intermediaryList := [][]int{{row, column}}
	for i := 0; i < k; i++ {
		tmp := make([][]int, 0)
		for _, mid := range intermediaryList {
			if mid[0] < n && mid[0] >= 0 && mid[1] < n && mid[1] >= 0 {
				for _, v := range direction {
					tmp = append(tmp, []int{mid[0] + v[0], mid[1] + v[1]})
				}
			}
		}
		if i < k-1 {
			intermediaryList = tmp
		} else {
			goalList = tmp
		}
	}
	interNum := 0
	for _, v := range goalList {
		if v[0] < n && v[0] >= 0 && v[1] < n && v[1] >= 0 {
			log.Printf(" inter %v ", v)
			interNum++
		}
	}
	log.Printf(" inter len %d \n", interNum)
	log.Printf(" goalList %v len %d \n", goalList, len(goalList))

	if interNum == len(goalList) {
		return float64(1)
	}

	res := float64(interNum) / math.Pow(8, float64(k))
	return res
}

var dirs = []struct{ i, j int }{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}}

func knightProbabilityLeetCode(n, k, row, column int) float64 {
	dp := make([][][]float64, k+1)
	for step := range dp {
		dp[step] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[step][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				if step == 0 {
					dp[step][i][j] = 1
				} else {
					for _, d := range dirs {
						if x, y := i+d.i, j+d.j; 0 <= x && x < n && 0 <= y && y < n {
							dp[step][i][j] += dp[step-1][x][y] / 8
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/knight-probability-in-chessboard/solutions/1264717/qi-shi-zai-qi-pan-shang-de-gai-lu-by-lee-2qhk/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func knightProbabilityChatGPT(n int, k int, row int, column int) float64 {
	directions := [][]int{{1, 2}, {-1, 2}, {-1, -2}, {1, -2}, {2, 1}, {2, -1}, {-2, -1}, {-2, 1}}

	// 当前和下一步的概率分布
	current := make([][]float64, n)
	for i := range current {
		current[i] = make([]float64, n)
	}
	current[row][column] = 1.0

	for step := 0; step < k; step++ {
		next := make([][]float64, n)
		for i := range next {
			next[i] = make([]float64, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if current[i][j] > 0 {
					for _, dir := range directions {
						ni, nj := i+dir[0], j+dir[1]
						if ni >= 0 && ni < n && nj >= 0 && nj < n {
							next[ni][nj] += current[i][j] / 8.0
						}
					}
				}
			}
		}
		current = next
	}

	// 计算总概率
	totalProbability := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			totalProbability += current[i][j]
		}
	}
	return totalProbability
}

/*
方法一：动态规划
思路

一个骑士有 8 种可能的走法，骑士会从中以等概率随机选择一种。部分走法可能会让骑士离开棋盘，另外的走法则会让骑士移动到棋盘的其他位置，并且剩余的移动次数会减少 1。

定义 dp[step][i][j] 表示骑士从棋盘上的点 (i,j) 出发，走了 step 步时仍然留在棋盘上的概率。特别地，当点 (i,j) 不在棋盘上时，dp[step][i][j]=0；当点 (i,j) 在棋盘上且 step=0 时，dp[step][i][j]=1。对于其他情况，dp[step][i][j]=
8
1 ×
di,dj
∑
​
 dp[step−1][i+di][j+dj]。其中 (di,dj) 表示走法对坐标的偏移量，具体为 (−2,−1),(−2,1),(2,−1),(2,1),(−1,−2),(−1,2),(1,−2),(1,2) 共 8 种。

作者：力扣官方题解
链接：https://leetcode.cn/problems/knight-probability-in-chessboard/solutions/1264717/qi-shi-zai-qi-pan-shang-de-gai-lu-by-lee-2qhk/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/*

### **分析原因和空间复杂度公式**

1. **内存爆炸原因**:
   - 在每次迭代中，`tmp` 中的状态数量会按指数级增长。对于第 `i` 次迭代，理论上可能有 \(8^i\) 个中间状态。这种增长是由于每个棋盘位置最多会生成 8 个可能的新位置。
   - 因为没有对 `tmp` 中的位置去重，也没有及时剔除越界位置，导致存储了大量冗余信息。
   - 当 `k` 较大时（如 \(k=30\)），存储的状态数接近 \(8^{30}\)，这是无法承受的。

2. **空间复杂度公式**:
   - **理论最大空间复杂度**：如果每次都存储所有可能的状态，空间复杂度为 \(O(8^k)\)。
   - **优化后的空间复杂度**：只需存储当前步和下一步的状态，可以优化到 \(O(n^2)\)，其中 \(n^2\) 是棋盘上最多可能的有效位置数。

---

### **如何优化算法**
使用动态规划避免指数级增长：

- **核心思想**：记录每个位置在某步的概率，不需要存储所有路径。
- 通过二维数组 `dp` 表示棋盘上的概率分布。
- 使用两个二维数组交替更新来避免内存爆炸。

代码优化如下：
```go
func knightProbability(n int, k int, row int, column int) float64 {
	directions := [][]int{{1, 2}, {-1, 2}, {-1, -2}, {1, -2}, {2, 1}, {2, -1}, {-2, -1}, {-2, 1}}

	// 当前和下一步的概率分布
	current := make([][]float64, n)
	for i := range current {
		current[i] = make([]float64, n)
	}
	current[row][column] = 1.0

	for step := 0; step < k; step++ {
		next := make([][]float64, n)
		for i := range next {
			next[i] = make([]float64, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if current[i][j] > 0 {
					for _, dir := range directions {
						ni, nj := i+dir[0], j+dir[1]
						if ni >= 0 && ni < n && nj >= 0 && nj < n {
							next[ni][nj] += current[i][j] / 8.0
						}
					}
				}
			}
		}
		current = next
	}

	// 计算总概率
	totalProbability := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			totalProbability += current[i][j]
		}
	}
	return totalProbability
}
```

**优化后的空间复杂度**：\
通过交替使用两个二维数组，空间复杂度降低到 \(O(n^2)\)。\
**时间复杂度**：每步计算需要枚举棋盘上的每个格子和 8 个方向，整体复杂度为 \(O(k \cdot n^2)\)。

---

### **如何在 Go 的测试程序中限制内存**

Go 提供了 `testing` 和 `runtime` 包，可以通过设置内存限制和检测内存使用。

1. **使用 `runtime.MemProfileRate` 控制内存分配分析**:
   ```go
   import (
       "runtime"
       "testing"
   )

   func TestKnightProbabilityMemory(t *testing.T) {
       runtime.MemProfileRate = 1 // 启用详细内存跟踪

       defer func() {
           var memStats runtime.MemStats
           runtime.ReadMemStats(&memStats)
           t.Logf("Memory Used: %v KB", memStats.Alloc/1024)
       }()

       result := knightProbability(8, 30, 4, 6)
       t.Logf("Result: %f", result)
   }
   ```

2. **限制最大内存**:
   使用 `testing` 包运行时，结合 `ulimit` 或内存限制工具执行测试：
   - 在 Linux/macOS 上，可以运行：
     ```bash
     ulimit -v 500000 # 设置最大内存为 500MB
     go test -v -run TestKnightProbabilityMemory
     ```

3. **通过 context 设置超时**:
   如果算法内存过高导致超时，可以设置 `timeout`：
   ```go
   import (
       "context"
       "time"
   )

   func TestKnightProbabilityTimeout(t *testing.T) {
       ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
       defer cancel()

       done := make(chan bool)
       go func() {
           knightProbability(8, 30, 4, 6)
           done <- true
       }()

       select {
       case <-ctx.Done():
           t.Fatal("Test timed out")
       case <-done:
           t.Log("Test completed")
       }
   }
   ```
*/

func knightProbabilityProgress(n int, k int, row int, column int) float64 {

	// current[i][j] 标识 i,j 位置起步到棋盘上的上的概率
	current := make([][]float64, n)
	for i := range current {
		current[i] = make([]float64, n)
	}
	current[row][column] = 1.0

	for step := 0; step < k; step++ {
		next := make([][]float64, n)
		for i := range next {
			next[i] = make([]float64, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if current[i][j] > 0 {
					for _, v := range direction {
						x, y := i+v[0], j+v[1]
						if x < n && x >= 0 && y < n && y >= 0 {
							next[x][y] += current[i][j] / 8.0
						}
					}
				}
			}
		}

		current = next
	}
	log.Printf("current  %v \n", current)
	var res float64

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += current[i][j]
		}
	}
	return res
}
