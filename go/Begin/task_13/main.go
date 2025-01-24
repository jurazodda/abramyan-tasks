package main

import "fmt"

func main() {
	const p = 3.14
	var R1, R2 float64
	fmt.Scan(&R1, &R2)

	S1 := p * (R1 * R1)
	S2 := p * (R2 * R2)
	S3 := S1 - S2

	fmt.Println(S1)
	fmt.Println(S2)
	fmt.Println(S3)
}
