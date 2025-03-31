package main

import (
	"bufio"
	"fmt"
	"github.com/exglegaming/d2-stats-repl/internal/api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load(".env")
	key := os.Getenv("API_KEY")
	agent := os.Getenv("USER_AGENT")

	cfg := &Config{
		apiConfig: &api.Client{
			ApiKey:    key,
			UserAgent: agent,
		},
	}

	fmt.Println("Welcome to Destiny REPL! A CLI tool to see your stats in Destiny 2!")
	fmt.Println("=================================================================================")
	fmt.Println()
	fmt.Println("Please enter the Bungie name of the character you want to lookup. (Ex: Name#1234)")
	fmt.Println("=================================================================================")
	fmt.Println()

	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("D2 > ")
	reader.Scan()
	fmt.Println()
	err := commandPlayer(cfg, reader.Text())
	if err != nil {
		log.Fatal(err)
	}

	go cfg.apiConfig.FetchCharacterIDs()

	startRepl(cfg)
}
