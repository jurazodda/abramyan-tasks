package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("X: ")
	fmt.Scan(&x)
	fmt.Print("Y: ")
	fmt.Scan(&y)

	res := (x > 0 && y > 0) || (x < 0 && y < 0)
	fmt.Printf("Result: %t\n", res)
}
