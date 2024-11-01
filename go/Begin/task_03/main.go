package main

import (
	"fmt"
)

func main() {
	var a, b float64

	fmt.Print("Enter the length of the rectangle: ")
	fmt.Scan(&a)
	fmt.Print("Enter the width of the ractangle: ")
	fmt.Scan(&b)

	s := a * b
	fmt.Printf("Area of the rectangle: %.2f\n", s)

	p := 2 * (a + b)
	fmt.Printf("Perimeter of the rectangle: %.2f\n", p)
}