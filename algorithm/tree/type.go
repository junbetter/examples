package tree

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

func CreateNode(val int) *Node {
	return &Node{Val: val}
}

func (node *Node) Print() {
	fmt.Print(node.Val, " ")
}

func (node *Node) SetVal(val int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.Val = val
}
