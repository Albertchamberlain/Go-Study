package main

import "fmt"
import "net/http"

type MyHandler struct {
}

func (mh *MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "输出内容")
}

func main() {
	myhandler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8090",
		Handler: &myhandler,
	}
	server.ListenAndServe()
}
