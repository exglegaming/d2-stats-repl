package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) SearchTrialsStats() (StatsResponse, error) {
	// Initialize a response to hold combined stats
	statsResponse := StatsResponse{
		Trials: &TrialsStats{},
	}

	foundStats := false

	for _, characterID := range c.CharacterIDs {
		// Create HTTP request
		url := fmt.Sprintf("%s/Destiny2/%d/Account/%s/Character/%s/Stats/Activities/?mode=84&count=250",
			baseURL,
			c.BungiePlayer.MembershipType,
			c.BungiePlayer.MembershipID,
			characterID)
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

		// Read the response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return StatsResponse{}, err
		}

		// Parse the raw response to access the data
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

		// Extract activities from response
		response, ok := rawResponse["Response"].(map[string]interface{})
		if !ok {
			continue
		}

		activities, ok := response["activities"].([]interface{})
		if !ok || len(activities) == 0 {
			fmt.Printf("No Trials activities found for character %s\n", characterID)
			continue
		}

		// Process activities to extract stats
		charStats := &TrialsStats{}

		for _, activity := range activities {
			activityData, ok := activity.(map[string]interface{})
			if !ok {
				continue
			}

			// Extract values from this activity
			values, ok := activityData["values"].(map[string]interface{})
			if !ok {
				continue
			}

			// Extract match result (win/loss)
			standing, ok := values["standing"].(map[string]interface{})
			if ok {
				if basic, ok := standing["basic"].(map[string]interface{}); ok {
					if value, ok := basic["value"].(float64); ok {
						if value == 0 { // 0 = win
							charStats.Wins++
						} else {
							charStats.Losses++
						}
					}
				}
			}

			// Extract kills, deaths, assists
			charStats.TotalKills += extractActivityStat(values, "kills")
			charStats.TotalDeaths += extractActivityStat(values, "deaths")
			charStats.TotalAssists += extractActivityStat(values, "assists")
			charStats.TimePlayed += extractActivityStat(values, "timePlayedSeconds")
			charStats.FlawlessCompletions += extractActivityStat(values, "flawlessCompletions")
			charStats.Wins += extractActivityStat(values, "activitiesWon")
			charStats.Losses += extractActivityStat(values, "activitiesLost")
		}

		// Only include this character's stats if they have played Trials
		if charStats.TotalKills > 0 || charStats.Wins > 0 || charStats.Losses > 0 {
			combineStats(statsResponse.Trials, charStats)
			foundStats = true
		}
	}

	// If we didn't find any stats, return an appropriate error
	if !foundStats {
		return StatsResponse{}, fmt.Errorf("no Trials activities found for any character")
	}

	// Calculate aggregate values like K/D ratio after combining all stats
	finalizeStats(statsResponse.Trials)

	return statsResponse, nil
}

// Helper to extract a stat from activity values
func extractActivityStat(values map[string]interface{}, statName string) int {
	if stat, ok := values[statName].(map[string]interface{}); ok {
		if basic, ok := stat["basic"].(map[string]interface{}); ok {
			if value, ok := basic["value"].(float64); ok {
				return int(value)
			}
		}
	}
	return 0
}

// Helper function to combine stats from multiple character
func combineStats(combined *TrialsStats, charStats *TrialsStats) {
	combined.Wins += charStats.Wins
	combined.Losses += charStats.Losses
	combined.TotalKills += charStats.TotalKills
	combined.TotalDeaths += charStats.TotalDeaths
	combined.TotalAssists += charStats.TotalAssists
	combined.FlawlessCompletions += charStats.FlawlessCompletions
	combined.TimePlayed += charStats.TimePlayed
}

// Calculate final ratios and metrics after combining all stats
func finalizeStats(stats *TrialsStats) {
	// Calculate K/D ratio
	if stats.TotalDeaths > 0 {
		stats.KD = float64(stats.TotalKills) / float64(stats.TotalDeaths)
	}

	// Calculate KDA
	if stats.TotalDeaths > 0 {
		stats.KDA = float64(stats.TotalKills+stats.TotalAssists/2) / float64(stats.TotalDeaths)
	}
}
