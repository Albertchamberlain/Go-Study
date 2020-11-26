package main

import (
	"html/template"
	"net/http"
)

func welcome33(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("GoWeb/Cookie/view/indexjq.html")
	t.Execute(w, nil)
}
func setCookie33(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{Name: "mykey", Value: "myvalue", HttpOnly: false}
	http.SetCookie(w, &c)
	t, _ := template.ParseFiles("GoWeb/Cookie/view/indexjq.html")
	t.Execute(w, nil)

}

func main() {
	server := http.Server{Addr: ":8090"}
	http.Handle("GoWeb/Cookie/static/", http.StripPrefix("GoWeb/Cookie/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome33)
	http.HandleFunc("/setCookie", setCookie33)
	_ = server.ListenAndServe()
}
