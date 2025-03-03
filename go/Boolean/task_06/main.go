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

	res := a < b && b < c
	fmt.Printf("Result: %t\n", res)
}
