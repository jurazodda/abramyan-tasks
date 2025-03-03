package main

import "fmt"

func main() {
	var x, y, x1, y1, x2, y2 int
	fmt.Print("X1: ")
	fmt.Scan(&x1)
	fmt.Print("Y1: ")
	fmt.Scan(&y1)
	fmt.Print("X2: ")
	fmt.Scan(&x2)
	fmt.Print("Y2: ")
	fmt.Scan(&y2)

	res := (x1 < x && x < x2) && (y2 < y && y < y1)
	fmt.Printf("Result: %t\n", res)
}