package main

import "fmt"

func main() {
	var num1, num2, num3, pos, neg int
	fmt.Scan(&num1, &num2, &num3)

	if num1 > 0 {
		pos++
	} else {
		neg++
	}

	if num2 > 0 {
		pos++
	} else {
		neg++
	}

	if num3 > 0 {
		pos++
	} else {
		neg++
	}
	fmt.Println(pos)
	fmt.Println(neg)
}