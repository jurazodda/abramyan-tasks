package main

import (
	"fmt"
	"math"
)

func main() {
	var x, f float64
	fmt.Scan(&x)

	if x > 0 {
		f = 2 * math.Sin(x)
	} else {
		f = 6 - x
	}
	fmt.Println(f)
}