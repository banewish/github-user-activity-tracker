package main

import "fmt"

func formatEvent(e Event) string {
	switch e.Type {
	case "PushEvent":
		return fmt.Sprintf("Pushed to %s at %s", e.Repo.Name, e.CreatedAt)
	case "CreateEvent":
		return fmt.Sprintf("Created something in %s at %s", e.Repo.Name, e.CreatedAt)
	case "IssuesEvent":
		return fmt.Sprintf("Updated issues in %s at %s", e.Repo.Name, e.CreatedAt)
	case "WatchEvent":
		return fmt.Sprintf("Starred %s at %s", e.Repo.Name, e.CreatedAt)
	default:
		return fmt.Sprintf("Did %s in %s at %s", e.Type, e.Repo.Name, e.CreatedAt)
	}
}
