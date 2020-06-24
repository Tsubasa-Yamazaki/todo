package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/post", post)
	http.ListenAndServe(":8080", nil)
}

func post(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Hallo world!")
}