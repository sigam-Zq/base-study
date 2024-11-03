package main

import "fmt"

var g = 100

func main() {
	// i := f1()
	// fmt.Printf("main: i  = %d \n main: g  = %d \n", i, g)
	/**
	f: g  = 100
	main: i  = 100
	 main: g  = 200
	 **/

	// i := f2()
	// fmt.Printf("main: i  = %d \n main: g  = %d \n", i, g)
	/**
		f: g  = 100
	main: i  = 100
	 main: g  = 200

	 ? 这里是在return 之后执行了defer
	 g = 100
	 r = g
	 g = 200 (defer)
	 return r
			// **/

	i := f3()
	fmt.Printf("main: i  = %d \n main: g  = %d \n", i, g)
	/**
		go  build -o run && ./run
	f: r  = 100
	main: i  = 200
	 main: g  = 100

	 g=100
	 r=g
	 r=0
	 r=200
	 return
		**/
}

//   将返回值保存在 栈上  -- defer 执行 --- 函数返回
func f1() int {
	// var r int

	defer func() {
		g = 200
	}()
	fmt.Printf("f: g  = %d \n", g)
	return g
}
func f2() (r int) {

	defer func() {
		g = 200
	}()

	fmt.Printf("f: g  = %d \n", g)
	return g
}

func f3() (r int) {
	r = g
	defer func() {
		r = 200
	}()

	r = 0
	fmt.Printf("f: r  = %d \n", r)
	return r
}
