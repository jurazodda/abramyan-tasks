package main

import "fmt"

func main() {
	var X, A, Y, B float64
	fmt.Scan(&X, &A, &Y, &B)

	res1 := A / X
	res2 := B / Y
	res := res1 / res2
	
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res)
}
