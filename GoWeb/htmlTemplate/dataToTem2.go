package main

import (
	"html/template"
	"net/http"
)

func welcome3(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("GoWeb/htmlTemplate/view/index2.html")
	_ = t.Execute(w, "xiaopang") //此处传递数据
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", welcome3)
	server.ListenAndServe()
}
