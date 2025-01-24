package main

import "fmt"

func main() {
	const p = 3.14
	var d float64
	fmt.Scan(&d)

	L := p * d
	fmt.Println(L)
}
