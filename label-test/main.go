package main

import "fmt"

func main() {
	// TestContinueLabel()
	// fmt.Println("ContinueLabel --------------------------------Break")
	// TestBreak()
	TestBreakLabel()

	TestGoto()
}

func TestContinueLabel() {

	// label1:
	for i := 0; i < 3; i++ {
		fmt.Printf("befor i %d \n", i)
	label2:
		for j := 0; j < 3; j++ {
			fmt.Printf("befor j %d \n", j)
			if j == 1 {
				// continue label1
				continue label2
			}
			fmt.Printf("after j %d \n", j)
		}
		fmt.Printf("after i %d \n", i)
	}
}

func TestBreak() {

	for i := 0; i < 3; i++ {
		fmt.Printf("befor i %d \n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("befor j %d \n", j)
			if j == 1 {
				break
			}
			fmt.Printf("after j %d \n", j)
		}
		fmt.Printf("after i %d \n", i)
	}
}

func TestBreakLabel() {

Label1:
	for i := 0; i < 3; i++ {
		fmt.Printf("befor i %d \n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("befor j %d \n", j)
			if j == 1 {
				break Label1
			}
			fmt.Printf("after j %d \n", j)
		}
		fmt.Printf("after i %d \n", i)
	}
}

func TestGoto() {

	// Label1:
	for i := 0; i < 3; i++ {
		fmt.Printf("befor i %d \n", i)
	Label2:
		for j := 0; j < 3; j++ {
			fmt.Printf("befor j %d \n", j)
			if j == 1 {
				goto Label2
			}
			fmt.Printf("after j %d \n", j)
		}
		fmt.Printf("after i %d \n", i)
	}
}
