package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the D2 Stats REPL!")
	fmt.Println("Command Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
