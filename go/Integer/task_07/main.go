package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	r1 := num / 10
	r2 := num % 10

	sum := r1 + r2
	prod := r1 * r2

	fmt.Println(sum)
	fmt.Println(prod)
}
