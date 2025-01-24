package main

import (
	"fmt"
	"math"
)

func main() {
	var V1, V2, T, S1 float64
	fmt.Print("V1: ")
	fmt.Scan(&V1)
	fmt.Print("V2: ")
	fmt.Scan(&V2)
	fmt.Print("T: ")
	fmt.Scan(&T)
	fmt.Print("S1: ")
	fmt.Scan(&S1)

	S := math.Abs(S1 - T * (V1 + V2))
	fmt.Printf("S: %f\n", S)
}
