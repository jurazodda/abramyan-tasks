package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	V := a * b * c
	S := 2 * (a*b + b*c + a*c)

	fmt.Println(V)
	fmt.Println(S)
}
