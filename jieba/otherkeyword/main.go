package main

import (
	"fmt"
	"log"

	"github.com/yanyiwu/gojieba"
)

func main() {

	log.Println("-------------------------------------")

	testJieba()
}

func testJieba() {

	var s string
	x := gojieba.NewJieba()
	defer x.Free()

	s = `我喜欢上班`

	keywords := x.ExtractWithWeight(s, 20)

	fmt.Println("Extract:", keywords)
}
