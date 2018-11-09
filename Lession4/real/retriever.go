package real

import (
	"time"
	"net/http"
	"net/http/httputil"
)

type Retriever struct {
	UserArgent string
	TimeOut    time.Duration
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()// 關閉
	if err != nil {
		panic(err)
	}
	return string(result)
}
