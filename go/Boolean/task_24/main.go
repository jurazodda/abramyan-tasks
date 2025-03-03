package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("A: ")
	fmt.Scan(&a)
	fmt.Print("B: ")
	fmt.Scan(&b)
	fmt.Print("C: ")
	fmt.Scan(&c)

	d := b*b - 4*a*c > 0
	fmt.Printf("Result: %t\n", d)
}