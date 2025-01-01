package deferre

import (
	"errors"
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {

	// 异常情况
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer t.Close()
	}

	defer fmt.Println(" 异常情况")
	// 正常情况
	ts2 := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts2 {
		defer Close(t)
	}

	defer fmt.Println(" 正常情况")
	// 更正 上面的异常

	ts3 := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts3 {
		t2 := t
		defer t2.Close()
	}

	defer fmt.Println(" 更正 上面的异常")

	//
	fmt.Println("=======================")
	/*

				Each time a "defer" statement executes,
				the function value and parameters to the call are evaluated as usualand saved anew but the actual function is not invoked.

		这句话。可以得出下面的结论：

		defer后面的语句在执行的时候，
		函数调用的参数会被保存起来，但是不执行。
		也就是复制了一份。
		但是并没有说struct这里的this指针如何处理，
		通过这个例子可以看出go语言并没有把这个明确写出来的this指针当作参数来看待。

				   foo(2, 0)
	*/
	foo(2, 0)
	fmt.Println("=======================")
	//

}

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

func Close(t Test) {
	t.Close()
}

func foo(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()
	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	i = a / b
	/*

		third defer err divided by zero!
		second defer err <nil>
		first defer err <nil>
	*/
	return
}
