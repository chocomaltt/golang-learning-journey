package main

import (
	"Jobsheet-1/task"
	"fmt"
)

const appName = "jobsheet-1"

func main() {
	fmt.Println("Hello, Go!")

	var umur int = 23
	nama := "Reza"
	var aktif bool = true
	var rate float64 = 50000

	fmt.Printf("%s | %s, %d th, aktif=%t, rate=%.0f\n", appName, nama, umur, aktif, rate)

	fmt.Println("==========================")

	for i := -1; i <= 10; i++ {
		fmt.Println(i, klasifikasi(i))
	}
	defer fmt.Println("\nTask 1 selesai!")

	fmt.Println("\n==========================")
	fmt.Println("+          TASK          +")
	fmt.Println("==========================")

	const degree float32 = 25
	convert := task.CelsiusConvertion(degree)

	fmt.Printf("Suhu %0.f Celsius sama dengan:\n%0.f Fahrenheit, %0.f Kelvin\n", degree, convert[0], convert[1])

	defer fmt.Println("\nTask 2 selesai!")

	fmt.Println("\n==========================")
	fmt.Println("+        FIZZBUZZ        +")
	fmt.Println("==========================")

	for i := 0; i <= 15; i++ {
		fmt.Print(task.FizzBuzz(i) + "\n")
	}

	defer fmt.Println("\nTask 3 selesai!")
}

func klasifikasi(n int) string {
	switch {
	case n < 0:
		return "negatif"
	case n == 0:
		return "nol"
	default:
		return "positif"
	}
}
