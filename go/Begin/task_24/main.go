package main

import "fmt"

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	A, B, C = B, C, A
	fmt.Println(A, B, C)
}
