package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("A: ")
	fmt.Scan(&a)
	fmt.Print("B: ")
	fmt.Scan(&b)

	res := a >= 0 || b < -2
	fmt.Printf("Result: %t\n", res)
}
