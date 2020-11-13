package main

import (
	"fmt"

	"github.com/captainlee1024/learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	fmt.Println(root)

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{
		Value: 5,
		Left:  nil,
		Right: nil}
	root.Right.Left = new(tree.Node)
	root.Right.Right = tree.CreateNode(2)

	nodes := []tree.Node{
		{
			Value: 4,
			Left:  nil,
			Right: nil},
		{},
		{
			Value: 1,
			Left:  nil,
			Right: &root},
	}
	fmt.Println(nodes)

	root.Right.Right.SetValue(3)
	root.Right.Right.Print()

	pRroot := &root
	pRroot.Print()
	pRroot.SetValue(100)
	pRroot.Print()
	root.Print()

	var nNode tree.Node
	nNode.SetValue(100)

	root.Traverse()

	myNode := myTreeNode{&root}
	myNode.postOrder()

	count := 0
	myNode.node.TraverseFunc(func(n *tree.Node) {
		count++
	})
	fmt.Printf("\nmyNode 一共 %d 个节点", count)

	cNode := root.TraverseWithChannel()
	maxNode := 0
	// 如果 TraverseWithChannel 里的channel out 不关闭,就要在这边使用 gorutine 接收T raverseWithChannel 通道里的数据
	// 并且要使用Sleep保证数据能够传输完毕
	// go func() {
	// 	for node := range cNode {
	// 		if node.Value > maxNode {
	// 			maxNode = node.Value
	// 		}
	// 	}
	// }()
	// time.Sleep(time.Second)

	// 如果关闭了可以不使用groutine进行接收
	for node := range cNode {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}

	fmt.Println("Max node value:", maxNode)
}
