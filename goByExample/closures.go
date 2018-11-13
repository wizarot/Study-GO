package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 另一个菲波拉契数列实现,这个比较快,然而可能算100还是有点大了.速度还可以,貌似整形溢出了?没验证..
func fibPack() func() uint64 {
	var i, j uint64 = 1, 1
	return func() uint64 {
		temp := uint64( i + j )
		i, j = j, temp
		return temp
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fibInt := fibPack()
	for i := 0; i < 100; i++ { // 循环100次
		fmt.Println(fibInt())
	}
}
