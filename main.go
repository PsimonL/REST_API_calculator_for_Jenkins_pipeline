package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	fmt.Println("GET requested")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
