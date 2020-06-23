package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/post", post)
	http.ListenAndServe(":8080", nil)
}

func post(w http.ResponseWriter, r *http.Request) {
	type data struct {
		importance string `json:"importance"`
    	task       string `json:"task"`
    	limit      string `json:"limit"`
	}

	fmt.Println("Hallo world!")
}