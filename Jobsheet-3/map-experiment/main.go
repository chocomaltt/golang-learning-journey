package main

import "fmt"

func main() {
	stock := map[string]int{
		"apple": 10,
		"orange": 5,
	}

	fmt.Println("Stock: ", stock)

	stock["mango"] = 7
	stock["apple"] = stock["apple"] + 2
	fmt.Println("Updated stock: ", stock)

	amount, ok := stock["banana"]
	fmt.Println("Banana: ", amount, "Available: ", ok)

	delete(stock, "orange")
	fmt.Println("After delete: ", stock)

	for name, amount := range stock {
		fmt.Printf("%s = %d\n", name, amount)
	}
}