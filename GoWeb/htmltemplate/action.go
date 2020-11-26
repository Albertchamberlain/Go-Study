package main

import (
	"html/template"
	"net/http"
)

func test(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("GoWeb/htmltemplate/view/ifelseaction.html")
	//第二个参数传递类型默认值:nil,"",0,false等都会导致if不成立
	t.Execute(rw, "")
}

func main() {

	server := http.Server{Addr: ":8090"}

	http.HandleFunc("/test", test)

	server.ListenAndServe()
}
