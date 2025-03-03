package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 float64
	fmt.Print("x1: ")
	fmt.Scan(&x1)
	fmt.Print("y1: ")
	fmt.Scan(&y1)
	fmt.Print("x2: ")
	fmt.Scan(&x2)
	fmt.Print("y2: ")
	fmt.Scan(&y2)

	// var i, k float64
	// for i = 0; i < 8; i++ {
	// 	for k = 0; k < 8; k++{
	// 		if math.Abs(x1 - k) * math.Abs(y1 - i) == 2 {
	// 			fmt.Print("* ")
	// 		} else if x1 == k && y1 == i {
	// 			fmt.Print("+ ")
	// 		} else {
	// 			fmt.Print(". ")
	// 		}
	// 	}
	// 	fmt.Println()
	//}

	res := math.Abs(x2 - x1) * math.Abs(y2 - y1) == 2
	fmt.Printf("Result: %t\n", res)
}
