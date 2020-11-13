package tree

import "fmt"

// Node 是一个二叉树节点结构
type Node struct {
	Value int
	Left, Right *Node
}

// CreateNode 方法用于控制初始化
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

// Print 打印节点的值
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// SetValue 设置节点 Value 的值
func (node *Node) SetValue(value int)  {
	if node == nil {
		fmt.Println("Setting value to nil "+
		"node. Ignord.")
		return
	}
	node.Value = value
}
