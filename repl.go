package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	name   string
	apiKey string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("D2 > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command. Please try again or use the 'help' command to see all commands.")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a list of all commands available",
			callback:    commandHelp,
		},
		"pve": {
			name:        "pve",
			description: "Displays your pve stats",
			callback:    commandPVE,
		},
		"pvp": {
			name:        "pvp",
			description: "Displays your crucible stats",
			callback:    commandPVP,
		},
		"trials": {
			name:        "trials",
			description: "See your Trials of Osiris stats",
			callback:    commandTrials,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    commandExit,
		},
	}
}
