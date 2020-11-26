package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("GoWeb/downloadfile/view/index.html")
	_ = t.Execute(w, nil)
}
func download(w http.ResponseWriter, r *http.Request) {
	filename := r.FormValue("filename")
	f, err := ioutil.ReadFile("D:/go/" + filename)
	if err != nil {
		_, _ = fmt.Fprintln(w, "文件下载失败,", err)
		return
	}
	h := w.Header()
	h.Set("Content-Type", "application/octet-stream")
	h.Set("Content-Disposition", "attachment;filename="+filename)
	_, _ = w.Write(f)
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", welcome)
	http.HandleFunc("/download", download)
	_ = server.ListenAndServe()
}
