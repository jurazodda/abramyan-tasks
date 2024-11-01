package main

import (
	"fmt"
)

func main() {
	var a float64

	fmt.Print("Enter the side of the square: ")
	fmt.Scan(&a)
	
	p := 4 * a
	fmt.Printf("Perimeter of the square: %.2f\n", p)
}
