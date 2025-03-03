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

	res := (a == b && a!= c) || (a == c && a != b) || (b == c && b != a)
	fmt.Printf("Result: %t\n", res)
}