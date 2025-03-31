package main

import "fmt"

func commandPVE(cfg *Config, args ...string) error {
	stats, err := cfg.apiConfig.SearchHistoricalStats()
	if err != nil {
		return err
	}

	cfg.apiConfig.StatsResponse = &stats

	fmt.Println("===PvE Stats===")
	fmt.Printf("Total Kills: %v\n", cfg.apiConfig.StatsResponse.PvE.TotalKills)
	fmt.Printf("Total Deaths: %v\n", cfg.apiConfig.StatsResponse.PvE.TotalDeaths)
	fmt.Println()
	return nil
}
