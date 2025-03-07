package main

import "fmt"

func main() {
	var a, sum float64
	s := 1.0
	var n int
	fmt.Scan(&a, &n)
	
	for i := 0; i <= n; i++ {
		sum = sum + s
		s *= a
	}
	fmt.Println(sum)
}