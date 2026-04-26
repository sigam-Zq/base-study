package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"plugin-sche/schema"
)

// Plugin entry point
func main() {
	// Read request from stdin
	// input, err := io.ReadAll(os.Stdin)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	// 	os.Exit(1)
	// }
	if len(os.Args) < 2 {
		fmt.Println("no input")
		return
	}

	input := os.Args[1]

	var req schema.Request
	if err := json.Unmarshal([]byte(input), &req); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling request: %v\n", err)
		os.Exit(1)
	}

	// Execute plugin logic
	resp := execute(req)

	// Write response to stdout
	output, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling response: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(output))
}

// execute contains the actual plugin logic
func execute(req schema.Request) schema.Response {
	// TODO: Implement your plugin logic here

	fmt.Fprintf(os.Stderr, "Plugin %s received request: %+v\n", "timeOut", req)

	time.Sleep(time.Second * 4)
	return schema.Response{
		Status: "success",
		Data:   map[string]interface{}{"data": fmt.Sprintf("Hello from %s! msg %s!", "timeOut", req.Data)},
		Error:  "",
	}
}
