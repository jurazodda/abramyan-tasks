package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	avg := math.Sqrt(a * b)
	fmt.Println(avg)
}
