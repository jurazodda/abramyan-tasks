package main

import "fmt"

func main() {
	var r float64
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&r)
	const p float64 = 3.14

	l := 2 * p * r
	fmt.Printf("Circumference of the circle: %v\n", l)

	s := p * r*r 
	fmt.Printf("Area of the circle: %.2f\n", s)
}