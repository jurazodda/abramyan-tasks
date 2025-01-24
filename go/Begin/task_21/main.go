package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3)

	a := math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
	b := math.Sqrt(math.Pow(x1-x3, 2) + math.Pow(y1-y3, 2))
	c := math.Sqrt(math.Pow(x2-x3, 2) + math.Pow(y2-y3, 2))

	p := (a + b + c) / 2
	S := math.Sqrt(p * (p - a) * (p - b) * (p - c))

	fmt.Printf("P: %f\n", p*2)
	fmt.Printf("S: %f\n", S)
}
