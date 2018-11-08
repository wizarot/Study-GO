package main

import (
	"Study-GO/Lession3/tree"
	"fmt"
)


// 通过别名,给tree实现一个后序遍历
type myTreeNode struct {
	node *tree.Node
}

// 实现我们自己的比遍历
func (myNode *myTreeNode)postOrder()  {
	if myNode == nil || myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root.Value = 3
	root.Right = tree.CreateNode(5)
	root.Left = &tree.Node{}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Println("中序遍历:")
	root.Traverse()
	fmt.Println("")
	fmt.Println("后序遍历:")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()

}
