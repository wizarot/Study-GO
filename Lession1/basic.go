package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

// 包内部变量
var a = 10
var b = "aaa"
//c := 1000 // 外面不能这样定义! 错误方式
var (
	c = 10
	d = 30
)

// 未复制变量
func variablesZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q \n", a, s)
}

// 初始化已复制变量
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 变量类型推断
func variableTypeDeduction() {
	var a, b, c, s = 4, 5, true, "deftt" //自动判断类型
	fmt.Println(a, b, c, s)
}

// 简短变量定义
func variableShorter() {
	a, b, c, s := 1, 2, true, "hello"
	fmt.Println(a, b, c, s)
}

// 验证下欧拉公式
func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1) // 浮点数
	fmt.Printf("%.3f \n", cmplx.Exp(1i*math.Pi)+1) // 浮点数,显示小数点后3位
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) // 隐式转换类型是不行的,必须强制转. int那一步转换,其实可能会有问题,如果float结果是4.99995...这样的,很可能转完变成了4! 那就麻烦了.需要注意!
	fmt.Println(c)
}

// 常量定义
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4 // 注意看这个类型 由于const仅做了相当于文本替换的动作
	var c int
	c = int(math.Sqrt(a*a + b*b)) // 这里如果没有定义类型,那么需要是什么类型,就可以当什么类型使用???
	fmt.Println(filename, a, b, c)
	//fmt.Printf("%s %d %s \n", a, a, b) // 但是当已经使用过一次,似乎这个变量类型就被定下来了.
}

// 枚举类型(并没有专用关键字)
func enums() {
	const (
		cpp    = iota // 自增表达式
		java
		_
		golang
		php
		python
	)
	fmt.Println(cpp, java, golang, php, python)
	// b ,kb,mb,gb,tb,pb
	const (
		b  = 1 << (10 * iota) // 这个用法碉堡了啊!
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

/*
变量要点回顾:
1. 变量类型写在变量名之后
2. 编译器可以推测变量类型
3. 没有char 只有rune
4. 原生支持复数
5. 枚举没有关键字,但可以使用const()来系列定义
 */
func main() {
	fmt.Println("Hello world! 你好 GO语言")
	variablesZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(a, b, c, d)
	/*
	内建变量类型:
	bool , string
	(u) + int,int8,int16,int32,int64     uintptr (指针ptr)
	byte , rune(字符型32位 4字节)
	float32,float64 complex64 ,complex128(复数,用于工程计算的好东西)
	 */
	// 验证欧拉公式
	euler()
	// GO只有强制类型转换,没有隐式类型转换
	triangle()
	// 常量
	consts()
	// 枚举
	enums()
}
