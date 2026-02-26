package main

import (
	"sync"
	"time"
)

// 针对 range 一个管道，循环什么时候结束

func main() {

	var it Iter[int] = make(Iter[int])
	// it2 := Iter2[int]{1, 2, 34, 3, 5, 6, 7, 8, 9, 10}
	var it2 Iter2[int]
	t := iterator[int]([]int{1, 2, 34, 3, 5, 6, 7, 8, 9, 10})
	it2 = t
	var l sync.WaitGroup
	l.Add(1)
	go func() {

		for v := range it2 {
			println(v)
			println("out2 ...")
		}

		for v := range it {
			println(v)
			println("out ...")
		}
	}()

	go func() {

		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 2)
			it <- i
			println("in ...")
		}
		close(it)
		l.Done()
	}()

	l.Wait()
	println("main exit")
}

type Iter[T any] chan T

type Iter2[T any] <-chan T

func iterator[T any](item []T) chan T {
	ch := make(chan T)
	go func() {
		time.Sleep(100 * time.Millisecond)
		for _, elm := range item {
			ch <- elm
		}
		close(ch)
	}()
	return ch
}
