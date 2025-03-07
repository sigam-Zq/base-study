package main

import (
	"bufio"
	"os"
)

func main() {
	f, err := os.Open("f/abb.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scan := bufio.NewScanner(f)

	for scan.Scan() {
		cnt := scan.Text()
		println(cnt)
	}

}
