package main

import (
	"fmt"
	"time"
)

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ts := time.Now().UnixNano()
// 	h := 1000000000
// 	sum := 0

// 	for i := 0; i <= h; i++ {
// 		sum = sum + i
// 	}

// 	fmt.Println(sum)
// 	fmt.Print("时间毫秒：")
// 	fmt.Println((time.Now().UnixNano() - ts) / 1000000)
// }

func Count1(start int, end int, ch chan int) {
	var cccc int
	for j := start; j < end; j++ {
		cccc = cccc + j
	}
	ch <- cccc
}

func main() {
	ts := time.Now().UnixNano()
	h := 1000000000
	sum := 0
	ch := make(chan int, 100)
	numLength := cap(ch)
	fmt.Println(numLength)

	for i := 0; i < numLength; i++ {
		num := h / numLength
		go Count1(num*i, num*i+num, ch)
	}
	for i := 0; i < numLength; i++ {
		select {
		case msg := <-ch:
			sum = sum + msg

		}
	}
	fmt.Println(sum + h)
	fmt.Print("时间毫秒：")
	fmt.Println((time.Now().UnixNano() - ts) / 1000000)
}
