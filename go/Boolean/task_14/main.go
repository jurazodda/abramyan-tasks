package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("A: ")
	fmt.Scan(&a)
	fmt.Print("B: ")
	fmt.Scan(&b)
	fmt.Print("C: ")
	fmt.Scan(&c)

	res := (a>0 && b<0 && c<0) || (a<0 && b>0 && c<0) || (a<0 && b<0 && c>0)
	fmt.Printf("Result: %t\n", res)
}
