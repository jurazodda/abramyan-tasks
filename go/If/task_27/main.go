package main

import "fmt"

func main() {
	var x, f float64
	fmt.Scan(&x)

	if x < 0 {
		f = 0
	} else if int(x) % 2 == 0{
		f = 1
	} else if int(x) % 2 == 1 {
		f = -1
	}
	fmt.Println(f)
}