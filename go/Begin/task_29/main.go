package main

import "fmt"

func main() {
	const p = 3.14
	var a float64
	fmt.Scan(&a)

	r := p * a / 180
	fmt.Printf("r: %f\n", r)
}
