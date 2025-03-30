package main

import (
	"fmt"
)

func commandPVP(cfg *Config, args ...string) error {
	stats, err := cfg.apiConfig.SearchHistoricalStats()
	if err != nil {
		return err
	}

	cfg.apiConfig.StatsResponse = &stats

	fmt.Println("Crucible stats:")
	fmt.Printf("Wins: %v\n", cfg.apiConfig.StatsResponse.Crucible.Wins)
	fmt.Printf("Losses: %v\n", cfg.apiConfig.StatsResponse.Crucible.Losses)
	fmt.Printf("KD: %.2f\n", cfg.apiConfig.StatsResponse.Crucible.KD)
	fmt.Printf("KDA: %v\n", cfg.apiConfig.StatsResponse.Crucible.KDA)
	fmt.Printf("Total Kills: %v\n", cfg.apiConfig.StatsResponse.Crucible.TotalKills)
	fmt.Printf("Total Deaths: %v\n", cfg.apiConfig.StatsResponse.Crucible.TotalDeaths)
	fmt.Printf("Total Assits: %v\n", cfg.apiConfig.StatsResponse.Crucible.TotalAssists)
	fmt.Printf("Longest Kill Spree: %v\n", cfg.apiConfig.StatsResponse.Crucible.KillSpree)
	fmt.Printf("Average Lifespan: %v\n", cfg.apiConfig.StatsResponse.Crucible.AverageLifespan)
	fmt.Printf("Total Time Played: %v\n", cfg.apiConfig.StatsResponse.Crucible.TimePlayed)
	fmt.Println()

	return nil
}
