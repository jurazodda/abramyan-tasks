package main

import "fmt"

func main() {
	var x1, x2, y1, y2 int
	fmt.Print("x1: ")
	fmt.Scan(&x1)
	fmt.Print("y1: ")
	fmt.Scan(&y1)
	fmt.Print("x2: ")
	fmt.Scan(&x2)
	fmt.Print("y2: ")
	fmt.Scan(&y2)

	res := (x1 == x2 && y1 == y2 - 1)||
	(y1 == y2 && x1 == x2 - 1)||
	(x1 == x2 && y1 == y2 + 1)||
	(y1 == y2 && x1 == x2 + 1)||
	(y1 == y2 - 1 && x1 == x2 + 1)||
	(y1 == y2 - 1 && x1 == x2 - 1)||
	(y1 == y2 + 1 && x1 == x2 - 1)||
	(y1 == y2 + 1 && x1 == x2 + 1)
	
	fmt.Printf("Result: %t\n", res)
}
