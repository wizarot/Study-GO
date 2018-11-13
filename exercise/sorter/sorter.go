package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
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
	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("Read values:", values)
	} else {
		fmt.Println(err)
	}
}

// 读取文件
func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Faild to open the input file ", file)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line,seems unexpected.")
			return
		}
		str := string(line) //转换字符数组为字符串
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}
