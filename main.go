package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", SayWelcome)
	fmt.Println("Server is running!")
	http.ListenAndServe(":8080", nil)
}

func SayWelcome(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello and welcome to Mutual Fund Portfolio Viewer!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}