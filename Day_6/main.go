package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About"))
	}).Methods("GET")

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}).Methods("GET")

	port := "5000"

	fmt.Print("Server sedang berjalan di port " + port)
	http.ListenAndServe("localhost:"+port, route)
}
