package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	fmt.Scan(&num1, &num2)

	sum := num1 + num2
	dif := num1 - num2
	prod := num1 * num2
	quotientOfSquare := (num1 * num1) / (num2 * num2)

	fmt.Println(sum)
	fmt.Println(dif)
	fmt.Println(prod)
	fmt.Println(quotientOfSquare)
}
