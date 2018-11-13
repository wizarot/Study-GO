package main

import "fmt"

type rector struct {
	width, height int
}

func (r *rector) Area() int {
	return r.width * r.height
}

func (r *rector) Perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rector{10, 5}
	fmt.Println(r)
	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())

	rp := &r
	fmt.Println(rp.Area())
	fmt.Println(rp.Perimeter())
}
