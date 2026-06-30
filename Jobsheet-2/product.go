package main

type Order struct {
	Audit
	ID int
	Amount float64
}

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}