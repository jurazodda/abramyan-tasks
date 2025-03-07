package main

import "fmt"

func main() {
	var n, sum float64
	res := 1.0
	fmt.Scan(&n)

	for i := 1.0; i <= n; i++ {
		res = res * i
		sum += 1 / res
	}
	fmt.Println(sum)
}