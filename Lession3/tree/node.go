package tree

import "fmt"

/**
封装为包:
命名,名字应该是CamelCase
首字母大写 Public
首字母小写 private
每个目录都是一个包. 每个目录里只能有一个main函数,一个包里可以有多个文件 main包包含的是一个可执行的入口
为结构定义的方法必须放在同一个包内,但可以不是一个文件
 */

type Node struct {
	Value       int
	Left, Right *Node
}

// 为struct定义方法(不写在结构体中,而是写在外面..)
func (node Node) Print() { // 在函数名前面加上 括号,接受者
	fmt.Print(node.Value," ")
}

// 那么这个方法是传值还是穿引用的呢?
func (node *Node) SetValue(value int) { // 如果是 node Node那么是传值的,因此这里修改是不会传递回原变量的.改成node *Node这样就是传指针了,其它不用动,但这样就可以正确被修改了!
	if node == nil {
		fmt.Println("禁止在nil节点设置值")
		return
	}
	node.Value = value
}

// 没构造函数,但可以设定工厂函数.注意需要返回局部变量的地址
func CreateNode(value int) *Node {
	return &Node{Value: value} // 产生的是一个局部变量,但是C这里是个典型错误,然而GO语言这样是可以的
	// 那么问题来了,这个局部变量放在堆上还是栈上? 答案是不需要管,反正go垃圾回收器足够好,可以自动判断
}

