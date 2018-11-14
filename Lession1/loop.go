package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"time"
	"io"
	"strings"
)

// 计算1一直加到100
func sum() int {
	sum := 0
	// for 不需要括号,三个表达式也可以省略
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum
}

// 转化成二进制
func convertToBin(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result // Itoa 整数转字符串 Atoi字串转整形. strconv 字符串转数字,布尔等的库
	}
	return result
}

// 打印文件内容,为了防止编译出来执行的地方和文件不在同一文件件的问题,建议还是给一个路径吧!
func printFile(filename string) {
	file, err := os.Open(filename) // 获取文件资源?
	if err != nil {
		panic(err)
	}
	printFileContants(file)
}

// 使用reader 对象可以传入各种允许的reader
func printFileContants(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() { // 没有初始条件,没有递增语句. 这样的写法就相当于while,因此go语言没有涉及while语法
		fmt.Println(scanner.Text())
	}
}

// 写个死循环,ctrl+z退出
func forever() {
	for { // 什么都不写,这样就可以死循环,不用while true这类语法了.
		fmt.Println("Hello")
		time.Sleep(1 * time.Second)
	}
}

/**
要点总结:
1. for if 后面没有括号
2. if 条件里也可以定义变量
3. 没有while ,可以用for的不同写法来完成
4. switch不需要break,可以直接写switch多个条件.(其实感觉就是一堆的if else if ..这种)
 */
func main() {

	fmt.Println(sum())
	fmt.Println(convertToBin(5), // 101
		convertToBin(13),        // 1101
		convertToBin(78348121),  // 1101
		convertToBin(0),         // 1101
	)

	printFile("/Users/will/go_productions/src/Study-GO/Lession1/abc.txt") // 需要根据实际地址手动更改!我这个项目就放这里了.
	s := `sdhfjekjk  
kdjkf
skdjkfj
aaa
bbb

中文
`
	printFileContants(strings.NewReader(s))
	// forever()
}
