package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchCharacterIDs() error {
	// Create a request body
	reqBody := BungieStatsRequest{
		MembershipType: c.BungiePlayer.MembershipType,
		MembershipID:   c.BungiePlayer.MembershipID,
	}

	// Create HTTP request
	url := fmt.Sprintf("%s/Destiny2/%d/Profile/%s/?components=200",
		baseURL,
		reqBody.MembershipType,
		reqBody.MembershipID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("X-API-Key", c.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse response
	var rawResponse map[string]interface{}
	err = json.Unmarshal(body, &rawResponse)
	if err != nil {
		return err
	}

	// Check for API errors
	errorCode, _ := rawResponse["ErrorCode"].(float64)
	if errorCode != 1 {
		errorStatus, _ := rawResponse["ErrorStatus"].(string)
		return fmt.Errorf("API error %s", errorStatus)
	}

	// Navigate to the characterIDs array
	response, ok := rawResponse["Response"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid response format")
	}

	characters, ok := response["characters"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("characters not found in response")
	}

	characterData, ok := characters["data"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("character data not found")
	}

	// Convert to string slice
	characterIDs := make([]string, 0, len(characterData))
	for characterID := range characterData {
		characterIDs = append(characterIDs, characterID)
	}

	c.CharacterIDs = characterIDs
	return nil
}
