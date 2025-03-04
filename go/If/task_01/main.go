package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	
	if num > 0 {
		num = num + 1 
	}
	fmt.Println(num)
	
}