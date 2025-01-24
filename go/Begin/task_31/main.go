package main

import "fmt"

func main() {
	var Tf float64
	fmt.Scan(&Tf)

	Tc := (Tf - 32) * 5 / 9
	fmt.Println(Tc)
}
