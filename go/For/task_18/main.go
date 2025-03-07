package main

import (
	"fmt"
)

func main() {	
	var a  float64 
	var n int
	var sum float64
	fmt.Scan(&a, &n) 
	var s float64 = 1

	for i := 0; i <= n; i++{
		sum = sum + s
		s *= a
	}
	fmt.Println(sum)
}