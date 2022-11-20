package classphoto

import "sort"

// Time: O(nlogn), 'n' being the number of students
// Space: O(N) or O(logN) depending on the sorting algorithm.
func classPhoto(redShirts []int, blueShirts []int) bool {
	sort.Slice(redShirts, func(i, j int) bool {
		return redShirts[i] > redShirts[j]
	})

	sort.Slice(blueShirts, func(i, j int) bool {
		return blueShirts[i] > blueShirts[j]
	})

	redInFrontRow := redShirts[0] < blueShirts[0]

	for i, redShirt := range redShirts {
		blueShirt := blueShirts[i]

		if redInFrontRow && redShirt >= blueShirt {
			return false
		}

		if !redInFrontRow && blueShirt >= redShirt {
			return false
		}
	}
	return true
}
