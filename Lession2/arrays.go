package main

import "fmt"

func printArr(arr [5]int) {
	arr[0] = 100 // 尝试修改数组中的值
	for i, v := range arr {
		fmt.Println(i, "=>", v)
	}
}

func printArr1(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, "=>", v)
	}
	arr[0] = 100 // 尝试修改数组中的值(不需要使用 *arr ,这点有点奇怪?)
}

/**
数组:
1. [10]int 和 [20]int 是不同的值类型(传参会报错)
2. func f(arr [5]int) 会 拷贝 数组
3. 整个数组来回折腾真的很麻烦,go语言中一般不直接使用数组,我们用切片slice
 */
func main() {
	// 声明方式
	var arr1 [5]int // 默认值全是0..
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10} // 防止[]被认为是切片语法,所以里面需要加...代表需要编译器来计算下数组的大小
	fmt.Println(arr1, arr2, arr3)

	// 多维数组
	var grid [4][5] int
	fmt.Println(grid)

	// 遍历数组,方式1
	for i := 0; i < len(arr3); i++ { // len函数获取数组长度
		fmt.Println(arr3[i])
	}
	// 通过range关键字 获得数组的下标,和值(我爱php foreach)
	for i, v := range arr3 {
		fmt.Println(i, arr3[i], v)
	}
	// 只需要值
	for _, v := range arr3 {
		fmt.Println(v)
	}
	// range 很好,意义明确,美观

	// 数组是值类型?
	printArr(arr1)
	printArr(arr3)
	// printArr(arr2) // 数组参数长度不对了! 直接报参数错误
	// 尝试修改数组中的值 ,再输出,那么结果是没变..
	fmt.Println(arr1, arr3)

	// 如果在函数里修改值呢?那么需要传指针了,现在arr1[0]的值就被函数修改了.
	printArr1(&arr3)
	fmt.Println(arr3)

}
