package main

import "fmt"

func fact(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fact(n-1)
}

func fib(n int) int  {
	if n <= 2{
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fact(7))

	// 尝试一个递归的菲波拉契数列计算(很慢)
	fmt.Println(fib(8)) // 别给太大数字了.100就跑很久了..恩,这个算法简单易理解,但是真的很慢啊!
}
