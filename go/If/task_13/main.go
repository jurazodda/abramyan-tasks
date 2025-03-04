package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Scan(&num1, &num2, &num3)

	if num1 < num2 && num2 < num3 {
		fmt.Println(num2)
	} else if num2 < num1 && num1 < num3 {
		fmt.Println(num1)
	} else {
		fmt.Println(num3)
	}
}