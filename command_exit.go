package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, args ...string) error {
	fmt.Printf("Thank you for using Destiny REPL!... Have a good day %s#%d!", cfg.apiConfig.BungiePlayer.DisplayName, cfg.apiConfig.BungiePlayer.BungieGlobalDisplayNameCode)

	os.Exit(0)
	return nil
}
