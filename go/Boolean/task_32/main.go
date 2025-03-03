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

	res := (a*a + b*b == c*c) || (a*a + c*c == b*b) || (c*c + b*b == a*a)
	fmt.Printf("Result: %t\n", res)
}
