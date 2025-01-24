package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)

	a := math.Abs(x1 - x2)
	b := math.Abs(y1 - y2)
	p := 2 * (a + b)
	s := a * b
	
	fmt.Println(p)
	fmt.Println(s)
}
