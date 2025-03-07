package main

import "fmt"

func main() {
	var k, i, res float64
	fmt.Scan(&k)

	for i = 1; i <= 10; i++ {
		res = i * k
		fmt.Printf("%.2f \n", res)
	}
}