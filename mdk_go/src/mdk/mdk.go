package main

import (
	//"html/template"
	//"io/ioutil"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>as</h1><div>asd</div>")
}

func phase1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>xxx</h1><div>mmm</div>")
	fmt.Println("answer")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", viewHandler)
	r.HandleFunc("/get", phase1).Methods("GET")
	http.Handle("/", r)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
