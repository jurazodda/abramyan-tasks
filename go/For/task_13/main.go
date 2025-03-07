package main

import (
	"fmt"
)

func main() {
	var n int
	var res float64
	s := 1.1
	fmt.Scan(&n)

	/*
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				res += s
			} else {
				res -= s
			}

			s += 0.1
		}
	*/

	for i := 0; i < n; i++ {
		res += ((-1.0) * s * float64((i+1)%2) + s * float64(i%2))
		s += 0.1
	}

	fmt.Printf("%.2f\n", res)
}
