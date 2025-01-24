package main

import "fmt"

func main() {
	var A, B float64
	fmt.Scan(&A, &B)

	x := -B / A
	fmt.Println(x)
}
