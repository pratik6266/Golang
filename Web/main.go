package main

import (
	"fmt"
	"net/http"
	"web/routers"
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
	r := routers.Router()
	r.HandleFunc("/", handler).Methods("GET")
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}  