package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("hello plug2.so", os.Args)
}

func PluginP(name string) {
	fmt.Println("Hello Plugin2 so", name)
}
