package main

import "fmt"

func main() {

	var map1 = map[chan int]int{make(chan int): 1, make(chan int): 2}

	fmt.Println(map1)
}
