package main

import (
	"html/template"
	"net/http"
)

//注意:只有首字母大写的属性才能在模版中访问到
type User struct {
	Name string
	Age  int
}

func welcome4(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("GoWeb/HtmlTemplate/view/index3.html")
	_ = t.Execute(w, User{"xiaopang", 20}) //此处传递数据
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", welcome4)
	server.ListenAndServe()
}
