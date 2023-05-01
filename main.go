package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
	fmt.Println("GET requested")
}

func main() {
	//fmt.Println("REST API listening...")
	//http.HandleFunc("/hello", helloHandler)
	//http.HandleFunc("/add", addHandler)
	//http.HandleFunc("/subtract", subtractHandler)
	//http.HandleFunc("/multiply", multiplyHandler)
	//http.HandleFunc("/divide", divideHandler)
	//http.HandleFunc("/exponentiate", exponentiateHandler)
	//http.HandleFunc("/squareroot", sqrtHandler)
	//log.Fatal(http.ListenAndServe(":3001", nil))
	nums := [2]int{16, 2}
	result := nums[0] / 2.0
	for i := 0; i < 10; i++ {
		result = result - ((result*result - nums[0]) / (2 * result))
	}
	fmt.Println("result = ", result)
}
