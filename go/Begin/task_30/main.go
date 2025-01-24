package main

import "fmt"

func main() {
	const p = 3.14
	var a float64
	fmt.Scan(&a)

	g := a * 180/p
	fmt.Printf("G: %f\n", g)
}
