package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Userr struct {
	Name string
	Age  int
}

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("GoWeb/Ajax/view/index.html")
	_ = t.Execute(w, nil)
}
func getUser(w http.ResponseWriter, r *http.Request) {
	users := make([]Userr, 0)
	users = append(users, Userr{"xiaopang", 20})
	users = append(users, Userr{"xiaoming", 21})
	users = append(users, Userr{"xiaohong", 22})
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	b, _ := json.Marshal(users)
	fmt.Fprintln(w, string(b))
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.Handle("GoWeb/Ajax/static/", http.StripPrefix("GoWeb/Ajax/static", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	http.HandleFunc("/getUser", getUser)
	_ = server.ListenAndServe()
}
