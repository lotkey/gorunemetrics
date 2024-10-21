// This example uses the Quests API to check if the player has completed all
// unlocked quests.

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/lotkey/gorunemetrics"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please pass in a player name as a command line argument.")
		return
	}

	if len(os.Args) != 2 {
		fmt.Println("Please pass in a SINGLE player name as a command line argument.")
		os.Exit(1)
	}

	playerName := os.Args[1]

	// Create a new RuneMetrics client using the default Golang HTTP client
	client := gorunemetrics.NewClient(http.DefaultClient)

	// Get the user's profile quest status
	quests, err := client.GetQuests(playerName)
	if err != nil {
		panic(err)
	}

	incompleteUnlockedP2PQuests := []*gorunemetrics.PlayerQuestStatus{}
	incompleteUnlockedF2PQuests := []*gorunemetrics.PlayerQuestStatus{}

	// Find all quests that the user has unlocked but has not completed
	for _, quest := range quests {
		if quest.UserEligible && quest.Status != gorunemetrics.Completed {
			if quest.Members {
				incompleteUnlockedP2PQuests = append(incompleteUnlockedP2PQuests, quest)
			} else {
				incompleteUnlockedF2PQuests = append(incompleteUnlockedF2PQuests, quest)
			}
		}
	}

	// Check if there are any incomplete but unlocked F2P quests
	if len(incompleteUnlockedF2PQuests) == 0 {
		fmt.Printf("%v has completed all unlocked free quests.\n", playerName)
	} else {
		fmt.Printf("%v can complete the following unlocked free quests:\n", playerName)
		for _, quest := range incompleteUnlockedF2PQuests {
			fmt.Printf("  - %v\n", quest.Title)
		}
	}

	// Check if there are any incomplete but unlocked P2P quests
	fmt.Println()
	if len(incompleteUnlockedP2PQuests) == 0 {
		fmt.Printf("%v has completed all members quests.\n", playerName)
	} else {
		fmt.Printf("%v can complete the following unlocked members quests:\n", playerName)
		for _, quest := range incompleteUnlockedP2PQuests {
			fmt.Printf("  - %v\n", quest.Title)
		}
	}

	fmt.Println()
	if len(incompleteUnlockedF2PQuests)+len(incompleteUnlockedP2PQuests) == 0 {
		fmt.Printf("Way to go %v!\n", playerName)
	} else {
		fmt.Printf("Get to work %v!\n", playerName)
	}
}
