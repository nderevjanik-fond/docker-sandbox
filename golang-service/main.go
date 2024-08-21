package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	now := time.Now()
	fmt.Fprintf(w, "%s - Hello World!\n", now)
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Server up and running...")
	http.ListenAndServe(":8080", nil)
}
