package main

import "fmt"
import "net/http"

func first(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "第一个")
}
func second(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "第二个")
}

func main() {
	server := http.Server{
		Addr: "localhost:8090",
	}
	//注意此处使用HandleFunc函数
	http.HandleFunc("/first", first)
	http.HandleFunc("/second", second)
	_ = server.ListenAndServe()
}
