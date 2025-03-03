package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	res1 := num / 10
	res2 := num % 10

	fmt.Println(res1, res2)
}
