package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Print("Enter the area of the circle: ")
	fmt.Scan(&s)
	const p float64 = 3.14

	r := math.Sqrt(s / p)
	d := 2 * r
	fmt.Printf("Diameter of the circle: %.2f\n", d)

	l := 2 * p * r
	fmt.Printf("The length of the circle: %.2f\n", l)
}