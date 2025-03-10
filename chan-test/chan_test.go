package chantest

import (
	"fmt"
	"sync"
	"testing"
)

func TestChan(t *testing.T) {

	var w sync.WaitGroup
	// var w2 sync.WaitGroup

	noBufCh := make(chan int)

	Buf1Ch := make(chan int, 1)

	fmt.Printf(" len(noBufCh) %d cap(noBufCh) %d len(Buf1Ch) %d  cap(Buf1Ch) %d \n", len(noBufCh), cap(noBufCh), len(Buf1Ch), cap(Buf1Ch))

	w.Add(1)
	go func() {
		noBufCh <- 1
		w.Done()
		fmt.Println(" noBufCh 1")

	}()

	go func() {
		Buf1Ch <- 1
		fmt.Println(" Buf1Ch 1")
	}()

	go func() {
		Buf1Ch <- 1
		fmt.Println(" Buf1Ch 2")
	}()

	// Gosched() 函数用于让出 CPU 时间，让其他 Goroutine 拥有运行的机会。其原理是将当前 Goroutine 放回到队列中，等待下一次调度。
	// runtime.Gosched()
	// 关闭的作用是告诉接收者该通道再无新数据发送
	// w2.Add(1)
	// go func() {
	// 	w.Wait()
	// 	fmt.Println(" noBufCh close ")
	// 	close(noBufCh)
	// 	// w2.Done()
	// }()
	// close(Buf1Ch)
	fmt.Printf(" len(noBufCh) %d cap(noBufCh) %d len(Buf1Ch) %d  cap(Buf1Ch) %d \n", len(noBufCh), cap(noBufCh), len(Buf1Ch), cap(Buf1Ch))

	if v, ok := <-Buf1Ch; ok {
		fmt.Println(" Buf1Ch out 1")
		fmt.Println(v)
	}
	// w2.Wait()
	if v, ok := <-noBufCh; ok {
		fmt.Println(" noBufCh out 1")
		fmt.Println(v)
	}
	go func() {
		fmt.Println(" noBufCh in 2")
		noBufCh <- 1
	}()
	selectFunc()

	sync.NewCond()

}

func selectFunc() {

	a := make(chan int, 1)
	a <- 1
	b := make(chan int, 1)
	b <- 1

	select {
	case <-a:
		fmt.Println("a")
	case <-b:
		fmt.Println("b")
	default:
		fmt.Println("default")
	}

}
