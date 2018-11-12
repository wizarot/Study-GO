package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Printf("Write %d as ", i)
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("今天周末")
	default:
		fmt.Println("今天工作日")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("上午")
	default:
		fmt.Println("下午")
	}

}
