package main

import "fmt"

func main() {

	a := map[int]bool{1: true}

	fmt.Println(a[2])

	b := map[int]int{1: 2}
	fmt.Println(b[2])
	c := map[int]string{1: "111"}
	fmt.Println(c[2])

	// map 本身是无序的
	noSortMap := map[string]string{"sfafsfd": "sfafsfd", "fdgsgsdgds": "fdgsgsdgds", "dfdghfg": "dfdghfg", "sdghrasg": "sdghrasg"}
	for k, v := range noSortMap {
		fmt.Printf("%s---%s \n", k, v)
	}

}
