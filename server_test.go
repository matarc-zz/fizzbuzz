package main

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
)

func TestGetFizzBuzz(t *testing.T) {
	res, err := http.Get("http://localhost:8080/fizzbuzz/fd/fd/0/1/1")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 422 {
		t.Errorf("The server should return an error since int1 is not > 0")
	}
	res.Body.Close()

	res, err = http.Get("http://localhost:8080/fizzbuzz/fd/fd/1/0/1")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 422 {
		t.Errorf("The server should return an error since int2 is not > 0")
	}
	res.Body.Close()

	res, err = http.Get("http://localhost:8080/fizzbuzz/fd/fd/1/1/0")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 422 {
		t.Errorf("The server should return an error since the limit is not > 0")
	}
	res.Body.Close()

	res, err = http.Get("http://localhost:8080/fizzbuzz/fd/fd/dfsd/1/0")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 404 {
		t.Errorf("The server should return a 404 error since int1 is not a valid number")
	}
	res.Body.Close()

	res, err = http.Get("http://localhost:8080/fizzbuzz/fd/fd/1/fds/0")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 404 {
		t.Errorf("The server should return a 404 error since int2 is not a valid number")
	}
	res.Body.Close()

	res, err = http.Get("http://localhost:8080/fizzbuzz/fd/fd//1/0/fsdd")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 404 {
		t.Errorf("The server should return a 404 error since limit is not a valid number")
	}
	res.Body.Close()

	res, err = http.Get("http://localhost:8080/fizzbuzz/fizz/buzz/3/5/20")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Status code should be 200, no error in this API call")
	}

	fields := FizzBuzzList{}
	err = json.NewDecoder(res.Body).Decode(&fields)
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
	expectedList := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13",
		"14", "fizzbuzz", "16", "17", "fizz", "19", "buzz"}
	compareFizzBuzzList(t, expectedList, fields.List, len(expectedList))
}

func TestMain(m *testing.M) {
	go StartServer(":8080")
	os.Exit(m.Run())
}
