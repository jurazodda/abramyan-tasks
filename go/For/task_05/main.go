package main

import "fmt"

func main() {
	var k, i, res float64
	fmt.Scan(&k)

	for i = 0.1; i <= 1; i += 0.1 {
		res = i * k
		fmt.Printf("%.2f \n", res)
	}

}