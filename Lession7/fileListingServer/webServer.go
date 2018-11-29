package main

import (
	"Study-GO/Lession7/fileListingServer/fileListing"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			// log.Warn("Error handling request: ", err.Error()) //为什么没有,这个是手动实现的
			log.Printf("Error handling request: %s", err.Error())
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code),
				code)
		}
	}
}

// 来个http服务器玩玩
func main() {
	http.HandleFunc("/list/",
		errWrapper(fileListing.HandleFileList))
	err := http.ListenAndServe(":8888", nil) // http://localhost:8888/list/fib.txt
	if err != nil {
		panic(err)
	}
}
