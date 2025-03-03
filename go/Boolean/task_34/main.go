package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("X: ")
	fmt.Scan(&x)
	fmt.Print("Y: ")
	fmt.Scan(&y)

	res := (x + y) % 2 == 1
	fmt.Printf("Result: %t\n", res)
}