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
	client := gorunemetrics.NewClient(http.DefaultClient)

	profile, err := client.GetProfile(playerName)
	if err != nil {
		panic(err)
	}

	fmt.Println(profile)
}
