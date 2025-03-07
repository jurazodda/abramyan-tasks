package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := 1.0

	for i := 1.1; i <= float64(n); i += 0.1 {
		res *= i
	}
	fmt.Printf("Result: %.2f\n", res)
}
