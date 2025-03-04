package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Scan(&num1, &num2, &num3)
	cnt := 0
	if num1 > 0 {
		cnt++
	} 
	if num2 > 0 {
		cnt++
	}
	if num3 > 0 {
		cnt++
	} 
	fmt.Println(cnt)
}