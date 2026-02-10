package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/tools/playwright/run", handleRun)

	log.Println("Playwright wrapper running on port :8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}

func handleRun(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed. Must be a POST request", http.StatusMethodNotAllowed)
	}
}
