package main

import "fmt"

func updateSlice(s []int) { // []不标长度,那么它就是一个slice
	s[0] = 100
}

/**
切片:
1. 左开右闭区间? 看是从几开始数数咯~
2. slice本质上是对array的一个view(视图)
3. cap() -> capacity容量[kəˈpæsɪti]
4. 直接通过下标取不到,然而通过切片是可以的:  ptr开头指针,len切片长度,cap指向实际数组剩余长度; slice可以向后扩展,但不能向前扩展 sa2[i]不可以超过len(sa2),向后扩展,不可以超过底层数组cap(sa2)
5. 如果想slice中添加元素append超越原来数组大小,那么系统将会重新分配一个更大的底层数组
6. 由于append采用值传递的关系,必须接收append的返回值
 */
func main() {
	// 定义一个数组
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:] =", arr[:])
	fmt.Println("arr[0:] =", arr[0:])
	fmt.Println("arr[1:] =", arr[1:])
	fmt.Println("对slice进行修改-------------------")
	s1 := arr[2:]
	fmt.Println("arr[2:] =", s1)
	s2 := arr[:]
	fmt.Println("arr[:] =", s2)
	fmt.Println("进行修改后-------------------------")
	updateSlice(s1)
	fmt.Println(s1)  // 视图被修改了
	fmt.Println(s2)  // 相关的视图也被修改了
	fmt.Println(arr) // 原始数据也被修改了

	// Reslice
	fmt.Println("Reslice :")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	// slice扩展
	fmt.Println("Extending slice :")
	arr1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	sa1 := arr1[2:6]
	sa2 := sa1[3:5] // sa2? 直接通过下标取不到,然而通过切片是可以的:  ptr开头指针,len切片长度,cap指向实际数组剩余长度; slice可以向后扩展,但不能向前扩展 sa2[i]不可以超过len(sa2),向后扩展,不可以超过底层数组cap(sa2)
	fmt.Println(arr1, sa1, sa2)
	fmt.Printf("sa1=%v, len(sa1)=%d, cap(sa1)=%d\n", sa1, len(sa1), cap(sa1))
	fmt.Printf("sa2=%v, len(sa2)=%d, cap(sa2)=%d\n", sa2, len(sa2), cap(sa2))

	//添加元素
	fmt.Println("Append data:")
	s3 := append(sa2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	// s4 和 s5 会开一个新的数组,不是原来的arr1了!!
	fmt.Println("s3 , s4 ,s5 =", s3, s4, s5)
	fmt.Println("arr=",arr1)

}
