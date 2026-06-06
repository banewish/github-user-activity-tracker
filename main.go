package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <github-username>")
		os.Exit(1)
	}

	username := strings.TrimSpace(os.Args[1])
	if username == "" {
		fmt.Println("Usage: go run . <github-username>")
		os.Exit(1)
	}

	fmt.Println("Checking GitHub user activity for:", username)

	events, err := fetchEvents(username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Request successful. Found %d events.\n", len(events))

	limit := 3
	if len(events) < limit {
		limit = len(events)
	}
	for i := 0; i < limit; i++ {
		fmt.Println(formatEvent(events[i]))
	}

}
