package main

import "fmt"

func main() {
	var k, n int
	fmt.Scan(&k, &n)

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", k)
	}
	fmt.Println()
}