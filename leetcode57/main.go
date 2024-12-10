package main

import (
	"fmt"
	"log"
)

func main() {

	log.Println(squareIsWhite("h3"))

}

func squareIsWhite(coordinates string) bool {
	diff := coordinates[0] - coordinates[1]
	fmt.Println(diff)
	return diff&1 > 0
}
