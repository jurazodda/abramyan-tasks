package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	AC := math.Abs(C - A)
	BC := math.Abs(B - C)
	res := AC * BC

	fmt.Println(res)
}
