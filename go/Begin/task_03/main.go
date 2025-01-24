package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	S := a * b
	P := 2 * (a + b)

	fmt.Println(S)
	fmt.Println(P)
}
