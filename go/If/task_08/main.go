package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	if num1 > num2 {
		fmt.Println(num1)
		fmt.Println(num2)
	} else if num2 > num1 {
		fmt.Println(num2)
		fmt.Println(num1)
	}
}