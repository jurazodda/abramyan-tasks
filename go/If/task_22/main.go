package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x > 0 && y > 0 {
		fmt.Println("I")
	} else if x < 0 && y > 0 {
		fmt.Println("II")
	} else if x < 0 && y < 0 {
		fmt.Println("III")
	} else if x > 0 && y < 0 {
		fmt.Println("IV")
	}
}