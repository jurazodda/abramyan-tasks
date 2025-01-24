package main

import "fmt"

func main() {
	var V1, V2, S, T float64
	fmt.Scan(&V1, &V2, &S, &T)

	res := S + T * (V1 + V2)
	fmt.Println(res)
}
