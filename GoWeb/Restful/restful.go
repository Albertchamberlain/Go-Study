package main

import (
	"fmt"
	"github.com/mux"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, _ = fmt.Fprintln(w, "dayinle", vars["key"])
}

func abc(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "abc")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{key}", hello)
	r.HandleFunc("/abc", abc)
	_ = http.ListenAndServe(":8090", r)

}
