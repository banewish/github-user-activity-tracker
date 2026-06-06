package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func buildURL(username string) string {
	return fmt.Sprintf("https://api.github.com/users/%s/events", username)
}

func fetchEvents(username string) ([]Event, error) {
	resp, err := http.Get(buildURL(username))
	if err != nil {
		return nil, fmt.Errorf("fetching data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("user not found")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned: %s", resp.Status)
	}

	var events []Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("decoding JSON: %w", err)
	}

	return events, nil
}
