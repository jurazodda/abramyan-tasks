package main

import "fmt"

func main() {
	var num1, num2, num3, num4 int
	fmt.Scan(&num1, &num2, &num3, &num4)

	if num1 != num2 && num1 != num3 && num1 != num4 {
		fmt.Println(1)
	} else if num2 != num1 && num2 != num3 && num2 != num4 {
		fmt.Println(2)
	} else if num3 != num1 && num3 != num2 && num3 != num4 {
		fmt.Println(3)
	} else {
		fmt.Println(4)
	}
}