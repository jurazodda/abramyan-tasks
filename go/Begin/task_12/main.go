package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Print("Enter the first leg of the triangle: ")
	fmt.Scan(&a)
	fmt.Print("Enter the second leg of the triangle: ")
	fmt.Scan(&b)

	c := math.Sqrt(a*a + b*b)
	fmt.Printf("Hypotenuse of the triangle: %.2f\n", c)

	p := a + b + c
	fmt.Printf("Perimeter of the triangle: %.2f\n", p)
}