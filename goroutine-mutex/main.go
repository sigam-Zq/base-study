package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	done := make(chan bool, 1)
	var mu sync.Mutex

	g1Count := 0
	g2Count := 0
	// goroutine 1
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				time.Sleep(100 * time.Microsecond)
				g1Count++
				mu.Unlock()
			}
		}
	}()
	// time.Sleep(110 * time.Microsecond)

	// goroutine 2
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		g2Count++
		//do sometinring
		mu.Unlock()
	}

	done <- true
	fmt.Printf("g1 count: %d\n", g1Count)
	fmt.Printf("g2 count: %d\n", g2Count)
}
