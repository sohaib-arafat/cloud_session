package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi from Cloud Run!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
