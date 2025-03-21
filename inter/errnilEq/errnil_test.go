package errnileq

import (
	"os"
	"testing"
)

func TestXxx(t *testing.T) {
	var err1 error
	t.Log(err1 == nil)
	var err2 *error
	t.Log(err2 == nil)
	var err3 *os.PathError
	t.Log(err3 == nil)
	t.Log("========================")
	t.Log("foo() == nil")
	t.Log(foo() == nil)
	t.Log("bar() == nil")
	t.Log(bar() == nil)
	t.Log("foo() == bar()")
	t.Log(foo() == bar())
	t.Log("baz() == nil")
	t.Log(baz() == nil)
	// t.Log((&inter1) == (&inter2))
	t.Log("========================")

	// 接口为 nil 和  接口值为  nil 的区别  （其中nil 内还包含类型信息？）

	var inter1 interface{}
	t.Log("inter1 == nil")
	t.Log(inter1 == nil)

	var inter2 interface {
		Func(string) string
	}
	t.Log("inter2 == nil")
	t.Log(inter2 == nil)
	t.Log("inter1 == inter2")
	t.Log(inter1 == inter2)
	// t.Log("(&inter1) == (&inter2)")
	// t.Log((&inter1) == (&inter2))
	// invalid operation: (&inter1) == (&inter2) (mismatched types *interface{} and *interface{Func(string) string})

	var x error = nil               // 普通赋值
	y := (*os.PathError)(nil)       // 类型转换为指针
	t.Log("error == *os.PathError") //
	t.Log(x == y)                   // true

}

func foo() error {
	var err *os.PathError
	return err
}

func bar() error {
	var err error
	return err
}

func baz() []string {
	var err []string
	return err
}
