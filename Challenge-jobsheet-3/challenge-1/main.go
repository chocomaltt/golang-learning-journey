package main

import "fmt"

func main() {
	items := []string{"Pencil", "Paper", "Sticky Notes"}
	
	stock := make(map[string]int)

	for _,name := range items {
		stock[name]++
	}

	for name, amount := range stock{
		fmt.Printf("%s: %d\n", name, amount)
	}
}