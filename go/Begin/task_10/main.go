package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	sum := num1 + num2
	fmt.Printf("Sum: %.2f\n", sum)

	difference := num1 - num2
	fmt.Printf("Difference: %.2f\n", difference)

	product := num1 * num2
	fmt.Printf("Product: %.2f\n", product)

	quotientOfSquares := (num1 * num1) / (num2 * num2)
	fmt.Printf("Quotient of squares: %.2f\n", quotientOfSquares)
}
