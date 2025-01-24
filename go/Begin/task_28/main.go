package main

import "fmt"

func main() {
	var A, B float64
	fmt.Scan(&A)

	B = A * A
	fmt.Println("A^2:", B)
	A = B * A
	fmt.Println("A^3:", A)
	B = A * B
	fmt.Println("A^5:", B)
	A = B * B
	fmt.Println("A^10:", A)
	B = A * B
	fmt.Println("A^15:", B)
}
