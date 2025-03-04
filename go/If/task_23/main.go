package main

import "fmt"

func main() {
	var x1, y1, x2, y2,	x3, y3, x4, y4 int
	fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3)

	if x1 != x2 && x2 == x3 {
		x4 = x1
	} else if x2 != x1 && x1 == x3 {
		x4 = x2
	} else {
		x4 = x3
	} 

	if y1 != y2 && y2 == y3 {
		y4 = y1
	} else if y2 != y1 && y1 == y3 {
		y4 = y2
	} else {
		y4 = y3
	}

	fmt.Println(x4, y4)
}