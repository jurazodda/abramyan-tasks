package main

import "fmt"

func main() {
	var num int 
	fmt.Scan(&num)

	res := num / 1000 % 10
	fmt.Println(res)
}
