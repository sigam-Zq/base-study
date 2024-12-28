package main

import (
	"fmt"
)

type abb struct {
	Name string
}

func main() {
	var user struct {
		Name string
		Age  int
	}
	// 这里 分配内存了么
	user.Name = "pprof.cn"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	// var  这里 会分配 内存阿
	var abb abb
	abb.Name = "aaa"
	fmt.Printf("%#v\n", abb)

}
