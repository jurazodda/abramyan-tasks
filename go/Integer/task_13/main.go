package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	r1 := num / 100
	r2 := num / 10 % 10
	r3 := num % 10

	res := r2*100 + r3*10 + r1
	fmt.Println(res)
}
