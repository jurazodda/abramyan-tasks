package main

import (
	"fmt"
)

func main() {
	var a, b, c float64

	fmt.Print("Enter the edge lengths of the rectangular parallelepiped: ")
	fmt.Scan(&a, &b, &c)

	v := a * b * c
	fmt.Printf("Volume of the rectangular parallelepiped: %.2f\n", v)

	s := 2 * (a*b + b*c + a*c)
	fmt.Printf("Surface area of the rectangular parallelepiped: %.2f\n", s)
}
