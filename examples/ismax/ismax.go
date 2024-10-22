// This example uses the Profile API to check if a given player has achieved
// at least Level 99 in all skills.

package main

import (
	"fmt"
	"net/http"
	"os"
	"slices"

	"github.com/lotkey/gorunemetrics"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please pass in a player name as a command line argument.")
		return
	}

	if len(os.Args) != 2 {
		fmt.Println(
			"Please pass in a SINGLE player name as a command line argument.",
		)
		os.Exit(1)
	}

	playerName := os.Args[1]

	// Create a new RuneMetrics client using the default Golang HTTP client
	client := gorunemetrics.NewClient(http.DefaultClient)

	// Get the user's profile data
	profile, err := client.GetProfile(playerName)
	if err != nil {
		panic(err)
	}

	// Find the skill with the minimum level
	minimumSkill := slices.MinFunc(
		profile.SkillValues,
		func(a, b *gorunemetrics.SkillValue) int {
			return compareInts(a.Level, b.Level)
		},
	)

	// If the minimum level is at least 99, they have maxed!
	if minimumSkill.Level >= 99 {
		fmt.Printf("%v has maxed. Way to go %v!\n", playerName, playerName)
	} else {
		fmt.Printf("%v has not maxed... yet. \U0001f60e\n", playerName)
	}
}

func compareInts(a, b int) int {
	if a == b {
		return 0
	}

	if a > b {
		return 1
	}

	return -1
}
