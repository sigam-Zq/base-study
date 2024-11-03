package benchtest

func fast() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 1000; j++ {
			for k := 0; k < 10000; k++ {

			}
		}
	}
}

func slow() {
	for i := 0; i < 10000; i++ {
		for j := 0; j < 1000; j++ {
			for k := 0; k < 100; k++ {

			}
		}
	}
}

// 斐波那契数列
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func sum(a, b int) int {
	return a + b
}
