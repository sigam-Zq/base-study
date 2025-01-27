package leetcode98

import (
	"math"
	"strconv"
	"testing"
)

// 45. 跳跃游戏 II

/*
给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

0 <= j <= nums[i]
i + j < n
返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。



示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
示例 2:

输入: nums = [2,3,0,1,4]
输出: 2


提示:

1 <= nums.length <= 104
0 <= nums[i] <= 1000
题目保证可以到达 nums[n-1]
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func([]int) int
		nums []int
		want int
	}{
		{
			f:    jump,
			nums: []int{2, 3, 1, 1, 4},
			want: 2,
		}, {
			f:    jump,
			nums: []int{2, 3, 0, 1, 4},
			want: 2,
		}} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := v.f(v.nums); got != v.want {
				t.Errorf("got %d   want %d  \n", got, v.want)
			}
		})

	}

}

func jump(nums []int) int {

	n := len(nums)

	dp := make([]int, n)
	dp[0] = 0

	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt
	}

	for i := 1; i < n; i++ {

		for j := i - 1; j >= 0; j-- {
			if nums[j] >= i-j {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}

	}

	return dp[n-1]
}

func jumpLeetCode1(nums []int) int {
	position := len(nums) - 1
	steps := 0
	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] >= position {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/jump-game-ii/solutions/230241/tiao-yue-you-xi-ii-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func jumpLeetCode2(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/jump-game-ii/solutions/230241/tiao-yue-you-xi-ii-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
