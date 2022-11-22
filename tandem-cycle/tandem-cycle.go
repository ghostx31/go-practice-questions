package tandemcycle

import "sort"

// Time: O(n - logn); n is the number of riders
// Space: O(logn) or O(n) depending on the sorting algorithm
func TandemCycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)
	sort.Ints(blueShirtSpeeds)

	if fastest {
		reverseArray(blueShirtSpeeds)
	}

	totalSpeed := 0

	for i, redSredShirtSpeed := range redShirtSpeeds {
		blueShirtSpeed := blueShirtSpeeds[i]
		totalSpeed += max(redSredShirtSpeed, blueShirtSpeed)
	}

	return totalSpeed
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func reverseArray(arr []int) {
	for f, e := 0, len(arr)-1; f < e; f, e = f+1, e-1 {
		arr[f], arr[e] = arr[e], arr[f]
	}
}
