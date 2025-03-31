package main

import "fmt"

func commandTrials(cfg *Config, args ...string) error {
	stats, err := cfg.apiConfig.SearchTrialsStats()
	if err != nil {
		return err
	}

	cfg.apiConfig.StatsResponse = &stats

	fmt.Println("===Trials of Osiris stats===")
	fmt.Printf("Flawless Completions: %v\n", cfg.apiConfig.StatsResponse.Trials.FlawlessCompletions)
	fmt.Printf("Wins: %v\n", cfg.apiConfig.StatsResponse.Trials.Wins)
	fmt.Printf("Losses: %v\n", cfg.apiConfig.StatsResponse.Trials.Losses)
	fmt.Printf("KD: %.2f\n", cfg.apiConfig.StatsResponse.Trials.KD)
	fmt.Printf("KDA: %.2f\n", cfg.apiConfig.StatsResponse.Trials.KDA)
	fmt.Printf("Total Kills: %v\n", cfg.apiConfig.StatsResponse.Trials.TotalKills)
	fmt.Printf("Total Deaths: %v\n", cfg.apiConfig.StatsResponse.Trials.TotalDeaths)
	fmt.Printf("Total Assits: %v\n", cfg.apiConfig.StatsResponse.Trials.TotalAssists)
	fmt.Printf("Longest Kill Spree: %v\n", cfg.apiConfig.StatsResponse.Trials.KillSpree)
	fmt.Printf("Average Lifespan: %v\n", cfg.apiConfig.StatsResponse.Trials.AverageLifespan)
	fmt.Printf("Total Time Played: %v\n", cfg.apiConfig.StatsResponse.Trials.TimePlayed)
	fmt.Println()
	return nil
}
