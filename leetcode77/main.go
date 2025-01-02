package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

//3280. 将日期转换为二进制表示

func main() {
	// log.Println(convertDateToBinary("2020-10-10")) // 10100-1010-1010

	log.Println(binary(20)) // 10100-1010-1010
	// log.Println(byte(1))       // 2025/01/01 10:47:26 1
	// log.Println(byte('0' + 1)) // 2025/01/01 10:47:26 49
}

func convertDateToBinary(date string) string {
	ans := make([]string, 0)
	dateList := strings.Split(date, "-")
	for _, num := range dateList {
		n, _ := strconv.Atoi(num)
		ans = append(ans, strconv.FormatInt(int64(n), 2))

	}

	return strings.Join(ans, "-")
}

func binary(x int) string {
	var s []byte
	for ; x != 0; x >>= 1 {
		s = append(s, '0'+byte(x&1))
		fmt.Println(s)
	}
	log.Println(string(s))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func convertDateToBinaryLeetCode(date string) string {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	return binary(year) + "-" + binary(month) + "-" + binary(day)
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/convert-date-to-binary/solutions/3030637/jiang-ri-qi-zhuan-huan-wei-er-jin-zhi-bi-nhll/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
