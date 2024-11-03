package main

import (
	"log"
	"unsafe"
)

func main() {

	var str1 string
	var i1 int
	var i8 int8
	var i16 int16
	var i3 int32
	var i6 int64

	str2 := "---222-2222222222222222222222222222222222222222222222222222222222222222222222222222222222222222"

	log.Println(len([]byte(str1)))
	log.Println(len([]byte(str2)))

	log.Println(unsafe.Sizeof(str1))
	log.Println(unsafe.Sizeof(str2))

	log.Println(unsafe.Sizeof(i1))
	log.Println(unsafe.Sizeof(i8))
	log.Println(unsafe.Sizeof(i16))
	log.Println(unsafe.Sizeof(i3))
	log.Println(unsafe.Sizeof(i6))
}
