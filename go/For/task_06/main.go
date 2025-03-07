package main

import "fmt"

func main() {
	var k, i, res float64
	fmt.Scan(&k)

	for i = 1.2; i <= 2; i += 0.2 {
		res = k * i
		fmt.Printf("%.2f\n", res)
	}
}