package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	key := os.Getenv("API_KEY")

	fmt.Println("Welcome to Destiny REPL! A CLI tool to see your stats in Destiny 2!")
	fmt.Println("===================================================================")
	fmt.Println()

	player := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your Bungie name. (ex: name#1234)")
	fmt.Print("D2 > ")
	player.Scan()

	fmt.Println()
	fmt.Println("Now that you have entered your Bungie name you can begin to pull your stats.")
	fmt.Println("To see a list of commands use the 'help' command.")
	fmt.Println()

	cfg := &config{
		name:   player.Text(),
		apiKey: key,
	}
	startRepl(cfg)
}
