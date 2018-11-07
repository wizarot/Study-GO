package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

/**
String字符串相关类型,了解下
使用range 遍历pos,rune对
使用utf8.RuneCountInString(s)获得字符数量,简单使用len(s)是不行的
使用len(s)获得字节长度
使用[]byte获得字节
 */
func main() {
	s := "Yes我爱GO语言!🏀⚽⚡👄👍🔥" // 视觉上是16个字符
	fmt.Println(len(s))         // 检查下长度 len为40
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) // 每个中文和特殊符号占3字节(byte) UTF-8编码
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune 按照正确形式进行了截取
		fmt.Printf("(%d , %X)", i, ch) // %c可以输出字符 %X输出大写16进制
	}
	fmt.Println()

	fmt.Println("Rune count in string:", utf8.RuneCountInString(s)) // 根据rune来算有多少个字符

	// 比较正确的循环字符串,并且定位到第几个字符的方式
	for i, ch := range []rune(s) { // 切片..rune(s)
		fmt.Printf("(%d , %c)", i, ch)
	}
	fmt.Println()

	// 其它字符串操作: 在包strings中
	fmt.Println(strings.Compare("aaak", "aaa")) // 需要好好看看strings包中的常见方法
}
