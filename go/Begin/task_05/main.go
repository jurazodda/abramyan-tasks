package main

import "fmt"

func main() {
	var a float64
	fmt.Print("Enter the edge of the cube: ")
	fmt.Scan(&a)

	v := a * a * a
	fmt.Printf("Volume of the cube: %v\n", v)

	s := 6 * a * a
	fmt.Printf("Surface area of the cube: %v\n", s)
}