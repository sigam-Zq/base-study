package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("request started at", time.Now())
	defer fmt.Println("request started at- defer", time.Now())
	defer func() { fmt.Println("request started end", time.Now()) }()
	time.Sleep(10 * time.Second)

	fmt.Println("sleep after", time.Now())
}

/**

main.go
z@zdeMacBook-Pro popquiz % go run main.go
2026-05-25 13:14:22.612553 +0800 CST m=+0.000059251
2026-05-25 13:14:22.612696 +0800 CST m=+0.000202751
**/

/*
这里的 defer fmt.Println("request started at- defer", time.Now())
 在到这里那一刻就执行了, 延迟了的仅仅是打印操作而已
*/
