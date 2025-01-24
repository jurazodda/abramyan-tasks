package main

import "fmt"

func main() {
	var A1, B1, C1, A2, B2, C2 float64
	fmt.Scan(&A1, &B1, &C1, &A2, &B2, &C2)

	D := A1*B2 - A2*B1
	x := (C1*B2 - C2*B1) / D
	y := (A1*C2 - A2*C1) / D

	fmt.Println("x =", x)
	fmt.Println("y =", y)
}
