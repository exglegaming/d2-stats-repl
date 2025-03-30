package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, args ...string) error {
	fmt.Printf("Thank you for using Destiny REPL!... Have a good day %s#%d!", cfg.CurrentPlayer.DisplayName, cfg.CurrentPlayer.BungieGlobalDisplayNameCode)

	os.Exit(0)
	return nil
}
