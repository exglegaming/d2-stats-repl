package api

import "encoding/json"

const (
	baseURL = "https://www.bungie.net/Platform"
)

type Client struct {
	ApiKey            string
	UserAgent         string
	BungieNameRequest *BungieNameRequest
	BungieResponse    *BungieResponse
	BungiePlayer      *BungiePlayer
	StatsResponse     *StatsResponse
	CharacterIDs      []string
}

// BungieNameRequest Api interactions
type BungieNameRequest struct {
	DisplayName     string `json:"displayName"`
	DisplayNameCode int    `json:"displayNameCode"`
}

type BungieResponse struct {
	Response    []BungiePlayer `json:"Response"`
	ErrorCode   int            `json:"ErrorCode"`
	ErrorStatus string         `json:"ErrorStatus"`
}

// BungiePlayer represents a player returned from the search API
type BungiePlayer struct {
	MembershipID                string `json:"membershipId"`
	MembershipType              int    `json:"membershipType"`
	DisplayName                 string `json:"displayName"`
	BungieGlobalDisplayName     string `json:"bungieGlobalDisplayName"`
	BungieGlobalDisplayNameCode int    `json:"bungieGlobalDisplayNameCode"`
	CrossSaveOverride           int    `json:"crossSaveOverride"`
}

// BungieStatsRequest API interaction to pull stats
type BungieStatsRequest struct {
	MembershipType int    `json:"membershipType"`
	MembershipID   string `json:"membershipId"`
}

// StatsResponse represents the stats pulled from the API
type StatsResponse struct {
	Trials   *TrialsStats
	Crucible *CrucibleStats
	RawJSON  json.RawMessage
}

type TrialsStats struct {
	Wins                int     `json:"activitiesWon"`
	Losses              int     `json:"activitiesLost"`
	KD                  float64 `json:"killsDeathsRatio"`
	KDA                 float64 `json:"killsDeathsAssistsRatio"`
	TotalKills          int     `json:"kills"`
	TotalDeaths         int     `json:"deaths"`
	TotalAssists        int     `json:"assists"`
	KillSpree           int     `json:"longestKillSpree"`
	AverageLifespan     int     `json:"averageLifespan"`
	FlawlessCompletions int     `json:"flawlessCompletions"`
	TimePlayed          int     `json:"secondsPlayed"`
}

type CrucibleStats struct {
	Wins            int     `json:"activitiesWon"`
	Losses          int     `json:"activitiesLost"`
	KD              float64 `json:"killsDeathsRatio"`
	KDA             float64 `json:"killsDeathsAssistsRatio"`
	TotalKills      int     `json:"kills"`
	TotalDeaths     int     `json:"deaths"`
	TotalAssists    int     `json:"assists"`
	KillSpree       int     `json:"longestKillSpree"`
	AverageLifespan int     `json:"averageLifespan"`
	TimePlayed      int     `json:"secondsPlayed"`
}
