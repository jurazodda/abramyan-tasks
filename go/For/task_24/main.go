package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var n int
	fmt.Scan(&x, &n)
	s, f, one, sum := 1.0, 1.0, 1.0, 1.0

	for i := 1.0; i < float64(n * 2); i++ {
		f = f * i
		s = s * x
		if int(i) % 2 == 0 {
			one *= -1
			sum = sum + one * s / f
		}
	}
	fmt.Println(sum)
	fmt.Println(math.Cos(x))
}