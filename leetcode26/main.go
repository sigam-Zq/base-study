package main

import "log"

/*
3216. 交换后字典序最小的字符串
给你一个仅由数字组成的字符串 s，在最多交换一次 相邻 且具有相同 奇偶性 的数字后，返回可以得到的
字典序最小的字符串
。

如果两个数字都是奇数或都是偶数，则它们具有相同的奇偶性。例如，5 和 9、2 和 4 奇偶性相同，而 6 和 9 奇偶性不同。



示例 1：

输入： s = "45320"

输出： "43520"

解释：

s[1] == '5' 和 s[2] == '3' 都具有相同的奇偶性，交换它们可以得到字典序最小的字符串。

示例 2：

输入： s = "001"

输出： "001"

解释：

无需进行交换，因为 s 已经是字典序最小的。



提示：

2 <= s.length <= 100
s 仅由数字组成。
*/

func main() {

	// for _, v := range []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'} {
	// 	log.Printf("---- %x , %d \n", v, v)

	// }
	// 2024/10/30 09:26:13 ---- 30 , 48
	// 2024/10/30 09:26:13 ---- 31 , 49
	// 2024/10/30 09:26:13 ---- 32 , 50
	// 2024/10/30 09:26:13 ---- 33 , 51
	// 2024/10/30 09:26:13 ---- 34 , 52
	// 2024/10/30 09:26:13 ---- 35 , 53
	// 2024/10/30 09:26:13 ---- 36 , 54
	// 2024/10/30 09:26:13 ---- 37 , 55
	// 2024/10/30 09:26:13 ---- 38 , 56
	// 2024/10/30 09:26:13 ---- 39 , 57

	log.Println(getSmallestString("001"))
	log.Println(getSmallestString("31"))
}

func getSmallestString(s string) string {
	strBytes := []byte(s)

	left, right := 0, 1
	log.Printf(" strBytes %v \n", strBytes)
	for right < len(s) {

		// v > 0 保证字典序  v%2==0 保证两个数字的差为偶数即为相同奇偶性
		log.Printf("--- strBytes[right] - strBytes[left]   %d \n", strBytes[right]-strBytes[left])
		// 这里不包裹int 会导致　v为 byte 类型 当 49 - 51 时候因为溢出导致的 结果为 254

		// 2024/10/30 09:44:31  strBytes [51 49]
		// 2024/10/30 09:44:31 --- strBytes[right] - strBytes[left]   254
		if v := int(strBytes[left]) - int(strBytes[right]); v > 0 && v%2 == 0 {
			strBytes[left], strBytes[right] = strBytes[right], strBytes[left]
			break
		}

		left++
		right++
	}

	return string(strBytes)
}
