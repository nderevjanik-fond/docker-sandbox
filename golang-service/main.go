package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Server up and running...")
	http.ListenAndServe(":8080", nil)
}
