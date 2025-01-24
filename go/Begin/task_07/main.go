package main

import "fmt"

func main() {
	const p = 3.14
	var R float64
	fmt.Scan(&R)

	L := 2 * p * R
	S := p * R * R

	fmt.Println(L)
	fmt.Println(S)
}
