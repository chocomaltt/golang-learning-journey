package main

import "fmt"

func Change[T any, U any] (in []T, f func(T) U) []U {
	out := make([]U, 0, len(in))
	for _, v := range in {
		out = append(out, f(v))
	}
	return out
}

type Number interface {
	~int | ~float64
}

func Sum[T Number](xs []T) T {
	var total T
	for _,x := range xs {
		total += x
	}
	return total
}

// Challenge No.2
func Filter[T any](in []T, pred func(T) bool) []T {
	filtered := make([]T, 0, len(in))
	for _,x := range in {
		if pred(x) {
			filtered = append(filtered,x)
		}
	}
	return filtered
}

func main() {
	number := []int{1,2,3}
	power := Change(number, func(n int) int { return n*n })
	fmt.Println("Power of 2: ", power)

	word := []string{"a", "bb", "ccc"}
	length := Change(word, func(s string) int { return len(s) })
	fmt.Println("Length each word: ", length)

	fmt.Println("Total: ", Sum(number))

	// Challenge No.2
	evens := Filter(number, func(n int) bool { return n%2==0})
 
	fmt.Printf("Even number: ")
	for _,n := range evens {
		if n == number[len(number)-1]{
			fmt.Printf("%d", n)
		} else {
			fmt.Printf("%d, ",n)
		}
	}
}