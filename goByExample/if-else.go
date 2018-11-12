package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7是奇数")
	} else {
		fmt.Println("7是偶数")
	}
	if 8%4 == 0 {
		fmt.Println("8可以被4整除")
	}
	if num := 8; num < 0 {
		fmt.Println("负数")
	} else if num < 10 {
		fmt.Println("有一位数字")
	} else {
		fmt.Println("有多位数字")
	}
}
