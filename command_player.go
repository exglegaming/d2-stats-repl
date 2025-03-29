package main

import "fmt"

func commandPlayer(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("You must enter a Bungie name.\n")
	}
	fmt.Println()

	name := args[0]
	cfg.name = name

	return nil
}
