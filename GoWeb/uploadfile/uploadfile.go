package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func welcome13(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("GoWeb/uploadfile/view/index.html")
	_ = t.Execute(w, nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	fileName := r.FormValue("name")
	file, fileHeader, _ := r.FormFile("file")
	b, _ := ioutil.ReadAll(file)

	_ = ioutil.WriteFile("D:/"+fileName+fileHeader.Filename[strings.LastIndex(fileHeader.Filename, "."):], b, 0777)
	t, _ := template.ParseFiles("GoWeb/uploadfile/view/success.html")
	_ = t.Execute(w, nil)
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", welcome13)
	http.HandleFunc("/upload", upload)
	_ = server.ListenAndServe()
}
