package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) SearchHistoricalStats() (StatsResponse, error) {
	// Create a request body
	reqBody := BungieStatsRequest{
		MembershipType: c.BungiePlayer.MembershipType,
		MembershipID:   c.BungiePlayer.MembershipID,
	}

	// Create HTTP request
	url := fmt.Sprintf("%s/Destiny2/%d/Account/%s/Stats/",
		baseURL,
		reqBody.MembershipType,
		reqBody.MembershipID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return StatsResponse{}, err
	}

	// Set headers
	req.Header.Set("X-API-Key", c.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return StatsResponse{}, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return StatsResponse{}, err
	}

	// Parse the raw response to access the nested structure
	var rawResponse map[string]interface{}
	if err := json.Unmarshal(body, &rawResponse); err != nil {
		return StatsResponse{}, err
	}

	// Check for API errors
	errorCode, _ := rawResponse["ErrorCode"].(float64)
	if errorCode != 1 {
		errorStatus, _ := rawResponse["ErrorStatus"].(string)
		return StatsResponse{}, fmt.Errorf("API error: %s", errorStatus)
	}

	// Create response object
	statsResponse := StatsResponse{
		RawJSON: body,
	}

	// Access the nested response structure to get the stats
	response, ok := rawResponse["Response"].(map[string]interface{})
	if !ok {
		return StatsResponse{}, fmt.Errorf("invalid response format")
	}

	merged, ok := response["mergedAllCharacters"].(map[string]interface{})
	if !ok {
		return StatsResponse{}, fmt.Errorf("no character stats found")
	}

	results, ok := merged["results"].(map[string]interface{})
	if !ok {
		return StatsResponse{}, fmt.Errorf("no stats results found")
	}

	// Process each requested mode
	for modeID, modeData := range results {
		modeStats, ok := modeData.(map[string]interface{})
		if !ok {
			fmt.Println("Failed to convert modeData to map")
			continue
		}

		allTime, ok := modeStats["allTime"].(map[string]interface{})
		if !ok {
			continue
		}

		// Look for PvP data
		if modeID == "allPvP" {
			fmt.Println("Found PvP data!")
			fmt.Println()
			crucible := extractCrucibleStats(allTime)
			statsResponse.Crucible = crucible
		}
	}

	// Check if we found any stats
	if statsResponse.Crucible == nil {
		return StatsResponse{}, fmt.Errorf("no trials stats found")
	}

	return statsResponse, nil
}
