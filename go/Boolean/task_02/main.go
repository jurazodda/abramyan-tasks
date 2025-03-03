package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scan(&num)

	res := num%2 == 1
	fmt.Printf("Result: %t\n", res)
}
