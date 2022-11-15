package moveelementtoend

// Time: O(n) | Space: O(1)
func MoveElementToEnd(arr []int, toMove int) []int {
	i := 0
	j := len(arr)

	for i < j {
		if arr[j] == toMove {
			arr[i], arr[j] = arr[j], arr[i]
		}
		i += 1
	}
	return arr
}
