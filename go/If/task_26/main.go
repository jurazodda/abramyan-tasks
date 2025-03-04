package main

import "fmt"

func main() {
	var x, f float64
	fmt.Scan(&x)

	if x <= 0 {
		f = -x
	} else if 0 < x && x < 2 {
		f = x*x
	} else if x >= 2 {
		f = 4
	}
	fmt.Println(f)
}