package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, value := range nums {
		sum = sum + value
	}
	fmt.Println("sum:", sum)

	for key, value := range nums {
		if value == 3 {
			fmt.Println("3 at index:", key)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s \n", key, value)
	}

	for key, value := range "go中文" {
		// fmt.Println(key, rune(value))
		fmt.Printf("%d => %c \n", key, value)// %c输出rune
	}
}
