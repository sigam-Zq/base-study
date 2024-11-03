package main

import (
	"embed"
	"log"
)

//go:embed dict
var dictDir embed.FS

func main() {
	dictDir.ReadDir("dict")
	log.Println()
}
