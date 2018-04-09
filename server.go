package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func StartServer(address string) error {
	router := mux.NewRouter()
	// Create a route for our REST API on the method GET for fizzbuzz.
	router.HandleFunc("/fizzbuzz/{str1:.+}/{str2:.+}/{int1:[0-9]+}/{int2:[0-9]+}/{limit:[0-9]+}",
		GetFizzBuzz).Methods("GET")
	return http.ListenAndServe(address, router)
}

// Endpoint GET Method '/fizzbuzz/str1/str2/int1/int2/limit'
// Generate a list of strings from 1 to limit where multiples of `int1` are replaced by `str1`,
// multiples of int2 replaced by str2, and multiples of `int1`*`int2` are replaced by
// the concatenation of `str1` and `str2`.
// Returns a list of strings encoded in JSON.
func GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	num1, err := strconv.Atoi(params["int1"])
	if err != nil {
		http.Error(w, "int1 is not a valid interger", http.StatusUnprocessableEntity)
		return
	}
	num2, err := strconv.Atoi(params["int2"])
	if err != nil {
		http.Error(w, "int2 is not a valid interger", http.StatusUnprocessableEntity)
		return
	}
	limit, err := strconv.Atoi(params["limit"])
	if err != nil {
		http.Error(w, "limit is not a valid interger", http.StatusUnprocessableEntity)
		return
	}
	fizzBuzzList, err := GenerateFizzBuzz(params["str1"], params["str2"], num1, num2, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	fields := FizzBuzzList{fizzBuzzList}
	err = json.NewEncoder(w).Encode(fields)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
