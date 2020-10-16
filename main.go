package main

import (
	"fmt"
	"time"
)

func getMatchIdsForPlayerId(playerId int64, maxTimeInSeconds int64) []int64 {
	recentMatches := GetRecentMatches(playerId)
	var recentMatchIds []int64

	for _, matchDetails := range recentMatches {
		if (matchDetails.StartTime + maxTimeInSeconds) >= time.Now().Unix() {
			recentMatchIds = append(recentMatchIds, matchDetails.MatchID)
		}
	}
	return recentMatchIds
}

func main() {
	team := GetTeam()
	var matches []int64

	for _, teammember := range team.Members {
		fmt.Printf("Lade Daten für %s (Id: %d) \r\n", teammember.Name, teammember.Id)
		matches = append(matches, getMatchIdsForPlayerId(teammember.Id, (60*60*24))...)
	}

	fmt.Println(matches)

	// recentMatches := GetRecentMatches(64989974)
	// for _, recentMatch := range recentMatches {

	// 	// Filter out games which were not all pick
	// 	if recentMatch.GameMode != 1 && recentMatch.GameMode != 22 {
	// 		continue
	// 	}

	// 	fmt.Print(recentMatch.MatchID)

	// 	if recentMatch.PlayerSlot <= 127 && recentMatch.RadiantWin {
	// 		fmt.Println(" --> Gewonnen!")
	// 	} else {
	// 		fmt.Println(" --> Verloren :(")
	// 	}
	// }
}
