package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v len=%d, cap=%d \n", s, len(s), cap(s))
}

/**
slice 的一些操作
 */
func main() {
	fmt.Println("Create slice")
	// 直接创建一个slice
	var s []int // Zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	fmt.Println(s1)

	s2 := make([]int, 16)
	printSlice(s2)
	s3 := make([]int, 10, 32) // 第三个参数是cap,方便我们预估一个长度
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1) //复制slice
	printSlice(s2)

	fmt.Println("Delete element from slice")
	// 删除中间的第4个元素 8
	s2 = append(s2[:3], s2[4:]...) // 第二个参数,本质上是一个可变参数,要求是一个一个的element,但直接写太麻烦了,这里用语法 s2[:4]...把slice中的element给拿出来了
	printSlice(s2)

	fmt.Println("Pop from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Poping from tail")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}
