package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func handler (w http.ResponseWriter, r * http.Request){
	w.Write([]byte("<h1>Hello, World!</h1>"))
	// another way to send data
	//! json.NewEncoder(w).Encode("This is the data to be send");
	// how to get parameters from the URL
	// fmt.Println(r.URL.Path)
	// fmt.Println(r.URL.Query())
	//! params := mux.Vars(r)	
	//* fmt.Println(params["name"])	
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")

	http.ListenAndServe(":8080", r)
}	