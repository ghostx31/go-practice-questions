package nodedepth

type BinaryTree struct {
	val   int
	Left  *BinaryTree
	Right *BinaryTree
}

type Depth struct {
	node  *BinaryTree
	depth int
}

// Recursive
// Time: O(N)
// Space: O(M) where M is the number of nodes.
func nodeDepth(root *BinaryTree) int {
	return findNodeDepth(root, 0)
}

func findNodeDepth(node *BinaryTree, depth int) int {
	if node == nil {
		return 0
	}

	return depth + findNodeDepth(node.Left, depth+1) + findNodeDepth(node.Right, depth+1)
}

// Iterative
// Time and space same as the recursive one
func iterativeNodeDepth(root *BinaryTree) int {
	sumOfDepth := 0
	depthLevel := []Depth{Depth{root, 0}}

	for len(depthLevel) > 0 {
		node, depth := depthLevel[len(depthLevel)-1].node, depthLevel[len(depthLevel)-1].depth
		depthLevel = depthLevel[:len(depthLevel)-1]

		sumOfDepth += depth

		if node.Left != nil {
			depthLevel = append(depthLevel, Depth{node.Left, depth + 1})
		}

		if node.Right != nil {
			depthLevel = append(depthLevel, Depth{node.Right, depth + 1})
		}
	}

	return sumOfDepth
}
