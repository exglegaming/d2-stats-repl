package main

import (
	"fmt"
	"strconv"
	"strings"
)

func commandPlayer(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("You must enter a Bungie name.\n")
	}
	fmt.Println()

	// parse Bungie name (format: name#1234)
	bungieName := args[0]
	parts := strings.Split(bungieName, "#")
	if len(parts) != 2 {
		return fmt.Errorf("You must enter a Bungie name format. Use: Name#1234\n")
	}

	displayName := parts[0]
	displayCode, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("The code after # must be a number.\n")
	}

	// Search for player using Bungie API
	fmt.Printf("Searching for player: %s#%d\n", displayName, displayCode)
	cfg.apiConfig.BungiePlayer, err = cfg.apiConfig.SearchPlayerByBungieName(displayName, displayCode)
	if err != nil {
		return fmt.Errorf("Error searching for player: %s\n", err)
	}

	// Store player info in the config
	//cfg.apiConfig.BungiePlayer = player
	fmt.Printf("Player found: %s#%d\n",
		cfg.apiConfig.BungiePlayer.DisplayName,
		cfg.apiConfig.BungiePlayer.BungieGlobalDisplayNameCode)
	fmt.Println()

	return nil
}
