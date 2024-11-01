package main

import (
	"fmt"
)

func main() {
	var d float64
	const p float64 = 3.14

	fmt.Print("Enter the diameter of the circle: ")
	fmt.Scan(&d)

	l := p * d
	fmt.Printf("Circumference of the circle: %.2f\n", l)
}