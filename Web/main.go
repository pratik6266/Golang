package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func handler (w http.ResponseWriter, r * http.Request){
	w.Write([]byte("<h1>Hello, World!</h1>"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")

	http.ListenAndServe(":8080", r)
}	