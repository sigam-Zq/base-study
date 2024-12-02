package main

import "fmt"

func main() {

	a := map[int]bool{1: true}

	fmt.Println(a[2])

	b := map[int]int{1: 2}
	fmt.Println(b[2])
	c := map[int]string{1: "111"}
	fmt.Println(c[2])
}
