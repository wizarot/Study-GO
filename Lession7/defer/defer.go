package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func tryDefer() {
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println(3)
	// // return
	// panic("error occured!")
	// fmt.Println(4)
	for i := 0; i < 30; i++ {
		defer fmt.Println(i)
		if i == 20 {
			panic("too many i")
		}
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 测试下资源
func writeFile(filename string) {
	// file, err := os.Create(filename) // 新建文件
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666) // 人为弄个错误出来

	// 自定义一个error
	err = errors.New("this is a custom error")
	if err != nil {
		// panic(err)
		// fmt.Println("Error:",err.Error())
		// 再来一种err是一种*PathError
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err) // *pathError以外的其它类型Error
		} else {
			fmt.Printf("%s , %s, %s \n",
				pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()
	// 写入文件这里,见 https://www.imooc.com/learn/927
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	// tryDefer()
	// 输出文件
	writeFile("fib.txt")
}
