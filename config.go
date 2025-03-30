package main

import "fmt"

// BungiePlayer represents a player returned from the search API
type BungiePlayer struct {
	MembershipID                string `json:"membershipId"`
	MembershipType              int    `json:"membershipType"`
	DisplayName                 string `json:"displayName"`
	BungieGlobalDisplayName     string `json:"bungieGlobalDisplayName"`
	BungieGlobalDisplayNameCode int    `json:"bungieGlobalDisplayNameCode"`
	CrossSaveOverride           int    `json:"crossSaveOverride"`
}

// Character represents a Destiny character
type Character struct {
	CharacterID string
	ClassType   int
	Light       int
	EmblemPath  string
}

// Config stores application configuration and state
type Config struct {
	// API Configuration
	ApiKey     string
	ApiBaseURL string

	// Current Session Data
	CurrentPlayer *BungiePlayer
	Characters    map[string]Character

	// User Preferences
	DefaultPlatform int
	OutputFormat    string // e.g., "text", "json", "table"

	// Application State
	LastApiCall int64 // timestamp of last API call (for rate limiting)
	CacheTTL    int   // Time-to-live for cached data in seconds

	// Cache paths
	ManifestPath string
	CachePath    string

	// Debug settings
	Verbose bool
}

// HasPlayer returns true if a player is currently selected
func (c *Config) HasPlayer() bool {
	return c.CurrentPlayer != nil
}

// GetFullPlayerName returns the formatted Bungie name
func (c *Config) GetFullPlayerName() string {
	if c.CurrentPlayer == nil {
		return "No player selected"
	}
	return fmt.Sprintf("%s#%d",
		c.CurrentPlayer.BungieGlobalDisplayName,
		c.CurrentPlayer.BungieGlobalDisplayNameCode)
}
