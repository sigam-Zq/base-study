package main

import (
	"fmt"

	"zqsrh.com/go.mutilmod/util"
	subUtil "zqsrh.com/one-submod/util"
)

func main() {
	fmt.Println("subTwo")
	util.NowVersion()
	subUtil.NowVersion()
}
