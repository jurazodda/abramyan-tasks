package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 float64
	fmt.Print("x1: ")
	fmt.Scan(&x1)
	fmt.Print("y1: ")
	fmt.Scan(&y1)
	fmt.Print("x2: ")
	fmt.Scan(&x2)
	fmt.Print("y2: ")
	fmt.Scan(&y2)

    res := math.Abs(x2 - x1) == math.Abs(y2 - y1)
    fmt.Printf("Result: %t\n", res)
}
