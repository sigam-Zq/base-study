package main

import "log"

// 3270. 求出数字答案

func main() {
	log.Println(generateKey(987, 879, 798))
}

func generateKey(num1 int, num2 int, num3 int) int {
	ans := 0

	for i := 1000; i >= 1; i /= 10 {
		minNum := 10
		for _, v := range []int{num1, num2, num3} {
			minNum = min(minNum, (v/i)%10)
		}
		ans += i * minNum
	}

	return ans
}

func generateKeyLeetCode(num1 int, num2 int, num3 int) int {
	key := 0
	for p := 1; num1 > 0 && num2 > 0 && num3 > 0; p *= 10 {
		key += min(num1%10, min(num2%10, num3%10)) * p
		num1, num2, num3 = num1/10, num2/10, num3/10
	}
	return key
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/find-the-key-of-the-numbers/solutions/3036325/qiu-chu-shu-zi-da-an-by-leetcode-solutio-84cv/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
