package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	fmt.Scan(&num1, &num2)

	absNum1 := math.Abs(num1)
	absNum2 := math.Abs(num2)

	sum := absNum1 + absNum2
	dif := absNum1 - absNum2
	prod := absNum1 * absNum2
	div := absNum1 / absNum2

	fmt.Println(sum)
	fmt.Println(dif)
	fmt.Println(prod)
	fmt.Println(div)
}
