package main

import (
	"errors"
	"fmt"
	"strconv"
	"challenge-2/result"
)

func parseAge(s string) result.Result[int] {
	n, err := strconv.Atoi(s)
	if err != nil {
		return result.Err[int](fmt.Errorf("Parse age: %w", err))
	}
	if n < 0 {
		return result.Err[int](errors.New("Age must be non-negative."))
	}
	return result.Ok(n)
}

func main() {
	ok := parseAge("23")
	bad := parseAge("abcde")

	fmt.Println(ok.IsOk(), ok.Unwrap())
	fmt.Println(bad.UnwrapOr(0))

	doubled := result.Map(ok, func (n int) int { return n*2 })

	fmt.Println(doubled.Unwrap())

	if age, err := bad.Get(); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Age: ", age)
	}
}