package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	AC := math.Abs(C - A)
	BC := math.Abs(C - B)
	sum := AC + BC
	
	fmt.Println(AC)
	fmt.Println(BC)
	fmt.Println(sum)
}
