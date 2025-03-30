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

	cfg := &Config{
		ApiKey:          key,
		ApiBaseURL:      baseURL,
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
