package schema

type Request struct {
	Data map[string]interface{} `json:"data"`
}

type Response struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
	Error  string                 `json:"error"`
}
