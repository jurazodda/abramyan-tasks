package main

import "fmt"

func main() {
	var a float64
	fmt.Print("Enter the side of the square: ")
	fmt.Scan(&a)
	
	s := a * a
	fmt.Printf("Area of the square: %.2f\n", s)
}