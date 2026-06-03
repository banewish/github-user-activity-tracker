package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Placeholder for GitHub user activity tracking logic
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

}
