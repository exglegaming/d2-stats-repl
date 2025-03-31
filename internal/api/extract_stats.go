package api

func extractCrucibleStats(allTime map[string]interface{}) *CrucibleStats {
	stats := &CrucibleStats{}

	// Extract each stat from the allTime data
	stats.Wins = extractStatValue(allTime, "activitiesWon")
	stats.Losses = extractStatValue(allTime, "activitiesLost")
	stats.TotalKills = extractStatValue(allTime, "kills")
	stats.TotalDeaths = extractStatValue(allTime, "deaths")
	stats.TotalAssists = extractStatValue(allTime, "assists")
	stats.KD = extractStatFloatValue(allTime, "killsDeathsRatio")
	stats.KDA = extractStatFloatValue(allTime, "killsDeathsAssistsRatio")
	stats.KillSpree = extractStatValue(allTime, "longestKillSpree")
	stats.AverageLifespan = extractStatValue(allTime, "averageLifespan")
	stats.TimePlayed = extractStatValue(allTime, "secondsPlayed")

	return stats
}

func extractPvEStats(allTime map[string]interface{}) *PvEStats {
	stats := &PvEStats{}

	// Extract each stat from the allTime data
	stats.TotalKills = extractStatValue(allTime, "kills")
	stats.TotalDeaths = extractStatValue(allTime, "deaths")
	return stats
}

func extractStatValue(allTime map[string]interface{}, statName string) int {
	stat, ok := allTime[statName].(map[string]interface{})
	if !ok {
		return 0
	}

	basic, ok := stat["basic"].(map[string]interface{})
	if !ok {
		return 0
	}

	val, ok := basic["value"].(float64)
	if !ok {
		return 0
	}
	return int(val)
}

func extractStatFloatValue(allTime map[string]interface{}, statName string) float64 {
	stat, ok := allTime[statName].(map[string]interface{})
	if !ok {
		return 0
	}

	basic, ok := stat["basic"].(map[string]interface{})
	if !ok {
		return 0
	}

	val, ok := basic["value"].(float64)
	if !ok {
		return 0
	}
	return val
}
