package main

import "log"

func main() {
	log.Printf("res %d \n", twoEggDrop(6))
}

func twoEggDrop(n int) int {

	// 转化为找寻递增数列的过程
	tmp := n
	// difference := 0
	i := 0
	for tmp > 0 {
		tmp -= i
		log.Printf(" %d ", tmp)
		i++
		log.Printf("i  %d ", i)
		// difference += i
	}
	log.Println("")

	return i - 1
}
