package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		n1, _ := strconv.Atoi(os.Args[1])
		operator := os.Args[2]
		n2, _ := strconv.Atoi(os.Args[3])

		fmt.Println(calc(n1, n2, operator))
	} else {
		fmt.Println("Add argument!")
	}
}

func calc(n1, n2 int, operator string) any {
	operand := map[string]int{
		"*": n1 * n2,
		"/": func() int {
			if n2 != 0 {
				return n1 / n2
			}
			return 0
		}(),
		"+": n1 + n2,
		"-": n1 - n2,
		"%": func() int {
			if n2 != 0 {
				return n1 % n2
			}
			return 0
		}(),
	}

	if result, availOption := operand[operator]; availOption {
		return result
	}

	return "invalid operator"
}
