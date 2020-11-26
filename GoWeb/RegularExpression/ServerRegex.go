package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

func welcome99(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("GoWeb/RegularExpression/view/index.html")
	t.Execute(w, nil)
}

func register99(w http.ResponseWriter, r *http.Request) {
	{
		u := r.FormValue("username")
		r, _ := regexp.MatchString(`^[0-9a-zA-Z]{6,12}$`, u)
		if r {
			fmt.Fprintln(w, "注册成功")
		} else {
			fmt.Fprintln(w, "用户名格式不正确")
		}
	}
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.Handle("GoWeb/RegularExpression/static/", http.StripPrefix("GoWeb/RegularExpression/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome99)
	http.HandleFunc("/register", register99)
	_ = server.ListenAndServe()
}
