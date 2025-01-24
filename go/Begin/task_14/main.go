package main

import "fmt"

func main() {
	const p = 3.14
	var L float64
	fmt.Scan(&L)

	R := L / (2 * p)
	S := p * R * R
	
	fmt.Println(R)
	fmt.Println(S)
}
