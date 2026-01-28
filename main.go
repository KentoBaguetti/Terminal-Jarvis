package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	copilot "github.com/github/copilot-sdk/go"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	client := copilot.NewClient(nil)
	if err := client.Start(); err != nil {
		log.Fatal("Error starting CoPilot client:", err)
	}
	defer client.Stop()

	session, err := client.CreateSession(&copilot.SessionConfig{Model: "gpt-4.1", Streaming: true, Tools: []copilot.Tool{}})
	if err != nil {
		log.Fatal("Error creating CLI Session:", err)
	}

	session.On(func(event copilot.SessionEvent) {
		if event.Type == "assistant.message_delta" {
			fmt.Print(*event.Data.DeltaContent)
		}
		if event.Type == "session.idle" {
			fmt.Println("\n")
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

		_, err = session.SendAndWait(copilot.MessageOptions{Prompt: userInput}, 0)
		if err != nil {
			log.Fatal("Error sending prompt to CLI:", err)
		}

	}

}
