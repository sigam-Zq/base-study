package closure

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {

	c := a()
	c()
	c()
	c()

	fmt.Println(":======")
	d := a()
	d()
	d()

	f := test()
	f()

	fmt.Println("add :======")
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	// 此时tmp1和tmp2不是一个实体了
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))

	fmt.Println("defer + closure :======")

	var whatever1 [5]struct{}
	// 正常
	for i := range whatever1 {
		defer fmt.Println(i)
	}

	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func test() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

// 外部引用函数参数局部变量
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}
