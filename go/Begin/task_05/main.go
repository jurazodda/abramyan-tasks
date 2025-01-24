package main

import "fmt"

func main() {
	var a float64
	fmt.Scan(&a)

	V := a * a * a
	S := 6 * a * a

	fmt.Println(V)
	fmt.Println(S)
}
