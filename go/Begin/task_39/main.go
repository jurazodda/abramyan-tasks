package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	D := B*B - 4*A*C
	x1 := ((-B) + math.Sqrt(D)) / (2 * A)
	x2 := ((-B) - math.Sqrt(D)) / (2 * A)

	fmt.Println(x1)
	fmt.Println(x2)
}
