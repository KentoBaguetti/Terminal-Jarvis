package main

// JSON response for now TEMP
// modify later to hold more data
type Resp struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

// temp parameters for the mcp server
type BadmintonParams struct {
	Query string `json:"query"`
}

// temp return type for the badminton stuff
type BadmintonResult struct {
	Result Resp `json:"result"`
}
