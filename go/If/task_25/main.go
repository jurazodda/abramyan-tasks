package main

import "fmt"

func main() {
	var x, f int
	fmt.Scan(&x)

	if x < -2 || x > 2 {
		f = 2 * x
	} else {
		f = -3 * x
	}
	fmt.Println(f)
}