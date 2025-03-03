package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	res := (N % 3600) / 60
	fmt.Println(res)
}
