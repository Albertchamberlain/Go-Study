package main

import "fmt"
import "net/http"

type MyHandler2 struct{}
type MyOtherHandler struct{}

func (mh *MyHandler2) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "第一个")
}
func (mh *MyOtherHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "第二个")
}

func main() {
	myhandler := MyHandler2{}
	myother := MyOtherHandler{}
	server := http.Server{
		Addr: "localhost:8090",
	}
	http.Handle("/myhandler", &myhandler)
	http.Handle("/myother", &myother)
	server.ListenAndServe()
}
