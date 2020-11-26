package main

import "fmt"
import "net/http"

func param(res http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintln(res, "第一个")
	header := req.Header
	_, _ = fmt.Fprintln(res, "Header全部数据:", header)
	var acc []string = header["Accept"]
	for _, n := range acc {
		_, _ = fmt.Fprintln(res, "Accepth内容:", n)
	}
}

func main() {
	server := http.Server{
		Addr: "localhost:8090",
	}
	http.HandleFunc("/param", param)
	_ = server.ListenAndServe()
}
