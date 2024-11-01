package main

import (
	"fmt"
)

func main() {
	var a, b float64

	fmt.Print("Enter the first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&b)

	avg := (a + b) / 2
	fmt.Printf("Avarage value: %.2f\n", avg)
}