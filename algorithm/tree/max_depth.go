package tree

import "math"

func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	return int(math.Max(float64(left), float64(right))) + 1
}
