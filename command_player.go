package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	player, err := searchPlayerByBungieName(cfg, displayName, displayCode)
	if err != nil {
		return fmt.Errorf("Error searching for player: %s\n", err)
	}

	// Store player info in the config
	cfg.CurrentPlayer = player
	fmt.Printf("Player found: %s#%d\n",
		player.BungieGlobalDisplayName,
		player.BungieGlobalDisplayNameCode)

	return nil
}

// Helper function for API communication
func searchPlayerByBungieName(cfg *Config, name string, code int) (*BungiePlayer, error) {
	// Create request body
	reqBody := BungieNameRequest{
		DisplayName:     name,
		DisplayNameCode: code,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	// Create HTTP request
	url := cfg.ApiBaseURL + "/Destiny2/SearchDestinyPlayerByBungieName/-1/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("X-API-Key", cfg.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", cfg.UserAgent)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse response
	var bungieResponse BungieResponse
	if err := json.Unmarshal(body, &bungieResponse); err != nil {
		return nil, err
	}

	// Check for API errors
	if bungieResponse.ErrorCode != 1 {
		return nil, fmt.Errorf("API error: %s\n", bungieResponse.ErrorStatus)
	}

	// Check if player was found
	if len(bungieResponse.Response) == 0 {
		return nil, fmt.Errorf("No player found with that Bungie name.\n")
	}

	// If there are multiple accounts, prefer the cross-save primary
	player := bungieResponse.Response[0]
	if len(bungieResponse.Response) > 1 {
		for _, p := range bungieResponse.Response {
			if p.CrossSaveOverride > 0 && p.MembershipType == p.CrossSaveOverride {
				player = p
				break
			}
		}
	}
	return &player, nil
}
