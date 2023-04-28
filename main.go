package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
	fmt.Println("GET requested")
}

func main() {
	fmt.Println("REST API listening...")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/subtract", subtractHandler)
	http.HandleFunc("/multiply", multiplyHandler)
	http.HandleFunc("/divide", divideHandler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
