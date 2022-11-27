package nonconstructiblechange

import "sort"

// Time: O(N - log(N)); N is the number of coins
// Space: O(log(N)) or O(N) based on the sorting algo
func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)
	change := 0
	for _, coin := range coins {
		if coin-change > 1 {
			break
		}
		change += coin
	}
	return change + 1
}
