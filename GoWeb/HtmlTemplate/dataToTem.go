package main

import (
	"html/template"
	"net/http"
)

func welcome2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("GoWeb/HtmlTemplate/view/index.html")
	_ = t.Execute(w, "xiaopang")
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", welcome2)
	_ = server.ListenAndServe()
}
