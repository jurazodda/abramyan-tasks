package main

import "fmt"

func main() {
	var x float64
	var n int
	res, s, sum := 1.0, 1.0, 1.0
	fmt.Scan(&x, &n)

	for i := 1.0; i < float64(n); i++ {
		s = s * x
		res =  res * i
		sum = sum + s / res
	}
	fmt.Println(sum)
}