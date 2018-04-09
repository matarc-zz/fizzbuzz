package main

import (
	"fmt"
)

var (
	ErrInvalidInt  = fmt.Errorf("num1, num2 and limit should be superior to 0")
	ErrEmptyString = fmt.Errorf("str1 and str2 shouldn't be empty strings")
)

type FizzBuzzList struct {
	List []string
}

// Generate a list of strings from 1 to `limit` where multiples of `num1` are replaced by `str1`,
// multiples of `num2` are replaced by `str2` and multiples of `num1`*`num2` are replaced by
// the concatenation of `str1` and `str2`. Returns an error if `limit` <= 0 or if the strings are empty.
func GenerateFizzBuzz(str1, str2 string, num1, num2, limit int) (fizzBuzzList []string, err error) {
	if limit <= 0 || num1 <= 0 || num2 <= 0 {
		err = ErrInvalidInt
		return
	}
	if str1 == "" || str2 == "" {
		err = ErrEmptyString
		return
	}

	fizzBuzzList = make([]string, limit)
	for i := 0; i < limit; i++ {
		if (i+1)%num1 == 0 {
			fizzBuzzList[i] = str1
		}
		if (i+1)%num2 == 0 {
			fizzBuzzList[i] += str2
		}
		if fizzBuzzList[i] == "" {
			fizzBuzzList[i] = fmt.Sprintf("%d", i+1)
		}
	}
	return
}
