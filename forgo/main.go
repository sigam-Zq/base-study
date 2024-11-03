package main

import (
	"fmt"
	"log"
	"sync"
)

type Abb struct {
	Acc string
}

func main() {

	ac := []Abb{{Acc: "111"}, {Acc: "222"}, {Acc: "3333"}}

	var wg sync.WaitGroup
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	var wg3 sync.WaitGroup

	wg.Add(3)

	log.Println("------------no params")
	for _, v := range ac {
		go func() {
			defer wg.Done()
			fmt.Println(v.Acc)
		}()

	}

	wg.Wait()

	wg1.Add(3)
	log.Println("------------params1")
	for _, v := range ac {
		go func(p Abb) {
			defer wg1.Done()
			fmt.Println(p.Acc)
		}(v)

	}
	wg1.Wait()

	wg2.Add(3)
	log.Println("------------params2")
	for _, v := range ac {
		go func(p string) {
			defer wg2.Done()
			fmt.Println(p)
		}(v.Acc)

	}
	wg2.Wait()

	wg3.Add(3)
	log.Println("------------ 压栈")
	for _, v := range ac {

		func(p2 Abb) {

			go func() {
				defer wg3.Done()
				fmt.Println(p2.Acc)
			}()
		}(v)

	}
	wg3.Wait()

}
