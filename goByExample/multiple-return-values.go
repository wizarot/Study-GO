package main

import "fmt"

func vals() (i, j int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)
}
