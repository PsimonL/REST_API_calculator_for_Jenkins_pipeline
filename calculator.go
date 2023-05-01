package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	nums, err := validate(w, r)
	if err != nil {
		return
	}

	result := nums[0] - nums[1]
	json.NewEncoder(w).Encode(result)
	fmt.Printf("Subtraction: %d - %d = %d\n", nums[1], nums[0], result)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	nums, err := validate(w, r)
	if err != nil {
		return
	}

	result := nums[0] + nums[1]
	json.NewEncoder(w).Encode(result)
	fmt.Printf("Addition: %d + %d = %d\n", nums[0], nums[1], result)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	nums, err := validate(w, r)
	if err != nil {
		return
	}

	result := nums[0] * nums[1]
	json.NewEncoder(w).Encode(result)
	fmt.Printf("Multiplication: %d * %d = %d\n", nums[0], nums[1], result)
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	nums, err := validate(w, r)
	if err != nil {
		return
	}

	if nums[1] == 0 {
		http.Error(w, "Division by zero!", http.StatusBadRequest)
		fmt.Println("Division by zero!")
		return
	}

	result := nums[0] / nums[1]
	json.NewEncoder(w).Encode(result)
	fmt.Printf("Division: %d / %d = %d\n", nums[0], nums[1], result)
}

func exponentiateHandler(w http.ResponseWriter, r *http.Request) {
	nums, err := validate(w, r)
	if err != nil {
		return
	}

	result := 1
	for i := 0; i < int(nums[1]); i++ {
		result *= nums[0]
	}
	json.NewEncoder(w).Encode(result)
	fmt.Printf("Exponentiation: %d ^ %d = %d\n", nums[0], nums[1], result)
}

func sqrtHandler(w http.ResponseWriter, r *http.Request) {
	nums, err := validate(w, r)
	if err != nil {
		return
	}

	result := nums[0] / 2.0
	for i := 0; i < 10; i++ {
		result = result - ((result*result - nums[0]) / (2 * result))
	}
	json.NewEncoder(w).Encode(result)
	fmt.Printf("Exponentiation: %d ^ 1/%d = %d\n", nums[0], nums[1], result)
}
