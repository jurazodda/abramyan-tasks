package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	r1 := num / 100
	r2 := num / 10 % 10
	r3 := num % 10

	sum := r1 + r2 + r3
	prod := r1 * r2 * r3

	fmt.Println(sum)
	fmt.Println(prod)
}
