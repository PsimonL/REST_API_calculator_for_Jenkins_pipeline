package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	fmt.Println("GET requested")
}

func main() {
	fmt.Println("REST API listening...")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/add", addHandler) // curl -X POST -H "Content-Type: application/json" -d "[5, 1]" http://localhost:3001/add
	http.HandleFunc("/subtract", subtractHandler)
	http.HandleFunc("/multiply", multiplyHandler)
	http.HandleFunc("/divide", divideHandler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
