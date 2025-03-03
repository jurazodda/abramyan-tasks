package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scan(&num)

	res := (num%2 == 0) && (num>= 10 && num<= 99)
	fmt.Printf("Result: %t\n", res)
}