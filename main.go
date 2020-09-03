package main

import "fmt"

func main() {
	recentMatches := GetRecentMatches(64989974)
	for _, recentMatch := range recentMatches {

		// Filter out games which were not all pick
		if recentMatch.GameMode != 1 && recentMatch.GameMode != 22 {
			continue
		}

		fmt.Print(recentMatch.MatchID)

		if recentMatch.PlayerSlot <= 127 && recentMatch.RadiantWin {
			fmt.Println(" --> Gewonnen!")
		} else {
			fmt.Println(" --> Verloren :(")
		}
	}
}
