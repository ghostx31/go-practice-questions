package threenumbersum

import (
	"sort"
)

func threeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)

	triplet := [][]int{}
	for something := 0; something < len(array)-2; something++ {
		TwoNumberSum(array, something, target, triplet)
	}
	return triplet
}

func TwoNumberSum(arr []int, i int, target int, triplet [][]int) {
	left := i + 1
	right := len(arr) - 1
	for left < right {
		currSum := arr[i] + arr[left] + arr[right]

		if currSum == target {
			triplet = append(triplet, []int{arr[i], arr[left], arr[right]})
			left++
			right--
		} else if currSum > target {
			right--
		} else {
			left++
		}
	}
}
