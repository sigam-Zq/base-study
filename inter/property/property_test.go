package property

import "testing"

// 类型的并集
type foo interface {
	~string | ~int
}

//  string 和 int 的交集
type bar interface {
	~string
	~int
}

// wrong signature for TestXxx, test functions cannot have type parameters
func TestXxx[T ReadWriter](t *testing.T) {
	var f string
	abb := "abb"
	f = abb
	_ = f

	// 验证实现
	var a T
	a = (StringReadWriter)("nil")
	// a = (ReadWriter)("nil")

	a = (BytesReadWriter)("nil")
	// a = (ReadWriter)("nil")
	_ = a
}

type ReadWriter interface {
	~string | ~[]rune

	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

// 类型 StringReadWriter 实现了接口 Readwriter
type StringReadWriter string

func (s StringReadWriter) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (s StringReadWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

//  类型BytesReadWriter 没有实现接口 Readwriter
type BytesReadWriter []byte

func (s BytesReadWriter) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (s BytesReadWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
