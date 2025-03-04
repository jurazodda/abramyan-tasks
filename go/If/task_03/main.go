package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if num > 0 {
		num = num + 1
	} else if num < 0 {
		num = num - 2 
	} else {
		num = 10
	}
	fmt.Println(num)
}