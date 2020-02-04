package main

import (
	"examples/algorithm/tree"
	"fmt"
)

func main() {
	var root *tree.Node
	root = &tree.Node{Val: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Val: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetVal(4)

	fmt.Println(tree.MaxDepth(root))
}
