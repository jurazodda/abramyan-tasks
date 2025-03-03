package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("A: ")
	fmt.Scan(&a)
	fmt.Print("B: ")
	fmt.Scan(&b)

	res := a%2 == 1 || b%2 == 1
	fmt.Printf("Result: %t\n", res)
}
