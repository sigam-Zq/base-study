package main

import (
	"encoding/json"
	"fmt"
)

// 可通过如下逻辑解析任意 json
func main() {

	humanStr := `{
		"content":"您有1条临期待办事件未处理，其中紧急事件0件，一般事件1件，请及时处理",
		"list":[{"abb":"aaa"},{"cc":"dd"}]
	}`

	anyJsonUnmarshal(humanStr)

	anyJsonUnmarshal2(humanStr)

}

func anyJsonUnmarshal(humanStr string) {

	var obj interface{}
	err := json.Unmarshal([]byte(humanStr), &obj)
	if err != nil {
		panic(err)
	}
	objMap, ok := obj.(map[string]interface{})
	fmt.Println("ok", ok)
	for k, v := range objMap {
		switch value := v.(type) {
		case string:
			fmt.Printf("type of %s is string, value is %v\n", k, value)
		case interface{}:
			fmt.Printf("type of %s is interface{}, value is %v\n", k, value)
		default:
			fmt.Printf("type of %s is wrong, value is %v\n", k, value)
		}
	}

}

func anyJsonUnmarshal2(humanStr string) {

	var obj map[string]interface{}
	err := json.Unmarshal([]byte(humanStr), &obj)
	if err != nil {
		panic(err)
	}
	// objMap, ok := obj.(map[string]interface{})
	// fmt.Println("ok", ok)
	for k, v := range obj {
		switch value := v.(type) {
		case string:
			fmt.Printf("type of %s is string, value is %v\n", k, value)
		case interface{}:
			fmt.Printf("type of %s is interface{}, value is %v\n", k, value)
		default:
			fmt.Printf("type of %s is wrong, value is %v\n", k, value)
		}
	}

}
