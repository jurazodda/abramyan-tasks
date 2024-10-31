package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	absNum1 := math.Abs(num1)
	absNum2 := math.Abs(num2)

	sum := absNum1 + absNum2
	fmt.Printf("Sum: %.2f\n", sum)

	difference := absNum1 - absNum2
	fmt.Printf("Difference: %.2f\n", difference)

	product := absNum1 * absNum2
	fmt.Printf("Product: %.2f\n", product)

	division := absNum1 / absNum2
	fmt.Printf("Division: %.2f\n", division)
}
