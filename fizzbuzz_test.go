package main

import (
	"testing"
)

func TestGenerateFizzBuzz(t *testing.T) {
	fizzBuzzList, err := GenerateFizzBuzz("", "buzz", 3, 5, 100)
	if err != ErrEmptyString {
		t.Errorf("GenerateFizzBuzz should have returned `ErrEmptyString` when `str1` is empty.")
	}
	fizzBuzzList, err = GenerateFizzBuzz("fizz", "", 3, 5, 100)
	if err != ErrEmptyString {
		t.Errorf("GenerateFizzBuzz should have returned `ErrEmptyString` when `str2` is empty.")
	}
	fizzBuzzList, err = GenerateFizzBuzz("fizz", "buzz", 0, 5, 100)
	if err != ErrInvalidInt {
		t.Errorf("GenerateFizzBuzz should have returned `ErrInvalidInt` when `int1` <= 0.")
	}
	fizzBuzzList, err = GenerateFizzBuzz("fizz", "buzz", 3, 0, 100)
	if err != ErrInvalidInt {
		t.Errorf("GenerateFizzBuzz should have returned `ErrInvalidInt` when `int2` <= 0.")
	}
	fizzBuzzList, err = GenerateFizzBuzz("fizz", "buzz", 3, 5, 0)
	if err != ErrInvalidInt {
		t.Errorf("GenerateFizzBuzz should have returned `ErrInvalidInt` when `limit` <= 0.")
	}
	fizzBuzzList, err = GenerateFizzBuzz("fizz", "buzz", 3, 5, 20)
	if err != nil {
		t.Errorf("GenerateFizzBuzz shouldn't have returned an error.")
	}
	if len(fizzBuzzList) != 20 {
		t.Errorf("fizzBuzzList should contain exactly 20 strings")
	}
	resultList := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13",
		"14", "fizzbuzz", "16", "17", "fizz", "19", "buzz"}
	for i := 0; i < 20; i++ {
		if resultList[i] != fizzBuzzList[i] {
			t.Errorf("#%d : expected '%s' is '%s'", i+1, resultList[i], fizzBuzzList[i])
		}
	}
}
