package main

import "fmt"

func main() {
	var n, res int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		res += 2 * i - 1
	}
	fmt.Println(res)
}