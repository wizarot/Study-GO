package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

// 为struct定义方法(不写在结构体中,而是写在外面..)
func (node treeNode) print() { // 在函数名前面加上 括号,接受者
	fmt.Println(node.value)
}

// 那么这个方法是传值还是穿引用的呢?
func (node *treeNode) setValue(value int) { // 如果是 node treeNode那么是传值的,因此这里修改是不会传递回原变量的.改成node *treeNode这样就是传指针了,其它不用动,但这样就可以正确被修改了!
	if node == nil {
		fmt.Println("禁止在nil节点设置值")
		return
	}
	node.value = value
}

// 没构造函数,但可以设定工厂函数.注意需要返回局部变量的地址
func createNode(value int) *treeNode {
	return &treeNode{value: value} // 产生的是一个局部变量,但是C这里是个典型错误,然而GO语言这样是可以的
	// 那么问题来了,这个局部变量放在堆上还是栈上? 答案是不需要管,反正go垃圾回收器足够好,可以自动判断
}

// 实现一个遍历各节点的函数
func (node *treeNode)traverse()  {
	// 防止nil节点情况
	if node == nil {
		return
	}
	// 中序遍历-先遍历左子树,再来右子树
	node.left.traverse()
	node.print()
	node.right.traverse()
}

/**
面向对象
go仅支持封装,不支持继承和多态
继承和多态使用面向接口编程来完成
go语言没有class,只有struct
如果包里面有指针接受者,那么建议全部都用指针接收者,这样可以保持一致性.
 */
func main() {
	// 新建node
	// var root treeNode
	root := treeNode{value: 3}
	root.left = &treeNode{}             // 指针类型,赋值指向的是地址
	root.right = &treeNode{5, nil, nil} // 不指定,直接写这样也可以
	root.right.left = new(treeNode)     // 用new函数,返回treeNode地址
	root.left.right = createNode(2)     // 自建构造函数
	fmt.Println(root)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	// 结构体方法的使用
	root.print()
	fmt.Println("set value?")
	root.setValue(4)
	root.print()

	// go语言的指针操作特殊性
	pRoot := &root
	pRoot.print()
	pRoot.setValue(4)
	pRoot.print()
	root.print()

	// nil指针也可以正常调用方法,但是会在使用过程中,由于找不到对应成员值而报个错误
	var nilNode *treeNode
	nilNode.setValue(100)

	fmt.Println("traverse tree:")
	// 中序遍历root树
	root.traverse()
}
