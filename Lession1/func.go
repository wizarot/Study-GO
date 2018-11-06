package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

// 返回错误信息
func eval(a, b int, op rune) (int, error) {
	var result int
	switch op {
	case '+':
		result = a + b // switch会自动break ,除非fallthrough
		//fallthrough 除非使用fallthrough, 节省很多代码啊!
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		result, _ = div(a, b)
	default:
		return 0, fmt.Errorf("不支持的操作符:" + string(op)) // 错误信息?的一种写法
	}
	return result, nil
}

// 返回多个值
func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 函数作为参数传入?
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()    // 取得操作函数的指针?
	opName := runtime.FuncForPC(p).Name() // 取得指针对应函数名称?
	fmt.Printf("Calling function %s with args (%d , %d) \n", opName, a, b)
	return op(a, b)
}

// 准备好一个作为参数传入的函数
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数
func sum1(numbers ... int) int {
	s := 0
	for i := range numbers { // range遍历可循环变量
		s += numbers[i] // 这个参数被认为是个数组?
	}
	return s
}

// golang函数不支持默认参数..这个有点奇怪,然而人家 要求必须是显式调用的似乎也有些道理?
/**
要点:
1. 返回值类型写在最后面
2. 可以返回多个值
3. 函数本身可以作为参数
4. 没有默认参数,可选参数
5. 可以使用可变参数
 */
func main() {
	if result, err := eval(123, 5, 'x'); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Println(q, r)

	// 传入满足定义的函数进去
	fmt.Println(apply(pow, 2, 6))
	// 直接作为匿名函数直接传入参数(有点函数式编程的意思在里面了.)
	fmt.Println(apply(func(a int, b int) int {
		return a * b
	}, 3, 5))

	fmt.Println(sum1(1, 2, 3, 4, 5, 6, 7))
}
