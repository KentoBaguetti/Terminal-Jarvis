package main

import (
	"encoding/json"
	"log"
	"net/http"

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
