package main

import "fmt"

func main() {
	var nyear uint
	fmt.Scan(&nyear)

	if nyear % 4 == 0 && nyear % 100 == 1 && nyear % 400 == 0 {
		fmt.Println(366)
	} else {
		fmt.Println(365)
	}
}