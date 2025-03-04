package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	if a < b && b < c {
		a *= 2
		b *= 2
		c *= 2 
	} else {
		a = -a
		b = -b
		c = -c
	}
	fmt.Println(a, b, c)
}