package main

func main() {

	a := 1

	defer func() {
		b := a + 1
		println("defer 1")
		println(b)
	}()

	defer func(a int) {
		b := a + 1
		println("defer 2")
		println(b)
	}(a)

	a = 99

}
