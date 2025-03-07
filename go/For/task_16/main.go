package main

import "fmt"

func main() {
	var a float64
	var n int
	res := 1.0
	fmt.Scan(&a, &n)

	for i := 0; i < n; i++ {
		res *= a
		fmt.Println(res)
	}
}