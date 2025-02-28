package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello, World!")
	})

	http.Handle("/", r)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
