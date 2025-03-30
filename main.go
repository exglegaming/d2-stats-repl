package main

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	baseURL = "https://www.bungie.net/Platform"
)

func main() {
	godotenv.Load(".env")
	key := os.Getenv("API_KEY")
	agent := os.Getenv("USER_AGENT")

	cfg := &Config{
		ApiKey:          key,
		ApiBaseURL:      baseURL,
		UserAgent:       agent,
		DefaultPlatform: -1,
		CurrentPlayer: &BungiePlayer{
			MembershipID:                "",
			MembershipType:              0,
			DisplayName:                 "",
			BungieGlobalDisplayName:     "",
			BungieGlobalDisplayNameCode: 0,
			CrossSaveOverride:           0,
		},
	}
	startRepl(cfg)
}
