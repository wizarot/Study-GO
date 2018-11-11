package main

import (
	"flag"
	"fmt"
)

var infile *string = flag.String("i", "infile", "File Contains values for sorting.")
var outfile *string = flag.String("o", "outfile", "file to receie sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort alogorithm")

/**
练习一个命令行可以输入文件,排序输出的功能
 */
func main() {
	// 处理命令行输入参数,使用flag包
	flag.Parse()
	if infile != nil {
		fmt.Println("infile=", *infile, "outfile=", *outfile, "algorithm=", *algorithm)
	}
}
