package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"plugin-sche/schema"
)

// writeError writes a structured error response to stdout and exits with code 1.
func writeError(msg string) {
	resp := schema.Response{Status: "error", Error: msg}
	out, _ := json.Marshal(resp)
	fmt.Println(string(out))
	os.Exit(1)
}

// Plugin entry point
func main() {
	if len(os.Args) < 2 {
		writeError("missing input argument")
	}

	input := os.Args[1]

	var req schema.Request
	if err := json.Unmarshal([]byte(input), &req); err != nil {
		writeError(fmt.Sprintf("invalid JSON input: %v", err))
	}

	// Execute plugin logic
	resp := execute(req)

	// Write response to stdout
	output, err := json.Marshal(resp)
	if err != nil {
		writeError(fmt.Sprintf("failed to marshal response: %v", err))
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
