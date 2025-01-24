package main

import "fmt"

func main() {
	var V, U, T1, T2 float64
	fmt.Scan(&V, &U, &T1, &T2)

	So := V * T1
	Sp := (V - U) * T2
	S := So + Sp
	fmt.Println(S)
}
