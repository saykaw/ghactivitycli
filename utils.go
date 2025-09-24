package main

import (
	"fmt"
	"ghactivitycli/types"
)

func formatEvent(events []types.GitHubEvent) {
	for i, event := range events {
		fmt.Printf("Event #%d:\n", i+1)
		fmt.Println("  Event ID:", event.ID)
		fmt.Println("  Event Type:", event.Type)
		fmt.Println("  Actor Login:", event.Actor.Login)
		fmt.Println("  Repository Name:", event.Repo.Name)
		fmt.Println("  Number of Commits:", len(event.Payload.Commits))
		fmt.Println("  Event Created At:", event.CreatedAt)
		fmt.Println("-----------------------------")
	}

}
