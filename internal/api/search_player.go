package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SearchPlayerByBungieName Helper function for API communication
func (c *Client) SearchPlayerByBungieName(name string, code int) (*BungiePlayer, error) {
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
	url := baseURL + "/Destiny2/SearchDestinyPlayerByBungieName/-1/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("X-API-Key", c.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

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
