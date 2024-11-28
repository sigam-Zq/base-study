package main

import "fmt"

func main() {

	hanoi(3, 'F', 'A', 'T')
}

func hanoi(n int, F, A, T byte) {
	if n == 1 {
		fmt.Printf("move %d from  %c to %c \n", n, F, T)
		return
	}

	hanoi(n-1, F, T, A)
	fmt.Printf("move %d from %c to %c", n, F, T)
	hanoi(n-1, A, F, T)

}
