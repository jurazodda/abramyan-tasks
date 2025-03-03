package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)
	fmt.Print("Enter the third number: ")
	fmt.Scan(&num3)

	res := (num1 + num2 == 0) || (num1 + num3 == 0) || (num2 + num3 == 0)
	fmt.Printf("Result: %t\n", res)
}
