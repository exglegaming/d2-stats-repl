package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	key := os.Getenv("API_KEY")

	cfg := &config{
		name:   "",
		apiKey: key,
	}
	startRepl(cfg)
}
