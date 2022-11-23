package tournamentwinner

const WINNER = 1
const POINTS = 3

// Time: O(n)
// Space: O(M); where M is the number of teams
func tournamentWinner(comp [][]string, results []int) string {
	currBestTeam := ""
	scores := map[string]int{currBestTeam: 0}

	for i, compIter := range comp {
		oneTeam, twoTeam := compIter[0], compIter[1]
		res := results[i]

		winnerTeam := twoTeam

		if res == WINNER {
			winnerTeam = oneTeam
		}

		scores[winnerTeam] += POINTS

		if scores[winnerTeam] > scores[currBestTeam] {
			currBestTeam = winnerTeam
		}
	}
	return currBestTeam
}
