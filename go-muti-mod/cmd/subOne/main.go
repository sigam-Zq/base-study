package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
	"zqsrh.com/one-submod/util"
)

func main() {
	fmt.Println(reverse.String("Hello"))

	util.NowVersion()
}
