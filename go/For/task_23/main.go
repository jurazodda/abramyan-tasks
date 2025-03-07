package main

import (
	"fmt"
	"math"
)

func main() {

	var x, sum float64
	var n int

	f, s, one := 1.0, 1.0, -1.0
	fmt.Scan(&x, &n)

	for i := 1.0; i < float64(2*n+1); i++ {
		f = f * i
		s = s * x
		if int(i)%2 == 1 {
			one *= -1.0
			sum = sum + one*s/f
		}
	}
	fmt.Println(sum)
	fmt.Println(math.Sin(x))
}
