package main

import (
	"html/template"
	"net/http"
)

func welcome12(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("GoWeb/htmltemplate/view/layout.html", "GoWeb/htmltemplate/view/head.html", "GoWeb/htmltemplate/view/foot.html")

	_ = t.ExecuteTemplate(w, "layout", nil)
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", welcome12)
	_ = server.ListenAndServe()
}
