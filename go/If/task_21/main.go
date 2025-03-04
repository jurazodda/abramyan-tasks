package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x == 0 && y == 0 {
		fmt.Println(0)
	} else if x == 0 && y != 0 {
		fmt.Println(1)
	} else if y == 0 && x != 0 {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}
}