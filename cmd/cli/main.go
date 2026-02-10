package main

// this is the main file for the CLI interface
// also contains the Copilot LLM

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	copilot "github.com/github/copilot-sdk/go"
	"github.com/kentobaguetti/terminaljarvis/internal/shared"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	BadmintonScheduleRetriever := copilot.DefineTool("ubc_badminton_schedule_retriever", "Return the badminton schedule from UBC for the next three days including today", func(params shared.BadmintonParams, inv copilot.ToolInvocation) (shared.BadmintonResult, error) {
		answer := GetBadmintonSchedule()
		return shared.BadmintonResult{
			Result: answer,
		}, nil
	})

	client := copilot.NewClient(nil)
	if err := client.Start(); err != nil {
		log.Fatal("Error starting CoPilot client:", err)
	}
	defer client.Stop()

	session, err := client.CreateSession(&copilot.SessionConfig{Model: "claude-haiku-4.5", Streaming: true, Tools: []copilot.Tool{BadmintonScheduleRetriever}})
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

		start := time.Now()

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

		t := time.Now()
		fmt.Println("Inference time:", t.Sub(start))

	}

}
