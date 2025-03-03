package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	r1 := num / 10
	r2 := num % 10
	res := r2*10 + r1

	fmt.Println(res)
}
