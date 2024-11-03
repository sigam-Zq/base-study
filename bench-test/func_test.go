package benchtest

import "testing"

func BenchmarkCPUfast(b *testing.B) {

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fast()
	}
}

func BenchmarkCPUsolw(b *testing.B) {

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slow()
	}
}

func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(10)
	}
}

func BenchmarkFib20(b *testing.B) {
	b.ResetTimer() // 重置计时器，忽略前面的准备时间
	for n := 0; n < b.N; n++ {
		fib(20)
	}
}

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum(1, 2)
	}
}
