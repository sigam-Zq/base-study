package main

import (
	"fmt"
)

/*
label test

(双层循环中才有效)

continue label
break label
*/

func main() {
	// log.Println("--")

label:
	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {
			fmt.Printf("i --- %d  j --- %d \n", i, j)

			if j == 2 {
				// 上下一致-- 这里双层循环的 continue 等价于 break
				continue label
				// break
			}
			// if j == 2 {
			// 	break label
			// }
		}

	}
}
