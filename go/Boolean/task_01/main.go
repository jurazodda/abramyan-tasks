package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scan(&num)

	res := num > 0
	fmt.Printf("Result: %t\n", res)
}
