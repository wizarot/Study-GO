package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 = generator(), generator()
	// n1 := <-c1
	// n2 := <-c2
	// 我想同时收c1,c2 哪个先到收哪个
	for {
		select {
		case n := <-c1:
			fmt.Println("Select from C1 ", n)
		case n := <-c2:
			fmt.Println("select from C2 ", n)
		}
	}

}
