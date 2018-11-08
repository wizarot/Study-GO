package main

import (
	"fmt"
	"Study-GO/Lession4/mooc"
	"Study-GO/Lession4/real"
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
	r = real.Retriever{}
	fmt.Println(download(r))
}
