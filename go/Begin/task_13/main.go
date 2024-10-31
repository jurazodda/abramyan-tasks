package main

import (
	"fmt"
)

func main() {
	var r1, r2 float64
	fmt.Print("Enter the radius of the first circle: ")
	fmt.Scan(&r1)
	fmt.Print("Enter the radius of the second circle: ")
	fmt.Scan(&r2)
	const p float64 = 3.14

	if r1 <= r2 {
		fmt.Println("Error: R1 must be greater than R2.")
		return
	}

	s1 := p * r1*r1
	fmt.Printf("Area of the first circle: %.2f\n", s1)

	s2 := p * r2*r2 
	fmt.Printf("Area of the second circle: %.2f\n", s2)

	s3 := s1 - s2
	fmt.Printf("Area of the ring: %.2f\n", s3)
}
