package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var nums []int
	err := decoder.Decode(&nums)
	if err != nil {
		http.Error(w, "Invalid request body.", http.StatusBadRequest)
		return
	}

	if len(nums) != 2 {
		http.Error(w, "Invalid number of operands.", http.StatusBadRequest)
		return
	}

	result := nums[0] + nums[1]
	json.NewEncoder(w).Encode(result)
	fmt.Printf("Added %d and %d to get %d\n", nums[0], nums[1], result)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	fmt.Println("GET requested")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/add", addHandler) // curl -d [2,3] -H "Content-Type: application/json" http://localhost:3001/add
	log.Fatal(http.ListenAndServe(":3001", nil))
}
