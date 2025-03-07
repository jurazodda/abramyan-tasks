package main

import "fmt"

func main() {
	var a, b, sum int
	fmt.Scan(&a, &b)

	for i := a; i <= b; i++ {
		sum = sum + i*i
	}
	fmt.Println(sum)
}