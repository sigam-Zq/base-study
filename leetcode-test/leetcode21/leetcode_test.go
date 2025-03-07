package leetcode

import (
	"log"
	"testing"
)

/*
3175. 找到连续赢 K 场比赛的第一位玩家

有 n 位玩家在进行比赛，玩家编号依次为 0 到 n - 1 。

给你一个长度为 n 的整数数组 skills 和一个 正 整数 k ，其中 skills[i] 是第 i 位玩家的技能等级。skills 中所有整数 互不相同 。

所有玩家从编号 0 到 n - 1 排成一列。

比赛进行方式如下：

队列中最前面两名玩家进行一场比赛，技能等级 更高 的玩家胜出。
比赛后，获胜者保持在队列的开头，而失败者排到队列的末尾。
这个比赛的赢家是 第一位连续 赢下 k 场比赛的玩家。

请你返回这个比赛的赢家编号。



示例 1：

输入：skills = [4,2,6,3,9], k = 2

输出：2

解释：

一开始，队列里的玩家为 [0,1,2,3,4] 。比赛过程如下：

玩家 0 和 1 进行一场比赛，玩家 0 的技能等级高于玩家 1 ，玩家 0 胜出，队列变为 [0,2,3,4,1] 。
玩家 0 和 2 进行一场比赛，玩家 2 的技能等级高于玩家 0 ，玩家 2 胜出，队列变为 [2,3,4,1,0] 。
玩家 2 和 3 进行一场比赛，玩家 2 的技能等级高于玩家 3 ，玩家 2 胜出，队列变为 [2,4,1,0,3] 。
玩家 2 连续赢了 k = 2 场比赛，所以赢家是玩家 2 。

示例 2：

输入：skills = [2,5,4], k = 3

输出：1

解释：

一开始，队列里的玩家为 [0,1,2] 。比赛过程如下：

玩家 0 和 1 进行一场比赛，玩家 1 的技能等级高于玩家 0 ，玩家 1 胜出，队列变为 [1,2,0] 。
玩家 1 和 2 进行一场比赛，玩家 1 的技能等级高于玩家 2 ，玩家 1 胜出，队列变为 [1,0,2] 。
玩家 1 和 0 进行一场比赛，玩家 1 的技能等级高于玩家 0 ，玩家 1 胜出，队列变为 [1,2,0] 。
玩家 1 连续赢了 k = 3 场比赛，所以赢家是玩家 1 。



提示：

n == skills.length
2 <= n <= 105
1 <= k <= 109
1 <= skills[i] <= 106
skills 中的整数互不相同。
*/

func TestMain(t *testing.T) {
	args := []struct {
		name   string
		f      func(skills []int, k int) int
		skills []int
		k      int
		want   int
	}{
		{
			name:   "oneTest",
			f:      findWinningPlayer,
			skills: []int{4, 2, 6, 3, 9},
			k:      2,
			want:   2,
		},
		{
			name:   "twoTest",
			f:      findWinningPlayer,
			skills: []int{2, 5, 4},
			k:      3,
			want:   1,
		},
		{
			name:   "threeTest",
			f:      findWinningPlayer,
			skills: []int{7, 11},
			k:      2,
			want:   1,
		},
		{
			name:   "fourTest",
			f:      findWinningPlayer,
			skills: []int{16, 4, 7, 17},
			k:      562084119,
			want:   1,
		},
	}

	for _, v := range args {
		t.Run(v.name, func(t *testing.T) {
			if got := v.f(v.skills, v.k); v.want != got {
				t.Errorf(" got %d  want %d \n", got, v.want)
			}
		})
	}

}

func findWinningPlayer(skills []int, k int) int {
	countVictories := 0
	for i, j := 0, 1; ; {

		log.Printf(" i %d , j %d, countVictories %d \n", i, j, countVictories)
		if skills[i] > skills[j] {
			countVictories++
			j++
		} else {
			countVictories = 1
			i = j
			j = i + 1
		}

		if countVictories == k {
			return i
		}

		// i 越界后给一个取余操作 防止越界
		// 保持一个环形结构

		if j >= len(skills) {
			// 当前的i 必定是 全局最大值 直接返回i 即可
			return i
			// j %= len(skills)
		}
		if i >= len(skills) {
			i = i % len(skills)
		}

	}

	// return 0
}
