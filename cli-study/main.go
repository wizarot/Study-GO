package main

import (
	"bufio"
	"fmt"
	"os"
)

var usage = `
Usage说明: -port=9000 指定端口号

其它... 
`

// 使用flag包获取参数
func main() {
	scanner := bufio.NewScanner(os.Stdin) // 有stdin获取数据
	fmt.Println("请输入内容: ")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}
		fmt.Println("line: ", line)
	}
	// 错误处理,为什么放循环外面呢?
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
