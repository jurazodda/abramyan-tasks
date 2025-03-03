package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scan(&num)

	r1 := num / 1000
	r2 := num / 100 % 10
	r3 := num / 10 % 10
	r4 := num % 10

	res := r1 == r4 && r2 == r3
	fmt.Printf("Result: %t\n", res)
}