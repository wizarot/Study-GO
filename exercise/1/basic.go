package main

import (
	"fmt"
	"strconv"
	os "os"
	"bufio"
	"math"
)

func enums() {
	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func sum(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func converToBin(num int) string {
	result := ""
	for ; num > 0; num /= 2 {
		result = strconv.Itoa(num%2) + result
	}
	return result
}

func printFile(filePath string) {
	// 获取文件资源 os.open
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	// bufio.newScanner
	for scanner.Scan() {
		// scan sc text
		fmt.Println(scanner.Text())
	}
}

// 函数作为参数传入
func apply(opFunction func(int, int) int, a, b int) int {
	return opFunction(a, b)
}

func sums(numbers ...int) {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	fmt.Println(sum)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

/**
需要反复练习
 */
func main() {
	fmt.Println("hello world!")
	enums()
	sum(10)
	// 转二进制
	fmt.Println(converToBin(4))
	printFile("/Users/will/go_productions/src/Study-GO/exercise/1/abc.txt")
	//
	fmt.Println(
		apply(func(i int, i2 int) int {
			return i * i2
		}, 3, 5))
	fmt.Println(
		apply(func(i int, i2 int) int {
			return int(math.Pow(float64(i), float64(i2)))
		}, 2, 6))

	sums(10, 10, 11)
	// 变量交换
	a, b := 1, 3
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

	// 训话字符串
	str := "🔥你好hello test多语言🔥"
	for key, value := range str {
		fmt.Println(key, "=>", string(value))
	}
	fmt.Println("======================")
	for key, value := range []rune(str) {
		fmt.Println(key, "=>", string(value))
	}

}
