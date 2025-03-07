package main

import (
	"fmt"
	"log"
)

/*
3274. 检查棋盘方格颜色是否相同

https://leetcode.cn/problems/check-if-two-chessboard-squares-have-the-same-color/description/?envType=daily-question&envId=2024-12-03

给你两个字符串 coordinate1 和 coordinate2，代表 8 x 8 国际象棋棋盘上的两个方格的坐标。

以下是棋盘的参考图。



如果这两个方格颜色相同，返回 true，否则返回 false。

坐标总是表示有效的棋盘方格。坐标的格式总是先字母（表示列），再数字（表示行）。



示例 1：

输入： coordinate1 = "a1", coordinate2 = "c3"

输出： true

解释：

两个方格均为黑色。

示例 2：

输入： coordinate1 = "a1", coordinate2 = "h3"

输出： false

解释：

方格 "a1" 是黑色，而 "h3" 是白色。

*/

func main() {
	// 位运算计算奇偶性
	fmt.Println(1 & 1)
	fmt.Println(2 & 1)
	fmt.Println(3 & 1)
	fmt.Println(4 & 1)
	fmt.Println(5 & 1)

	log.Println(checkTwoChessboards("a1", "c3"))
	log.Println(checkTwoChessboards("a1", "h3"))
}

// 位运算
func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	x := coordinate1[0] + coordinate1[1]
	y := coordinate2[0] + coordinate2[1]
	return (x & 1) == (y & 1)
}

/*
coordinate1[0] + coordinate1[1] 确定棋盘 从左上到右下的 对角线是第几个

 & 1 是为了判断这个序号的奇偶性

 两个如果奇偶性相同 一定是一个颜色的


*/
