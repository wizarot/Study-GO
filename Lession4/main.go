package main

import (
	"fmt"
	"Study-GO/Lession4/mooc"
	"Study-GO/Lession4/real"
	"time"
)

type Retriver interface {
	Get(url string) string
}

func download(r Retriver) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriver
	r = mooc.Retriever{"this is a fake imooc.com"}
	inspect(r)
	r = &real.Retriever{
		UserArgent: "Chrome/12.0",
		TimeOut:    time.Minute,
	}
	inspect(r)
	// fmt.Println(download(r))
	// Type assertion类型断言?
	// realRetriever := r.(*real.Retriever)
	// mookRetriever := r.(mooc.Retriever)// TODO: 这段会报错,但不是很明白Type assertion的具体作用.
	// inspect(mookRetriever)
	if mockRetriever, ok := r.(mooc.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		inspect(mockRetriever)
		fmt.Println("not a mock retriever")
	}
	// inspect(realRetriever)
}

func inspect(r Retriver) {
	fmt.Printf("%T %v \n", r, r) // 显示下r是什么东西  %T显示变量类型, %v是变量值?
	switch v := r.(type) {
	case mooc.Retriever:
		fmt.Println("Contents", v.Contents)
	case *real.Retriever:
		fmt.Println("UserArgent", v.UserArgent)

	}
}
