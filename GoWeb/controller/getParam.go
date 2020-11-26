package main

import "fmt"
import "net/http"

func param2(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	fmt.Fprintln(res, req.Form)

	name := req.FormValue("name")
	age := req.FormValue("age")

	fmt.Fprintln(res, name, age)
}

func main() {
	server := http.Server{
		Addr: "localhost:8090",
	}
	http.HandleFunc("/param", param2)
	_ = server.ListenAndServe()
}
