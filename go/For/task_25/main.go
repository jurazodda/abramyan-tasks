package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var n int
	s, res, one := 1.0, 0.0, 1.0
	fmt.Scan(&x, &n)

	for i := 1; i <= n; i++ {
		s *= x
		res += one * s / float64(i)
		one *= -1
	}
	fmt.Println(res)
	fmt.Println(math.Log(1 + x))
}
