package main

import "fmt"

// 折半搜索和插入合成新切片.. 先不用这个了,太麻烦了..
// func halfFind(arr []int,find int) int  {
// 	k := len(arr) / 2
//
// 	return k
// }

// 插入排序
func insertSort(arr []int) []int {
	for i, v := range arr {
		// 取出,并循环i前面的切片, 前面认为是一个已排序切片,因此使用折半算法查找来˚减少点计算量
		// arr = append(halfFind(arr[:i-1],v) , arr[i:]...) // 注意! ...
		// fmt.Println(halfFind(arr[:i],v))
		for j, val := range arr[:i] { // 如果数组长度很大,还是值得做一个优化算法的
			if val > v {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// 插入排序,但这个似乎在插入过程中,怎么操作比较好呢?各种循环就很丑了.恩,用切片操作似乎是个比较好的主意
func main() {
	arr := []int{9, 10, 8, 7, 5, 6, 4, 3, 2, 1, 2, 3, 5, 6}
	fmt.Println(insertSort(arr))

}
