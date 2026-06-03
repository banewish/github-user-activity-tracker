package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt string `json:"created_at"`
}

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

	resp, err := http.Get(buildURL(username))
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		fmt.Println("User not found. Status code:", resp.StatusCode)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Github API returned:", resp.StatusCode)
		return
	}
	fmt.Println("Request successful:", resp.StatusCode)
	var events []Event
	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	limit := 3
	if len(events) < limit {
		limit = len(events)
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("Type: %s, Repo: %s, CreatedAt: %s\n", events[i].Type, events[i].Repo.Name, events[i].CreatedAt)
	}

}

func buildURL(username string) string {
	return fmt.Sprintf("https://api.github.com/users/%s/events", username)
}
