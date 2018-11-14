package mooc

import "fmt"

type Retriever struct {
	Contents string
}

func Stringer(r *Retriever) string {
	return fmt.Sprintf("this is r __toString(): {Contents= %s}", r.Contents)
}

func (r Retriever) Get(url string) string {
	return r.Contents
}
