package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Print("Enter the first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&b)

	geomAvg := math.Sqrt(a * b)
	fmt.Printf("Geometric mean: %.2f\n", geomAvg)
}