package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Score float64
}

func (s Monster) Print() {
	fmt.Println("-----")
	fmt.Println(s)
	fmt.Println("-----")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s *Monster) Set(name string, score float64) {
	s.Name = name
	s.Score = score
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()

	if kd != reflect.Struct && kd != reflect.Ptr {
		panic("expect struct")
	}

	// If it's a pointer, get the element type
	if kd == reflect.Ptr {
		// val = val.Elem()
		typ = typ.Elem()
		val = reflect.Indirect(val)
	}

	num := val.NumField()
	fmt.Printf("Field Number %d \n", num)

	for i := 0; i < num; i++ {
		fmt.Printf("Field %d : value %v \n", i, val.Field(i))

		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d : tag %v \n", i, tagVal)
		}
	}

	numOfMethod := val.NumMethod()
	fmt.Printf("Method Number %d \n", numOfMethod)

	for i := 0; i < numOfMethod; i++ {
		fmt.Printf("Method %d : Name %v \n", i, typ.Method(i).Name)
	}
	val.Method(1).Call(nil)
	var param []reflect.Value
	param = append(param, reflect.ValueOf(1), reflect.ValueOf(2))

	res := val.Method(0).Call(param)
	fmt.Println("res=", res[0].Int())

	if reflect.ValueOf(a).Kind() == reflect.Ptr {
		valPtr := reflect.ValueOf(a)
		typPtr := reflect.TypeOf(a)

		otherMethodNum := valPtr.NumMethod()
		fmt.Printf("otherMethodNum Number %d \n", otherMethodNum)
		for i := 0; i < otherMethodNum; i++ {
			fmt.Printf("Method %d : Name %v \n", i, typPtr.Method(i).Name)
		}

		// Call the Set method
		setMethod := reflect.ValueOf(a).MethodByName("Set")

		if setMethod.IsValid() {
			params := []reflect.Value{reflect.ValueOf("NewName"), reflect.ValueOf(95.5)}
			setMethod.Call(params)
			fmt.Println("After Set method call:", val.Interface())
		} else {
			fmt.Println("Set method not found")
		}
		val.Method(1).Call(nil)
	}

}

func main() {

	var a Monster = Monster{
		Name:  "大黄",
		Score: 90.1,
	}

	TestStruct(&a)

}
