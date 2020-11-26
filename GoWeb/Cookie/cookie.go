package main

import (
	"html/template"
	"net/http"
)

func welcome55(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("GoWeb/Cookie/view/index.html")
	t.Execute(w, nil)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{Name: "mykey", Value: "myvalue"}
	http.SetCookie(w, &c)
	t, _ := template.ParseFiles("GoWeb/Cookie/view/index.html")
	t.Execute(w, nil)

}
func getCookie(w http.ResponseWriter, r *http.Request) {
	cs := r.Cookies()
	t, _ := template.ParseFiles("GoWeb/Cookie/view/index.html")
	t.Execute(w, cs)
}
func main() {
	server := http.Server{Addr: ":8090"}
	http.Handle("GoWeb/Cookie/static/", http.StripPrefix("GoWeb/Cookie/static/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome55)
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	server.ListenAndServe()
}
