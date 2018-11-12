package main

import "fmt"

func sum(numbers ...int) { // 可变参数
	fmt.Println(numbers, "")
	total := 0
	for _, value := range numbers {
		total += value
	}
	fmt.Println("sum:", total)
}

func main() {
	sum(1, 2, 3)
	sum(1, 2, 3, 4, 5)

	nums := []int{1, 2, 3, 4}
	sum(nums...) // slice作为参数时,有点特殊
}
