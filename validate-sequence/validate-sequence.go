package validatesequence

// Time: O(n) | Space: O(1)
func IsValidSubsequence(arr []int, seq []int) bool {

	seqId := 0

	for _, num := range arr {
		if num == seq[seqId] {
			seqId++
		}

		if seqId == len(seq) {
			return true
		}
	}

	return false
}
