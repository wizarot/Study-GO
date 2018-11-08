package tree

/**
各种遍历方法放一个文件里
 */

// 实现一个遍历各节点的函数
func (node *Node)Traverse()  {
	// 防止nil节点情况
	if node == nil {
		return
	}
	// 中序遍历-先遍历左子树,再来右子树
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}