package main

import (
	"html/template"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("GoWeb/HtmlTemplate/view/index.html")
	t.Execute(w, nil)
}

func main() {
	server := http.Server{Addr: ":8090"}

	http.Handle("/hello/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}
