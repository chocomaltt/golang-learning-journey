package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	panjang, lebar float64
}

type circle struct {
	radius float64
}

func (r rectangle) area () float64 {
	return r.panjang*r.lebar
}

func (r rectangle) perimeter () float64 {
	return 2 *  r.panjang + 2 * r.lebar
}

func (c circle) area () float64 {
	return math.Phi * math.Pow(c.radius, 2)
}

func (c circle) perimeter () float64 {
	return 2 * math.Phi * c.radius
}

func measure (s Shape) {
	if  c, ok := s.(circle);ok {
		fmt.Println("========== Lingkaran ==========")
		fmt.Println("Jari-jari/Radius: ", c.radius)
	} else {
		fmt.Println("========== Persegi Panjang ==========")
	}
	fmt.Println(s)
	fmt.Println("Luas/Area: ",s.area())
	fmt.Printf("Keliling/Perimiter: %0.f \n",s.perimeter())
}

func main() {
	rect := rectangle{panjang: 10, lebar: 5}
	circ := circle{radius: 7}

	measure(rect)
	measure(circ)
}