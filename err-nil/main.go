package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {

	var err1 error

	println("err1 == nil")
	println(err1 == nil)

	var err2 *error
	println("err2 == nil")
	println(err2 == nil)

	var err3 *os.PathError
	println("err3 == nil")
	println(err3 == nil)

	var wg sync.WaitGroup
	fmt.Printf("%v\n", wg)

	err4 := foo()
	println("err4 == nil")
	println(err4 == nil) // false   ??因为接口 error 具有动态类型 *os.PathError
	fmt.Printf("%v\n", err4)
}

func foo() error {
	var err *os.PathError

	println("inner err == nil")
	println(err == nil) //true
	fmt.Printf("%v\n", err)

	return err
}
