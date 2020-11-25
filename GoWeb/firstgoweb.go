package main

import "fmt"
import "net/http"

func welcome(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Golang")
}

func main() {
	http.HandleFunc("/", welcome)
	fmt.Println("服务已启动")
	_ = http.ListenAndServe("localhost:8090", nil)

}
