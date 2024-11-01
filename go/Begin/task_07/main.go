package main

import (
	"fmt"
)

func main() {
	var r float64
	const p float64 = 3.14

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&r)

	l := 2 * p * r
	fmt.Printf("Circumference of the circle: %.2f\n", l)

	s := p * r*r 
	fmt.Printf("Area of the circle: %.2f\n", s)
}