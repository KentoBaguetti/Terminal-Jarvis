package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

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

		fmt.Println(userInput)

	}

}
