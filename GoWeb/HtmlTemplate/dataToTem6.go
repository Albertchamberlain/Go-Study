package main

import (
	"html/template"
	"net/http"
	"time"
)

func MyFormat(s string) string {
	t, _ := time.Parse("2020-11-26 11:02:03", s)
	t = t.Add(60e9)
	return t.Format("2020-11-26 11:02:03")
}

func html(res http.ResponseWriter, req *http.Request) {
	funcMap := template.FuncMap{"mf": MyFormat}
	t := template.New("GoWeb/HtmlTemplate/view/index6.html").Funcs(funcMap)
	t, _ = t.ParseFiles("GoWeb/HtmlTemplate/view/index6.html")
	s := "2020-11-26 11:02:03"
	_ = t.Execute(res, s)
}

func main() {
	server := http.Server{
		Addr: "localhost:8090",
	}
	http.HandleFunc("/time", html)
	_ = server.ListenAndServe()
}
