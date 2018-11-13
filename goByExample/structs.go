package main

import "fmt"

type persion struct {
	name string
	age  int
}

func main() {
	fmt.Println(persion{"Bob", 20})
	fmt.Println(persion{name: "Alice", age: 30})
	fmt.Println(persion{name: "Frand"})
	fmt.Println(&persion{name: "Ann", age: 40})

	sean := persion{name: "Sean", age: 50}
	fmt.Println(sean.name)

	sp := &sean
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sean.age)
}
