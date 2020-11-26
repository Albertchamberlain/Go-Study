package main

import (
	"html/template"
	"net/http"
	"time"
)

func welcome6(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("GoWeb/HtmlTemplate/view/index5.html")
	time := time.Date(2020, 11, 26, 10, 54, 5, 0, time.Local)
	_ = t.Execute(w, time)
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", welcome6)
	_ = server.ListenAndServe()
}
