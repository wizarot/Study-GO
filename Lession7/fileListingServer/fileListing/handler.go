package fileListing

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):] // 获取到字符串 "/list/fib.txt",因此进行下截取
	file, err := os.Open(path)
	if err != nil {
		// http.Error(writer,
		// 	err.Error(),
		// 	http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all) // response
	return nil
}
