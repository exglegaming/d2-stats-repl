package main

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
