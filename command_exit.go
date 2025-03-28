package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Printf("Thank you for using Destiny REPL!... Have a good day %s!", cfg.name)

	os.Exit(0)
	return nil
}
