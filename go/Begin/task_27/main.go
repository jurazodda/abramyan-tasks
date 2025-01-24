package main

import "fmt"

func main() {
	var A float64
	fmt.Scan(&A)

	A = A * A
	fmt.Println("A^2:", A)
	A = A * A
	fmt.Println("A^4:", A)
	A = A * A
	fmt.Println("A^8:", A)
}
