package main

import (
	"fmt"
	"os"
)

// os.Exit(1)  这里defer 不会被调用到
func main() {

	defer fmt.Println("---")

	os.Exit(1)
}
