package shared

type ToolRequest struct {
	Task   string         `json:"task"`
	Params map[string]any `json:"params"`
}

// JSON response for now TEMP
// modify later to hold more data
type ToolResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data,omitempty"`
	Error  string `json:"error,omitempty"`
}

// temp parameters for the mcp server
type BadmintonParams struct {
	Query string `json:"query"`
}

// temp return type for the badminton stuff
type BadmintonResult struct {
	Result ToolResponse `json:"result"`
}

type PlaywrightParams struct {
	URL         string `json:"url"`
	Instruction string `json:"instruction"`
}

type PlaywrightResult struct {
	Result ToolResponse `json:"result"`
}
