package main

import "fmt"

func main() {
	var n float64
	var res = 1.0
	fmt.Scan(&n)

	for i := 1.0; i <= n; i++ {
		res = i * res
	}
	fmt.Println(res)
}