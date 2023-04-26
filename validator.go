package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

func validate(w http.ResponseWriter, r *http.Request) ([]int, error) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return nil, errors.New("invalid request method")
	}

	decoder := json.NewDecoder(r.Body)
	var nums []int
	err := decoder.Decode(&nums)
	if err != nil {
		http.Error(w, "Invalid request body.", http.StatusBadRequest)
		return nil, errors.New("invalid request body")
	}

	if len(nums) != 2 {
		http.Error(w, "Invalid number of operands.", http.StatusBadRequest)
		return nil, errors.New("invalid number of operands")
	}

	return nums, nil
}
