package main

import (
	"fmt"
)

func main() {
	var l float64
	const p float64 = 3.14

	fmt.Print("Enter the lenght of the circumference: ")
	fmt.Scan(&l)
	
	r := l / 2 * p
	fmt.Printf("Circle radius: %.2f\n", r)

	s := p * r*r
	fmt.Printf("Circle area: %.2f\n", s)
}