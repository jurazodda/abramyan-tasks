package main

import "fmt"

func main() {
	var a, b, n int
	fmt.Scan(&a, &b)

	for i := a-1; i > b; i-- {
		fmt.Printf("%d ", i)
		n++
	}
	fmt.Println()
	fmt.Println(n)
}
