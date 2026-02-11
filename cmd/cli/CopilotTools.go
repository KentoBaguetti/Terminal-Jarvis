package main

// contains tool definition

import (
	"encoding/json"
	"log"
	"net/http"

	copilot "github.com/github/copilot-sdk/go"
	"github.com/kentobaguetti/terminaljarvis/internal/shared"
)

func GetBadmintonSchedule() shared.ToolResponse {
	resp, err := http.Get("http://localhost:8000/api/badminton-schedule")
	if err != nil {
		log.Fatal("Error getting response for badminton schedule:", err)
	}
	defer resp.Body.Close()

	var out shared.ToolResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		log.Fatal("Error decoding response body:", err)
	}

	return out
}

var BadmintonScheduleRetriever copilot.Tool = copilot.DefineTool("ubc_badminton_schedule_retriever", "Return the badminton schedule from UBC for the next three days including today", func(params shared.BadmintonParams, inv copilot.ToolInvocation) (shared.BadmintonResult, error) {
	answer := GetBadmintonSchedule()
	return shared.BadmintonResult{
		Result: answer,
	}, nil
})

func BrowserAutomation(params shared.PlaywrightParams) shared.ToolResponse {
	var out shared.ToolResponse
	return out
}

var CopilotTools = [...]copilot.Tool{BadmintonScheduleRetriever}
