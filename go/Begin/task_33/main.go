package main

import "fmt"

func main() {
	var X, A, Y float64
	fmt.Scan(&X, &A, &Y)

	kg := A / X
	res := kg * Y
	
	fmt.Println(kg)
	fmt.Println(res)
}
