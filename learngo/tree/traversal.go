package tree

import "fmt"

// Traverse 中序遍历tree
// func (node *Node) Traverse() {
// 	if node == nil {
// 		return
// 	}

// 	node.Left.Traverse()
// 	node.Print()
// 	node.Right.Traverse()
// }

// Traverse 中序遍历
func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

// TraverseFunc 中序工具
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

// TraverseWithChannel 使用 channel 进行遍历
func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
