package main

import (
	"fmt"
	"reflect"
)

func main() {
	var z = 123
	var y = &z
	var x interface{} = y

	v := reflect.ValueOf(&x)
	println(v.Kind())     //22 Pointer
	fmt.Println(v.Kind()) // ptr
	vx := v.Elem()
	println(vx.Kind())     //20
	fmt.Println(vx.Kind()) // interface
	vy := vx.Elem()
	println(vy.Kind())     //22
	fmt.Println(vy.Kind()) // ptr
	vz := vy.Elem()
	println(vz.Kind())     //2
	fmt.Println(vy.Kind()) // ptr
}
