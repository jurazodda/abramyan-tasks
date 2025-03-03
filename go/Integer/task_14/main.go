package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	r1 := num / 100
	r2 := num / 10 % 10
	r3 := num % 10

	res := r3*100 + r1*10 + r2
	fmt.Println(res)
}
