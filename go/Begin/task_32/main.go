package main

import "fmt"

func main() {
	var Tc float64
	fmt.Scan(&Tc)

	Tf := Tc * 9 / 5 + 32
	fmt.Println(Tf)
}
