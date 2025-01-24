package main

import (
	"fmt"
)

func main() {
	var A, B float64
	fmt.Scan(&A, &B)

	A, B = B, A
	fmt.Println(A, B)
}
