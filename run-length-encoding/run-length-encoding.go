package runlengthencoding

// Time: O(N)
// Space: O(N)
func runLengthEncoding(str string) []string {
	encodeChars := []string{}
	currRunLength := 1

	for i := 1; i <= len(str); i++ {
		prevChar := string(str[i-1])

		if i == len(str) || string(str[i]) != prevChar || currRunLength == 9 {
			encodeChars = append(encodeChars, string(currRunLength))
			encodeChars = append(encodeChars, prevChar)
			currRunLength = 0
		}
		currRunLength++
	}
	return encodeChars
}
