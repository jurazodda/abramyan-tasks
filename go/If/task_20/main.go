package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, ab, ac float64
	fmt.Scan(&a, &b, &c)

	ac = math.Abs(a - c)
	ab = math.Abs(a - b)
	if  ac > ab {
		fmt.Println(b, ab)
	} else {
		fmt.Println(c, ac)
	}
}