package removeduplicatesfromll

type LinkedList struct {
	Val  int
	Next *LinkedList
}

// Time: O(N)
// Space: O(1)
func RemoveDuplicatesFromLL(ll *LinkedList) *LinkedList {
	currentNode := ll

	for currentNode != nil && currentNode.Next != nil {
		nextNode := currentNode.Next

		if currentNode.Val == nextNode.Val {
			currentNode.Next = nextNode.Next
		} else {
			currentNode = nextNode
		}
	}
	return ll
}
