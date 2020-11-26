package main

import (
	"html/template"
	"net/http"
)

//注意:只有首字母大写的属性才能在模版中访问到
type User2 struct {
	Name string
	Age  int
}

func welcome5(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("GoWeb/htmltemplate/view/index4.html")
	m := make(map[string]interface{})
	m["user"] = User2{"xiaopang", 20}
	m["money"] = 10001
	_ = t.Execute(w, m)
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", welcome5)
	_ = server.ListenAndServe()
}
