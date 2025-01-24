package main

import (
	"fmt"
	"math"
)

func main() {
	const p = 3.14
	var S float64
	fmt.Scan(&S)

	R := math.Sqrt(S / p)
	D := 2 * R
	L := 2 * p * R
	
	fmt.Println(D)
	fmt.Println(L)
}
