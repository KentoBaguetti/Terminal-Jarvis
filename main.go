package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	copilot "github.com/github/copilot-sdk/go"
)

// json response
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

func GetBadmintonSchedule() Resp {
	resp, err := http.Get("http://localhost:8000/api/badminton-schedule")
	if err != nil {
		log.Fatal("Error getting response for badminton schedule:", err)
	}
	defer resp.Body.Close()

	var out Resp
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		log.Fatal("Error decoding response body:", err)
	}

	return out
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	BadmintonScheduleRetriever := copilot.DefineTool("ubc_badminton_schedule_retriever", "Return the badminton schedule from UBC for the next three days including today", func(params BadmintonParams, inv copilot.ToolInvocation) (BadmintonResult, error) {
		answer := GetBadmintonSchedule()
		return BadmintonResult{
			Result: answer,
		}, nil
	})

	client := copilot.NewClient(nil)
	if err := client.Start(); err != nil {
		log.Fatal("Error starting CoPilot client:", err)
	}
	defer client.Stop()

	session, err := client.CreateSession(&copilot.SessionConfig{Model: "gpt-4.1", Streaming: true, Tools: []copilot.Tool{BadmintonScheduleRetriever}})
	if err != nil {
		log.Fatal("Error creating CLI Session:", err)
	}

	session.On(func(event copilot.SessionEvent) {
		if event.Type == "assistant.message_delta" {
			fmt.Print(*event.Data.DeltaContent)
		}
		if event.Type == "session.idle" {
			fmt.Print("\n\n")
		}
	})

	for {

		fmt.Print("Enter message here (enter 'exit' or click ctrl+c to exit): ")

		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error with user input:", err)
		}

		userInput = strings.TrimSpace(userInput)

		if userInput == "" {
			continue
		} else if userInput == "exit" {
			break
		}

		_, err = session.SendAndWait(copilot.MessageOptions{Prompt: userInput}, 2*time.Minute)
		if err != nil {
			log.Fatal("Error sending prompt to CLI:", err)
		}

	}

}
