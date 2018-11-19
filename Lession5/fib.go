package main

import (
	"fmt"
	"io"
	"bufio"
	"strings"
)

// 经典的斐波那契数列又来了!
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}

}

type intGen func() int

// 给函数实现接口.. GO的一个特点
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		// 不能无限循环下去,这里设定一个上限值
		return 0,io.EOF // 文件截止符号,read相当于将函数当做文件来读取?
	}
	s := fmt.Sprintf("%d \n",next)
	// TODO: 注意這裡沒處理p太小的情況??? 貌似也沒聽懂
	return strings.NewReader(s).Read(p) // 代理实现? 没看懂这里..
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()

	// fmt.Println(f()) // 1
	// fmt.Println(f()) // 1
	// fmt.Println(f()) // 2
	// fmt.Println(f()) // 3
	// fmt.Println(f()) // 5
	// fmt.Println(f()) // 8
	// fmt.Println(f()) // 13
	// fmt.Println(f()) // 21

	// 实现文件的read接口后,直接用函数读取全部数据
	printFileContents(f)
}
