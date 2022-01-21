package main

import (
	"Golang/quereis/handlers"
	"net/http"
	_"github.com/gorilla/mux"
)
func main() {
	port := ":2004"
	//r := mux.NewRouter()
	//http.HandleFunc("/books", handlers.Books)
	http.HandleFunc("/users", handlers.Users)
	//http.HandleFunc("/", handlers.QueryParams)
	http.ListenAndServe(port, nil)
}
