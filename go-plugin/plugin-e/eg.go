package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

//export PluginP
func PluginP(name *C.char) {
	goStr := C.GoString(name)
	fmt.Println("Hello Plugin", goStr)
}

func main() {}
